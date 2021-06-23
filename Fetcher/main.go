package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"http://www.cnn.com",
		"http://www.golang.org",
		"http://www.yahoo.com",
	}

	var wg sync.WaitGroup

	resc := make(chan string)
	errc := make(chan error)

	for _, url := range urls {
		wg.Add(1)
		go func(url string, wg *sync.WaitGroup) {
			defer wg.Done()
			body, err := fetch(url)
			if err != nil {
				errc <- err
				return
			}
			resc <- string(body)
		}(url, &wg)
	}

	fmt.Println("Before For Select")
	for i := 0; i < len(urls); i++ {
		select {
		case res := <-resc:
			fmt.Println(res[:50])
		case err := <-errc:
			fmt.Println(err)
		}
	}
	fmt.Println("Before Wait")
	wg.Wait()
	fmt.Println("After Wait")

}

func fetch(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	return string(body), nil
}
