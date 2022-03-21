package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"time"

	tcps "test.com/fuzz/tcp"
	wordlist "test.com/fuzz/wordlist"
)

var wL2 *wordlist.Wordlist

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

	if *multiWLdir != "" {
		// if !strings.HasSuffix(*multiWLdir, "/") {
		// 	*multiWLdir = *multiWLdir + "/"
		// }
		wL2 = wordlist.NewWordlist(*worldlist)
		okunacakdosya := wL2.ReadDirectory(*multiWLdir)

		for _, paths := range okunacakdosya {

			wL2.Data = append(wL2.Data, scanFile(paths)...)

			//TODO Bakılacak scanFile

		}
		fmt.Printf(tcps.Red+"Kullanılan Dosya: %s\n"+tcps.White, okunacakdosya)
	} else {

		wL2 = wordlist.NewWordlist(*worldlist)
		wL2.Data = append(wL2.Data, scanFile(*worldlist)...)
	}

	if !strings.HasSuffix(*hedefwebsite, "/") {
		*hedefwebsite = *hedefwebsite + "/"
	}

	// if strings.Contains(*hedefwebsite, "http") {
	// 	*hedefwebsite = *hedefwebsite + "/"
	// }

	fmt.Printf("Bulunan path sayısı : %d\n", len(wL2.Data))

	//TODO Paths değişkeninin ismini değiştir.
	start := time.Now()
	http := tcps.Httpinit()

	for i := 0; i < len(wL2.Data); i++ {

		tcps.Wg.Add(1)

		go tcps.Request(http, (wL2.Data[i]), *hedefwebsite)
		//go tcps.Saldır(string(test[i]), *hedefwebsite)

	}
	tcps.Wg.Wait()

	duration := time.Since(start)

	fmt.Printf(tcps.White+"Executed Time :"+tcps.Red+" %d ms"+tcps.White, duration.Milliseconds())

}
