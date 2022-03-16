package main

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"sqlUtil"
)

var (
	db  *sql.DB
	err error
)

const (
	MaxOpenConnes = 2000
	MaxIdlwConns  = 1000
)

type Car struct {
	Id      string `json:"car_id"` //json tag
	Mileage int    `json:"car_mileage"`
}

// type Carslice struct {
// 	Cars []Car `json:"cars"`
// }

func checkErr(err error) {
	if err != nil {
		log.Fatalf("error: %v\n", err)
		return
	}
}

func main() {
	//open database
	db, err = sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
	checkErr(err)
	db.SetMaxOpenConns(MaxOpenConnes)
	db.SetMaxIdleConns(MaxIdlwConns)

	//Initialize the web framework
	router := gin.Default()
	router.GET("/cars", handleGetAll)
	router.GET("/cars/:id", handleGetOneCar)
	router.POST("/post", handlePost)
	router.PUT("/put", handlePut)
	router.DELETE("/cars/:id", handleDelete)

	//Initialize the server
	pool := x509.NewCertPool()
	caCertPath := "./ca.crt"
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}
	err = srv.ListenAndServeTLS("./server.crt", "./server.key")
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}

//router.GET("/cars", handleGetAll)
func handleGetAll(c *gin.Context) {
	rows, err := db.Query(`SELECT * from car_info`)
	defer rows.Close()
	checkErr(err)

	cars := make([]Car, 0)
	for rows.Next() {
		var car Car
		rows.Scan(&car.Id, &car.Mileage)
		cars = append(cars, car)
	}
	err = rows.Err()
	checkErr(err)
	c.String(http.StatusOK, "All the information in table car_info is here:\n")

	for _, car := range cars {
		carId := car.Id
		carMileage := car.Mileage
		c.String(http.StatusOK, "Car Id=%s,Car Mileage=%d\n", carId, carMileage)
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"cars": cars,
	// })
}

//router.GET("/cars/:id", handleGetOneCar)
func handleGetOneCar(c *gin.Context) {
	id := c.Param("id")
	rows, err := db.Query(`select *from car_info where id=?`, id)
	defer rows.Close()
	checkErr(err)
	cars := make([]Car, 0)
	for rows.Next() {
		var car Car
		rows.Scan(&car.Id, &car.Mileage)
		cars = append(cars, car)
	}
	err = rows.Err()
	checkErr(err)
	c.JSON(http.StatusOK, gin.H{
		"cars": cars,
	})
	c.String(http.StatusOK, "\n")
}

//router.POST("/post", handlePost)
func handlePost(c *gin.Context) {
	//var carslice Carslice
	//body, err := ioutil.ReadAll(c.Request.Body)
	//checkErr(err)
	// json.Unmarshal([]byte(body), &carslice)
	// fmt.Println("Here is the JSON data from client:")
	// fmt.Println(carslice)

	// //add the data to database
	// for _, car := range carslice.Cars {
	// 	carId := car.Id
	// 	carMileage := car.Mileage
	// 	fmt.Printf("Car Id=%s,Car Mileage=%d\n", carId, carMileage)
	// 	sqlUtil.InsertData(db, carId, carMileage)
	// }

	var car Car
	body, err := ioutil.ReadAll(c.Request.Body)
	checkErr(err)
	//json.Unmarshal([]byte(body), &car)
	json.Unmarshal(body, &car)
	fmt.Println("Here is the JSON data from client:")
	fmt.Println(car)
	sqlUtil.InsertData(db, car.Id, car.Mileage)
}

//router.PUT("/put", handlePut)
func handlePut(c *gin.Context) {
	var car Car
	body, err := ioutil.ReadAll(c.Request.Body)
	checkErr(err)
	json.Unmarshal([]byte(body), &car)
	fmt.Println("Here is the JSON data from client:")
	fmt.Println(car)
	sqlUtil.UpdateData(db, car.Id, car.Mileage)
}

//router.DELETE("/cars/:id", handleDelete)
func handleDelete(c *gin.Context) {
	id := c.Param("id")
	sqlUtil.DeleteData(db, id)
}
