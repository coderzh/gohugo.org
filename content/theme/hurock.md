+++
screenshot = "/images/hurock.screenshot.png"
thumbnail = "/images/hurock.tn.png"
demo = "/theme/hurock/"
source = "https://github.com/TiTi/hurock"
name = "hurock"
description = "A classic sidebar hugo theme"
license = "MIT"
licenselink = "https://github.com/titi/hurock/LICENSE"
homepage = "https://github.com/titi/hurock"
tags = ["blog", "personal"]
features = ["sidebar", "TableOfContents", "highlight", "404"]

[author]
  name = "Thibault ROHMER"
  homepage = "http://titi.github.io"

+++

# hurock
A classic sidebar Hugo theme

[My personnal Hugo blog](https://titi.github.io/)

# Screenshots

index.html:

<img src="https://raw.githubusercontent.com/titi/hurock/master/images/screenshot1.png" width="800" />

List view:

<img src="https://raw.githubusercontent.com/titi/hurock/master/images/screenshot2.png" width="800" />

Content view:

<img src="https://raw.githubusercontent.com/titi/hurock/master/images/screenshot3.png" width="800" />

404 error page:

<img src="https://raw.githubusercontent.com/titi/hurock/master/images/screenshot4.png" width="800" />

# Concept

* No external ressources (js/css/fonts)
* Easy to add stuff in header/footer
* Syntax highlight on client side using [highlight.js](https://highlightjs.org/)
* Shortcodes for vimeo/youtube/gist
* W3C valid
 * Clean HTML
 * Sidebar after content
 * CSS in head, JS before end of body

Theme mostly insipred on theme `nofancy` but also: `hyde-x`, `lanyon`, `liquorice`, `purehugo`, `redlounge`.

# Config

``` toml
baseurl = "http://example.com"
title = "Your site title"

[author]
  name = "Your Name"
  email = ""
  
  github = ""
  twitter = ""
  facebook = ""
  googleplus = ""
  linkedin = ""
  gaID = "Google Analytics Tracking ID"

[params]
  Description = "text under logo"
  # Optional
  disqusShortname = ""
  highlight_theme = "tomorrow-night"
  notoc = true


# Optional sections:
[permalinks]
  posts = "/:year/:month/:filename/"

[blackfriday]
  plainIdAnchors = true
```

* You can change the `highlight.js` theme, see `highlight_theme` parameter, but you'll have to add the corresponding css in your `static/highlight/` folder.
* `notoc` means no TableOfContents. It is used to hide the table of contents. Indeed, in case you're using the permalinks parameter, the table of contents links are not working with current Hugo version (0.13).
