package scrape

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"github.com/imroc/req/v3"
	"mole/utils"
)

// function to create a colly collector
func createCollector(domain string, depth int, threads int, proxies string, timeout time.Duration) *colly.Collector {
	fakeChrome := req.DefaultClient().ImpersonateChrome()
	c := colly.NewCollector(
		colly.UserAgent(fakeChrome.Headers.Get("user-agent")), // default user agents
		colly.AllowedDomains(domain, "*."+domain),             // subdomains can be removed
		colly.IgnoreRobotsTxt(),
		colly.MaxDepth(depth),
		colly.Async(true),
	)

	// set proxy rotator if proxies were specified
	if proxies != "" {
		proxyList, err := utils.ReadFile(proxies)
		if err != nil {
			log.Fatal(err) // exit if we couldn't read the file
		}
		rotator, err := proxy.RoundRobinProxySwitcher(proxyList...)
		if err != nil {
			log.Fatal(err) // exit if failed to rotate proxies
		}
		c.SetProxyFunc(rotator)
	}

	c.SetClient(&http.Client{
		Transport: fakeChrome.Transport,
		Timeout:   timeout,
	})

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: threads,
	})

	return c
}

// function to set the behavior on request and error (add as necessary)
func setCollyBehavior(c *colly.Collector, agent string, headers []string) {
	c.OnRequest(func(r *colly.Request) {
		if agent != "" {
			// set user agent if specified
			r.Headers.Set("User-Agent", agent)
		}

		// custom headers if specified
		if len(headers) != 0 {
			for _, header := range headers {
				parts := strings.Split(header, ":")
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					value := strings.TrimSpace(parts[1])
					r.Headers.Set(key, value)
				}
			}
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("[-] Error Occurred: %v\n -> URL: %v\n", err, r.Request.URL.String()) // no worries, just keep strolling
	})
}
