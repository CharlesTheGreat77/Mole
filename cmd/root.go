package cmd

import (
	"flag"
	"mole/internal/scrape"
)

func Execute() {
	url := flag.String("url", "", "specify a url [https://example.com]")
	agent := flag.String("user-agent", "", "specify a custom user agent")
	headers := flag.String("custom-headers", "", "specify a file containing headers to send on request [separated by line]")
	proxies := flag.String("proxies", "", "specify a file containing http/https/socks5 proxies [separated by line]")
	threads := flag.Int("threads", 2, "specify the number of threads [default: 2]")
	depth := flag.Int("depth", 2, "specify the max depth for crawling [default: 2]")
	timeout := flag.Int("timeout", 3, "specify a timeout (seconds) [default: 3]")

	help := flag.Bool("h", false, "show usage")

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	scrape.EndPoint(
		url,
		agent,
		headers,
		proxies,
		threads,
		depth,
		timeout,
	)
}
