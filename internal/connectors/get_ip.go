package connectors

import (
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"time"
)

func GetIPAddress() (ip string, addressType string, err error) {
	//Get IP on ifconfig.me
	doConfigCall, _ := http.NewRequest("GET", "https://ifconfig.me", nil)
	client := &http.Client{Timeout: time.Second * 20}

	ipAnswer, err := client.Do(doConfigCall)
	if err != nil {
		logrus.Warnf("An error occurred while fetching IP from ifconfig.me: %s.", err.Error())
		return
	}
	if ipAnswer.StatusCode != 200 {
		logrus.Warnf("ifconfig.me returned invalid status code: %d.", ipAnswer.StatusCode)
		return "", "", errors.New("invalid status code")
	}

	body, _ := io.ReadAll(ipAnswer.Body)
	ip = string(body)

	//Define addressType
	addressType = "AAAA"
	if len(strings.Split(ip, ".")) > 1 {
		addressType = "A"
	}

	logrus.Infof("Detected IP for your home is %s (which will create and/or update %s records)", ip, addressType)

	return
}
