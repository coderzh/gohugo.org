+++
screenshot = "/images/hugo-bootswatch.screenshot.png"
thumbnail = "/images/hugo-bootswatch.tn.png"
demo = "/theme/hugo-bootswatch/"
source = "https://github.com/nilproductions/hugo-bootswatch.git"
name = "Hugo Bootswatch"
license = "MIT"
licenselink = "https://github.com/nilproductions/hugo-bootswatch/LICENSE.md"
description = "Bootswatch Theme (Bootstrap 3) for Hugo"
homepage = "https://github.com/nilproductions/hugo-bootswatch"
tags = ["bootstrap", "bootswatch"]
features = []
min_version= 0.13

[author]
  name = "Nicholas Whittier"
  homepage = "http://imperialwicket.com"

# If porting an existing theme
[original]
  name = "Thomas Park"
  homepage = "http://bootswatch.com"
  repo = "https://github.com/thomaspark/bootswatch"

+++

# Hugo-bootswatch

Hugo-bootswatch is a [Hugo](http://gohugo.io) theme based on [Bootswatch](http://bootswatch.com/), which themes [Bootstrap 3](http://getbootstrap.com/). As such, Hugo-bootswatch will work as a Bootstrap theme, and also allows simple updating to support any Bootswatch customizations. Hugo-bootswatch comes with Bootstrap 3.3.2 and the bootswatch.js is already included. Customizing the Bootswatch theme is as easy as downloading the Bootswatch-customized bootstrap.min.css file (details below).

Hugo-bootswatch expects [section organization](http://gohugo.io/content/sections/), which are loaded as navbar links in the header, and placed in a "Sections" dropdown.

Hugo-bootswatch was originally built for documentation, so it is missing a lot of blog-type support (most notably the concept of authors). A 'post' archetype may make sense to accommodate this if there is interest.

## Customization

Hugo-bootswatch is made for customization, and it is easiest to customize by using your own `static/css/hugo-bootswatch.css` and `static/css/bootstrap.min.css` files. Note that these files are not empty, so the overriding file will need to accommodate any existing content by inclusion or replacement.

### Bootswatch

Visit [Bootswatch](http://bootswatch.com/) and choose the theme you like. Download the available `bootstrap.min.css` replacement and place it in `static/css/bootstrap.min.css` to override the default (Bootstrap 3.3.2) file.

### Hugo-bootswatch

Make any changes that are particular to your site in `static/css/hugo-bootswatch.css`. This will automatically apply your customizations to the site.
