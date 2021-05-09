package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"net/http"
)

func main() {
	var g errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"https://www.baidu.com/",
		"http://www.google.com/",
	}
	for i := range urls {
		url := urls[i]
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			bodyC, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(url, string(bodyC))
			return err
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Printf("error %v", err)
		return
	}
	fmt.Println("Successfully fetched all URLs.")
}
