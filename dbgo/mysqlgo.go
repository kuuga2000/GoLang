package main
import (
  "database/sql"
  "net/http"
  "text/template"

  _ "github.com/go-sql-driver/mysql"
)

type Employee struct{
  Id int
  Name string
  City string
}

func dbConn()(db *sql.DB){
  dbDriver := "mysql"
  dbUser   := "root"
  dbPass   := "root"
  dbName   := "schooling_go"
  db, err  := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
  if(err != nil){
    panic(err.Error())
  }
  return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request){
  db := dbConn()
  selDB, err := db.Query("SELECT * FROM employee ORDER BY id DESC");
  if(err != nil){
    panic(err.Error())
  }
  emp := Employee{}
  res := []Employee{}
  for selDB.Next(){
    var id int
    var name,city string
    err = selDB.Scan(&id,&name,&city)
    if(err!=nil){
      panic(err.Error())
    }
    emp.Id = id
    emp.Name = name
    emp.City = city
    res = append(res,emp)
    tmpl.ExecuteTemplate(w,"Index",res)
    defer db.Close()
  }
}

func main(){
  http.HandleFunc("/", Index)
  http.ListenAndServe(":8080",nil)
}
