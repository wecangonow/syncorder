package task

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	//DB "syncorder/database"
)

func Start() {

	ticker := time.NewTicker(time.Millisecond * 3000)

	go func() {
		for t := range ticker.C {
			log.Println("Tick at", t)
			//postJson := `{"fulfillment":{"tracking_number":"987654321","notify_customer": true}}`
			//shopify_authorization := "Basic " + base64.StdEncoding.EncodeToString([]byte(config.AppConfig.Shopify.AppKey+":"+config.AppConfig.Shopify.AppPassword))
			//httpclient.DoRequest("POST", config.AppConfig.Shopify.ApiUrl+"/orders/173639532576/fulfillments.json", postJson, shopify_authorization)
			names := getDirFileName("/tmp/syncorder/unfulfilled")
			if len(names) > 0 {
				fmt.Println("length of names is ", len(names))
				for _, name := range names {
					go func(name string) {
						fmt.Printf("printing in grountine %v number is %d\n", name, runtime.NumGoroutine())
						time.Sleep(100 * time.Microsecond)

					}(name)
				}
			}
		}
	}()
}

func getDirFileName(dir string) []string {
	fileList := []string{}
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil

	})

	ret := make([]string, 0)
	for _, file := range fileList {
		file = strings.Replace(file, dir, "", -1)
		file = strings.Replace(file, "/", "", -1)
		file = strings.Replace(file, ".json", "", -1)
		fmt.Println(file)
		if file != "" {
			ret = append(ret, file)
		}

	}
	return ret
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
