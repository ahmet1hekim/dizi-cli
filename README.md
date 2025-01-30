# dizi-cli 
![Go Version](https://img.shields.io/badge/go-%3E%3D1.23-blue.svg?logo=go&logoColor=white)
[![Colly](https://img.shields.io/badge/colly-v2.8.1-green.svg)](https://github.com/gocolly/colly)
[![ChromeDP](https://img.shields.io/badge/chromedp-v0.9.1-green.svg)](https://github.com/chromedp/chromedp)

A cli tool for watching TV shows directly from your terminal. Built with Go using Colly for web scraping and ChromeDP for headless browsing.Scraping some **pirated** show sites and watching them from your terminal.

**Disclaimer**: This project is for educational purposes only. Use it at your own risk. I do not endorse or support piracy.İ just made this to learn more about go and web scraping.

## Features ✨
<div style="display: flex; flex-wrap: wrap; gap: 20px; justify-content: center; align-items: flex-start;">
  <!-- First Row - Side by Side -->
  <div style="flex: 1 1 45%; min-width: 300px; box-sizing: border-box;">
    <img src="./readmegifs/1.gif" style="width: 100%; height: auto; display: block;">
    <p align="center"><b>Search and select shows</b></p>
  </div>
  
  <div style="flex: 1 1 45%; min-width: 300px; box-sizing: border-box;">
    <img src="./readmegifs/2.gif" style="width: 100%; height: auto; display: block;">
    <p align="center"><b>Search and select seasons and episodes</b></p>
  </div>

  <!-- Second Row - Full Width -->
  <div style="flex: 0 1 100%; max-width: 800px; margin: 0 auto;">
    <img src="./readmegifs/3.gif" style="width: 100%; height: auto; display: block;">
    <p align="center"><b>Stream the chosen episode</b></p>
  </div>
</div>

## Installation ⚙️
### releases
You can get the precompiled version from the releases tab.
### Build From Source
#### Prerequisites
- Go 1.23+
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
