package request

import (
	"fmt"
	"net/http"
	"strconv"
)

type webRequest struct {
	url string
}

type responseMap struct {
	twoHundreds   int
	threeHundreds int
	fourHundreds  int
}

func Fire(url, uri string, port, count int) (returnData responseMap) {
	var twoHundreds int = 0
	var threeHundreds int = 0
	var fourHundreds int = 0

	for counter := 0; count > counter; counter++ {
		reqData := renderRequest(url, uri, port)
		var failedReq int = 0

		fmt.Printf("Firing request #: %d out of %d \n", counter, count)

		resp, err := http.Get(reqData)
		if err != nil {
			failedReq++
			fmt.Println("Error:", err)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode < 300 {
			twoHundreds++
		} else if resp.StatusCode < 400 && resp.StatusCode > 299 {
			threeHundreds++
		} else {
			fourHundreds++
		}
	}
	returnData = responseMap{twoHundreds: twoHundreds, threeHundreds: threeHundreds, fourHundreds: fourHundreds}

	return returnData
}

func renderRequest(url, uri string, port int) (reqUrl string) {
	webPort := strconv.Itoa(port)
	fullUrl := webRequest{url: "https://" + url + ":" + webPort + uri}

	return fullUrl.url
}
