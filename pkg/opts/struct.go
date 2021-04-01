package opts

import (
	"github.com/p9c/opts/binary"
	"github.com/p9c/opts/duration"
	"github.com/p9c/opts/float"
	"github.com/p9c/opts/integer"
	"github.com/p9c/opts/list"
	"github.com/p9c/opts/text"
)

// Config defines the configuration items used by pod along with the various components included in the suite
//go:generate go run genopts/main.go
type Config struct {
	AddCheckpoints         *list.Opt
	AddPeers               *list.Opt
	AddrIndex              *binary.Opt
	AutoListen             *binary.Opt
	AutoPorts              *binary.Opt
	BanDuration            *duration.Opt
	BanThreshold           *integer.Opt
	BlockMaxSize           *integer.Opt
	BlockMaxWeight         *integer.Opt
	BlockMinSize           *integer.Opt
	BlockMinWeight         *integer.Opt
	BlockPrioritySize      *integer.Opt
	BlocksOnly             *binary.Opt
	CAFile                 *text.Opt
	CPUProfile             *text.Opt
	ClientTLS              *binary.Opt
	ConfigFile             *text.Opt
	ConnectPeers           *list.Opt
	Controller             *binary.Opt
	DarkTheme              *binary.Opt
	DataDir                *text.Opt
	DbType                 *text.Opt
	DisableBanning         *binary.Opt
	DisableCheckpoints     *binary.Opt
	DisableDNSSeed         *binary.Opt
	DisableListen          *binary.Opt
	DisableRPC             *binary.Opt
	Discovery              *binary.Opt
	ExternalIPs            *list.Opt
	FreeTxRelayLimit       *float.Opt
	GenThreads             *integer.Opt
	Generate               *binary.Opt
	Hilite                 *list.Opt
	LAN                    *binary.Opt
	LimitPass              *text.Opt
	LimitUser              *text.Opt
	Locale                 *text.Opt
	LogDir                 *text.Opt
	LogFilter              *list.Opt
	LogLevel               *text.Opt
	MaxOrphanTxs           *integer.Opt
	MaxPeers               *integer.Opt
	MinRelayTxFee          *float.Opt
	MulticastPass          *text.Opt
	Network                *text.Opt
	NoCFilters             *binary.Opt
	NoInitialLoad          *binary.Opt
	NoPeerBloomFilters     *binary.Opt
	NoRelayPriority        *binary.Opt
	NodeOff                *binary.Opt
	OneTimeTLSKey          *binary.Opt
	OnionEnabled           *binary.Opt
	OnionProxyAddress      *text.Opt
	OnionProxyPass         *text.Opt
	OnionProxyUser         *text.Opt
	P2PConnect             *list.Opt
	P2PListeners           *list.Opt
	Password               *text.Opt
	PipeLog                *binary.Opt
	Profile                *text.Opt
	ProxyAddress           *text.Opt
	ProxyPass              *text.Opt
	ProxyUser              *text.Opt
	RPCCert                *text.Opt
	RPCConnect             *text.Opt
	RPCKey                 *text.Opt
	RPCListeners           *list.Opt
	RPCMaxClients          *integer.Opt
	RPCMaxConcurrentReqs   *integer.Opt
	RPCMaxWebsockets       *integer.Opt
	RPCQuirks              *binary.Opt
	RejectNonStd           *binary.Opt
	RelayNonStd            *binary.Opt
	RunAsService           *binary.Opt
	Save                   *binary.Opt
	ServerPass             *text.Opt
	ServerTLS              *binary.Opt
	ServerUser             *text.Opt
	SigCacheMaxSize        *integer.Opt
	Solo                   *binary.Opt
	TLSSkipVerify          *binary.Opt
	TorIsolation           *binary.Opt
	TrickleInterval        *duration.Opt
	TxIndex                *binary.Opt
	UPNP                   *binary.Opt
	UUID                   *integer.Opt
	UseWallet              *binary.Opt
	UserAgentComments      *list.Opt
	Username               *text.Opt
	WalletFile             *text.Opt
	WalletOff              *binary.Opt
	WalletPass             *text.Opt
	WalletRPCListeners     *list.Opt
	WalletRPCMaxClients    *integer.Opt
	WalletRPCMaxWebsockets *integer.Opt
	WalletServer           *text.Opt
	Whitelists             *list.Opt
}
