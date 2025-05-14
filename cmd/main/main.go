package main

import (
	"fmt"
	"github.com/haobaojiang/diexample/common"
	orderPkg "github.com/haobaojiang/diexample/order/pkg"
	paymentPkg "github.com/haobaojiang/diexample/payment/pkg"
	"github.com/samber/do/v2"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func loadPlugins(inj do.Injector) {

	dir := `plugin-dist`

	entries, err := os.ReadDir(dir)
	checkErr(err)
	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".so") {
			continue
		}
		p, err := plugin.Open(filepath.Join(dir, entry.Name()))
		checkErr(err)
		s, err := p.Lookup("InitPlugin")
		checkErr(err)
		s.(func(do.Injector))(inj)
	}
}

func main() {

	// create a new injector
	inj := do.New()

	// load the plugins first
	loadPlugins(inj)

	// load the modules, if the plugin has already loaded the module with same functionalities, it will skip
	orderPkg.InitModule(inj)
	paymentPkg.InitModule(inj)

	// invoke the objects
	orders := do.MustInvoke[common.Orders](inj)

	// do the work
	order, _ := orders.GetById()
	fmt.Println("order id:", order.Id())
	checkErr(order.Pay())
}
