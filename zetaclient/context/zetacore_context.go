package context

import (
	"sort"
	"sync"

	"github.com/rs/zerolog"

	"github.com/zeta-chain/zetacore/pkg/chains"
	lightclienttypes "github.com/zeta-chain/zetacore/x/lightclient/types"
	observertypes "github.com/zeta-chain/zetacore/x/observer/types"
	"github.com/zeta-chain/zetacore/zetaclient/config"
)

// ZetacoreContext contains zetacore context params
// these are initialized and updated at runtime at every height
type ZetacoreContext struct {
	coreContextLock    *sync.RWMutex
	keygen             observertypes.Keygen
	chainsEnabled      []chains.Chain
	evmChainParams     map[int64]*observertypes.ChainParams
	bitcoinChainParams *observertypes.ChainParams
	currentTssPubkey   string
	crosschainFlags    observertypes.CrosschainFlags

	// blockHeaderEnabledChains is used to store the list of chains that have block header verification enabled
	// All chains in this list will have Enabled flag set to true
	blockHeaderEnabledChains []lightclienttypes.HeaderSupportedChain
}

// NewZetacoreContext creates and returns new ZetacoreContext
// it is initializing chain params from provided config
func NewZetacoreContext(cfg config.Config) *ZetacoreContext {
	evmChainParams := make(map[int64]*observertypes.ChainParams)
	for _, e := range cfg.EVMChainConfigs {
		evmChainParams[e.Chain.ChainId] = &observertypes.ChainParams{}
	}

	var bitcoinChainParams *observertypes.ChainParams
	_, found := cfg.GetBTCConfig()
	if found {
		bitcoinChainParams = &observertypes.ChainParams{}
	}

	return &ZetacoreContext{
		coreContextLock:          new(sync.RWMutex),
		chainsEnabled:            []chains.Chain{},
		evmChainParams:           evmChainParams,
		bitcoinChainParams:       bitcoinChainParams,
		crosschainFlags:          observertypes.CrosschainFlags{},
		blockHeaderEnabledChains: []lightclienttypes.HeaderSupportedChain{},
	}
}

// GetKeygen returns the current keygen
func (c *ZetacoreContext) GetKeygen() observertypes.Keygen {
	c.coreContextLock.RLock()
	defer c.coreContextLock.RUnlock()

	var copiedPubkeys []string
	if c.keygen.GranteePubkeys != nil {
		copiedPubkeys = make([]string, len(c.keygen.GranteePubkeys))
		copy(copiedPubkeys, c.keygen.GranteePubkeys)
	}

	return observertypes.Keygen{
		Status:         c.keygen.Status,
		GranteePubkeys: copiedPubkeys,
		BlockNumber:    c.keygen.BlockNumber,
	}
}

// GetCurrentTssPubkey returns the current tss pubkey
func (c *ZetacoreContext) GetCurrentTssPubkey() string {
	c.coreContextLock.RLock()
	defer c.coreContextLock.RUnlock()
	return c.currentTssPubkey
}

// GetEnabledChains returns all enabled chains including zetachain
func (c *ZetacoreContext) GetEnabledChains() []chains.Chain {
	c.coreContextLock.RLock()
	defer c.coreContextLock.RUnlock()

	copiedChains := make([]chains.Chain, len(c.chainsEnabled))
	copy(copiedChains, c.chainsEnabled)
	return copiedChains
}

// GetEnabledExternalChains returns all enabled external chains
func (c *ZetacoreContext) GetEnabledExternalChains() []chains.Chain {
	c.coreContextLock.RLock()
	defer c.coreContextLock.RUnlock()

	externalChains := make([]chains.Chain, 0)
	for _, chain := range c.chainsEnabled {
		if chain.IsExternal {
			externalChains = append(externalChains, chain)
		}
	}
	return externalChains
}

// GetEVMChainParams returns chain params for a specific EVM chain
func (c *ZetacoreContext) GetEVMChainParams(chainID int64) (*observertypes.ChainParams, bool) {
	c.coreContextLock.RLock()
	defer c.coreContextLock.RUnlock()

	evmChainParams, found := c.evmChainParams[chainID]
	return evmChainParams, found
}

// GetAllEVMChainParams returns all chain params for EVM chains
func (c *ZetacoreContext) GetAllEVMChainParams() map[int64]*observertypes.ChainParams {
	c.coreContextLock.RLock()
	defer c.coreContextLock.RUnlock()

	// deep copy evm chain params
	copied := make(map[int64]*observertypes.ChainParams, len(c.evmChainParams))
	for chainID, evmConfig := range c.evmChainParams {
		copied[chainID] = &observertypes.ChainParams{}
		*copied[chainID] = *evmConfig
	}
	return copied
}

// GetBTCChainParams returns (chain, chain params, found) for bitcoin chain
func (c *ZetacoreContext) GetBTCChainParams() (chains.Chain, *observertypes.ChainParams, bool) {
	c.coreContextLock.RLock()
	defer c.coreContextLock.RUnlock()

	if c.bitcoinChainParams == nil { // bitcoin is not enabled
		return chains.Chain{}, nil, false
	}

	chain := chains.GetChainFromChainID(c.bitcoinChainParams.ChainId)
	if chain == nil {
		return chains.Chain{}, nil, false
	}

	return *chain, c.bitcoinChainParams, true
}

// GetCrossChainFlags returns crosschain flags
func (c *ZetacoreContext) GetCrossChainFlags() observertypes.CrosschainFlags {
	c.coreContextLock.RLock()
	defer c.coreContextLock.RUnlock()
	return c.crosschainFlags
}

// GetAllHeaderEnabledChains returns all verification flags
func (c *ZetacoreContext) GetAllHeaderEnabledChains() []lightclienttypes.HeaderSupportedChain {
	c.coreContextLock.RLock()
	defer c.coreContextLock.RUnlock()
	return c.blockHeaderEnabledChains
}

// GetBlockHeaderEnabledChains checks if block header verification is enabled for a specific chain
func (c *ZetacoreContext) GetBlockHeaderEnabledChains(chainID int64) (lightclienttypes.HeaderSupportedChain, bool) {
	c.coreContextLock.RLock()
	defer c.coreContextLock.RUnlock()
	for _, flags := range c.blockHeaderEnabledChains {
		if flags.ChainId == chainID {
			return flags, true
		}
	}
	return lightclienttypes.HeaderSupportedChain{}, false
}

// Update updates zetacore context and params for all chains
// this must be the ONLY function that writes to zetacore context
func (c *ZetacoreContext) Update(
	keygen *observertypes.Keygen,
	newChains []chains.Chain,
	evmChainParams map[int64]*observertypes.ChainParams,
	btcChainParams *observertypes.ChainParams,
	tssPubKey string,
	crosschainFlags observertypes.CrosschainFlags,
	blockHeaderEnabledChains []lightclienttypes.HeaderSupportedChain,
	init bool,
	logger zerolog.Logger,
) {
	c.coreContextLock.Lock()
	defer c.coreContextLock.Unlock()

	// Ignore whatever order zetacore organizes chain list in state
	sort.SliceStable(newChains, func(i, j int) bool {
		return newChains[i].ChainId < newChains[j].ChainId
	})

	if len(newChains) == 0 {
		logger.Warn().Msg("UpdateChainParams: No chains enabled in ZeroCore")
	}

	// Add some warnings if chain list changes at runtime
	if !init {
		if len(c.chainsEnabled) != len(newChains) {
			logger.Warn().Msgf(
				"UpdateChainParams: ChainsEnabled changed at runtime!! current: %v, new: %v",
				c.chainsEnabled,
				newChains,
			)
		} else {
			for i, chain := range newChains {
				if chain != c.chainsEnabled[i] {
					logger.Warn().Msgf(
						"UpdateChainParams: ChainsEnabled changed at runtime!! current: %v, new: %v",
						c.chainsEnabled,
						newChains,
					)
				}
			}
		}
	}

	if keygen != nil {
		c.keygen = *keygen
	}

	c.chainsEnabled = newChains
	c.crosschainFlags = crosschainFlags
	c.blockHeaderEnabledChains = blockHeaderEnabledChains

	// update chain params for bitcoin if it has config in file
	if c.bitcoinChainParams != nil && btcChainParams != nil {
		c.bitcoinChainParams = btcChainParams
	}

	// update core params for evm chains we have configs in file
	for _, params := range evmChainParams {
		_, found := c.evmChainParams[params.ChainId]
		if !found {
			continue
		}
		c.evmChainParams[params.ChainId] = params
	}

	if tssPubKey != "" {
		c.currentTssPubkey = tssPubKey
	}
}

// IsOutboundObservationEnabled returns true if the chain is supported and outbound flag is enabled
func IsOutboundObservationEnabled(c *ZetacoreContext, chainParams observertypes.ChainParams) bool {
	flags := c.GetCrossChainFlags()
	return chainParams.IsSupported && flags.IsOutboundEnabled
}

// IsInboundObservationEnabled returns true if the chain is supported and inbound flag is enabled
func IsInboundObservationEnabled(c *ZetacoreContext, chainParams observertypes.ChainParams) bool {
	flags := c.GetCrossChainFlags()
	return chainParams.IsSupported && flags.IsInboundEnabled
}
