package flags

import (
	"github.com/urfave/cli"
)

const envVarPrefix = "BATCH_SUBMITTER_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	/* Required Flags */

	L1EthRpcFlag = cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   "L1_ETH_RPC",
	}
	L2EthRpcFlag = cli.StringFlag{
		Name:     "l2-eth-rpc",
		Usage:    "HTTP provider URL for L2",
		Required: true,
		EnvVar:   "L2_ETH_RPC",
	}
	SROAddressFlag = cli.StringFlag{
		Name:     "sro-address",
		Usage:    "Address of the SCC contract",
		Required: true,
		EnvVar:   "SRO_ADDRESS",
	}
	MaxBatchSubmissionTimeFlag = cli.DurationFlag{
		Name: "max-batch-submission-time",
		Usage: "Maximum amount of time that we will wait before " +
			"submitting an under-sized batch",
		Required: true,
		EnvVar:   prefixEnvVar("MAX_BATCH_SUBMISSION_TIME"),
	}
	PollIntervalFlag = cli.DurationFlag{
		Name: "poll-interval",
		Usage: "Delay between querying L2 for more transactions and " +
			"creating a new batch",
		Required: true,
		EnvVar:   prefixEnvVar("POLL_INTERVAL"),
	}
	NumConfirmationsFlag = cli.Uint64Flag{
		Name: "num-confirmations",
		Usage: "Number of confirmations which we will wait after " +
			"appending a new batch",
		Required: true,
		EnvVar:   prefixEnvVar("NUM_CONFIRMATIONS"),
	}
	ResubmissionTimeoutFlag = cli.DurationFlag{
		Name: "resubmission-timeout",
		Usage: "Duration we will wait before resubmitting a " +
			"transaction to L1",
		Required: true,
		EnvVar:   prefixEnvVar("RESUBMISSION_TIMEOUT"),
	}
	MnemonicFlag = cli.StringFlag{
		Name: "mnemonic",
		Usage: "The mnemonic used to derive the wallets for either the " +
			"sequencer or the l2output",
		Required: true,
		EnvVar:   prefixEnvVar("MNEMONIC"),
	}
	L2OutputHDPathFlag = cli.StringFlag{
		Name: "l2-output-hd-path",
		Usage: "The HD path used to derive the l2output wallet from the " +
			"mnemonic. The mnemonic flag must also be set.",
		Required: true,
		EnvVar:   prefixEnvVar("L2_OUTPUT_HD_PATH"),
	}

	/* Optional Flags */

	LogLevelFlag = cli.StringFlag{
		Name:   "log-level",
		Usage:  "The lowest log level that will be output",
		Value:  "info",
		EnvVar: prefixEnvVar("LOG_LEVEL"),
	}
)

var requiredFlags = []cli.Flag{
	L1EthRpcFlag,
	L2EthRpcFlag,
	SROAddressFlag,
	SROAddressFlag,
	MaxBatchSubmissionTimeFlag,
	PollIntervalFlag,
	NumConfirmationsFlag,
	ResubmissionTimeoutFlag,
	MnemonicFlag,
	L2OutputHDPathFlag,
}

var optionalFlags = []cli.Flag{
	LogLevelFlag,
}

// Flags contains the list of configuration options available to the binary.
var Flags = append(requiredFlags, optionalFlags...)
