package scrape

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"mole/utils"
)

// EndPoint uses colly to scrape URLs from a given domain.
func EndPoint(domain *string, agent *string, headers *string, proxies *string, threads *int, depth *int, timeout *int) {
	hostURL, err := url.Parse(*domain)
	if err != nil {
		log.Fatalf("invalid domain: %v", err)
	}
	baseDomain := hostURL.Hostname()

	var customHeaders []string
	if *headers != "" {
		customHeaders, err = utils.ReadFile(*headers)
		if err != nil {
			log.Fatalf("failed to read headers file: %v", err)
		}
	}

	c := createCollector(baseDomain, *depth, *threads, *proxies, time.Duration(*timeout)*time.Second)
	setCollyBehavior(c, *agent, customHeaders)

	visited := make(map[string]bool) // use a map for efficient visited tracking

	// helper function to check if a URL should be visited
	shouldVisit := func(link string) bool {
		if link == "" {
			return false
		}
		u, err := url.Parse(link)
		if err != nil {
			return false // skip invalid URLs
		}
		hostname := u.Hostname()
		if hostname == baseDomain || strings.HasSuffix(hostname, "."+baseDomain) {
			if !visited[link] {
				visited[link] = true
				return true
			}
		}
		return false
	}

	process := func(e *colly.HTMLElement, urlAttr string) {
		link := e.Request.AbsoluteURL(e.Attr(urlAttr))
		if shouldVisit(link) {
			fmt.Println(link)
			e.Request.Visit(link)
		}
	}

	c.OnHTML("form[action]", func(e *colly.HTMLElement) {
		process(e, "action")
	})

	c.OnHTML("a[href], link[href]", func(e *colly.HTMLElement) {
		process(e, "href")
	})

	c.OnHTML("script[src], iframe[src], img[src]", func(e *colly.HTMLElement) {
		process(e, "src")
	})

	c.OnHTML("meta[http-equiv=refresh][content]", func(e *colly.HTMLElement) {
		content := e.Attr("content")
		if urlIdx := strings.Index(content, "url="); urlIdx != -1 {
			link := e.Request.AbsoluteURL(content[urlIdx+4:])
			if shouldVisit(link) {
				fmt.Println(link)
				e.Request.Visit(link)
			}
		}
	})

	err = c.Visit(*domain)
	if err != nil {
		log.Fatalf("failed to start crawl: %v", err) // Use log.Fatalf
	}
	c.Wait()
}
