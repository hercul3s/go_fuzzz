package wordlist

import (
	"flag"
	"fmt"
	"net/url"
	"strings"

	"test.com/fuzz/file"
	tcps "test.com/fuzz/tcp"
)

type Wordlist struct {
	Data      []string
	Filenames []string
}

var wL2 Wordlist
var deneme file.File

func NewWordlist() *Wordlist {
	//server := &Server{Port: port, Connections: make(map[string]net.Conn)}
	//Wordlist := &Wordlist{Data:data[], Rfiles: make([]string, 0)}
	wL2 = Wordlist{Data: make([]string, 0), Filenames: make([]string, 0)}
	return &wL2
}

func (*Wordlist) WordListinit() {

	multiWLdir := flag.String("dir", "", " > -dir ./wordlists/")
	worldlist := flag.String("w", "", " > -w ./wordlists/wordlist.txt")
	hedefwebsite := flag.String("url", "", " > -url http://www.example.com/")
	// WordListGithub := flag.String("github", "", "a string")

	flag.Parse()
	url.Parse(*hedefwebsite)

	if *multiWLdir != "" {

		wL2.Filenames = deneme.ReadDirectory(*multiWLdir)

		//wL2 = NewWordlist(*worldlist)
		//okunacakdosya := deneme.ReadDirectory(*multiWLdir)
		for _, paths := range wL2.Filenames {

			wL2.Data = append(wL2.Data, deneme.ReadFiles(paths)...)

			//TODO Bakılacak scanFile

		}
		fmt.Printf(tcps.Red+"Kullanılan Dosya: %s\n"+tcps.White, wL2.Filenames)
	} else {

		//wL2 = wordlist.NewWordlist(*worldlist)
		wL2.Data = append(wL2.Data, deneme.ReadFiles(*worldlist)...)
	}
	if !strings.HasSuffix(*hedefwebsite, "/") {
		*hedefwebsite = *hedefwebsite + "/"
	}
	fmt.Printf("Bulunan path sayısı : %d\n", len(wL2.Data))
	http := tcps.Httpinit()
	for i := 0; i < len(wL2.Data); i++ {

		tcps.Wg.Add(1)

		go tcps.Request(http, (wL2.Data[i]), *hedefwebsite)
		//go tcps.Saldır(string(test[i]), *hedefwebsite)

	}
	tcps.Wg.Wait()

}
