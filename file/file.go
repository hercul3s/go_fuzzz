package file

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	//wordlist "test.com/fuzz/wordlist"
)

//var wL2 *wordlist.Wordlist

type File struct {
	Data      []string
	Filenames []string
}

func (w *File) ReadDirectory(path string) []string {

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

func (w *File) ReadFiles(logfile string) []string {
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
