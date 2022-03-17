package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var White = "\033[1;37m\033[0m"
var Red = "\033[31m"

var wg sync.WaitGroup
var testMutex sync.Mutex

// type Ayarlar struct {

// location ,err := os

// }

// func AyarlarıOku()(*Ayarlar, error){

// }

// type DefaultAyarlar struct{

// 	Filter FiltreAyarları

// }

// func ParseArguments()

type Wordlist struct {
	data      [][]byte
	position  int
	fuzzInurl string
}

// func (w *Wordlist) readFile(path string) error {
// 	var file *os.File
// 	var err error

// 	file, err = os.Open(path)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
// 	var data [][]byte

// 	reader := bufio.NewScanner(file)

// }

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

// func Dosyaoku(path string) []byte {

// 	data, err := ioutil.ReadFile(path)
// 	check(err)
// 	fmt.Print(string(data))

// 	return data
// }

func exampleReadDir(path string) string {
	var xx string
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		log.Panicf("failed reading directory: %s", err)
	}
	for _, entry := range entries {

		//fileExtension := filepath.Ext(entry.Name())
		//fmt.Println(fileExtension) // /uzantıları yazdır.
		xx = entry.Name()
	}
	fmt.Printf("\nNumber of files in current directory: %d\n", len(entries))
	//fmt.Printf("\nError: %v\n", err)
	return xx
}

func scanFile(logfile string) []string {
	abc := []string{}
	f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return nil
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		_ = sc.Text() // GET the line string
		//fmt.Println(sc.Text())
		abc = append(abc, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return nil
	}
	return abc
}

func Request(paths string, url string) {

	defer wg.Done()

	//for _, path := range paths {

	resp, err := http.Get(url + paths)

	//resp, err := http.NewRequest("GET", url+path, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	//defer resp.Body.Close()
	//time.Sleep(time.Millisecond)
	//fmt.Println(resp.Request.TLS.PeerCertificates[0].Signature)

	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(body))

	//}
	//wg.Wait()

}

func main() {

	// start := time.Now()

	wordListDir := flag.String("dir", "", "a string")
	hedefwebsite := flag.String("url", "", "a string")

	// WordListGithub := flag.String("github", "", "a string")
	flag.Parse()

	okunacakdosya := exampleReadDir(*wordListDir)

	abcd := scanFile(okunacakdosya)

	fmt.Printf("Kullanılan Dosya: %s\n", okunacakdosya)

	start := time.Now()

	if !strings.HasSuffix(*hedefwebsite, "/") {
		*hedefwebsite = *hedefwebsite + "/"
	}

	for _, path := range abcd {
		wg.Add(1)
		go Request( /*[]*/ path, *hedefwebsite)
	}
	wg.Wait()

	//go Request(abcd, *hedefwebsite)

	duration := time.Since(start)

	fmt.Printf("Executed Time : %d ms", duration.Milliseconds())

	// for _, website := range abcd {

	// 	respons, err := http.Get(website)
	// 	fmt.Println(respons.Status) // txt

	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}

	// 	defer respons.Body.Close()
	// }

	// WordListGithub := flag.String("github", "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Discovery/DNS/subdomains-top1million-10-million.txt", "a string")
	// dirWordlist, err := os.ReadFile(*wordListDir)

	// respons, err := http.Get("https://www.google.com/abcd")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// body, err := ioutil.ReadAll(respons.Body)
	// if err != nil {

	// 	log.Fatalln(err)
	// }

	// statuscode := respons.Status
	// if err != nil {

	// 	log.Fatalln(err)
	// }

	// strbody := string(statuscode)
	// duration := time.Since(start)

	// log.Printf("Verilen dizin %s\n", *wordListDir)

	// log.Printf("Verilen github Repo %s\n", *WordListGithub)
	// log.Printf(Red+"Executed %v ms : Status Code = %v\n"+White, duration.Milliseconds(), statuscode)

}
