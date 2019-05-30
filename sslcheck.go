package main

import (
	"fmt"
	"net/http"
	"flag"
	"encoding/json"
	"time"
	"log"
)

type SSLInfo struct {
	Domain string `json:"domain"`
	Issuer string `json:"issuer"`
	ExpireDate string	`json:"expire_date"`
	Remain int		`json:"days_remain"`
}

func main() {

	var target_domain = flag.String("domain", "", "check target domain")
	flag.Parse()

	target_url := fmt.Sprintf("https://%s/", *target_domain)
	resp, err := http.Get(target_url)
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	expireUTCTime := resp.TLS.PeerCertificates[0].NotAfter
	expireDate := expireUTCTime.Format("2006/01/02 15:04")
	day_remain := int(expireUTCTime.Sub(now).Hours()/24)
	issuer := fmt.Sprintf("%s",resp.TLS.PeerCertificates[0].Issuer)
	info := SSLInfo {
		Domain: *target_domain,
		ExpireDate: expireDate,
		Remain: day_remain, 
		Issuer: issuer }
	jsonByte, _ := json.Marshal(info)

	fmt.Println(string(jsonByte))

}
