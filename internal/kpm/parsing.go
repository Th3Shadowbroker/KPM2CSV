package kpm

import (
	"bufio"
	"os"
	"strings"
)

func ParseFile(filepath string) ([]Website, error) {
	var websites []Website

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var website = Website{}
	for scanner.Scan() {
		var key string
		var value string

		currentLine := scanner.Text()

		pair := strings.SplitN(currentLine, ":", 2)
		if len(pair) < 2 && currentLine != "---" {
			continue
		}

		key = strings.TrimSpace(pair[0])
		if key != "---" {
			value = strings.TrimSpace(pair[1])
		}

		switch key {

		case "Website name":
		case "Application":
			website.Name = value
			break

		case "Website URL":
			website.URL = value
			break

		case "Login name":
			website.LoginName = value
			break

		case "Login":
			website.Login = value
			break

		case "Password":
			website.Password = value
			break

		case "Comment":
			website.Comment = value
			break

		case "---":
			websites = append(websites, website)
			website = Website{}
			break
		}
	}

	return websites, nil
}
