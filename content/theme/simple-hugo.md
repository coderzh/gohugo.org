+++
screenshot = "/images/simple-hugo.screenshot.png"
thumbnail = "/images/simple-hugo.tn.png"
demo = "/theme/simple-hugo/"
source = "https://github.com/druzza/simple-hugo"
name = "simple-hugo"
license = "MIT"
licenselink = "https://github.com/druzza/simple-hugo/blob/master/LICENSE.md"
description = "An elegant, open source and simple theme"
homepage = "https://github.com/druzza/simple-hugo"
tags = ["blog", "dev"]
features = ["blog", "dev"]
min_version = 0.14

[author]
  name = "Daniel Ruzza"
  homepage = "https://github.com/druzza/simple-hugo"
  repo = "https://github.com/druzza/simple-hugo"

+++

# Alert
Work in progress.

# Simple Hugo
A simple [hugo](http://gohugo.io/) theme.

#### Params
- For add new links on navbar see ```[params.links]```.
```toml
[params]
  github = "username"
  twitter = "username"
  email = "email"
  shortDescription = "yada yada yada"
  [params.links]
    [params.links.About]
      "url" = "/about/"
      "title" = "About"
    [params.links.ContactMe]
      "url" = "/contact/"
      "title" = "Contact Me"
```
