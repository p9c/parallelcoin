module github.com/p9c/parallelcoin

go 1.16

require (
	github.com/VividCortex/ewma v1.1.1
	github.com/bitbandi/go-x11 v0.0.0-20171024232457-5fddbc9b2b09
	github.com/btcsuite/go-socks v0.0.0-20170105172521-4720035b7bfd
	github.com/btcsuite/golangcrypto v0.0.0-20150304025918-53f62d9b43e8
	github.com/btcsuite/goleveldb v1.0.0
	github.com/coreos/bbolt v1.3.3
	github.com/davecgh/go-spew v1.1.1
	github.com/enceve/crypto v0.0.0-20160707101852-34d48bb93815
	github.com/jackpal/gateway v1.0.7
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
	github.com/niubaoshu/gotiny v0.0.3
	github.com/p9c/log v0.0.6
	github.com/p9c/opts v0.0.6
	github.com/p9c/pod v1.9.25 // indirect
	github.com/p9c/qu v0.0.3
	github.com/programmer10110/gostreebog v0.0.0-20170704145444-a3e1d28291b2
	go.uber.org/atomic v1.7.0
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	gopkg.in/src-d/go-git.v4 v4.13.1
	lukechampine.com/blake3 v1.0.0
)

//replace gioui.org => ./gio
