package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`i=dick&from=AUTO&to=AUTO&smartresult=dict&client=fanyideskweb&salt=16519927951637&sign=d9be0e9a4fd7e0197a9945488b1e934d&lts=1651992795163&bv=b2402fb6d0c7c6d318c684a709d8d03d&doctype=json&version=2.1&keyfrom=fanyi.web&action=FY_BY_REALTlME`)
	req, err := http.NewRequest("POST", "https://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", `_ga=GA1.2.1968799229.1643000798; OUTFOX_SEARCH_USER_ID_NCOO=143297926.5785821; OUTFOX_SEARCH_USER_ID="-321987288@10.169.0.102"; JSESSIONID=abcf20_JDZiXPhomS2Ccy; fanyi-ad-id=305838; fanyi-ad-closed=0; ___rl__test__cookies=1651992795158`)
	req.Header.Set("Origin", "https://fanyi.youdao.com")
	req.Header.Set("Referer", "https://fanyi.youdao.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edg/101.0.1210.32")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Microsoft Edge";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
