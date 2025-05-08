<div align="center">
  <img src="./assets/mole.png" alt="Mole Logo" />
  <h1><strong>Mole</strong></h1>
  <p>🔎 Dig through web applications to find all endpoints 🔍</p>
</div>

Utilizing the Colly web scraping library, Mole efficiently navigates through websites to provide a comprehensive list of endpoints, allowing for deeper insight into the structure and available resources of a site.
This was made with a bounty in mind 💰


## Installation ⚙️
### Prerequisite
| Prerequisite | Version |
|--------------|---------|
| Go           |  <=1.22 |

```bash
# apt
apt install golang-go
# brew
brew install golang
```

### Clone Repo
```bash
git clone https://github.com/CharlesTheGreat77/mole
cd mole
# download(s) dependencies
go mod init mole
go mod tidy
go build -o mole main.go
```

## Usage 🔨

```bash
╰─ mole -h
Usage of mole:
  -custom-headers string
      specify a file containing headers to send on request [separated by line]
  -depth int
      specify the max depth for crawling [default: 2] (default 2)
  -h	show usage
  -proxies string
      specify a file containing http/https/socks5 proxies [separated by line]
  -threads int
      specify the number of threads [default: 2] (default 2)
  -timeout int
      specify a timeout (seconds) [default: 3] (default 3)
  -url string
      specify a url [https://example.com]
  -user-agent string
      specify a custom user agent
```

You will be **required** to specify a *url*.

# Demo 📹
![Demo](https://github.com/user-attachments/assets/f53b2a57-3697-46d4-a75e-ba4eb845c425)


## Mole crawler 
Out of the box, mole crawls html content for paths [endpoints] with the *selectors* specified into **crawl.go**. One can add *selectors* as they see fit.

Here's a quick example `./internal/scrape/crawl.go`:

```go
// selector to search for
c.OnHTML("form[action]", func(e *colly.HTMLElement) {
    // process the URL
		process(e, "action")
})
```
→ See <a href="https://go-colly.org/docs/introduction/start/">Colly Docs</a> for more details.

# Contributing
Contributions are welcome! Please feel free to submit a Pull Request.
