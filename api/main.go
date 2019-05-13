package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/JulzDiverse/eirinidotcf/api/api"
	feedb "github.com/JulzDiverse/eirinidotcf/api/db"
)

const defaultPort = "8080"

type DBCredentials struct {
	Address  string `json:"db_address"`
	Name     string `json:"db_name"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserProvided struct {
	Credentials DBCredentials `json:"credentials"`
}

type VCAPServices struct {
	UserProvided []UserProvided `json:"user-provided"`
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = defaultPort
	}

	services := os.Getenv("VCAP_SERVICES")
	if len(services) == 0 {
		panic(errors.New("No VCAP_SERVICES"))
	}

	var vcap VCAPServices
	if err := json.Unmarshal([]byte(services), &vcap); err != nil {
		fmt.Println("FAILED TO UNMARSHAL")
		panic(err)
	}

	fmt.Printf("%+v\n", vcap)

	dbURL := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		vcap.UserProvided[0].Credentials.Username,
		vcap.UserProvided[0].Credentials.Password,
		vcap.UserProvided[0].Credentials.Address,
		vcap.UserProvided[0].Credentials.Name,
	)

	fmt.Println("DB URL: ", dbURL)
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS photos(
title varchar(100) NOT NULL,
author varchar(100) NOT NULL,
hero varchar(50) NOT NULL,
data MEDIUMBLOB NOT NULL,
timestamp TIMESTAMP NOT NULL
);`)
	if err != nil {
		panic(err)
	}

	photobase := feedb.NewSQLPhotobase(db)
	handler := api.New(&photobase)

	http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
