package httpserver

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"syncorder/config"

	"io/ioutil"
)

func orderCreate(rw http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	if r.Method == "POST" {
		hmac246 := r.Header.Get("X-Shopify-Hmac-Sha256")
		result, _ := ioutil.ReadAll(r.Body)
		caculateHmac256 := ComputeHmac256(string(result), config.AppConfig.Shopify.SharedSecret)
		fmt.Printf("%v\n", caculateHmac256)
		fmt.Printf("%v\n", hmac246)

		rw.Write([]byte(config.AppConfig.Fetchx.Authorization))

	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte("method not allowed"))
	}

}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func Start() {
	http.HandleFunc("/orderCreate", orderCreate)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println(err.Error())
	}

}
