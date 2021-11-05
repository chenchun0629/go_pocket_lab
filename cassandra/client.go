package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

func main() {
	var cluster = gocql.NewCluster("localhost:9042")
	cluster.Authenticator = &gocql.PasswordAuthenticator{
		Username:              "cassandra",
		Password:              "cassandra",
		AllowedAuthenticators: nil,
	}

	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	m := make(map[string]interface{})
	if _, err = session.Query("describe keyspaces").MapScanCAS(m); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v \n", m)
}