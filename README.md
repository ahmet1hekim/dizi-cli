# Dizi-CLI 🎬

[![Go Report Card](https://goreportcard.com/badge/github.com/ahmet1hekim/dizi-cli)](https://goreportcard.com/report/github.com/ahmet1hekim/dizi-cli)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A cli tool for watching TV shows directly from your terminal. Built with Go using Colly for web scraping and ChromeDP for headless browsing.Scraipng some **pirated** show sites and watching them from your terminal.

**Disclaimer**: This project is for educational purposes only. Use it at your own risk. The developers do not endorse or support piracy.

## Features ✨
<div style="display: flex; flex-wrap: wrap; gap: 20px; justify-content: center">
![searching and selecting from shows](./readmegifs/1.gif)
![searching and selecting from episodes and seasons](./readmegifs/2.gif)
![aaand just streaming them](./readmegifs/3.gif)
</div>

## Installation ⚙️
### releases
You can get the precompiled version from the releases tab.
### Build From Source
#### Prerequisites
- Go 1.16+
- Chrome/Chromium browser

```bash
# Clone the repository
git clone https://github.com/ahmet1hekim/dizi-cli.git
cd dizi-cli

# Install dependencies
go get github.com/gocolly/colly/v2
go get github.com/chromedp/chromedp

# Build and install
go build -o dizi-cli
sudo mv dizi-cli /usr/local/bin/
