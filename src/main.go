package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "host"
	port     = port
	user     = "user"
	password = "password"
	dbname   = "dbname"
)

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	// CheckError(err)

	// defer db.Close()

	for a := 0; a < 1; a++ {
		for b := 0; b < 256; b++ {
			for c := 0; c < 256; c++ {
				for d := 0; d < 256; d++ {

					response := strconv.Itoa(a) + "." + strconv.Itoa(b) + "." + strconv.Itoa(c) + "." + strconv.Itoa(d)
					fmt.Println(response) //Output: Variable string 14 content

					insertCicle := `insert into "dns3"("ip") values('` + response + `')`
					// fmt.Println(insertCicle)
					// _, e := db.Exec(insertStmt)
					_, kal := db.Exec(insertCicle)
					// CheckError(e)
					// CheckError(kal)
					CheckError(kal)

				}
			}
		}
	}

	// // connection string
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// // open database
	// db, err := sql.Open("postgres", psqlconn)
	// CheckError(err)

	// close database
	// defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	// INSERT INTO dnsrecord ( ip ) VALUES ($a.$b.$c.$d);"
	// insertStmt := `insert into "dns"("ip") values(1)`
	// _, e := db.Exec(insertStmt)
	// CheckError(e)

	fmt.Println("Connected!")

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
