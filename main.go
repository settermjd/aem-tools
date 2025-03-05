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
	aemLinkTemplate := os.Getenv("AEM_LINK_TEMPLATE")
	linksFile := os.Getenv("LINKS_FILE")

	f, err := os.OpenFile(linksFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("could not open links file (%s); error: %v", linksFile, err)
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

		link, err := createAEMEditLink(aemLinkTemplate, line)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(link)
	}
}

// createAEMEditLink creates an AEM edit link using the provided link template and URL path
func createAEMEditLink(linkTemplate string, urlPath string) (string, error) {
	// This matches a path in a URL/string ending with /blog/*
	regex := `(?m)(?:\/blog\/)([\w$-_.+!*'(),;?:@=&]*)$`

	urlDetails, err := url.Parse(strings.TrimSuffix(urlPath, "\n"))
	if err != nil {
		return "", fmt.Errorf("could not extract the sub-path from the provided path (%s), because: %s", urlPath, err)
	}

	path := urlDetails.Path
	var re = regexp.MustCompile(regex)
	if re.Match([]byte(path)) {
		if matches := re.FindStringSubmatch(path); matches != nil {
			return fmt.Sprintf(linkTemplate, matches[1]), nil
		}
	}

	return "", fmt.Errorf("could not create AEM edit link with the provided URL path (%s)", urlPath)
}
