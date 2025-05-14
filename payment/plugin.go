package main

import (
	"github.com/haobaojiang/diexample/order/pkg"
	"github.com/samber/do/v2"
)

// ------------------------------

// InitPlugin : the entry
func InitPlugin(injector do.Injector) {
	pkg.InitModule(injector)
}

// main : this is required for building order
func main() {}
