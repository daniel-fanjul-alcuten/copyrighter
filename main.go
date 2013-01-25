package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Copyrighter struct {
	search  *regexp.Regexp
	replace string
}

func NewCopyrigher(corp, year string) *Copyrighter {
	search := regexp.MustCompile("(Copyright *(\\d+ *(,|-) *)?)(\\d+)( *" + corp + ")")
	replace := "${1}" + year + "${5}"
	return &Copyrighter{search, replace}
}

func (c *Copyrighter) Replace(source string) string {
	return c.search.ReplaceAllString(source, c.replace)
}

func main() {
	flag.Parse()
	args := flag.Args()

	year := time.Now().Year()
	c := NewCopyrigher("Spotify", strconv.Itoa(year))

	wait := make(chan bool, len(args))
	for _, arg := range args {
		filename := arg

		go func() {
			if data, err := ioutil.ReadFile(filename); err == nil {

				source := string(data)
				target := c.Replace(source)

				if source != target {
					tmp := filename + ".tmp"
					if err := ioutil.WriteFile(tmp, []byte(target), 0666); err != nil {
						log.Println(err)
					} else if err := os.Rename(tmp, filename); err != nil {
						log.Println(err)
					}
				}
			} else {
				log.Println(err)
			}

			wait <- true
		}()
	}

	for i := len(args); i > 0; i-- {
		<-wait
	}
}
