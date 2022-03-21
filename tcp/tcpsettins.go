package tcps

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var White = "\033[1;37m\033[0m"
var Red = "\033[31m"
var Wg sync.WaitGroup

type Tcpsettings struct {
	client *http.Client
}

func Httpinit() *Tcpsettings {

	//var s Tcpsettings
	s := &Tcpsettings{
		client: &http.Client{
			Timeout: time.Millisecond * 3000,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 100,

				// ExpectContinueTimeout: time.Millisecond * 800,
			}},
	}
	//s.client.Timeout = time.Millisecond * 1500
	return s
}

// func Httpinit() *Tcpsettings {
// 	readTimeout, _ := time.ParseDuration("50ms")
// 	writeTimeout, _ := time.ParseDuration("500ms")
// 	//timeout := time.Duration(500 * time.Millisecond)
// 	maxIdleConnDuration, _ := time.ParseDuration("10000ms")

// 	s := &Tcpsettings{
// 		client: fasthttp.Client{
// 			ReadTimeout:                   readTimeout,
// 			WriteTimeout:                  writeTimeout,
// 			MaxIdleConnDuration:           maxIdleConnDuration,
// 			NoDefaultUserAgentHeader:      true,  // Don't send: User-Agent: fasthttp
// 			DisableHeaderNamesNormalizing: false, // If you set the case on your headers correctly you can enable this
// 			DisablePathNormalizing:        true,
// 			MaxResponseBodySize:           1024,
// 			// increase DNS cache time to an hour instead of default minute

// 			Dial: (&fasthttp.TCPDialer{
// 				Concurrency:      4096,
// 				DNSCacheDuration: time.Hour * 1,
// 			}).Dial,
// 		},
// 	}
// 	return s

// }

// func Saldır(paths string, url string) {

// 	go Request(Httpinit(), paths, url)

// }

func Request(s *Tcpsettings, paths string, url string) {

	resp, err := s.client.Get(url + paths)

	if err != nil {
		log.Print(err)
	}

	defer resp.Body.Close()
	defer Wg.Done()

	if resp.StatusCode == 200 {
		// //resp, err := http.NewRequest("GET", url+path, nil)
		// defer resp.Body.Close()

		fmt.Printf(White+"Path :"+Red+"%s  "+White+"   Status Code :"+Red+" %s\n", paths, resp.Status)
	}
}
