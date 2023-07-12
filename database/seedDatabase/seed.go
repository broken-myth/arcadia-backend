package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"

	"github.com/delta/arcadia-backend/config"
	"github.com/delta/arcadia-backend/database"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
)

// printing the seeded rows
func PrintSeededRow(tableName string, seedStatus string, element map[string]interface{}) {
	delete(element, "created_at")
	delete(element, "updated_at")
	count := 0
	last := len(element)

	fmt.Print(color.HiCyanString(tableName + " : " + seedStatus + " seed : "))
	for k, v := range element {
		count++
		fmt.Print(color.HiMagentaString(k) + ":")
		fmt.Print(v)
		if count == last {
			fmt.Print("")
		} else {
			fmt.Print(", ")
		}
	}
}

// seeding the table
func seedTable(name string, result map[string][]map[string]interface{}) {
	fmt.Println(color.BlueString("Sarted seeding the table : " + name))

	seedContent := result[name]
	db := database.GetDB()

	for _, element := range seedContent {
		id := element["id"]
		var count int64

		db.Table(name).Where("id = ?", id).Count(&count)

		if count == 0 {
			element["created_at"] = time.Now()
			element["updated_at"] = time.Now()
			if err := db.Table(name).Create(element).Error; err != nil {
				PrintSeededRow(name, "error in creating", element)
				fmt.Println("\n", color.RedString("Error:"), err)
			} else {
				PrintSeededRow(name, "created", element)
			}
		} else {
			element["updated_at"] = time.Now()
			if err := db.Table(name).Where("id = ?", id).Updates(element).Error; err != nil {
				PrintSeededRow(name, "error in updating", element)
				fmt.Println("\n", color.RedString("Error:"), err)
			} else {
				PrintSeededRow(name, "updated", element)
			}
		}
		fmt.Println("")
	}
}

// seeding data
func SeedData(seeds []string) {
	for _, v := range seeds {
		result := utils.ReadJSON("database/seedDatabase/content/" + v + ".json")
		seedTable(v, result)
	}
}

func main() {

	config.InitConfig()

	utils.InitLogger()

	database.ConnectMySQLdb()

	model.MigrateDB()

	appEnv := os.Getenv("APP_ENV")

	// common for both prod and dev
	var seeds = []string{
		"constants",
		"arenas",
		"perks",
		"minicons",
		"characters",
	}

	// only for dev
	var devSeeds = []string{
		"user_registrations",
		"users",
	}

	if appEnv == "DEV" {
		// do not seed these tables for DOCKER env
		seeds = append(seeds, devSeeds...)
	}

	SeedData(seeds)
}
