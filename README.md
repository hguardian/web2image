
# web2image

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/suntong/web2image?status.svg)](http://godoc.org/github.com/suntong/web2image)
[![Go Report Card](https://goreportcard.com/badge/github.com/suntong/web2image)](https://goreportcard.com/report/github.com/suntong/web2image)
[![travis Status](https://travis-ci.org/suntong/web2image.svg?branch=master)](https://travis-ci.org/suntong/web2image)

## TOC
- [web2image - Web to image](#web2image---web-to-image)
- [Usage](#usage)
  - [$ web2image](#-web2image)
  - [chrome-headless](#chrome-headless)
- [Examples](#examples)
  - [Film info from cmore.se](#film-info-from-cmorese)
- [Installation](#installation)
  - [Author](#author)

## web2image - Web to image

The `web2image` will take a screenshot of a given web page.

# Usage

### $ web2image
```sh
Web to image
built on 2017-07-23

Tool to take screenshot from a web page

Usage:
  web2image OPTIONS... URL IMAGE-FILE

Options:

  -h, --help       display help information
  -d, --headless   use chrome-headless docker as client instead of chrome
  -c, --css       *the CSS selector for the region to take the screenshot of
  -v, --verbose    verbose mode (multiple -v options increase the verbosity.)
```

### chrome-headless

The following examples assumes that you've [setup chrome-headless](https://github.com/knq/chromedp/blob/master/examples/headless/README.md) already. Else, take a look at [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/) on how to start a local chrome with the remote-debugging-port opening at `9222`.

# Examples

## Film info from cmore.se

    web2image -d -c '#main-wrapper' 'http://www.cmore.se/film/3643033-deadpool' deadpool.png

and the result is:

![deadpool](deadpool.png "deadpool.png")

# Installation

```
go get github.com/suntong/web2image
```

Debian package will be provided shortly.

## Author

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

All patches welcome. 
