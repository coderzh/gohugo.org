+++
screenshot = "/images/herring-cove.screenshot.png"
thumbnail = "/images/herring-cove.tn.png"
demo = "/theme/herring-cove/"
source = "https://github.com/spf13/herring-cove.git"
name = "Herring Cove"
description = "Herring Cove is ported from the jekyll theme of the same name"
license = "MIT"
tags = ["blog", "company"]
min_version = 0.14

[author]
    name = "spf13"
    homepage = "http://spf13.com"

# If Porting existing theme
[original]
    author =  "arnp"
    repo = "http://www.github.com/arnp/herring-cove"

+++

Herring Cove
============

Herring Cove is a clean and responsive [Hugo](//gohugo.io) theme based on the [Jekyll](http://jekyllrb.com) theme of the same name.


### Overview

* Fixed Sidebar with social links
* Minimal design 
* Comments by Disqus
* Social Sharing abilities 
* Easy to configure
* Based on Bootstrap

### Screenshots

![screenshot](/images/screenshot1.png)
![screenshot](/images/screenshot2.png)
![screenshot](/images/screenshot-landing.png)

### Setup

1. Install Hugo
2. Fork or [download](//github.com/spf13/herring-cove/archive/master.zip) this theme repo
3. Edit the config file in the root directory of your site

### Download

[download](//github.com/spf13/herring-cove/archive/master.zip)

### License
* [MIT](//opensource.org/licenses/MIT)

## Author
**Ravi Patel**
- <https://github.com/arnp>

## Ported By
**Steve Francia**
- <https://github.com/spf13>

-------------
Herring Cove is always a work in progress and as such, I hope to clean up the code and add features as time permits. Feel free to add your own additions. 

*What's with the name?*

[Herring Cove](//www.capecodbeachchair.com/beachguide/index.cfm?page=3&BeachID=5) is a beach on Cape Cod in Provincetown, Massachusetts. If you're in the area, be sure to check it out!

## Additional features by
**Maikel Bollemeijer**
- <https://github.com/mbollemeijer>

-------------

1. Profile picture parameter
2. Links in the menus are generated through params in the config
3. Landingpage

Below an example of the yaml config
```yaml
params:
  ProfilePicture: "pathOrUrlToImage"
  links:
    Home : "/"
    Blog: "/blog/"
    About: "/about/"
```

## Landing page
If you want to enable the landing page make sure you have the following params in your config.

```yaml
params:
  ProfilePicture: "pathToImageOrUrlToImage"
  enableLandingPage: true
  landingPageRedirectUrl: "/about/"
  AuthorName: "John Doe"
```
