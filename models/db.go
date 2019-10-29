package models

import(
  "fmt"
  "log"
  _ "os"
  "database/sql"
  _ "github.com/lib/pq"
)

const(
  USER = "postgres"
  PASS = "postgres"
  DBNAME = "go_psql_demo_webapp_db"
)

func Connect() *sql.DB {
  URL := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", USER, PASS, DBNAME)
  db, err := sql.Open("postgres", URL)
  //db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  if err != nil {
    log.Fatal(err)
  }
  return db
}

func TestConnection() {
  con := Connect()
  defer con.Close()
  err := con.Ping()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Successfully connected to database!")
}
