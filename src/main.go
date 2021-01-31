package main

/*
	Mantainer: Pin-Yu, Wang
	Email: pywang@datalab.cs.nthu.edu.tw
*/

import (
	"flag"

	"github.com/pinyu/datalab-drinks-backend/src/api/routes"
	"github.com/pinyu/datalab-drinks-backend/src/infrastructure/orm"
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
		orm.MigrateDB()
	} else if dropFlag {
		orm.DropDB()
	} else {
		// starts the server
		routes.Run()
	}
}
