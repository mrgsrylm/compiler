package main

import (
	"github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi2/database"
	"github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi2/routes"
)

func main() {
	database.StartDB()

	r := routes.StartApp()
	r.Run(":4000")
}
