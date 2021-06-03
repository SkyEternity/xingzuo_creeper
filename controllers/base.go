package controllers

import (
	"log"
	"net/http"
)

//http
func RequestFn(url string) (err error, resp *http.Response) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	req.Header.Add("Cookie", "Hm_lvt_d96f3af1d84a28ce7598db6f236958bc=1621837279,1622097792,1622099454,1622441576; Hm_lpvt_d96f3af1d84a28ce7598db6f236958bc=1622444901")
	req.Header.Add("Referer", "https://www.d1xz.net")

	resp, err = client.Do(req)
	req.Close = true
	if err != nil {
		log.Fatal(err)
	}
	return
}
