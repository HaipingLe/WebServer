package sqlUtil

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalf("error: %v\n", err)
		return
	}
}

func InsertData(db *sql.DB, carId string, carMileage int) {
	fmt.Printf("Enter insertData()\n")
	stmt, err := db.Prepare("insert into car_info values(?,?)")
	defer stmt.Close()
	checkErr(err)

	fmt.Printf("start to insert data (%s,%d).\n", carId, carMileage)
	res, err := stmt.Exec(carId, carMileage)
	checkErr(err)
	last_insert_id, err := res.LastInsertId()
	checkErr(err)

	fmt.Printf("insert data successfully.\n")
	fmt.Printf("Last Insert Id=%d\n", last_insert_id)
}

func UpdateData(db *sql.DB, carId string, carMileage int) {
	stmt, err := db.Prepare("UPDATE car_info SET mileage=? WHERE id=?")
	defer stmt.Close()
	checkErr(err)

	res, err := stmt.Exec(carMileage, carId)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Printf("%d", num)

}

func DeleteData(db *sql.DB, carId string) {
	fmt.Println("delete")
	stmt, err := db.Prepare("DELETE FROM car_info WHERE id=?")
	defer stmt.Close()
	checkErr(err)

	row, err := stmt.Query(carId)
	defer row.Close()
	checkErr(err)
}
