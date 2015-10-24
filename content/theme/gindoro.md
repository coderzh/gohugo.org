+++
screenshot = "/images/gindoro.screenshot.png"
thumbnail = "/images/gindoro.tn.png"
demo = "/theme/gindoro/"
source = "https://github.com/cdipaolo/gindoro.git"
# theme.toml template for a Hugo theme
# See https://github.com/spf13/hugoThemes#themetoml for an example

name = "gindoro"
license = "MIT"
licenselink = "https://github.com/cdipaolo/gindoro/blob/master/LICENSE.md"
description = "minimalistic and _extremely_ lightweight hugo theme"
homepage = "http://dipaolo.conner.sh/"
tags = ["minimal", "hugo", "lightweight", "low-latency"]
features = ["very lightweight", "renders really, really quickly", "super minimal"]
min_version = 0.13

[author]
  name = "Conner DiPaolo"
  homepage = "http://dipaolo.conner.sh"

+++

#### gindoro
##### a minimalistic Hugo blogging theme in < 6K

[example site](http://dipaolo.conner.sh)

only two pages – index and post

no syntax highlighting (but code blocks do look nice regardless!)

LaTeX support built in (using KaTeX via CDN for fast rendering)

![front page screenshot](images/gindoro-index-screenshot.png)
![screenshot](images/gindoro-screenshot.png)

### installation

**configuration**

config.toml can include the following parameters (below is a complete example that will function as-is.) The example results in the above site with footer links to github and contact:

```toml
baseurl = "http://dipaolo.conner.sh"
languageCode = "en-us"
title = "conner dipaolo"

[params]
  Description = "build cool shit."
  Author = "conner dipaolo"
  GithubUser = "cdipaolo"
  Email = "cdipaolo96@gmail.com"
```

**starting a new blog**

these instructions are strikingly similar to [the Hugo quickstart guide](http://gohugo.io/overview/quickstart/)!

```bash
$ hugo new site path/to/site
$ cd path/to/site

# make a post (can edit later!)
$ hugo new hello-world.md

# install theme
$ git clone https://github.com/cdipaolo/gindoro themes

# run site!
$ hugo server -t gindoro --buildDrafts
2 pages created
0 tags created
0 categories created
in 5 ms
Serving pages from exampleHugoSite/public
Web Server is available at http://localhost:1313
Press Ctrl+C to stop
```
