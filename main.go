package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //using sqlite
)

//EventsDataAVG table
type EventsDataAVG struct {
	Slno            string  `gorm:"column:serial_number"`
	EventTimestamp  string  `gorm:"column:event_timestamp"`
	TempAvg         string  `gorm:"column:temp-avg"`
	ActivePwrAvg    string  `gorm:"column:active_pwr-avg"`
	WindDirAvg      string  `gorm:"column:wind_dir-avg"`
	AvailablePwrAvg string  `gorm:"column:available_pwr-avg"`
	EventAvg        float64 `gorm:"column:events-avg"`
}

//Events table
type Events struct {
	EventName string `gorm:"column:event_name"`
}

//Filenames contains the csv filenames
var Filenames []string

func init() {
	Filenames = []string{"test.db", "SQL-selector_data", "SQL-events_data", "SQL-events", "OPTION"}
}
func sqlFileToDB(filename string, dbname string) {
	var flag error

	db, err := gorm.Open("sqlite3", dbname)
	if err != nil {
		panic("can't connect to database")
	}
	defer db.Close()
	db.LogMode(true)

	file, err := os.Open(string("sql/" + filename + ".sql"))
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
func script(filename string, dbname string) {
	var flag error

	db, err := gorm.Open("sqlite3", dbname)
	if err != nil {
		panic("can't connect to database")
	}
	defer db.Close()
	db.LogMode(true)

	file, err := os.Open(string("sql/" + filename + ".sql"))
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
		result := db.Raw(eachln)
		flag = result.Error
		if flag != nil {
			log.Fatalf("Error : %f", flag)
		}
	}
}
func opt10(dbname string) []EventsDataAVG {
	var (
		flag error
		//res  []EventsDataAVG
		evnt []Events
		data []EventsDataAVG
	)
	db, err := gorm.Open("sqlite3", dbname)
	if err != nil {
		panic("can't connect to database")
	}
	defer db.Close()

	//db.LogMode(true)

	result := db.Raw("SELECT * FROM EventsData ORDER BY serial_number, event_timestamp;").Scan(&data)
	flag = result.Error
	if flag != nil {
		log.Fatalf("Error : %f", flag)
	}

	for i := range data {
		var (
			temp string  = ""
			sum  float64 = 0.0
			j    float64 = 0.0
		)

		result = db.Raw("SELECT event_name FROM Events WHERE serial_number = ?", data[i].Slno).Scan(&evnt)
		flag = result.Error
		if flag != nil {
			log.Fatalf("Error : %f", flag)
		}
		for _, word := range evnt {
			temp += (word.EventName + " ")
			if word.EventName == "available_pwr-avg" {
				t, err := strconv.ParseFloat(data[i].AvailablePwrAvg, 64)
				if err != nil {
					log.Fatalf("Error : %v", err)
				}
				sum += t
			}
			if word.EventName == "active_pwr-avg" {
				t, err := strconv.ParseFloat(data[i].ActivePwrAvg, 64)
				if err != nil {
					log.Fatalf("Error : %v", err)
				}
				sum += t
			}
			if word.EventName == "temp-avg" {
				t, err := strconv.ParseFloat(data[i].TempAvg, 64)
				if err != nil {
					log.Fatalf("Error : %v", err)
				}
				sum += t
			}
			if word.EventName == "wind_dir-avg" {
				t, err := strconv.ParseFloat(data[i].WindDirAvg, 64)
				if err != nil {
					log.Fatalf("Error : %v", err)
				}
				sum += t
			}
			j++

		}
		avg := sum / j
		data[i].EventAvg = avg
	}
	return data
}

func createEventsDataAVG(event EventsDataAVG, dbname string) {
	db, err := gorm.Open("sqlite3", dbname)
	if err != nil {
		panic("can't connect to database")
	}
	defer db.Close()
	db.LogMode(true)

	db.Create(&event)
}

func main() {
	var option int
	for {
		fmt.Print("\n\n\nOPTIONS :-\n\n")
		fmt.Print("\t1. Import the data to database.\n")
		fmt.Print("\t2. Query all the data from three tables individually.\n")
		fmt.Print("\t3. Query data for any one serial_number from events_data table.\n")
		fmt.Print("\t4. Show Count of records in all three tables.\n")
		fmt.Print("\t5. Show count of records in events_data table for each serial_number.\n")
		fmt.Print("\t6. Show count of records in events_data table for every super_id.\n")
		fmt.Print("\t7. Show count of records in events_data table for every serial_number and event_timestamp.\n")
		fmt.Print("\t8. Show if there is two records exist for same serial_number and event_timestamp.\n")
		fmt.Print("\t9. Show the minimum and maximum event_date for every serial_number.\n")
		fmt.Print("\t10. Calculate the average of all events for each serial_number and event_date [check events table for distinct event_name] and name the column as generation.(*)\n")
		fmt.Print("\t11. Sort all the data with serial_number and event_timestamp.\n")
		fmt.Print("\t12. Use all types of Joins to join all three tables and describe the difference between every Join type.\n")
		fmt.Print("\t13. Normalise all the tables to 2nd Normalise form. Define the Primary Key to every Table and use as foreign key in other tables [where ever possible].\n")
		fmt.Print("\t14. Update all the null/empty values/strings to 0 in all tables.\n")
		fmt.Print("\t15. Change Data type of event_timestamp column to timestamp [or long].\n")
		fmt.Print("\t16. Describe which columns are of no use.\n")
		fmt.Print("\t17. Calculate the average of active_pwr-avg and available_pwr-avg for each serial number and event number as name as actual_generation and expected_generation respectively.\n")
		fmt.Print("\t18. Show the records where actual_generation > expected_generation and expected_generation < actual_generation.\n")
		fmt.Print("\t19. EXIT\n")
		fmt.Print("\n\tEnter your choice?")

		fmt.Scanf("%v\n", &option)
		switch option {
		case 1:
			//cmd.Exec()
			fmt.Print("sql/SQL-" + Filenames[0] + ".sql\n")
			sqlFileToDB(Filenames[1], Filenames[0])
			sqlFileToDB(Filenames[2], Filenames[0])
			sqlFileToDB(Filenames[3], Filenames[0])
		case 2:
			script((Filenames[4] + "2"), Filenames[0])
		case 3:
			script((Filenames[4] + "3"), Filenames[0])
		case 4:
			script((Filenames[4] + "4"), Filenames[0])
		case 5:
			script((Filenames[4] + "5"), Filenames[0])
		case 6:
			script((Filenames[4] + "6"), Filenames[0])
		case 7:
			script((Filenames[4] + "7"), Filenames[0])
		case 8:
			script((Filenames[4] + "8"), Filenames[0])
		case 9:
			script((Filenames[4] + "9"), Filenames[0])
		case 10:
			table := opt10(Filenames[0])
			for i := range table {
				fmt.Println(table[i])
				createEventsDataAVG(table[i], Filenames[0])
			}
		case 11:
			script((Filenames[4] + "11"), Filenames[0])
		case 13:
			sqlFileToDB((Filenames[4] + "13_1"), Filenames[0])
			sqlFileToDB((Filenames[4] + "13_2"), Filenames[0])
			sqlFileToDB((Filenames[4] + "13_3"), Filenames[0])
			sqlFileToDB((Filenames[4] + "13_4"), Filenames[0])
		case 19:
			os.Exit(0)
		default:
			fmt.Println("Wrong Choice!")
		}
	}

}
