package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)
var testQueries *Queries 
 var testDb   *sql.DB
const (
	dbDriver = "postgres"
	dbSourcename=  "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)
 

func TestMain(m *testing.M){
	testDb , err := sql.Open(dbDriver , dbSourcename) 

	 if err!=nil{
		log.Fatalf("%s" , err) 
	 } 


	
	 testQueries  =  New(testDb) 
  
	 os.Exit(m.Run())
}