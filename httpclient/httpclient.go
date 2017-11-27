package httpclient

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// PUT POST GET
func DoRequest(method string, url string, data string, authorization string) (string, int) {
	fmt.Println(" Request url :  ", url)
	client := &http.Client{}
	request, err := http.NewRequest(method, url, strings.NewReader(data))
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", authorization)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("StatusCode: ", response.StatusCode)
	fmt.Println(string(contents))
	return string(contents), response.StatusCode
}

func DoPost(url string, data string) {

}
