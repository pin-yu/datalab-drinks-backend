package main

/*
	Mantainer: Pin-Yu, Wang
	Email: pywang@datalab.cs.nthu.edu.tw
*/

import (
	"flag"

	"github.com/pin-yu/datalab-drinks-backend/src/application/services"
	"github.com/pin-yu/datalab-drinks-backend/src/interface/routers"
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
		services.MigrateTable()
		routers.Run()
	}
}
