package spec

import (
	"github.com/p9c/opts/binary"
	"github.com/p9c/opts/duration"
	"github.com/p9c/opts/float"
	"github.com/p9c/opts/integer"
	"github.com/p9c/opts/list"
	"github.com/p9c/opts/meta"
	"github.com/p9c/opts/text"
	"github.com/p9c/parallelcoin/pkg/appdata"
	"github.com/p9c/parallelcoin/pkg/base58"
	"github.com/p9c/parallelcoin/pkg/chaincfg"
	"github.com/p9c/parallelcoin/pkg/constant"
	"github.com/p9c/parallelcoin/pkg/hdkeychain"
	"github.com/p9c/parallelcoin/pkg/opts"
	uberatomic "go.uber.org/atomic"
	"math/rand"
	"net"
	"path/filepath"
	"sync/atomic"
	"time"
)

// GetConfigs returns configuration options for ParallelCoin Pod
func GetConfigs() (c opts.Configs) {
	network := "mainnet"
	rand.Seed(time.Now().Unix())
	var datadir = &atomic.Value{}
	datadir.Store([]byte(appdata.Dir(constant.Name, false)))
	c = opts.Configs{
		"AddCheckpoints": list.New(meta.Data{
			Aliases: []string{"AC"},
			Group:   "debug",
			Label:   "Add Checkpoints",
			Description:
			"add custom checkpoints",
			Widget: "multi",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			[]string{},
		),
		"AddPeers": list.New(meta.Data{
			Aliases: []string{"AP"},
			Group:   "node",
			Label:   "Add Peers",
			Description:
			"manually adds addresses to try to connect to",
			Type:   "ipaddress",
			Widget: "multi",
			// Hook:        "addpeer",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			[]string{},
			// []string{"127.0.0.1:12345", "127.0.0.1:12345", "127.0.0.1:12345", "127.0.0.1:12344"},
		),
		"AddrIndex": binary.New(meta.Data{
			Aliases: []string{"AI"},
			Group:   "node",
			Label:   "Address Index",
			Description:
			"maintain a full address-based transaction index which makes the searchrawtransactions RPC available",
			Widget: "toggle",
			// Hook:        "dropaddrindex",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"AutoPorts": binary.New(meta.Data{
			Group: "debug",
			Label: "Automatic Ports",
			Description:
			"RPC and controller ports are randomized, use with controller for automatic peer discovery",
			Widget: "toggle",
			// Hook: "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"AutoListen": binary.New(meta.Data{
			Aliases: []string{"AL"},
			Group:   "node",
			Label:   "Automatic Listeners",
			Description:
			"automatically update inbound addresses dynamically according to discovered network interfaces",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			true,
		),
		"BanDuration": duration.New(meta.Data{
			Aliases: []string{"BD"},
			Group:   "debug",
			Label:   "Ban Opt",
			Description:
			"how long a ban of a misbehaving peer lasts",
			Widget: "duration",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			time.Hour*24,
		),
		"BanThreshold": integer.New(meta.Data{
			Aliases: []string{"BT"},
			Group:   "debug",
			Label:   "Ban Threshold",
			Description:
			"ban score that triggers a ban (default 100)",
			Widget: "integer",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.DefaultBanThreshold,
		),
		"BlockMaxSize": integer.New(meta.Data{
			Aliases: []string{"BMXS"},
			Group:   "mining",
			Label:   "Block Max Size",
			Description:
			"maximum block size in bytes to be used when creating a block",
			Widget: "integer",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.BlockMaxSizeMax,
		),
		"BlockMaxWeight": integer.New(meta.Data{
			Aliases: []string{"BMXW"},
			Group:   "mining",
			Label:   "Block Max Weight",
			Description:
			"maximum block weight to be used when creating a block",
			Widget: "integer",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.BlockMaxWeightMax,
		),
		"BlockMinSize": integer.New(meta.Data{
			Aliases: []string{"BMS"},
			Group:   "mining",
			Label:   "Block Min Size",
			Description:
			"minimum block size in bytes to be used when creating a block",
			Widget: "integer",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.BlockMaxSizeMin,
		),
		"BlockMinWeight": integer.New(meta.Data{
			Aliases: []string{"BMW"},
			Group:   "mining",
			Label:   "Block Min Weight",
			Description:
			"minimum block weight to be used when creating a block",
			Widget: "integer",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.BlockMaxWeightMin,
		),
		"BlockPrioritySize": integer.New(meta.Data{
			Aliases: []string{"BPS"},
			Group:   "mining",
			Label:   "Block Priority Size",
			Description:
			"size in bytes for high-priority/low-fee transactions when creating a block",
			Widget: "integer",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.DefaultBlockPrioritySize,
		),
		"BlocksOnly": binary.New(meta.Data{
			Aliases: []string{"BO"},
			Group:   "node",
			Label:   "Blocks Only",
			Description:
			"do not accept transactions from remote peers",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"CAFile": text.New(meta.Data{
			Aliases: []string{"CA"},
			Group:   "tls",
			Label:   "Certificate Authority File",
			Description:
			"certificate authority file for TLS certificate validation",
			Type:   "path",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			filepath.Join(string(datadir.Load().([]byte)), "ca.cert"),
		),
		"ConfigFile": text.New(meta.Data{
			Aliases: []string{"CF"},
			Label:   "Configuration File",
			Description:
			"location of configuration file, cannot actually be changed",
			Type:   "path",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			filepath.Join(string(datadir.Load().([]byte)), constant.PodConfigFilename),
		),
		"ConnectPeers": list.New(meta.Data{
			Aliases: []string{"CPS"},
			Group:   "node",
			Label:   "Connect Peers",
			Description:
			"connect ONLY to these addresses (disables inbound connections)",
			Type:   "address",
			Widget: "multi",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			[]string{},
		),
		"Controller": binary.New(meta.Data{
			Aliases: []string{"CN"},
			Group:   "node",
			Label:   "Enable Controller",
			Description:
			"delivers mining jobs over multicast",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"CPUProfile": text.New(meta.Data{
			Aliases: []string{"CPR"},
			Group:   "debug",
			Label:   "CPU Profile",
			Description:
			"write cpu profile to this file",
			Type:   "path",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			"",
		),
		"DarkTheme": binary.New(meta.Data{
			Aliases: []string{"DT"},
			Group:   "config",
			Label:   "Dark Theme",
			Description:
			"sets dark theme for GUI",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"DataDir": &text.Opt{
			Value: datadir,
			Data: meta.Data{
				Aliases: []string{"DD"},
				Label:   "Data Directory",
				Description:
				"root folder where application data is stored",
				Type:          "directory",
				Widget:        "string",
				Documentation: "<placeholder for detailed documentation>",
				OmitEmpty:     true,
			},
			Def: appdata.Dir(constant.Name, false),
		},
		"DbType": text.New(meta.Data{
			Aliases: []string{"DB"},
			Group:   "debug",
			Label:   "Database Type",
			Description:
			"type of database storage engine to use (only one right now, ffldb)",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.DefaultDbType,
		),
		"DisableBanning": binary.New(meta.Data{
			Aliases: []string{"NB"},
			Group:   "debug",
			Label:   "Disable Banning",
			Description:
			"disables banning of misbehaving peers",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"DisableCheckpoints": binary.New(meta.Data{
			Aliases: []string{"NCP"},
			Group:   "debug",
			Label:   "Disable Checkpoints",
			Description:
			"disables all checkpoints",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"DisableDNSSeed": binary.New(meta.Data{
			Aliases: []string{"NDS"},
			Group:   "node",
			Label:   "Disable DNS Seed",
			Description:
			"disable seeding of addresses to peers",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"DisableListen": binary.New(meta.Data{
			Aliases: []string{"NL"},
			Group:   "node",
			Label:   "Disable Listen",
			Description:
			"disables inbound connections for the peer to peer network",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"DisableRPC": binary.New(meta.Data{
			Aliases: []string{"NRPC"},
			Group:   "rpc",
			Label:   "Disable RPC",
			Description:
			"disable rpc servers, as well as kopach controller",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"Discovery": binary.New(meta.Data{
			Aliases: []string{"DI"},
			Group:   "node",
			Label:   "Disovery",
			Description:
			"enable LAN peer discovery in GUI",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"ExternalIPs": list.New(meta.Data{
			Aliases: []string{"EI"},
			Group:   "node",
			Label:   "External IP Addresses",
			Description:
			"extra addresses to tell peers they can connect to",
			Type:   "address",
			Widget: "multi",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			[]string{},
		),
		"FreeTxRelayLimit": float.NewFloat(meta.Data{
			Aliases: []string{"LR"},
			Group:   "policy",
			Label:   "Free Tx Relay Limit",
			Description:
			"limit relay of transactions with no transaction fee to the given amount in thousands of bytes per minute",
			Widget: "float",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.DefaultFreeTxRelayLimit,
		),
		"Generate": binary.New(meta.Data{
			Aliases: []string{"GB"},
			Group:   "mining",
			Label:   "Generate Blocks",
			Description:
			"turn on Kopach CPU miner",
			Widget: "toggle",
			// Hook:        "generate",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"GenThreads": integer.New(meta.Data{
			Aliases: []string{"GT"},
			Group:   "mining",
			Label:   "Generate Threads",
			Description:
			"number of threads to mine with",
			Widget: "integer",
			// Hook:        "genthreads",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			-1,
		),
		"Hilite": list.New(meta.Data{
			Aliases: []string{"HL"},
			Group:   "debug",
			Label:   "Hilite",
			Description:
			"list of packages that will print with attention getters",
			Type:   "string",
			Widget: "multi",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			[]string{},
		),
		"LAN": binary.New(meta.Data{
			Group: "debug",
			Label: "LAN Testnet Mode",
			Description:
			"run without any connection to nodes on the internet (does not apply on mainnet)",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"Locale": text.New(meta.Data{
			Aliases: []string{"LC"},
			Group:   "config",
			Label:   "Language",
			Description:
			"user interface language i18 localization",
			Widget: "string",
			// Hook:        "language",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			"en",
		),
		"LimitPass": text.New(meta.Data{
			Aliases: []string{"LP"},
			Group:   "rpc",
			Label:   "Limit Password",
			Description:
			"limited user password",
			Widget: "password",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			genPassword(),
		),
		"LimitUser": text.New(meta.Data{
			Aliases: []string{"LU"},
			Group:   "rpc",
			Label:   "Limit Username",
			Description:
			"limited user name",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			"limit",
		),
		"LogDir": text.New(meta.Data{
			Aliases: []string{"LD"},
			Group:   "config",
			Label:   "Log Directory",
			Description:
			"folder where log files are written",
			Type:   "directory",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			string(datadir.Load().([]byte)),
		),
		"LogFilter": list.New(meta.Data{
			Aliases: []string{"LF"},
			Group:   "debug",
			Label:   "Log Filter",
			Description:
			"list of packages that will not print logs",
			Type:   "string",
			Widget: "multi",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			[]string{},
		),
		"LogLevel": text.New(meta.Data{
			Aliases: []string{"LL"},
			Group:   "config",
			Label:   "Log Level",
			Description:
			"maximum log level to output",
			Widget: "radio",
			Options: []string{"off",
				"fatal",
				"error",
				"info",
				"check",
				"debug",
				"trace",
			},
			// Hook:        "loglevel",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			"info",
		),
		"MaxOrphanTxs": integer.New(meta.Data{
			Aliases: []string{"MO"},
			Group:   "policy",
			Label:   "Max Orphan Txs",
			Description:
			"max number of orphan transactions to keep in memory",
			Widget: "integer",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.DefaultMaxOrphanTransactions,
		),
		"MaxPeers": integer.New(meta.Data{
			Aliases: []string{"MP"},
			Group:   "node",
			Label:   "Max Peers",
			Description:
			"maximum number of peers to hold connections with",
			Widget: "integer",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.DefaultMaxPeers,
		),
		"MulticastPass": text.New(meta.Data{
			Aliases: []string{"PM"},
			Group:   "config",
			Label:   "Multicast Pass",
			Description:
			"password that encrypts the connection to the mining controller",
			Widget: "password",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			"pa55word",
		),
		"MinRelayTxFee": float.NewFloat(meta.Data{
			Aliases: []string{"MRTF"},
			Group:   "policy",
			Label:   "Min Relay Transaction Fee",
			Description:
			"the minimum transaction fee in DUO/kB to be considered a non-zero fee",
			Widget: "float",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.DefaultMinRelayTxFee.ToDUO(),
		),
		"Network": text.New(meta.Data{
			Aliases: []string{"NW"},
			Group:   "node",
			Label:   "Network",
			Description:
			"connect to this network: (mainnet, testnet)",
			Widget: "radio",
			Options: []string{"mainnet",
				"testnet",
				"regtestnet",
				"simnet",
			},
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			network,
		),
		"NoCFilters": binary.New(meta.Data{
			Aliases: []string{"NCF"},
			Group:   "node",
			Label:   "No CFilters",
			Description:
			"disable committed filtering (CF) support",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"NodeOff": binary.New(meta.Data{
			Aliases: []string{"NO"},
			Group:   "debug",
			Label:   "Node Off",
			Description:
			"turn off the node backend",
			Widget: "toggle",
			// Hook:        "node",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"NoInitialLoad": binary.New(meta.Data{
			Aliases: []string{"NIL"},
			Label:   "No Initial Load",
			Description:
			"do not load a wallet at startup",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"NoPeerBloomFilters": binary.New(meta.Data{
			Aliases: []string{"NPBF"},
			Group:   "node",
			Label:   "No Peer Bloom Filters",
			Description:
			"disable bloom filtering support",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"NoRelayPriority": binary.New(meta.Data{
			Aliases: []string{"NRPR"},
			Group:   "policy",
			Label:   "No Relay Priority",
			Description:
			"do not require free or low-fee transactions to have high priority for relaying",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"OneTimeTLSKey": binary.New(meta.Data{
			Aliases: []string{"OTK"},
			Group:   "wallet",
			Label:   "One Time TLS Key",
			Description:
			"generate a new TLS certificate pair at startup, but only write the certificate to disk",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"OnionEnabled": binary.New(meta.Data{
			Aliases: []string{"OE"},
			Group:   "proxy",
			Label:   "Onion Enabled",
			Description:
			"enable tor proxy",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"OnionProxyAddress": text.New(meta.Data{
			Aliases: []string{"OPA"},
			Group:   "proxy",
			Label:   "Onion Proxy Address",
			Description:
			"address of tor proxy you want to connect to",
			Type:   "address",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			"",
		),
		"OnionProxyPass": text.New(meta.Data{
			Aliases: []string{"OPW"},
			Group:   "proxy",
			Label:   "Onion Proxy Password",
			Description:
			"password for tor proxy",
			Widget: "password",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			"",
		),
		"OnionProxyUser": text.New(meta.Data{
			Aliases: []string{"OU"},
			Group:   "proxy",
			Label:   "Onion Proxy Username",
			Description:
			"tor proxy username",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			"",
		),
		"P2PConnect": list.New(meta.Data{
			Aliases: []string{"P2P"},
			Group:   "node",
			Label:   "P2P Connect",
			Description:
			"list of addresses reachable from connected networks",
			Type:   "address",
			Widget: "multi",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			[]string{},
		),
		"P2PListeners": list.New(meta.Data{
			Aliases: []string{"LA"},
			Group:   "node",
			Label:   "P2PListeners",
			Description:
			"list of addresses to bind the node listener to",
			Type:   "address",
			Widget: "multi",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			[]string{net.JoinHostPort("0.0.0.0",
				chaincfg.MainNetParams.DefaultPort,
			),
			},
		),
		"Password": text.New(meta.Data{
			Aliases: []string{"PW"},
			Group:   "rpc",
			Label:   "Password",
			Description:
			"password for client RPC connections",
			Type:   "password",
			Widget: "password",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			genPassword(),
		),
		"PipeLog": binary.New(meta.Data{
			Aliases: []string{"PL"},
			Label:   "Pipe Logger",
			Description:
			"enable pipe based logger IPC",
			Widget: "toggle",
			// Hook:        "",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"Profile": text.New(meta.Data{
			Aliases: []string{"HPR"},
			Group:   "debug",
			Label:   "Profile",
			Description:
			"http profiling on given port (1024-40000)",
			// Type:        "",
			Widget: "integer",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			"",
		),
		"ProxyAddress": text.New(meta.Data{
			Aliases: []string{"PA"},
			Group:   "proxy",
			Label:   "Proxy",
			Description:
			"address of proxy to connect to for outbound connections",
			Type:   "url",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			"",
		),
		"ProxyPass": text.New(meta.Data{
			Aliases: []string{"PPW"},
			Group:   "proxy",
			Label:   "Proxy Pass",
			Description:
			"proxy password, if required",
			Type:   "password",
			Widget: "password",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			genPassword(),
		),
		"ProxyUser": text.New(meta.Data{
			Aliases: []string{"PU"},
			Group:   "proxy",
			Label:   "ProxyUser",
			Description:
			"proxy username, if required",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			"proxyuser",
		),
		"RejectNonStd": binary.New(meta.Data{
			Aliases: []string{"REJ"},
			Group:   "node",
			Label:   "Reject Non Std",
			Description:
			"reject non-standard transactions regardless of the default settings for the active network",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"RelayNonStd": binary.New(meta.Data{
			Aliases: []string{"RNS"},
			Group:   "node",
			Label:   "Relay Nonstandard Transactions",
			Description:
			"relay non-standard transactions regardless of the default settings for the active network",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"RPCCert": text.New(meta.Data{
			Aliases: []string{"RC"},
			Group:   "rpc",
			Label:   "RPC Cert",
			Description:
			"location of RPC TLS certificate",
			Type:   "path",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			filepath.Join(string(datadir.Load().([]byte)), "rpc.cert"),
		),
		"RPCConnect": text.New(meta.Data{
			Aliases: []string{"RA"},
			Group:   "wallet",
			Label:   "RPC Connect",
			Description:
			"full node RPC for wallet",
			Type:   "address",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			net.JoinHostPort("127.0.0.1", chaincfg.MainNetParams.RPCClientPort),
		),
		"RPCKey": text.New(meta.Data{
			Aliases: []string{"RK"},
			Group:   "rpc",
			Label:   "RPC Key",
			Description:
			"location of rpc TLS key",
			Type:   "path",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			filepath.Join(string(datadir.Load().([]byte)), "rpc.key"),
		),
		"RPCListeners": list.New(meta.Data{
			Aliases: []string{"RL"},
			Group:   "rpc",
			Label:   "RPC Listeners",
			Description:
			"addresses to listen for RPC connections",
			Type:   "address",
			Widget: "multi",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			[]string{net.JoinHostPort("127.0.0.1", chaincfg.MainNetParams.RPCClientPort),
			},
		),
		"RPCMaxClients": integer.New(meta.Data{
			Aliases: []string{"RMXC"},
			Group:   "rpc",
			Label:   "Maximum RPC Clients",
			Description:
			"maximum number of clients for regular RPC",
			Widget: "integer",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.DefaultMaxRPCClients,
		),
		"RPCMaxConcurrentReqs": integer.New(meta.Data{
			Aliases: []string{"RMCR"},
			Group:   "rpc",
			Label:   "Maximum RPC Concurrent Reqs",
			Description:
			"maximum number of requests to process concurrently",
			Widget: "integer",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.DefaultMaxRPCConcurrentReqs,
		),
		"RPCMaxWebsockets": integer.New(meta.Data{
			Aliases: []string{"RMWS"},
			Group:   "rpc",
			Label:   "Maximum RPC Websockets",
			Description:
			"maximum number of websocket clients to allow",
			Widget: "integer",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.DefaultMaxRPCWebsockets,
		),
		"RPCQuirks": binary.New(meta.Data{
			Aliases: []string{"RQ"},
			Group:   "rpc",
			Label:   "RPC Quirks",
			Description:
			"enable bugs that replicate bitcoin core RPC's JSON",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"RunAsService": binary.New(meta.Data{
			Aliases: []string{"RS"},
			Label:   "Run As Service",
			Description:
			"shuts down on lock timeout",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"Save": binary.New(meta.Data{
			Aliases: []string{"SV"},
			Label:   "Save Configuration",
			Description:
			"save opts given on commandline",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"ServerPass": text.New(meta.Data{
			Aliases: []string{"SPW"},
			Group:   "rpc",
			Label:   "Server Pass",
			Description:
			"password for server connections",
			Type:   "password",
			Widget: "password",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			genPassword(),
		),
		"ServerTLS": binary.New(meta.Data{
			Aliases: []string{"ST"},
			Group:   "wallet",
			Label:   "Server TLS",
			Description:
			"enable TLS for the wallet connection to node RPC server",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			true,
		),
		"ServerUser": text.New(meta.Data{
			Aliases: []string{"SU"},
			Group:   "rpc",
			Label:   "Server User",
			Description:
			"username for chain server connections",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			"client",
		),
		"SigCacheMaxSize": integer.New(meta.Data{
			Aliases: []string{"SCM"},
			Group:   "node",
			Label:   "Signature Cache Max Size",
			Description:
			"the maximum number of entries in the signature verification cache",
			Widget: "integer",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.DefaultSigCacheMaxSize,
		),
		"Solo": binary.New(meta.Data{
			Group: "mining",
			Label: "Solo Generate",
			Description:
			"mine even if not connected to a network",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"ClientTLS": binary.New(meta.Data{
			Aliases: []string{"CT"},
			Group:   "tls",
			Label:   "TLS",
			Description:
			"enable TLS for RPC client connections",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			true,
		),
		"TLSSkipVerify": binary.New(meta.Data{
			Aliases: []string{"TSV"},
			Group:   "tls",
			Label:   "TLS Skip Verify",
			Description:
			"skip TLS certificate verification (ignore CA errors)",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"TorIsolation": binary.New(meta.Data{
			Aliases: []string{"TI"},
			Group:   "proxy",
			Label:   "Tor Isolation",
			Description:
			"makes a separate proxy connection for each connection",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"TrickleInterval": duration.New(meta.Data{
			Aliases: []string{"TKI"},
			Group:   "policy",
			Label:   "Trickle Interval",
			Description:
			"minimum time between attempts to send new inventory to a connected peer",
			Widget: "duration",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.DefaultTrickleInterval,
		),
		"TxIndex": binary.New(meta.Data{
			Aliases: []string{"TXI"},
			Group:   "node",
			Label:   "Tx Index",
			Description:
			"maintain a full hash-based transaction index which makes all transactions available via the getrawtransaction RPC",
			Widget: "toggle",
			// Hook:        "droptxindex",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"UPNP": binary.New(meta.Data{
			Aliases: []string{"UP"},
			Group:   "node",
			Label:   "UPNP",
			Description:
			"enable UPNP for NAT traversal",
			Widget: "toggle",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"UserAgentComments": list.New(meta.Data{
			Aliases: []string{"UA"},
			Group:   "policy",
			Label:   "User Agent Comments",
			Description:
			"comment to add to the user agent -- See BIP 14 for more information",
			Widget: "multi",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			[]string{},
		),
		"Username": text.New(meta.Data{
			Aliases: []string{"UN"},
			Group:   "rpc",
			Label:   "Username",
			Description:
			"password for client RPC connections",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			"username",
		),
		"UUID": &integer.Opt{Data: meta.Data{
			Label: "UUID",
			Description:
			"instance unique id (64bit random value)",
			Widget:        "string",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			Value: uberatomic.NewInt64(rand.Int63()),
		},
		"UseWallet": binary.New(meta.Data{
			Aliases: []string{"WC"},
			Group:   "debug",
			Label:   "Connect to Wallet",
			Description:
			"set ctl to connect to wallet instead of chain server",
			Widget:        "toggle",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"WalletFile": text.New(meta.Data{
			Aliases: []string{"WF"},
			Group:   "config",
			Label:   "Wallet File",
			Description:
			"wallet database file",
			Type:   "path",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			filepath.Join(string(datadir.Load().([]byte)), "mainnet", constant.DbName),
		),
		"WalletOff": binary.New(meta.Data{
			Aliases: []string{"WO"},
			Group:   "debug",
			Label:   "Wallet Off",
			Description:
			"turn off the wallet backend",
			Widget: "toggle",
			// Hook:        "wallet",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			false,
		),
		"WalletPass": text.New(meta.Data{
			Aliases: []string{"WPW"},
			Label:   "Wallet Pass",
			Description:
			"password encrypting public data in wallet - hash is stored so give on command line",
			Type:   "password",
			Widget: "password",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			"",
		),
		"WalletRPCListeners": list.New(meta.Data{
			Aliases: []string{"WRL"},
			Group:   "wallet",
			Label:   "Wallet RPC Listeners",
			Description:
			"addresses for wallet RPC server to listen on",
			Type:   "address",
			Widget: "multi",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			[]string{net.JoinHostPort("0.0.0.0",
				chaincfg.MainNetParams.WalletRPCServerPort,
			),
			},
		),
		"WalletRPCMaxClients": integer.New(meta.Data{
			Aliases: []string{"WRMC"},
			Group:   "wallet",
			Label:   "Legacy RPC Max Clients",
			Description:
			"maximum number of RPC clients allowed for wallet RPC",
			Widget: "integer",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.DefaultRPCMaxClients,
		),
		"WalletRPCMaxWebsockets": integer.New(meta.Data{
			Aliases: []string{"WRMWS"},
			Group:   "wallet",
			Label:   "Legacy RPC Max Websockets",
			Description:
			"maximum number of websocket clients allowed for wallet RPC",
			Widget: "integer",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			constant.DefaultRPCMaxWebsockets,
		),
		"WalletServer": text.New(meta.Data{
			Aliases: []string{"WS"},
			Group:   "wallet",
			Label:   "Wallet Server",
			Description:
			"node address to connect wallet server to",
			Type:   "address",
			Widget: "string",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			net.JoinHostPort("127.0.0.1",
				chaincfg.MainNetParams.WalletRPCServerPort,
			),
		),
		"Whitelists": list.New(meta.Data{
			Aliases: []string{"WL"},
			Group:   "debug",
			Label:   "Whitelists",
			Description:
			"peers that you don't want to ever ban",
			Type:   "address",
			Widget: "multi",
			// Hook:        "restart",
			Documentation: "<placeholder for detailed documentation>",
			OmitEmpty:     true,
		},
			[]string{},
		),
	}
	for i := range c {
		c[i].SetName(i)
	}
	return
}

func genPassword() string {
	s, e := hdkeychain.GenerateSeed(16)
	if e != nil {
		panic("can't do nothing without entropy! " + e.Error())
	}
	return base58.Encode(s)
}
