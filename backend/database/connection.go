package database

import (
	//"encoding/json"
	"errors"
	"fmt"
	"os"

	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func CreateDatabaseConnection() {
	dsn := os.Getenv("DATABASE_DSN")
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db.AutoMigrate(&Menu{}, &PlanType{}, &Plan{})
}

//
// func GetAllMenus() {
// 	var menus []Menu
// 	db.Find(&menus)
//
// 	res, _ := json.Marshal(menus)
// 	fmt.Println(string(res))
//
// }

func GetTodaysMenu() (*Menu, error) {
	var menu Menu
	currentTime := time.Now()

	res := db.Where("active_date = ?", currentTime.Format("2006-01-02")).Find(&menu)
	if res.RowsAffected != 1 {
		return nil, errors.New("Today's menu not uploaded")
	}
	return &menu, nil
}

func TakeMessout(day time.Time) {
	day_as_string := day.Format("2006-01-02")
	fmt.Println(day_as_string)
}
