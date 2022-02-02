package l2output

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum-optimism/optimistic-specs/bss/bindings/sro"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

var bigOne = big.NewInt(1)

type Config struct {
	Name     string
	L1Client *ethclient.Client
	L2Client *ethclient.Client
	SROAddr  common.Address
	ChainID  *big.Int
	PrivKey  *ecdsa.PrivateKey
}

type Driver struct {
	cfg            Config
	sroContract    *sro.StateRootOracle
	rawSroContract *bind.BoundContract
	walletAddr     common.Address
}

func NewDriver(cfg Config) (*Driver, error) {
	sroContract, err := sro.NewStateRootOracle(
		cfg.SROAddr, cfg.L1Client,
	)
	if err != nil {
		return nil, err
	}

	parsed, err := abi.JSON(strings.NewReader(
		sro.StateRootOracleABI,
	))
	if err != nil {
		return nil, err
	}

	rawSroContract := bind.NewBoundContract(
		cfg.SROAddr, parsed, cfg.L1Client, cfg.L1Client, cfg.L1Client,
	)

	walletAddr := crypto.PubkeyToAddress(cfg.PrivKey.PublicKey)

	return &Driver{
		cfg:            cfg,
		sroContract:    sroContract,
		rawSroContract: rawSroContract,
		walletAddr:     walletAddr,
	}, nil
}

// Name is an identifier used to prefix logs for a particular service.
func (d *Driver) Name() string {
	return d.cfg.Name
}

// WalletAddr is the wallet address used to pay for batch transaction fees.
func (d *Driver) WalletAddr() common.Address {
	return d.walletAddr
}

// GetBatchBlockRange returns the start and end L2 block heights that need to be
// processed. Note that the end value is *exclusive*, therefore if the returned
// values are identical nothing needs to be processed.
func (d *Driver) GetBatchBlockRange(
	ctx context.Context) (*big.Int, *big.Int, error) {

	name := d.cfg.Name

	callOpts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	// Determine the next uncommitted L2 block number. We do so by transforming
	// the timestamp of the latest committed L2 block into its block number and
	// adding one.
	sroTimestamp, err := d.sroContract.LatestBlockTimestamp(callOpts)
	if err != nil {
		log.Error(name+" unable to get latest block timestamp", "err", err)
		return nil, nil, err
	}
	start, err := d.sroContract.ComputeL2BlockNumber(callOpts, sroTimestamp)
	if err != nil {
		log.Error(name+" unable to compute latest l2 block number", "err", err)
		return nil, nil, err
	}
	start.Add(start, bigOne)

	// Next we need to obtain the current timestamp and the next timestamp at
	// which we will need to submit a state root. The former is done by simply
	// adding the submission interval to the latest committed block's timestamp;
	// the latter inspects the timestamp of the latest block.
	nextTimestamp, err := d.sroContract.NextTimestamp(callOpts)
	if err != nil {
		log.Error(name+" unable to get next block timestamp", "err", err)
		return nil, nil, err
	}
	latestHeader, err := d.cfg.L1Client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Error(name+" unable to retrieve latest header", "err", err)
		return nil, nil, err
	}
	currentTimestamp := big.NewInt(int64(latestHeader.Time))

	// If the submission window has yet to elapsed, we must wait before
	// submitting our L2 output commitment. Return start as the end value which
	// will signal that there is no work to be done.
	if currentTimestamp.Cmp(nextTimestamp) < 0 {
		log.Info(name+" submission interval has not elapsed",
			"currentTimestamp", currentTimestamp, "nextTimestamp", nextTimestamp)
		return start, start, nil
	}

	log.Info(name+" submission interval has elapsed",
		"currentTimestamp", currentTimestamp, "nextTimestamp", nextTimestamp)

	// Otherwise the submission interval has elapsed. Transform the next
	// expected timestamp into its L2 block number, and add one since end is
	// exclusive.
	end, err := d.sroContract.ComputeL2BlockNumber(callOpts, nextTimestamp)
	if err != nil {
		log.Error(name+" unable to compute next l2 block number", "err", err)
		return nil, nil, err
	}
	end.Add(end, bigOne)

	return start, end, nil
}

// CraftBatchTx transforms the L2 blocks between start and end into a batch
// transaction using the given nonce.
//
// NOTE: This method SHOULD NOT publish the resulting transaction.
func (d *Driver) CraftBatchTx(
	ctx context.Context,
	start, end, nonce *big.Int,
) (*types.Transaction, error) {

	name := d.cfg.Name

	log.Info(name+" crafting batch tx", "start", start, "end", end,
		"nonce", nonce)

	// Fetch the final block in the range, as this is the only state root we
	// need to submit.
	nextCheckpointBlock := new(big.Int).Sub(end, bigOne)
	checkpointBlock, err := d.cfg.L2Client.HeaderByNumber(
		ctx, nextCheckpointBlock,
	)
	if err != nil {
		return nil, err
	}

	// Fetch the next expected timestamp that we will submit along with the
	// state root.
	callOpts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	timestamp, err := d.sroContract.NextTimestamp(callOpts)
	if err != nil {
		return nil, err
	}

	// Sanity check that we are submitting against the same expected timestamp.
	expCheckpointBlock, err := d.sroContract.ComputeL2BlockNumber(
		callOpts, timestamp,
	)
	if err != nil {
		return nil, err
	}
	if nextCheckpointBlock.Cmp(expCheckpointBlock) != 0 {
		panic(fmt.Sprintf("next expected checkpoint block has changed, "+
			"want: %d, found: %d", nextCheckpointBlock.Uint64(),
			expCheckpointBlock.Uint64()))
	}

	numElements := new(big.Int).Sub(start, end).Uint64()
	log.Info(name+" batch constructed", "num_state_roots", numElements)

	opts, err := bind.NewKeyedTransactorWithChainID(
		d.cfg.PrivKey, d.cfg.ChainID,
	)
	if err != nil {
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.NoSend = true

	return d.sroContract.AppendStateRoot(
		opts, checkpointBlock.Root, timestamp,
	)
}

// SubmitBatchTx using the passed transaction as a template, signs and
// publishes the transaction unmodified apart from sampling the current gas
// price. The final transaction is returned to the caller.
func (d *Driver) SubmitBatchTx(
	ctx context.Context,
	tx *types.Transaction,
) (*types.Transaction, error) {

	opts, err := bind.NewKeyedTransactorWithChainID(
		d.cfg.PrivKey, d.cfg.ChainID,
	)
	if err != nil {
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(tx.Nonce())

	return d.rawSroContract.RawTransact(opts, tx.Data())
}
