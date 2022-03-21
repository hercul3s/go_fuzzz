package main

import (
	"fmt"
	"os"

	"time"

	tcps "test.com/fuzz/tcp"
	wordlist "test.com/fuzz/wordlist"
)

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

//########################################################################

func main() {
	if len(os.Args) == 1 {
		welcome()

	}
	// if strings.Contains(*hedefwebsite, "http") {
	// 	*hedefwebsite = *hedefwebsite + "/"
	// }
	start := time.Now()
	wordL := wordlist.NewWordlist()
	wordL.WordListinit()
	//TODO Paths değişkeninin ismini değiştir.

	duration := time.Since(start)
	fmt.Printf(tcps.White+"Executed Time :"+tcps.Red+" %d ms"+tcps.White, duration.Milliseconds())

}
