package task

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syncorder/config"
	DB "syncorder/database"
	"syncorder/httpserver"
	"syncorder/logger"
	"time"
)

type OrderTrackingInfo struct {
	TrackingInformation struct {
		StatusName        string `json:"status_name"`
		StatusCode        string `json:"status_code"`
		StatusDate        string `json:"status_date"`
		Source            string `json:"source"`
		StatusDateLocal   string `json:"status_date_local"`
		ToHub             string `json:"to_hub"`
		StatusDescription string `json:"status_description"`
		CurrentHub        string `json:"current_hub"`
	} `json:"tracking_information"`
	OrderInformation struct {
		TrackingNo string `json:"tracking_no"`
		SoNumber   string `json:"so_number"`
		ClientRef  string `json:"client_ref"`
	} `json:"order_information"`
}
type FulfillmentReqData struct {
	Fulfillment struct {
		TrackingNumber  string `json:"tracking_number"`
		TrackingCompany string `json:"tracking_company"`
		TrackingURL     string `json:"tracking_url"`
		NotifyCustomer  bool   `json:"notify_customer"`
	} `json:"fulfillment"`
}

func Start() {

	ticker := time.NewTicker(time.Second * config.AppConfig.TickSeconds)

	go func() {
		for t := range ticker.C {
			log.Println("Tick at", t)
			names := getDirFileName(config.AppConfig.FilePath.DbFilePath + "/unfulfilled")
			if len(names) > 0 {
				for _, name := range names {
					go syncorder(name)
				}
			}
		}
	}()
}

func syncorder(name string) {
	//postJson := `{"fulfillment":{"tracking_number":"987654321","notify_customer": true}}`
	//shopify_authorization := "Basic " + base64.StdEncoding.EncodeToString([]byte(config.AppConfig.Shopify.AppKey+":"+config.AppConfig.Shopify.AppPassword))
	//httpclient.DoRequest("POST", config.AppConfig.Shopify.ApiUrl+"/orders/173639532576/fulfillments.json", postJson, shopify_authorization)
	get_url := config.AppConfig.Fetchx.ApiUrl + "/" + name + "?reference_type=client_ref"
	get_url = strings.Replace(get_url, "#", "%23", -1)
	Logger := logger.GetInstance()
	Logger.Printf("printing in grountine %v order id is %s  get url is %s \n", name, getShopifyOrderIdByName(name), get_url)

	time.Sleep(100 * time.Microsecond)

}

func createFulfillment(order_id string) {

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
		if file != "" {
			ret = append(ret, file)
		}

	}
	return ret
}

func getShopifyOrderIdByName(name string) string {
	db := DB.GetInstance()
	orderInfo := httpserver.OrderInfo{}

	db.Read("unfulfilled", name, &orderInfo)

	return strconv.FormatInt(orderInfo.ID, 10)
}
