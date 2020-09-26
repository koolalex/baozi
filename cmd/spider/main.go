package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	//call_chromedp()

	//京东价格服务
	priceUrl := "http://p.3.cn/prices/mgets?type=1&skuIds=100013068434"

	c := colly.NewCollector()

	//colly 扩展增强
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("爬取中", r.URL)
	})

	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//	link := e.Attr("href")
	//
	//	// Print link
	//	fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	//
	//	// Visit link found on page
	//
	//	// Only those links are visited which are in AllowedDomains
	//	//c.Visit(e.Request.AbsoluteURL(link))
	//})

	c.OnHTML(".sku-name", func(e *colly.HTMLElement) {
		fmt.Println(strings.TrimSpace(e.Text))
	})

	c.OnResponse(func(resp *colly.Response) {
		//响应返回之后调用
		httpClient := http.Client{Timeout: 20 * time.Second}
		r, err := httpClient.Get(priceUrl)
		if err != nil {
			fmt.Println(err)
			return
		}

		respBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var dat []struct {
			P   string `json:"p"`
			Op  string `json:"op"`
			Cbf string `json:"cbf"`
			Id  string `json:"id"`
			M   string `json:"m"`
		}
		if err = json.Unmarshal(respBody, &dat); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(fmt.Sprintf("价格：%v", dat[0].P))
	})

	t := time.Now()
	if err := c.Visit("https://item.jd.com/100013068434.html"); err != nil {
		fmt.Println(err)
	}
	c.Wait()
	fmt.Printf("花费时间:%s", time.Since(t))
}

func call_chromedp() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.cnblogs.com/"),
		chromedp.WaitVisible(`#footer`, chromedp.ByID),
		chromedp.Nodes(`.//a[@class="titlelnk"]`, &nodes),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("get nodes:", len(nodes))
	// print titles
	for _, node := range nodes {
		fmt.Println(node.Children[0].NodeValue, ":", node.AttributeValue("href"))
	}
}
