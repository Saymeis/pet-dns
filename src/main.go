package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "host"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "dbname"
)

/*
	psql:
	create table go (id integer, ip varchar(15), A varchar(255), CNAME varchar(255), NS varchar(255), SOA varchar(255), SPF varchar(255), SRV varchar(255), TXT varchar(255));
*/

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)

	for a := 0; a < 1; a++ {
		for b := 0; b < 256; b++ {
			for c := 0; c < 256; c++ {
				for d := 0; d < 256; d++ {

					response := strconv.Itoa(a) + "." + strconv.Itoa(b) + "." + strconv.Itoa(c) + "." + strconv.Itoa(d)

					insertCicle := `insert into "go"("ip") values('` + response + `')`

					_, kal := db.Exec(insertCicle)
					CheckError(kal)

				}
			}
		}
	}
	CheckError(err)
	fmt.Println("Done!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
