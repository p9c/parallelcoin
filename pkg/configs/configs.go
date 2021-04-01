package configs

import (
	"github.com/p9c/opts/opt"
)

// Configs is the source location for the Config items, which is used to generate the Config struct
type Configs map[string]opt.Option
type ConfigSliceElement struct {
	Opt  opt.Option
	Name string
}
type ConfigSlice []ConfigSliceElement

func (c ConfigSlice) Len() int           { return len(c) }
func (c ConfigSlice) Less(i, j int) bool { return c[i].Name < c[j].Name }
func (c ConfigSlice) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
