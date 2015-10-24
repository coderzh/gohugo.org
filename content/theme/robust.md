+++
screenshot = "/images/robust.screenshot.png"
thumbnail = "/images/robust.tn.png"
demo = "/theme/robust/"
source = "https://github.com/dim0627/hugo_theme_robust.git"
name = "Robust"
license = "MIT"
licenselink = "https://github.com/dim0627/hugo_theme_robust/blob/master/LICENSE.md"
description = "'Robust' is two column theme for blogger."
homepage = ""
tags = ["blog"]
features = ["blog"]

[author]
    name = "Daisuke Tsuji"
    homepage = "http://yet.unresolved.xyz"


+++

# Robust

Robust is a theme for [Hugo](http://gohugo.io/).

![Robust Screenshot](https://raw.githubusercontent.com/dim0627/hugo_theme_robust/master/images/screenshot.png)

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
    git clone https://github.com/dim0627/hugo_theme_robust robust

## Build

    hugo server -t robust

# Configuration

**config.yaml**

``` yaml
BaseURL: "http://example.com"
LanguageCode: "en-us"
Title: "Robust"

Params:
  Author: "Your name."
  Birth: "Sun, Feb 26, 1989"
  DateForm: "Mon, Jan 2, 2006"
  GoogleAnalyticsUserID: "Your ID."
  GravatarHash: "Your Hash."
  Facebook: "Your ID."
  Twitter: "Your ID."
  Github: "Your ID."
  Disqus: "Your Disqus."
  SyntaxHighlightTheme: "solarized_dark.min.css"
  ShowRelatedPost: True
  ShowTagCloud: True

Indexes:
  tag: "tags"

permalinks:
  post: /blog/:year/:month/:day/:title/

MetadataFormat: "yaml"
```

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

# Contact us

Please mail to `dim0627@gmail.com` or SNS.

[https://www.facebook.com/daisuke.tsuji.735](https://www.facebook.com/daisuke.tsuji.735)
