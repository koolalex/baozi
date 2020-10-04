package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func GetGoodPrice(url string) string {
	re := regexp.MustCompile(`com/(.*?).html`)
	keynum := re.FindAllStringSubmatch(url, -1)
	keynum0 := keynum[0][1]
	key0 := strings.Split(keynum0, "/")[0]
	key1 := strings.Split(keynum0, "/")[1]
	priceurl := "http://pas.suning.com/nspcsale_0_000000000" + key1 + "_000000000" + key1 + "_" + key0 + "_20_021_0210101_500353_1000267_9264_12113_Z001___R9006849_3.3_1___000278188__.html?callback=pcData&_=1558663936729"
	if len(key1) == 11 {
		priceurl = "http://pas.suning.com/nspcsale_0_0000000" + key1 + "_0000000" + key1 + "_" + key0 + "_20_021_0210101_500353_1000267_9264_12113_Z001___R9006849_3.3_1___000278188__.html?callback=pcData&_=1558663936729"
	}

	resp, err := http.Get(priceurl)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("err")
	}
	s, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	re0 := regexp.MustCompile(`"netPrice":"(.*?)","warrantyList`)
	price := re0.FindAllStringSubmatch(string(s), -1)
	fmt.Println(price)
	// fmt.Println(priceurl)
	return price[0][1]
}
