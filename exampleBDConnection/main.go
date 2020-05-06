package main

import (
  "database/sql"
  "fmt"
  _ "github.com/godror/godror"
)
func main() {
  db, err := sql.Open("godror", "go_user/1234@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=tcp)(HOST=localhost)(PORT=1521)))(CONNECT_DATA=(SERVICE_NAME=xe)))")

  if err != nil {
    fmt.Println(err)
  }
  var col string
  sqlStatement := `select  1 as id from dual`
  row := db.QueryRow(sqlStatement)
  err2 := row.Scan(&col)
    if err2 != nil {
      fmt.Println(err2)
    }
  fmt.Printf("col=%s" , col)
  value, err := db.Query(`select 1 from dual
                          union 
                          select 2 from dual`)
  var table_name string
  for value.Next() {
    err := value.Scan(&table_name)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Printf("table_name=%s" , table_name)
  }
}