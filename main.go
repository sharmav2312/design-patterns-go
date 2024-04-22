package main

import (
	"design-patterns/patterns"
)

func main() {
	sqlCon := patterns.SQLConnection{ConnectionString: "sql connection string"}
	con := patterns.Connection{DB: sqlCon}
	con.DBConnect()

	oracleCon := patterns.OracleConnection{ConnectionString: "oracle connection string"}
	con2 := patterns.Connection{DB: oracleCon}
	con2.DBConnect()
}
