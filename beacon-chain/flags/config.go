package flags

import (
	"github.com/prysmaticlabs/prysm/shared/cmd"
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
)

// GlobalFlags specifies all the global flags for the
// beacon node.
type GlobalFlags struct {
	UnsafeSync       bool
	DisableDiscv5    bool
	MinimumSyncPeers int
	MaxPageSize      int
	DeploymentBlock  int
}

var globalConfig *GlobalFlags

// Get retrieves the global config.
func Get() *GlobalFlags {
	if globalConfig == nil {
		return &GlobalFlags{}
	}
	return globalConfig
}

// Init sets the global config equal to the config that is passed in.
func Init(c *GlobalFlags) {
	globalConfig = c
}

// ConfigureGlobalFlags initializes the global config.
// based on the provided cli context.
func ConfigureGlobalFlags(ctx *cli.Context) {
	cfg := &GlobalFlags{}
	if ctx.Bool(UnsafeSync.Name) {
		cfg.UnsafeSync = true
	}
	if ctx.Bool(DisableDiscv5.Name) {
		cfg.DisableDiscv5 = true
	}
	cfg.MaxPageSize = ctx.Int(RPCMaxPageSize.Name)
	cfg.DeploymentBlock = ctx.Int(ContractDeploymentBlock.Name)
	configureMinimumPeers(ctx, cfg)

	Init(cfg)
}

func configureMinimumPeers(ctx *cli.Context, cfg *GlobalFlags) {
	cfg.MinimumSyncPeers = ctx.Int(MinSyncPeers.Name)
	maxPeers := int(ctx.Int64(cmd.P2PMaxPeers.Name))
	if cfg.MinimumSyncPeers > maxPeers {
		log.Warnf("Changing Minimum Sync Peers to %d", maxPeers)
		cfg.MinimumSyncPeers = maxPeers
	}
}
