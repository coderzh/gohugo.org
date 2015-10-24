+++
screenshot = "/images/hugo-incorporated.screenshot.png"
thumbnail = "/images/hugo-incorporated.tn.png"
source = "https://github.com/nilproductions/hugo-incorporated.git"
name = "Hugo Incorporated"
license = "MIT"
licenselink = "https://github.com/nilproductions/hugo-incorporated/blob/master/LICENSE"
description = "Modern blog, great for companies, products or anything."
homepage = "https://github.com/nilproductions/hugo-incorporated"
tags = ["company", "product"]
features = ["blog", ]
min_version = 0.14

[author]
    name = "Nicholas Whittier"
    homepage = "http://github.com/imperialwicket"

[original]
    author =  "Karri Saarinen and Jori Lallo"
    homepage = "https://sendtoinc.com/"

+++

# Hugo Incorporated
Modern Hugo based blog, based entirely on [jekyll-incorporated](https://github.com/kippt/jekyll-incorporated). Great for companies, products or anything. See jekyll-incorporated live at [blog.sendtoinc.com](http://blog.sendtoinc.com). See hugo-incorporated live at [blog.nilproductions.com](http://blog.nilproductions.com).

## Installation & Usage

Hugo-incorporated requires Hugo 0.12-DEV or newer. If your content is not displaying after running `hugo server`, be sure that your version (`hugo version`) is at least 0.12-DEV.

Markdown Content goes in content/post/title.md. Then:

    hugo server -w

The stylesheet is included in static/css. If you want substantial customizations you should use the scss directory and install sass (and Ruby). At the moment, this is entirely optional, since I'm commiting the final stylesheet. Note that if you make customizations to main.css, then rebuild with sass, those changes could get lost.

    bundle install
    sass scss/main.scss static/css/main.css

More hugo-specific details available at [hugo](http://hugo.spf13.com/).

## Configuration
Hugo-incorporated specific configuration options are available in [config.yaml].

Edit: config.yaml (general options), main.css (theme colors &amp; fonts)

```
hugo-incorporated/
├── config.yaml
├── _scss/
    ├── main.scss
```

_Note: when editing config.yaml, you need to restart hugo to see the changes._

## Authors

Originally built as a Jekyll theme for [sendtoinc.com](https://sendtoinc.com).

**Karri Saarinen**

+ [http://twitter.com/karrisaarinen](http://twitter.com/karrisaarinen)
+ [http://github.com/ksaa](http://github.com/ksaa)

**Jori Lallo**

+ [http://twitter.com/jorilallo](http://twitter.com/jorilallo)
+ [http://github.com/jorde](http://github.com/jorilallo)

Ported to Hugo:

**Nicholas Whittier**

+ [http://twitter.com/nw_iw]
+ [http://github.com/imperialwicket]

## Todo:

+ Document publishing
+ Offer sass-less version (Remove ruby entirely?)
+ Fix feeds and 404 pages
+ Confirm core functionality
