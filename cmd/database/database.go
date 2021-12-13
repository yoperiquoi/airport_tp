package database

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func CreateConnexion() redis.Conn {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")
	return conn
}

func CloseConnection(conn redis.Conn) {
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("disconnected")
	}(conn)
}
