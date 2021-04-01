package main

var commands = map[string][]string{
	"build": {
		"go build -v",
	},
	"install": {
		"go install -v",
	},
	"headless": {
		"go install -v -tags headless",
	},
	"builder": {
		"go install -v ./cmd/p9build/.",
	},
}
