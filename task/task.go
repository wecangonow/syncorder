package task

import (
	"time"
	"log"

	DB "syncorder/database"
)

func Start() {

	ticker := time.NewTicker(time.Millisecond * 1000)

	go func() {
		for t := range ticker.C {
			log.Println("Tick at", t)
		}
	}()
}

func getShopifyOrders() {

	//// Read all fish from the database, unmarshaling the response.
	//records, err := db.ReadAll("fish")
	//if err != nil {
	//	fmt.Println("Error", err)
	//}
	//
	//fishies := []Fish{}
	//for _, f := range records {
	//	fishFound := Fish{}
	//	if err := json.Unmarshal([]byte(f), &fishFound); err != nil {
	//		fmt.Println("Error", err)
	//	}
	//	fishies = append(fishies, fishFound)
	//}
	//
	//fmt.Printf("%v", fishies)

}

