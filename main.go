package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"databaseTask/cmd"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //using sqlite
)

//Filenames contains the csv filenames
var Filenames []string

func init() {
	Filenames = []string{"test.db", "selector_data", "events_data", "events"}
}
func sqlFileToDB(filename string, dbname string) {
	var flag error

	db, err := gorm.Open("sqlite3", dbname)
	if err != nil {
		panic("can't connect to database")
	}
	defer db.Close()
	db.LogMode(true)

	file, err := os.Open(string("sql/SQL-" + filename + ".sql"))
	if err != nil {
		log.Fatalf("failed to open")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for _, eachln := range text {
		fmt.Println(eachln)
		result := db.Exec(eachln)
		flag = result.Error
		if flag != nil {
			log.Fatalf("Error : %f", flag)
		}
	}
}

func main() {
	cmd.Exec()

	fmt.Print("sql/SQL-" + Filenames[0] + ".sql\n")
	sqlFileToDB(Filenames[1], Filenames[0])
	sqlFileToDB(Filenames[2], Filenames[0])
	sqlFileToDB(Filenames[3], Filenames[0])
}
