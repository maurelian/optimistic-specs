package bss

import (
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/ethereum-optimism/optimistic-specs/bss/flags"
)

type Config struct {
	/* Required Params */

	// L1EthRpc is the HTTP provider URL for L1.
	L1EthRpc string

	// L2EthRpc is the HTTP provider URL for L1.
	L2EthRpc string

	// SROAddress is the SCC contract address.
	SROAddress string

	// MaxBatchSubmissionTime is the maximum amount of time that we will
	// wait before submitting an under-sized batch.
	MaxBatchSubmissionTime time.Duration

	// PollInterval is the delay between querying L2 for more transaction
	// and creating a new batch.
	PollInterval time.Duration

	// NumConfirmations is the number of confirmations which we will wait after
	// appending new batches.
	NumConfirmations uint64

	// ResubmissionTimeout is time we will wait before resubmitting a
	// transaction.
	ResubmissionTimeout time.Duration

	// Mnemonic is the HD seed used to derive the wallet private keys for both
	// the sequence and proposer. Must be used in conjunction with
	// SequencerHDPath and ProposerHDPath.
	Mnemonic string

	// L2OutputHDPath is the derivation path used to obtain the private key for
	// the l2output transactions.
	L2OutputHDPath string

	/* Optional Params */

	// LogLevel is the lowest log level that will be output.
	LogLevel string
}

// NewConfig parses the Config from the provided flags or environment variables.
// This method fails if ValidateConfig deems the configuration to be malformed.
func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		/* Required Flags */
		L1EthRpc:               ctx.GlobalString(flags.L1EthRpcFlag.Name),
		L2EthRpc:               ctx.GlobalString(flags.L2EthRpcFlag.Name),
		SROAddress:             ctx.GlobalString(flags.SROAddressFlag.Name),
		MaxBatchSubmissionTime: ctx.GlobalDuration(flags.MaxBatchSubmissionTimeFlag.Name),
		PollInterval:           ctx.GlobalDuration(flags.PollIntervalFlag.Name),
		NumConfirmations:       ctx.GlobalUint64(flags.NumConfirmationsFlag.Name),
		ResubmissionTimeout:    ctx.GlobalDuration(flags.ResubmissionTimeoutFlag.Name),
		Mnemonic:               ctx.GlobalString(flags.MnemonicFlag.Name),
		L2OutputHDPath:         ctx.GlobalString(flags.L2OutputHDPathFlag.Name),
		/* Optional Flags */
		LogLevel: ctx.GlobalString(flags.LogLevelFlag.Name),
	}

	err := ValidateConfig(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// ValidateConfig ensures additional constraints on the parsed configuration to
// ensure that it is well-formed.
func ValidateConfig(cfg *Config) error {
	// Sanity check log level.
	_, err := log.LvlFromString(cfg.LogLevel)
	return err
}
