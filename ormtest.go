package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
		"fmt"
)

type Rol struct {
  ID int `gorm:"primary_key"`
  Nombre string `gorm:"column:nombre"`
  Sistema_id int `gorm:"column:sistema_id"`
}

func (Rol) TableName() string {
  return "roles"
}

func main(){
	db, err := gorm.Open("sqlite3", "db/db_accesos.db")
  if err != nil {
    panic("failed to connect database")
  }
  //defer db.Close()
  db.Create(&Rol{Nombre: "Rol ORM", Sistema_id: 8})
  var roles []Rol
  if err := db.Find(&roles).Error; err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(roles)
  }
    
  db.Close()
}