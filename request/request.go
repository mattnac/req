package request

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

type webRequest struct {
	url string
}

type responseMap struct {
	TwoHundreds   int
	ThreeHundreds int
	FourHundreds  int
}

func Fire(url, uri string, port, count int, insecure bool) (returnData responseMap) {
	var (
		twoHundreds   = 0
		threeHundreds = 0
		fourHundreds  = 0
	)

	for counter := 0; count > counter; counter++ {
		reqData := renderRequest(url, uri, port)
		var failedReq int = 0

		//if insecure = true {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		//}

		resp, err := http.Get(reqData)
		if err != nil {
			failedReq++
			fmt.Printf("Request number: %d failed with error %s", count, err)
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
	returnData = responseMap{TwoHundreds: twoHundreds, ThreeHundreds: threeHundreds, FourHundreds: fourHundreds}

	return returnData
}

func renderRequest(url, uri string, port int) (reqUrl string) {
	fullUrl := webRequest{url: fmt.Sprintf("https://%s:%d/%s", url, port, uri)}

	return fullUrl.url
}
