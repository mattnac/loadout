package request

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type webRequest struct {
	url string
}

type responseMap struct {
	TwoHundreds   int
	ThreeHundreds int
	FourHundreds  int
}

// var responseTimes []int

func Fire(url, uri string, port, count int, insecure bool, delay int) (returnData responseMap, responseTimes []int) {
	var (
		twoHundreds   = 0
		threeHundreds = 0
		fourHundreds  = 0
	)

	var (
		twoHundredsResp = promauto.NewCounter(prometheus.CounterOpts{
			Name: "two_hundred_responses",
			Help: "Total number of 200 responses",
		})
		threeHundredsResp = promauto.NewCounter(prometheus.CounterOpts{
			Name: "three_hundred_responses",
			Help: "Total number of 300 responses",
		})
		fourHundredsResp = promauto.NewCounter(prometheus.CounterOpts{
			Name: "four_hundred_responses",
			Help: "Total number of 300 responses",
		})
	)

	for counter := 0; count > counter; counter++ {
		reqData := renderRequest(url, uri, port)
		var failedReq int = 0

		if insecure == true {
		} else {
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		}
		start := time.Now()
		resp, err := http.Get(reqData)
		// time.Sleep(time.Duration(delay) * time.Millisecond)
		if err != nil {
			failedReq++
			fmt.Printf("Request number: %d failed with error %s", count, err)
			continue
		}

		elapsed := time.Since(start)
		responseTimes = append(responseTimes, int(elapsed.Milliseconds()))

		defer resp.Body.Close()
		if resp.StatusCode < 300 {
			twoHundreds++
			twoHundredsResp.Inc()
		} else if resp.StatusCode < 400 && resp.StatusCode > 299 {
			threeHundreds++
			threeHundredsResp.Inc()
		} else {
			fourHundreds++
			fourHundredsResp.Inc()
		}
	}
	returnData = responseMap{TwoHundreds: twoHundreds, ThreeHundreds: threeHundreds, FourHundreds: fourHundreds}
	return returnData, responseTimes
}

func renderRequest(url, uri string, port int) (reqUrl string) {
	fullUrl := webRequest{url: fmt.Sprintf("https://%s:%d/%s", url, port, uri)}

	return fullUrl.url
}
