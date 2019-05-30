package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type SSLInfo struct {
	Domain     string `json:"domain"`
	Issuer     string `json:"issuer"`
	ExpireDate string `json:"expire_date"`
	Remain     int    `json:"days_remain"`
}

func checkSSL(domain string) (result string, err error) {

	target_url := fmt.Sprintf("https://%s/", domain)

	resp, err := http.Get(target_url)
	if err != nil {
		return "", err
	}

	now := time.Now()
	expireUTCTime := resp.TLS.PeerCertificates[0].NotAfter
	expireDate := expireUTCTime.Format("2006/01/02 15:04")
	day_remain := int(expireUTCTime.Sub(now).Hours() / 24)
	issuer := fmt.Sprintf("%s", resp.TLS.PeerCertificates[0].Issuer)
	info := SSLInfo{
		Domain:     domain,
		ExpireDate: expireDate,
		Remain:     day_remain,
		Issuer:     issuer}
	jsonByte, err := json.Marshal(info)
	if err != nil {
		return "", err
	}

	return string(jsonByte), nil
}

func main() {

	var target_domain = flag.String("domain", "", "check target domain")
	flag.Parse()

	json, err := checkSSL(*target_domain)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(json)

}
