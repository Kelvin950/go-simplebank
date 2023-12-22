package main

import (
	"database/sql"
	"log"

	"github.com/kelvin950/simplebank/api"
	"github.com/kelvin950/simplebank/db"
	"github.com/kelvin950/simplebank/util"
	_ "github.com/lib/pq"
)

//***sqlc_01HH5KJ4HM1QQRJYY0GJB0AJN3**
//export SQLC_AUTH_TOKEN=sqlc_xxxxxxxx
func main() {

	config , err:= util.LoadConfig(".")
 
	if err!=nil{
		log.Fatal("cannot load  config:",err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatalf("%s", err)
	}

	store :=db.NewStore(conn) 
	server ,err:= api.NewServer(config ,store )

	if err!=nil{
		log.Fatalf("%s" , err) 
	} 
	
	err =  server.Start(config.ServerAddress) 
	if err!=nil{
		log.Fatalf("%s" , err) 
	}



}