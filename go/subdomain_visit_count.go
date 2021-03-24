package main

import (
	"strconv"
	"strings"
)

// question: https://leetcode.com/problems/subdomain-visit-count/
func subdomainVisits(cpdomains []string) []string {
	domains := make(map[string]int)
	res := make([]string, 0)

	for _, d := range cpdomains {
		count, _ := strconv.Atoi(strings.Split(d, " ")[0])
		address := strings.Split(d, " ")[1]
		components := strings.Split(address, ".")

		subdomains := components[len(components)-1]
		if val, contained := domains[subdomains]; contained {
			domains[subdomains] = val + count
		} else {
			domains[subdomains] = count
		}
		for i := len(components) - 2; i >= 0; i-- {
			subdomains = components[i] + "." + subdomains
			if val, contained := domains[subdomains]; contained {
				domains[subdomains] = val + count
			} else {
				domains[subdomains] = count
			}
		}
	}

	for k, v := range domains {
		res = append(res, strconv.Itoa(v)+" "+k)
	}

	return res
}
