# dizi-cli 
![Go Version](https://img.shields.io/badge/go-%3E%3D1.23-blue.svg?logo=go&logoColor=white)
[![Colly](https://img.shields.io/badge/colly-v2.8.1-green.svg)](https://github.com/gocolly/colly)
[![ChromeDP](https://img.shields.io/badge/chromedp-v0.9.1-green.svg)](https://github.com/chromedp/chromedp)

A cli tool for watching TV shows directly from your terminal. Built with Go using Colly for web scraping and ChromeDP for headless browsing.Scraping some **pirated** show sites and watching them from your terminal.

**Disclaimer**: This project is for educational purposes only. Use it at your own risk. I do not endorse or support piracy.I just made this to learn more about go and web scraping.

##Features and Usage

<table border="0" cellspacing="0" cellpadding="0" style="border: none; border-collapse: collapse;">
  <tr style="border: none;">
    <td width="50%" style="border: none;">
      <img src="./readmegifs/1.gif" style="width: 100%">
      <p align="center"><b>Search and select shows</b></p>
    </td>
    <td width="50%" style="border: none;">
      <img src="./readmegifs/2.gif" style="width: 100%">
      <p align="center"><b>Browse seasons & episodes</b></p>
    </td>
  </tr>
  <tr style="border: none;">
    <td colspan="2" style="border: none;">
      <img src="./readmegifs/3.gif" style="width: 100%; max-width: 800px; display: block; margin: 0 auto">
      <p align="center"><b>Stream the chosen episode</b></p>
    </td>
  </tr>
</table>

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
