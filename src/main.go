package main

/*
	Mantainer: Pin-Yu, Wang
	Email: pywang@datalab.cs.nthu.edu.tw
*/

import (
	"flag"

	"github.com/pinyu/datalab-drinks-backend/src/application/services"
	"github.com/pinyu/datalab-drinks-backend/src/interface/routes"
)

var migrateFlag bool
var dropFlag bool

func parseFlag() {
	flag.BoolVar(&migrateFlag, "m", false, "migrate database")
	flag.BoolVar(&dropFlag, "d", false, "drop database")
	flag.Parse()
}

func main() {
	parseFlag()

	if migrateFlag {
		services.MigrateTable()
	} else if dropFlag {
		services.DropTable()
	} else {
		// starts the server
		routes.Run()
	}
}
