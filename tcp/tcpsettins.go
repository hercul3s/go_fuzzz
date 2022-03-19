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
	client http.Client
}

func Httpinit() *Tcpsettings {

	var s Tcpsettings
	s.client.Timeout = time.Millisecond * 1500
	return &s
}

func Saldır(paths string, url string) {

	go Request(Httpinit(), paths, url)

}

func Request(s *Tcpsettings, paths string, url string) {

	defer Wg.Done()

	// eskisi
	//#############################################
	// client := &http.Client{
	// 	Timeout: time.Millisecond * 1500,
	// }
	resp, err := s.client.Get(url + paths)
	// resp, err := client.Get(url + paths)

	//resp, err := http.NewRequest("GET", url+path, nil)
	if err != nil {
		log.Println(err)

	}
	defer resp.Body.Close()

	//resp.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36")

	fmt.Printf(White+"Path :"+Red+"%s  "+White+"   Status Code :"+Red+" %s\n", paths, resp.Status)

}
