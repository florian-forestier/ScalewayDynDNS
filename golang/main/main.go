package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Connect struct {
	Name string `json:"name"`
	Type string `json:"type"`
	ChangeType string `json:"changeType"`
	Records []Records `json:"records"`
}

type Records struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Priority int `json:"priority"`
	TTL int `json:"ttl"`
	Data string `json:"data"`
}

func main() {
	//Get data
	var onlineDomain string
	var onlineAPI string

	flag.StringVar(&onlineDomain, "domain", "", "Define the domain or subdomain to update (eg github.com or example.github.com)")
	flag.StringVar(&onlineAPI, "key", "", "Set the key for your Online API access")
	flag.Parse()

	if onlineDomain=="" || onlineAPI=="" {
		log.Fatal("No domain or key given, abort")
	}

	//Exploding onlineDomain
	exploded := strings.Split(onlineDomain, ".")

	domain := exploded[len(exploded)-2]+"."+ exploded[len(exploded)-1]
	subdomain := ""
	if len(exploded) > 2 {
		for i:= len(exploded)-3; i>=0; i-- {
			subdomain = exploded[i]+"."+subdomain
		}
		subdomain = subdomain[:len(subdomain)-1]
	}


	//Get your IP on ifconfig.me
	doConfigCall, _ := http.NewRequest("GET", "http://ifconfig.me", nil)
	client := &http.Client{Timeout: time.Second * 20}

	ipAnswer, err := client.Do(doConfigCall)
	if err != nil || ipAnswer.StatusCode != 200 {
		log.Fatal("Error while trying to GET ifconfig.me !")
	}

	body, _ := ioutil.ReadAll(ipAnswer.Body)
	ip := string(body)

	//Define addressType
	addressType := "AAAA"
	if len(strings.Split(ip, "."))>1 {
		addressType = "A"
	}

	//Building request to Scaleway API
	var contentToSend []Connect
	var records []Records

	records = append(records, Records{Name:subdomain, Type:addressType, Priority:0, TTL:3600, Data:ip})
	contentToSend = append(contentToSend, Connect{Name: subdomain, ChangeType:"REPLACE", Type:addressType, Records:records})

	marshalled, _ := json.Marshal(contentToSend)
	fmt.Println(string(marshalled))

	readyToSend := strings.NewReader(string(marshalled))

	res2, _ := http.NewRequest("PATCH", "https://api.online.net/api/v1/domain/"+domain+"/version/active", readyToSend)
	res2.Header.Set("Authorization", "Bearer "+onlineAPI)

	client = &http.Client{Timeout: time.Second * 20}

	_, err = client.Do(res2)
	if err != nil {
		log.Fatal("Error while trying to PATCH online DNS !")
	}
}