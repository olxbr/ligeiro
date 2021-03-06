package envcfg_test

import (
	"fmt"

	"github.com/olxbr/ligeiro/envcfg"
)

func ExampleLoad() {
	Config := envcfg.Load(envcfg.Map{
		"APPLICATION": "myappname",
		"LISTEN_PORT": ":8080",
		"MAX_BYTES":   "300000",
	})

	fmt.Println(Config.Get("listenPort"))
	fmt.Println(Config.Get("logLevel"))
	fmt.Println(Config.Get("maxBytes"), fmt.Sprintf("%T", Config.GetInt("maxBytes")))
	// Output:
	// :8080
	// debug
	// 300000 int
}

func ExampleLoadBundled() {
	Config := envcfg.LoadBundled()

	fmt.Println(Config.Get("version"))
	fmt.Println(Config.Get("logLevel"))
	// Output:
	// detached
	// debug
}
