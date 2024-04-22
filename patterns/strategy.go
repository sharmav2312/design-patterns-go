package patterns

import "fmt"

type DBConnection interface {
	Connect()
}

type SQLConnection struct {
	ConnectionString string
}

func (con SQLConnection) Connect() {
	fmt.Println("SQL Connection:" + con.ConnectionString)
}

type OracleConnection struct {
	ConnectionString string
}

func (con OracleConnection) Connect() {
	fmt.Println("Oracle Connection:" + con.ConnectionString)
}

type Connection struct {
	DB DBConnection
}

func (con Connection) DBConnect() {
	con.DB.Connect()
}
