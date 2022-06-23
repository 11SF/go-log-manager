package main

import "11sf/go_log_manager/router"

func main() {
	e := router.New()
	e.Logger.Fatal(e.Start(":8080"))
}
