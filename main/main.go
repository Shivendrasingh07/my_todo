package main

import (
	"github.com/my_todo/Database"
	"github.com/my_todo/Routes"
)

func main() {

	Database.Connect()
	Routes.Route()

	// Mount the admin sub-router
	//	r.Mount("/admin", Handlers.adminRouter())

}
