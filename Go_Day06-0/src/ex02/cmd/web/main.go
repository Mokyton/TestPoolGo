package main

import (
	"encoding/gob"
	"ex01/pkg/db/psql"
	"flag"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	thoughts *psql.Model
}

func main() {
	addr := flag.String("addr", ":8888", "Сетевой адрес HTTP")
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := psql.NewDbConnection(host, port, userdb, password, dbname)
	if err != nil {
		errorLog.Fatal(err)
	}
	err = db.CreateTable()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		thoughts: db,
	}

	gob.Register(sesKey(0))

	srv := &http.Server{
		Addr:      *addr,
		Handler:   app.Router(),
		ErrorLog:  errorLog,
		ConnState: connections.OnstateChange,
	}

	infoLog.Printf("Запуск сервера на %s\n", *addr)
	srv.ListenAndServe()
}
