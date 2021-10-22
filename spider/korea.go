package spider

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const douURL = "https://warehouse13.tistory.com/?page="
const limit = 25

func SpiderDou() {
	var wg sync.WaitGroup
	var data []*Article
	page := 1

	// queueChan := make(chan struct{}, 28)
	for k := 0; k <= page; k++ {
		wg.Add(1)
		// queueChan <- struct{}{}
		go func(n int) {
			fmt.Println("====", n)
			defer func() {
				wg.Done()
				// <-queueChan
			}()
			url := fmt.Sprintf("%s%d", douURL, n)
			req, _ := http.NewRequest("Get", url, nil)
			// cookie1 := &http.Cookie{}

			// req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
			// req.Header.Add("Host", "www.douban.com")
			// req.Header.Add("Referer", "https://www.douban.com/group/")
			// req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
			// req.Header.Add("Upgrade-Insecure-Requests", "1")
			req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.54 Safari/537.36")
			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				fmt.Println(url)
				log.Fatal(err)
			}

			// res, err := http.Get(url)
			// if err != nil {
			// 	fmt.Println(url)
			// 	log.Fatal(err)
			// }

			defer res.Body.Close()

			if res.StatusCode != 200 {
				fmt.Println(url)
				log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
			}

			//将html生成goquery的Document
			dom, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				fmt.Println(url)
				log.Fatalln(err)
			}

			// 筛选class为grid-16-8的元素
			dom.Find("#content").Each(func(i int, selection *goquery.Selection) {
				selection.Find(".post-item").Each(func(i int, trSelection *goquery.Selection) {

					title := trSelection.Find(".title").Text()
					link, ok := trSelection.Find("a").Attr("href")
					if ok {
						article := &Article{
							Title: title,
							Link:  link,
						}

						fmt.Println("====", article.Title)
						// go appendFile(n, article)
						time.Sleep(10 * time.Second)
						// resInfo, err := http.Get(link)
						// if err != nil {
						// 	log.Fatal(err)
						// }
						// defer resInfo.Body.Close()

						// if resInfo.StatusCode != 200 {
						// 	fmt.Println(link)
						// 	log.Fatalf("status code error: %d %s", resInfo.StatusCode, resInfo.Status)
						// }

						// domInfo, err := goquery.NewDocumentFromReader(resInfo.Body)
						// if err != nil {
						// 	fmt.Println(link)
						// 	log.Fatalln(err)
						// }

						// domInfo.Find(".article").Each(func(i int, selection *goquery.Selection) {
						// 	title := selection.Find("h1").Text()
						// 	created := selection.Find(".create-time").Text()

						// 	var content string
						// 	selection.Find("p").Each(func(i int, selection *goquery.Selection) {
						// 		tmp := selection.Text()
						// 		content += fmt.Sprintf("%s\r\n", tmp)
						// 	})
						// 	article.Created = created
						// 	article.Content = content
						// 	article.Title = title

						// })
						// fmt.Println("====", article.Title)
						// //存入文件
						// go WriteDou(n, article)

						// data = append(data, article)
					}

				})
			})
		}(k)
		time.Sleep(20 * time.Second)
	}

	fmt.Println(len(data))
	wg.Wait()
}
