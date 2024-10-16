package main

import (
	"fmt"
	"sort"
)

func (cfg *config) printReport(baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	sortedPages := cfg.sortPagesByEmbeddedLinks()

	for k, v := range sortedPages {
		fmt.Printf("Found %v internal links to %s\n", v, k)
	}
}

func (cfg *config) sortPagesByEmbeddedLinks() map[string]int {
	var sortedPages = map[string]int{}

	keys := make([]string, 0, len(cfg.pages))
	for k := range cfg.pages {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return cfg.pages[keys[i]] < cfg.pages[keys[j]]
	})

	for _, k := range keys {
		sortedPages[k] = cfg.pages[k]
	}

	return sortedPages
}
