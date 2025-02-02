package main

import (
	"context"
	"errors"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/zeta-chain/zetacore/app"
	zetae2econfig "github.com/zeta-chain/zetacore/cmd/zetae2e/config"
	"github.com/zeta-chain/zetacore/e2e/config"
	"github.com/zeta-chain/zetacore/e2e/runner"
)

// NewShowTSSCmd returns the show TSS command
// which shows the TSS address in the network
func NewShowTSSCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-tss [config]",
		Short: "Show address of the TSS",
		RunE:  runShowTSS,
		Args:  cobra.ExactArgs(1),
	}
	return cmd
}

func runShowTSS(_ *cobra.Command, args []string) error {
	// read the config file
	conf, err := config.ReadConfig(args[0])
	if err != nil {
		return err
	}

	// initialize logger
	logger := runner.NewLogger(true, color.FgHiCyan, "")

	// set config
	app.SetConfig()

	// initialize context
	ctx, cancel := context.WithCancel(context.Background())

	// get EVM address from config
	evmAddr := conf.Accounts.EVMAddress
	if !ethcommon.IsHexAddress(evmAddr) {
		cancel()
		return errors.New("invalid EVM address")
	}

	// initialize deployer runner with config
	testRunner, err := zetae2econfig.RunnerFromConfig(
		ctx,
		"tss",
		cancel,
		conf,
		ethcommon.HexToAddress(evmAddr),
		conf.Accounts.EVMPrivKey,
		logger,
	)
	if err != nil {
		cancel()
		return err
	}

	// fetch the TSS address
	if err := testRunner.SetTSSAddresses(); err != nil {
		return err
	}

	// print the TSS address
	logger.Info("TSS EVM address: %s\n", testRunner.TSSAddress.Hex())
	logger.Info("TSS BTC address: %s\n", testRunner.BTCTSSAddress.EncodeAddress())

	return nil
}
