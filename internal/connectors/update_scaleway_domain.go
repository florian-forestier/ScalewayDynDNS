package connectors

import (
	"encoding/json"
	"github.com/Artheriom/ScalewayDynDNS/internal/helpers"
	"github.com/Artheriom/ScalewayDynDNS/internal/structures"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func UpdateScalewayDomain(domainName string, subdomains []string, newIp string, ipKind string) (err error) {

	logrus.Infof("Updating domain %s and its subdomains.", domainName)

	//Building request to Scaleway API
	var contentToSend []structures.Connect
	var records []structures.Records

	for _, k := range subdomains {
		records = append(records, structures.Records{Name: k, Type: ipKind, Priority: 0, TTL: 3600, Data: newIp})
		contentToSend = append(contentToSend, structures.Connect{Name: k, ChangeType: "REPLACE", Type: ipKind, Records: records})
		records = []structures.Records{}
	}

	marshalled, _ := json.Marshal(contentToSend)

	readyToSend := strings.NewReader(string(marshalled))
	res2, _ := http.NewRequest("PATCH", "https://api.online.net/api/v1/domain/"+domainName+"/version/active", readyToSend)
	res2.Header.Set("Authorization", "Bearer "+helpers.Key)

	client := &http.Client{Timeout: time.Second * 20}

	resp, err := client.Do(res2)
	if err != nil {
		logrus.Warnf("Error while trying to PATCH online DNS. Error was: %s", err.Error())
	} else if resp.StatusCode != 204 {
		logrus.Warnf("Server responded %d on PATCH for domain %s.", resp.StatusCode, domainName)
	} else {
		logrus.Infof("%s and its subdomains have been updated to new IP %s.", domainName, newIp)
	}

	return
}
