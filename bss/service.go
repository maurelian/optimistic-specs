package bss

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum-optimism/optimistic-specs/bss/txmgr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

// Driver is an interface for creating and submitting batch transactions for a
// specific contract.
type Driver interface {
	// Name is an identifier used to prefix logs for a particular service.
	Name() string

	// WalletAddr is the wallet address used to pay for batch transaction
	// fees.
	WalletAddr() common.Address

	// GetBatchBlockRange returns the start and end L2 block heights that
	// need to be processed. Note that the end value is *exclusive*,
	// therefore if the returned values are identical nothing needs to be
	// processed.
	GetBatchBlockRange(ctx context.Context) (*big.Int, *big.Int, error)

	// CraftBatchTx transforms the L2 blocks between start and end into a batch
	// transaction using the given nonce. A dummy gas price is used in the
	// resulting transaction to use for size estimation.
	//
	// NOTE: This method SHOULD NOT publish the resulting transaction.
	CraftBatchTx(
		ctx context.Context,
		start, end, nonce *big.Int,
	) (*types.Transaction, error)

	// SubmitBatchTx using the passed transaction as a template, signs and
	// publishes the transaction unmodified apart from sampling the current gas
	// price. The final transaction is returned to the caller.
	SubmitBatchTx(
		ctx context.Context,
		tx *types.Transaction,
	) (*types.Transaction, error)
}

type ServiceConfig struct {
	Context         context.Context
	Driver          Driver
	PollInterval    time.Duration
	L1Client        *ethclient.Client
	TxManagerConfig txmgr.Config
}

type Service struct {
	cfg    ServiceConfig
	ctx    context.Context
	cancel func()

	txMgr txmgr.TxManager

	wg sync.WaitGroup
}

func NewService(cfg ServiceConfig) *Service {
	ctx, cancel := context.WithCancel(cfg.Context)

	txMgr := txmgr.NewSimpleTxManager(
		cfg.Driver.Name(), cfg.TxManagerConfig, cfg.L1Client,
	)

	return &Service{
		cfg:    cfg,
		ctx:    ctx,
		cancel: cancel,
		txMgr:  txMgr,
	}
}

func (s *Service) Start() error {
	s.wg.Add(1)
	go s.eventLoop()
	return nil
}

func (s *Service) Stop() error {
	s.cancel()
	s.wg.Wait()
	return nil
}

func (s *Service) eventLoop() {
	defer s.wg.Done()

	name := s.cfg.Driver.Name()

	for {
		select {
		case <-time.After(s.cfg.PollInterval):
			// Determine the range of L2 blocks that the batch submitter has not
			// processed, and needs to take action on.
			log.Info(name + " fetching current block range")
			start, end, err := s.cfg.Driver.GetBatchBlockRange(s.ctx)
			if err != nil {
				log.Error(name+" unable to get block range", "err", err)
				continue
			}

			// No new updates.
			if start.Cmp(end) == 0 {
				log.Info(name+" no updates", "start", start, "end", end)
				continue
			}
			log.Info(name+" block range", "start", start, "end", end)

			// Query for the submitter's current nonce.
			nonce64, err := s.cfg.L1Client.NonceAt(
				s.ctx, s.cfg.Driver.WalletAddr(), nil,
			)
			if err != nil {
				log.Error(name+" unable to get current nonce",
					"err", err)
				continue
			}
			nonce := new(big.Int).SetUint64(nonce64)

			tx, err := s.cfg.Driver.CraftBatchTx(
				s.ctx, start, end, nonce,
			)
			if err != nil {
				log.Error(name+" unable to craft batch tx",
					"err", err)
				continue
			}

			// Construct the transaction submission clousure that will attempt
			// to send the next transaction at the given nonce and gas price.
			sendTx := func(ctx context.Context) (*types.Transaction, error) {
				log.Info(name+" attempting batch tx", "start", start,
					"end", end, "nonce", nonce)

				tx, err := s.cfg.Driver.SubmitBatchTx(ctx, tx)
				if err != nil {
					return nil, err
				}

				log.Info(
					name+" submitted batch tx",
					"start", start,
					"end", end,
					"nonce", nonce,
					"tx_hash", tx.Hash(),
				)

				return tx, nil
			}

			// Wait until one of our submitted transactions confirms. If no
			// receipt is received it's likely our gas price was too low.
			receipt, err := s.txMgr.Send(s.ctx, sendTx)
			if err != nil {
				log.Error(name+" unable to publish batch tx",
					"err", err)
				continue
			}

			// The transaction was successfully submitted.
			log.Info(name+" batch tx successfully published",
				"tx_hash", receipt.TxHash)

		case err := <-s.ctx.Done():
			log.Error(name+" service shutting down", "err", err)
			return
		}
	}
}
