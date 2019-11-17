package main

import (
	"Kronos/boot"
	"Kronos/routes"
)

func main() {
	routers := routes.InitRouters()
	boot.Run(routers)
}
