package pod

import (
	"github.com/p9c/parallelcoin/version"
)

func Init() int {
	I.Ln(version.Get())
	
	
	
	return 0
}
