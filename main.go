package main

import (
	"Kronos/bootstrap"
	"Kronos/routes"
)

func main() {
	routers := routes.InitRouters()
	bootstrap.Run(routers)
}
