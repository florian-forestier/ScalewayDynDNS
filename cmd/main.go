package main

import (
	"github.com/Artheriom/ScalewayDynDNS/internal/connectors"
	"github.com/Artheriom/ScalewayDynDNS/internal/helpers"
	"github.com/sirupsen/logrus"
)

func main() {

	ip, kind, err := connectors.GetIPAddress()
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	for k, v := range helpers.Domains {
		_ = connectors.UpdateScalewayDomain(k, v, ip, kind)
	}
}
