package wordlist

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	"test.com/fuzz/file"
	tcps "test.com/fuzz/tcp"
)

var Fileopr file.File

type Wordlist struct {
	Data      []string
	FilePaths []string
	//TestingFile *file.File
}

var wL2 Wordlist

func NewWordlist() *Wordlist {
	//server := &Server{Port: port, Connections: make(map[string]net.Conn)}
	//Wordlist := &Wordlist{Data:data[], Rfiles: make([]string, 0)}
	wL2 = Wordlist{Data: make([]string, 0), FilePaths: make([]string, 0)} //TestingFile: &file.File{}}
	return &wL2
}

func (*Wordlist) WordListinit() {

	multiWLdir := flag.String("dir", "", " > -dir ./wordlists/")
	worldlist := flag.String("w", "", " > -w ./wordlists/wordlist.txt")
	hedefwebsite := flag.String("url", "", " > -url http://www.example.com/")
	//WordListGithub := flag.String("github", "", "a string")
	WordlistGithub := flag.String("github", "", " > -github https://x.x.x.x/wordlist.txt")

	flag.Parse()
	url.Parse(*hedefwebsite)

	if *multiWLdir != "" {

		wL2.FilePaths = Fileopr.ReadDirectory(*multiWLdir)

		//wL2 = NewWordlist(*worldlist)
		//okunacakdosya := deneme.ReadDirectory(*multiWLdir)
		for _, paths := range wL2.FilePaths {

			wL2.Data = append(wL2.Data, Fileopr.ReadFiles(paths)...)
		}

		//TODO Bakılacak scanFile

		fmt.Printf(tcps.Red+"Kullanılan Dosya: %s\n"+tcps.White, wL2.FilePaths)
	} else if *WordlistGithub != "" {

		keys := tcps.SimpleRequest(*WordlistGithub)
		wL2.Data = append(wL2.Data, keys...)
		//wL2.Data = append(wL2.Data, Fileopr.ReadFiles(*WordlistGithub)...)

		fmt.Fprintf(os.Stderr, tcps.Red+"Github Wordlist: %s\n"+tcps.White, *WordlistGithub)

	} else {
		//wL2 = wordlist.NewWordlist(*worldlist)
		wL2.Data = append(wL2.Data, Fileopr.ReadFiles(*worldlist)...)

	}

	if !strings.HasSuffix(*hedefwebsite, "/") && *WordlistGithub == "" {
		*hedefwebsite = *hedefwebsite + "/"
	}
	fmt.Printf("Bulunan path sayısı : %d\n", len(wL2.Data))
	fmt.Printf("Url : %s\n", *hedefwebsite)
	http := tcps.Httpinit()

	for i := 0; i < len(wL2.Data); i++ {

		tcps.Wg.Add(1)

		go tcps.Request(http, wL2.Data[i], *hedefwebsite)
		//go tcps.Saldır(string(test[i]), *hedefwebsite)

	}
	tcps.Wg.Wait()

}
