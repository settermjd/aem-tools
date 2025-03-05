package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	aemEditLink := os.Getenv("AEM_EDIT_LINK")
	linksFile := os.Getenv("LINKS_FILE")

	f, err := os.OpenFile(linksFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return
		}

		urlDetails, err := url.Parse(strings.TrimSuffix(line, "\n"))
		if err != nil {
			log.Fatal(fmt.Errorf("could not parse link: %s. reason: %s", line, err))
		}

		path := urlDetails.Path
		var re = regexp.MustCompile(`(?m)(?:\/blog\/)([\w$-_.+!*'(),;?:@=&]*)$`)
		if re.Match([]byte(path)) {
			matches := re.FindStringSubmatch(path)
			if matches != nil {
				fmt.Printf(aemEditLink+"\n", matches[1])
			}
		}
	}
}
