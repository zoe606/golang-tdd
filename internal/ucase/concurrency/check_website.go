package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChan := make(chan result)
	//var results []bool

	for _, url := range urls {
		go func(u string) {
			//sering terjadi race condition
			results[u] = wc(u)

			// Send statement
			// ngirim data result ke channel
			resultChan <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		//time.Sleep(2 * time.Second)
		fmt.Println("menunggu")

		// Receive expression
		// nerima data dari channel
		r := <-resultChan
		results[r.string] = r.bool
	}

	//time.Sleep(2 * time.Second)
	fmt.Println(results)
	return results
}

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a string, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

/*
1.Well, a chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool
2.Notice how we have to use make when creating a channel; rather than say var ch chan struct{}. When you use var the variable will be initialised with the "zero" value of the type. So for string it is "", int it is 0, etc.
For channels the zero value is nil and if you try and send to it with <- it will block forever because you cannot send to nil channels
*/
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
