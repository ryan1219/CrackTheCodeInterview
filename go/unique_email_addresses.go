package main

import "strings"

// question: https://leetcode.com/problems/unique-email-addresses/
func numUniqueEmails(emails []string) int {
	dic := make(map[string]map[string]int)
	count := 0

	for _, email := range emails {
		components := strings.Split(email, "@")
		localName := parseLocalName(components[0])
		domainName := components[1]
		if _, contained := dic[domainName]; contained {
			if _, contained = dic[domainName][localName]; !contained {
				dic[domainName][localName] = 1
				count++
			}
		} else {
			dic[domainName] = make(map[string]int)
			dic[domainName][localName] = 1
			count++
		}
	}

	return count
}

func parseLocalName(localName string) string {
	localName = strings.ReplaceAll(localName, ".", "")

	index := strings.Index(localName, "+")

	if index != -1 {
		return localName[:index]
	}

	return localName
}
