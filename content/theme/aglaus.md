+++
screenshot = "/images/aglaus.screenshot.png"
thumbnail = "/images/aglaus.tn.png"
source = "https://github.com/dim0627/hugo_theme_aglaus"
name = "Aglaus"
license = "MIT"
licenselink = "https://github.com/dim0627/hugo_theme_aglaus/blob/master/LICENSE.md"
description = "Aglaus is list design with eyecatch theme.Inspired by lanyon."
homepage = "http://yet.unresolved.xyz/hugo_theme_aglaus"
tags = ["blog", "list", "eyecatch"]
features = ["blog"]
min_version = 0.14

[author]
    name = "Daisuke Tsuji"
    homepage = "http://yet.unresolved.xyz"


+++

# Aglaus

    Sorry, I do not have a good at English.
    It is naive English, but please acknowledge.

Aglaus is a single-column theme for [Hugo](http://gohugo.io/).

![Aglaus Screenshot](https://raw.githubusercontent.com/dim0627/hugo_theme_aglaus/master/images/top.png)

![Aglaus Screenshot](https://raw.githubusercontent.com/dim0627/hugo_theme_aglaus/master/images/bottom.png)

![Aglaus Screenshot](https://raw.githubusercontent.com/dim0627/hugo_theme_aglaus/master/images/post.png)

## Features

* Google Analytics
* Gravatar Profile
* Disqus
* SNS Links(Facebook, Twitter, GitHub)
* Share Button
* Eye-catching Image
* Tagging
* Related Post

# Installation

Referred from [hugoThemes#Installing Themes](https://github.com/spf13/hugoThemes#installing-themes).

## Installing with other all themes

If you would like to install all of the available hugo themes, simply clone the entire repository from within your working directory.

    git clone --recursive https://github.com/spf13/hugoThemes.git themes

## Installing a single theme

    mkdir themes
    cd themes
    git clone https://github.com/dim0627/hugo_theme_aglaus aglaus
    
## Build with Aglaus

    hugo server -t aglaus

# Configuration

**config.yaml**

``` yaml
BaseURL: "http://example.com"
LanguageCode: "en-us"
Title: "Aglaus"

Params:
  Author: "Your name."
  Birth: "Sun, Feb 26, 1989"
  DateForm: "Mon, Jan 2, 2006"
  GoogleAnalyticsUserID: "Your ID."
  GravatarHash: "Your Hash."
  Facebook: "Your ID."
  Twitter: "Your ID."
  Github: "Your ID."
  ShowRelatedPost: True
  Disqus: "Your Disqus."
  SyntaxHighlightTheme: "solarized_dark.min.css"

Indexes:
  tag: "tags"

permalinks:
  post: /blog/:year/:month/:day/:title/

MetadataFormat: "yaml"
```

Example : [My config.yaml](https://github.com/dim0627/dim0627.github.io/blob/source/config.yaml)

**example post**

``` markdown
---
title: "Post title here"
eyecatch: "hugo.png" # Eye-cathinc image from [static/images/***]
date: 2014-09-17
comments: true
tags: [gitHub, octopress, jekyll]
---

Contents here
```

# Read more details

Please read this.

http://yet.unresolved.xyz/hugo_theme_aglaus/blog/2015/01/06/about-aglaus/

# Contact us

Please mail to `dim0627@gmail.com` or SNS.

[https://www.facebook.com/daisuke.tsuji.735](https://www.facebook.com/daisuke.tsuji.735)

[https://twitter.com/dim0627](https://twitter.com/dim0627)
