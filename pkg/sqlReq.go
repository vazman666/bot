package pkg

import (
	"bot/models"
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//var database *sql.DB

func SqlReq(partnum string) []models.Sklad {
	db, err := sql.Open("mysql", "vazman:rbhgbxb1@unix(/var/run/mysqld/mysqld.sock)/japautozap")

	if err != nil {
		log.Println(err)
	}
	var req []models.Sklad
	defer db.Close()
	partnum = strings.ToUpper(partnum)
	fmt.Printf("Request to base, partnum=%v\n", partnum)
	rows, err := db.Query("select * from sklad where partnum = ? ", partnum)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tmp models.Sklad

		err := rows.Scan(&tmp.Id, &tmp.Name, &tmp.Firm, &tmp.Qtym, &tmp.Qtyt, &tmp.Price, &tmp.Cellm, &tmp.Cellt, &tmp.Partnum)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Нашлось в базе %v\n", tmp)
		req = append(req, tmp)
	}
	//fmt.Printf("Nothing found tmp.Id=%v\n", tmp.Id)
	if len(req) == 0 {
		req = append(req, models.Sklad{Name: "nothing found"})
	}
	return req

}
