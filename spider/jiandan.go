package spider

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func SpiderJan() {
	var wg sync.WaitGroup
	var data []*Article
	k := 28
	// queueChan := make(chan struct{}, 28)
	// for k := 1; k <= 27; k++ {
	wg.Add(1)
	// queueChan <- struct{}{}
	go func(n int) {
		fmt.Println("====", n)
		defer func() {
			wg.Done()
			// <-queueChan
		}()
		url := fmt.Sprintf("http://jandan.net/p/tag/故事/page/%d", n)
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(url)
			log.Fatal(err)
		}

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

		// 筛选class为list-post的元素
		dom.Find(".list-post").Each(func(i int, selection *goquery.Selection) {
			selection.Find(".indexs").Each(func(i int, selection *goquery.Selection) {
				title := selection.Find("a").Text()
				link, ok := selection.Find("a").Attr("href")
				if ok {
					article := &Article{
						Title: title,
						Link:  link,
					}

					// go appendFile(n, article)
					time.Sleep(10 * time.Second)
					resInfo, err := http.Get(link)
					if err != nil {
						log.Fatal(err)
					}
					defer resInfo.Body.Close()

					if resInfo.StatusCode != 200 {
						fmt.Println(link)
						log.Fatalf("status code error: %d %s", resInfo.StatusCode, resInfo.Status)
					}

					domInfo, err := goquery.NewDocumentFromReader(resInfo.Body)
					if err != nil {
						fmt.Println(link)
						log.Fatalln(err)
					}

					domInfo.Find(".post").Each(func(i int, selection *goquery.Selection) {
						created := selection.Find(".time_s").Text()

						if len(created) > 0 {

							var content string
							selection.Find("p").Each(func(i int, selection *goquery.Selection) {
								tmp := selection.Text()
								content += fmt.Sprintf("%s\r\n", tmp)
							})
							article.Created = created
							article.Content = content
						}
					})
					fmt.Println("====", article.Title)
					//存入文件
					go WriteDan(n, article)

					data = append(data, article)
				}
			})
		})
	}(k)
	// time.Sleep(20 * time.Second)
	// }

	fmt.Println(len(data))
	wg.Wait()
}

func Spider() {
	var wg sync.WaitGroup
	for i := 1; i <= 28; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			res, err := http.Get(fmt.Sprintf("http://jandan.net/p/tag/故事/page/%d", i))
			if err != nil {
				log.Fatal(err)
			}

			defer res.Body.Close()

			if res.StatusCode != 200 {
				log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
			}

			//将html生成goquery的Document
			dom, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				log.Fatalln(err)
			}

			var data []*Article
			// 筛选class为list-post的元素
			dom.Find(".list-post").Each(func(i int, selection *goquery.Selection) {
				selection.Find(".indexs").Each(func(i int, selection *goquery.Selection) {
					title := selection.Find("a").Text()
					link, ok := selection.Find("a").Attr("href")
					if ok {
						article := &Article{
							Title: title,
							Link:  link,
						}
						//获取详情
						resInfo, err := http.Get(link)
						if err != nil {
							log.Fatal(err)
						}
						defer resInfo.Body.Close()

						if resInfo.StatusCode != 200 {
							log.Fatalf("status code error: %d %s", resInfo.StatusCode, resInfo.Status)
						}

						domInfo, err := goquery.NewDocumentFromReader(resInfo.Body)
						if err != nil {
							log.Fatalln(err)
						}

						domInfo.Find(".post").Each(func(i int, selection *goquery.Selection) {
							created := selection.Find(".time_s").Text()

							if len(created) > 0 {

								var content string
								selection.Find("p").Each(func(i int, selection *goquery.Selection) {
									tmp := selection.Text()
									content += fmt.Sprintf("%s\r\n", tmp)
								})
								article.Created = created
								article.Content = content
							}
						})
						//存入文件
						go WriteDan(i, article)

						data = append(data, article)
					}
				})
			})
			fmt.Println(len(data))
		}()
	}

	wg.Wait()
}
