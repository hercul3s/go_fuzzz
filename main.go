package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"

	"time"

	tcps "test.com/fuzz/tcp"
)

//var testMutex sync.Mutex

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

func welcome() {

	fmt.Printf("#############" + tcps.Red + " Welcome to the Fuzzer" + tcps.White + " #############\n")
	fmt.Printf("# " + tcps.Red + "Available Options 	" + tcps.White + "		        #\n")
	fmt.Printf("#" + tcps.Red + " -dir : Directory of wordlists      " + tcps.White + "           #\n")
	fmt.Printf("#" + tcps.Red + " -w : Wordlist              " + tcps.White + "                   #\n")
	fmt.Printf("#" + tcps.Red + " -url : Target URL       " + tcps.White + "                      #\n")
	//fmt.Printf("#" + Red + " -github : Wordlist from Github  " + White + "              #\n")
	fmt.Printf("#" + tcps.Red + " -help : Help                   " + tcps.White + "               #\n")
	fmt.Printf("####################################################")
	os.Exit(0)

}

func readDirectory(path string) []string {
	var rfiles []string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Panicf("failed reading directory: %s", err)
	}
	for _, entry := range files {

		if strings.HasSuffix(entry.Name(), ".txt") {

			rfiles = append(rfiles, path+entry.Name())
		}

	}
	//fileExtension := filepath.Ext(entry.Name())
	//fmt.Println(fileExtension) // /uzantıları yazdır.
	fmt.Printf("\nNumber of files in current directory: %d\n", len(rfiles))

	return rfiles
}

func scanFile(logfile string) []string {
	words := []string{}
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
		if sc.Text() != "" {
			words = append(words, sc.Text())
		}
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return nil
	}
	return words
}

//########################################################################

func main() {
	if len(os.Args) == 1 {
		welcome()

	}

	multiWLdir := flag.String("dir", "", " > -dir ./wordlists/")
	worldlist := flag.String("w", "", " > -w ./wordlists/wordlist.txt")
	hedefwebsite := flag.String("url", "", " > -url http://www.example.com/")
	// WordListGithub := flag.String("github", "", "a string")

	flag.Parse()
	url.Parse(*hedefwebsite)
	var test []string
	var okunacakdosya []string

	//var test []string

	if *multiWLdir != "" {
		//test = dizinoku(*multiwordLists)
		if !strings.HasSuffix(*multiWLdir, "/") {
			*multiWLdir = *multiWLdir + "/"
		}
		okunacakdosya = readDirectory(*multiWLdir)

		for _, paths := range okunacakdosya {

			test = append(test, scanFile(paths)...)
			//TODO Bakılacak scanFile

		}

	} else {
		//test = []string{*worldlist}
		test = append(test, scanFile(*worldlist)...)
	}
	//okunacakdosya := dizinoku(*multiwordLists)

	fmt.Printf(tcps.Red+"Kullanılan Dosya: %s\n"+tcps.White, okunacakdosya)

	if !strings.HasSuffix(*hedefwebsite, "/") {
		*hedefwebsite = *hedefwebsite + "/"
	}

	// if strings.Contains(*hedefwebsite, "http") {
	// 	*hedefwebsite = *hedefwebsite + "/"
	// }

	fmt.Printf("Bulunan path sayısı : %d\n", len(test))

	//TODO Paths değişkeninin ismini değiştir.
	start := time.Now()
	x := tcps.Httpinit()

	for i := 0; i < len(test); i++ {

		tcps.Wg.Add(1)

		go tcps.Request(x, (test[i]), *hedefwebsite)
		//go tcps.Saldır(string(test[i]), *hedefwebsite)

	}
	tcps.Wg.Wait()

	duration := time.Since(start)

	fmt.Printf(tcps.White+"Executed Time :"+tcps.Red+" %d ms"+tcps.White, duration.Milliseconds())

}
