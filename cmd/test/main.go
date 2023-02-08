package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"xmirror.cn/iast/goat/model"
)

var (
	BaseUrl string
	ApiInfo *model.ApiInfo
	Count   int64

	Num  = flag.Int("n", 10, "number of rounds per worker run")
	Cnt  = flag.Int("c", 10, "count of concurrent worker(s)")
	Ip   = flag.String("ip", "", "IP of the target host")
	Port = flag.Int("p", 0, "Port of the target site")
)

func main() {
	flag.Parse()

	var err error
	ApiInfo, err = ParseApiInfo()
	if err != nil {
		log.Printf("ParseApiInfo err:%s\n", err)
		return
	}

	var ok bool
	if BaseUrl, ok = getBaseUrl(ApiInfo.Variable); !ok {
		log.Printf("baseurl can't be empty\n")
		return
	}
	log.Printf("access %s:", BaseUrl)
	startTime := time.Now()
	wait := sync.WaitGroup{}
	for i := 0; i < *Cnt; i++ {
		wait.Add(1)
		log.Printf("the %d concurrent request begins \n", i)
		go func() {
			traverse(*Num)
			wait.Done()
		}()
	}
	wait.Wait()
	log.Printf("request ended, %d requests per target, It takes : %s \n", Count, time.Since(startTime).String())
}

func traverse(number int) {
	for i := 0; i < number; i++ {
		atomic.AddInt64(&Count, 1)
		for _, v := range ApiInfo.Item {
			if v.Request.Method == "POST" {
				requestPost(&v)
			} else if v.Request.Method == "GET" {
				requestGet(&v)
			}
		}
	}
}

func getBaseUrl(variable []model.Variable) (string, bool) {
	if len(*Ip) > 0 {
		if parsed := net.ParseIP(*Ip); parsed != nil {
			if *Port == 0 {
				return fmt.Sprintf("http://%s/", parsed.String()), true
			} else {
				return fmt.Sprintf("http://%s:%d/", parsed.String(), *Port), true
			}
		}
	}
	for _, v := range variable {
		if v.Key == "baseUrl" {
			return v.Value, true
		}
	}
	return "", false
}

func requestGet(v *model.Item) {
	urls := getURl(v.Request.URL)
	for _, url := range urls {
		if url == "" {
			log.Printf("failed to get url from %s \n", v.Name)
			continue
		}
		httpResult, err := httpDo(url, "GET", nil, nil)
		if err != nil {
			continue
		}
		log.Printf("Request API %s returns %s ", url, httpResult)
	}
}

func requestPost(v *model.Item) {
	urls := getURl(v.Request.URL)
	urlencoded := v.Request.Body.Urlencoded[0]
	header := http.Header{}
	header.Add("Content-Type", "application/x-www-form-urlencoded")
	for _, url := range urls {
		if url == "" {
			log.Printf("failed to get url from %s \n", v.Name)
			continue
		}
		payload := strings.NewReader(fmt.Sprintf("%s=%s", urlencoded.Key, urlencoded.Value))
		httpResult, err := httpDo(url, "POST", payload, header)
		if err != nil {
			continue
		}
		log.Printf("Request API %s returns %s ", url, httpResult)
	}
}

func getURl(urlInterface interface{}) [2]string {
	value := reflect.ValueOf(urlInterface)
	var urlStr string
	switch value.Kind() {
	case reflect.String:
		urlStr = urlInterface.(string)
	case reflect.Map:
		urlStr = urlInterface.(map[string]interface{})["raw"].(string)
	}

	urlStr = strings.Replace(urlStr, "{{baseUrl}}", BaseUrl, -1)
	var urlStr2 string
	if strings.Contains(urlStr, "/unsafe/") {
		urlStr2 = strings.Replace(urlStr, "/unsafe/", "/safe/", 1)
		return [2]string{urlStr, urlStr2}
	} else {
		urlStr2 = strings.Replace(urlStr, "/safe/", "/unsafe/", 1)
		return [2]string{urlStr2, urlStr}
	}

}

func ParseApiInfo() (*model.ApiInfo, error) {

	file, err := os.Open("postman_collection.json")
	if err != nil {

		return nil, err
	}
	jsonData, err := ioutil.ReadAll(file)
	if err != nil {

		return nil, err
	}
	var toolTT = &model.ApiInfo{}
	err = json.Unmarshal(jsonData, toolTT)
	if err != nil {

		return nil, err
	}
	return toolTT, nil
}

func httpDo(urlStr, method string, reader io.Reader, header http.Header) (string, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(method, urlStr, reader)
	if header != nil {
		req.Header = header
	}
	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}
