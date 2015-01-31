package main

import (
	"github.com/meatballhat/packer-builder-onlinelabs/builder/onlinelabs"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(onlinelabs.NewBuilder())
	server.Serve()
}
