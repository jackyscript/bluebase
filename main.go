package main

import (
	"bluebase/security"
	"log"

	"github.com/dop251/goja"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
)

func main() {

	app := pocketbase.New()

	jsvm.MustRegister(app, jsvm.Config{
		OnInit: func(vm *goja.Runtime) {
			// register sanitize binding
			vm.Set("sanitize", security.Sanitize)
		},
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
