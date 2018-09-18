package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/prometheus/common/log"
	"github.com/valyala/fasthttp"
	_ "github.com/chenhg5/go-admin/adapter/fasthttp"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
)

func main() {
	router := fasthttprouter.New()

	eng := engine.Default()

	cfg := config.Config{
		DATABASE: []config.Database{
			{
				HOST:         "127.0.0.1",
				PORT:         "3306",
				USER:         "root",
				PWD:          "root",
				NAME:         "godmin",
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       "mysql",
			},
		},
		DOMAIN: "localhost",
		PREFIX: "admin",
		INDEX:  "/",
	}

	adminPlugin := admin.NewAdmin(datamodel.Generators)
	examplePlugin := example.NewExample()

	log.Fatal(eng.AddConfig(cfg).AddPlugins(adminPlugin, examplePlugin).Use(router))

	var waitChan chan int
	go func() {
		fasthttp.ListenAndServe(":8897", router.Handler)
	}()
	<-waitChan
}
