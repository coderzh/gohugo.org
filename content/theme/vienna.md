+++
screenshot = "/images/vienna.screenshot.png"
thumbnail = "/images/vienna.tn.png"
demo = "/theme/vienna/"
source = "https://github.com/keichi/vienna.git"
author = "Keichi Takahashi"
description = "Simple and clean blog theme for hugo"
license = "MIT"
name = "vienna"
source_repo = "git@github.com:keichi/vienna.git"
tags = ["blog", "tags", "bootstrap"]

+++

# Vienna

## Overview

Vienna is a simple and clean blog theme for [Hugo](http://gohugo.io/).
Notable features are:

- Simple and clean design
- Client side source code highlighting
- Social links (Twitter, Facebook, GitHub, LinkedIn, Instagram, Keybase)
- Support for tags
- Analytics with Google Analytics or Mixpanel
- Responsive design
- Font Awesome icons

## Installation

In your hugo site directory, run:

```shell
$ mkdir themes
$ cd themes
$ git clone https://github.com/keichi/vienna
```

Vienna is also included in the `spf13/hugoThemes` repository.

## Configuration

You may specify following options in `config.toml` of your site to make use of
this theme's features.

```toml
baseurl = "Your site URL"
languageCode = "en-us"
title = "Your site title"
# Copyright notice. This is displayer in the footer.
copyright = "&copy; Copyright notice"

[params]
    # Social accounts. Link to these accounts are displayed in the header and
    # footer.
    twitter = "Your Twitter username"
    github = "Your GitHub username"
    linkedin = "Your LinkedIn username"
    googleplus = "Your Google+ user id"
    facebook = "Your Facebook username"
    stackoverflow = "Your Stackoverflow user id (number)"
    keybase = "Your keybase.io username"
    instagram = "Your Instagram username"
    # Disqus shortname
    disqus = "Your disqus shortname"
    # Google Analytics API key.
    ga_api_key = "Your Google Analytics tracking id"
    # Mixpanel API key.
    mixpanel_api_key = "Your Mixpanel API key"
    author = "Your Name"
    avatar = "/path/to/avatar"
    contact = "Your contact link (ex. mailto:foo@example.com)"
    bio = "Your short bio"
    # Short subtitle/tagline. This is displayed in the header.
    subtitle = "Short subtitle/tagline of your blog"
    themecolor = "#hexcolor" # Defines the tab color in Chrome for Android.
```

## Usage

Use hugo's `-t vienna` or `--theme=vienna` option with hugo commands.
Example:

```shell
$ hugo server -t vienna -w -D
```
