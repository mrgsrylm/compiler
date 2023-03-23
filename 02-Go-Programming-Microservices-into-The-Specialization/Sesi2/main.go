package main

import (
	"github.com/gusrylmubarok/test/tree/main/02-Go-Programming-Microservices-into-The-Specialization/Sesi2/routes"
)

func main() {
	var PORT = ":4000"

	routes.StartServer().Run(PORT)
}
