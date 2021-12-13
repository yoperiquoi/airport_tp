package main

import "airport_tp/cmd/database"

func main() {
	conn := database.CreateConnexion()

	database.CloseConnection(conn)
}
