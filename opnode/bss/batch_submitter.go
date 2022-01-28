package bss

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ethereum-optimism/optimistic-specs/opnode/bss/txmgr"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/urfave/cli"
)

const (
	// defaultDialTimeout is default duration the service will wait on
	// startup to make a connection to either the L1 or L2 backends.
	defaultDialTimeout = 5 * time.Second
)

// Main is the entrypoint into the batch submitter service. This method returns
// a closure that executes the service and blocks until the service exits. The
// use of a closure allows the parameters bound to the top-level main package,
// e.g. GitVersion, to be captured and used once the function is executed.
func Main(gitVersion string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		cfg, err := NewConfig(ctx)
		if err != nil {
			return err
		}

		log.Info("Initializing batch submitter")

		batchSubmitter, err := NewBatchSubmitter(cfg, gitVersion)
		if err != nil {
			log.Error("Unable to create batch submitter", "error", err)
			return err
		}

		log.Info("Starting batch submitter")

		if err := batchSubmitter.Start(); err != nil {
			return err
		}
		defer batchSubmitter.Stop()

		log.Info("Batch submitter started")

		<-(chan struct{})(nil)

		return nil
	}
}

// BatchSubmitter is a service that configures the necessary resources for
// running the TxBatchSubmitter and StateBatchSubmitter sub-services.
type BatchSubmitter struct {
	ctx context.Context
}

// NewBatchSubmitter initializes the BatchSubmitter, gathering any resources
// that will be needed by the TxBatchSubmitter and StateBatchSubmitter
// sub-services.
func NewBatchSubmitter(cfg Config, gitVersion string) (*BatchSubmitter, error) {
	ctx := context.Background()

	// Set up our logging to stdout.
	logHandler := log.StreamHandler(os.Stdout, log.TerminalFormat(true))

	logLevel, err := log.LvlFromString(cfg.LogLevel)
	if err != nil {
		return nil, err
	}

	log.Root().SetHandler(log.LvlFilterHandler(logLevel, logHandler))

	// Parse l2output wallet private key and SRO contract address.
	wallet, err := hdwallet.NewFromMnemonic(cfg.Mnemonic)
	if err != nil {
		return nil, err
	}

	_, err = wallet.PrivateKey(accounts.Account{
		URL: accounts.URL{
			Path: cfg.L2OutputHDPath,
		},
	})
	if err != nil {
		return nil, err
	}

	_, err = parseAddress(cfg.SROAddress)
	if err != nil {
		return nil, err
	}

	// Connect to L1 and L2 providers. Perform these last since they are the
	// most expensive.
	l1Client, err := dialL1EthClientWithTimeout(ctx, cfg.L1EthRpc)
	if err != nil {
		return nil, err
	}

	_, err = dialL1EthClientWithTimeout(ctx, cfg.L2EthRpc)
	if err != nil {
		return nil, err
	}

	_, err = l1Client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	_ = txmgr.Config{
		ResubmissionTimeout:  cfg.ResubmissionTimeout,
		ReceiptQueryInterval: time.Second,
		NumConfirmations:     cfg.NumConfirmations,
	}

	return &BatchSubmitter{
		ctx: ctx,
	}, nil
}

func (b *BatchSubmitter) Start() error {
	return nil
}

func (b *BatchSubmitter) Stop() {
}

// dialL1EthClientWithTimeout attempts to dial the L1 provider using the
// provided URL. If the dial doesn't complete within defaultDialTimeout seconds,
// this method will return an error.
func dialL1EthClientWithTimeout(ctx context.Context, url string) (
	*ethclient.Client, error) {

	ctxt, cancel := context.WithTimeout(ctx, defaultDialTimeout)
	defer cancel()

	return ethclient.DialContext(ctxt, url)
}

// parseAddress parses an ETH addres from a hex string. This method will fail if
// the address is not a valid hexidecimal address.
func parseAddress(address string) (common.Address, error) {
	if common.IsHexAddress(address) {
		return common.HexToAddress(address), nil
	}
	return common.Address{}, fmt.Errorf("invalid address: %v", address)
}
