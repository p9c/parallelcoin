// This generator reads a podcfg.Configs map and spits out a podcfg.Config struct
package main

import (
	"fmt"
	"github.com/p9c/parallelcoin/pkg/opts"
	"github.com/p9c/parallelcoin/pkg/spec"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
)

func main() {
	c := spec.GetConfigs()
	var o string
	var cc opts.ConfigSlice
	for i := range c {
		cc = append(cc, opts.ConfigSliceElement{Opt: c[i], Name: i})
	}
	sort.Sort(cc)
	for i := range cc {
		t := reflect.TypeOf(cc[i].Opt).String()
		// W.Ln(t)
		// split := strings.Split(t, "podcfg.")[1]
		o += fmt.Sprintf("\t%s\t%s\n", cc[i].Name, t)
	}
	var e error
	var out []byte
	var wd string
	generated := fmt.Sprintf(configBase, o)
	if out, e = format.Source([]byte(generated)); e != nil {
		// panic(e)
		fmt.Println(e)
	}
	if wd, e = os.Getwd(); e != nil {
		// panic(e)
	}
	// fmt.Println(string(out), wd)
	if e = ioutil.WriteFile(filepath.Join(wd, "struct.go"), out, 0660); e != nil {
		panic(e)
	}
}

var configBase = `package opts

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
%s}
`
