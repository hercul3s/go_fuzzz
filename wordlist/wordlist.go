package wordlist

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Wordlist struct {
	Data      []string
	Filenames []string
}

var wL2 Wordlist

func (w *Wordlist) ReadDirectory(path string) []string {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Panicf("failed reading directory: %s", err)
	}
	for _, entry := range files {

		if strings.HasSuffix(entry.Name(), ".txt") {

			w.Filenames = append(w.Filenames, path+entry.Name())
		}

	}
	//fileExtension := filepath.Ext(entry.Name())
	//fmt.Println(fileExtension) // /uzantıları yazdır.
	fmt.Printf("\nNumber of files in current directory: %d\n", len(w.Filenames))

	return w.Filenames
}

func NewWordlist(data string) *Wordlist {
	//server := &Server{Port: port, Connections: make(map[string]net.Conn)}
	//Wordlist := &Wordlist{Data:data[], Rfiles: make([]string, 0)}
	Wordlist := &Wordlist{Data: make([]string, 0), Filenames: make([]string, 0)}
	return Wordlist
}
