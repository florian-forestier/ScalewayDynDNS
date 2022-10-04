package helpers

import (
	"flag"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strings"
)

var Domains = map[string][]string{}
var Key = ""

func init() {
	logrus.Infof("Acquiring configuration...")
	var onlineDomain string

	flag.StringVar(&onlineDomain, "domain", "", "Define the domain or subdomain to update (eg github.com or example.github.com)")
	flag.StringVar(&Key, "key", "", "Set the key for your Online API access")
	flag.Parse()

	if onlineDomain == "" {
		// Try to get variable through environment
		onlineDomain = os.Getenv("SCALEWAY_DYNDNS_DOMAIN")
		if onlineDomain == "" {
			log.Fatal("No domain given, abort")
		}
	}
	if Key == "" {
		// Try to get variable through environment
		Key = os.Getenv("SCALEWAY_API_KEY")
		if Key == "" {
			log.Fatal("No API key given, abort")
		}
	}

	// Explode domains through ";"
	list := strings.Split(onlineDomain, ";")
	for _, k := range list {
		exploded := strings.Split(k, ".")
		domain := exploded[len(exploded)-2] + "." + exploded[len(exploded)-1]
		Domains[domain] = append(Domains[domain], k+".")
	}

	logrus.Infof("Configuration loaded.")
}
