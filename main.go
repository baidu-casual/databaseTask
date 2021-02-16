package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

type cred struct {
	username string
	password string
}

//Events table
type Events struct {
	EventName string `gorm:"column:event_name"`
}

//Filenames contains the csv filenames
var (
	Filenames []string
	user      cred
)

func init() {
	Filenames = []string{"test.db", "SQL-selector_data", "SQL-events_data", "SQL-events", "OPTION"}
	user.username = "baidu"
	user.password = "IxEMu4p5s3ApbbQD"
}
func sqlFileToDB(filename string, db *gorm.DB) {
	var flag error

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
func script(filename string, db *gorm.DB) {
	var flag error

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
		fmt.Printf("Rows Affected : %v\n", result.RowsAffected)
		flag = result.Error
		if flag != nil {
			log.Fatalf("Error : %f", flag)
		}
	}
}
func opt10(db *gorm.DB) []EventsDataAVG {
	var (
		flag error
		//res  []EventsDataAVG
		evnt []Events
		data []EventsDataAVG
	)

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

func main() {
	login := user.username + ":" + user.password
	dsn := "@tcp(localhost:3306)/xenonstack?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(login+dsn), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}
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
			sqlFileToDB(Filenames[1], db)
			sqlFileToDB(Filenames[2], db)
			sqlFileToDB(Filenames[3], db)
		case 2:
			script((Filenames[4] + "2"), db)
		case 3:
			script((Filenames[4] + "3"), db)
		case 4:
			script((Filenames[4] + "4"), db)
		case 5:
			script((Filenames[4] + "5"), db)
		case 6:
			script((Filenames[4] + "6"), db)
		case 7:
			script((Filenames[4] + "7"), db)
		case 8:
			script((Filenames[4] + "8"), db)
		case 9:
			script((Filenames[4] + "9"), db)
		case 10:
			table := opt10(db)
			for i := range table {
				fmt.Println(table[i])
				db.Create(&table[i])
			}
		case 11:
			script((Filenames[4] + "11"), db)
		case 13:
			sqlFileToDB((Filenames[4] + "13_1"), db)
			sqlFileToDB((Filenames[4] + "13_2"), db)
			sqlFileToDB((Filenames[4] + "13_3"), db)
			sqlFileToDB((Filenames[4] + "13_4"), db)
		case 19:
			os.Exit(0)
		default:
			fmt.Println("Wrong Choice!")
		}
	}
}
