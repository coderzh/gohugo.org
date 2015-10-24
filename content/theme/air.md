+++
screenshot = "/images/air.screenshot.png"
thumbnail = "/images/air.tn.png"
demo = "/theme/air/"
source = "https://github.com/syui/hugo-theme-air.git"
name = "Air"
license = "MIT"
licenselink = "https://github.com/syui/hugo-theme-air/blob/master/LICENSE.md"
description = "Casper theme from Ghost Foundation, Plus particles.js"
homepage = "https://syui.github.io/hugo-theme-air"
tags = ["blog"]
features = ["disqus", "google analytics","share"]
min_version=0.01

[author]
  name = "Ichinose Syuu"
  homepage = "https://syui.github.io"

[original]
  name = "hugo-theme-casper"
  homepage = "http://vjeantet.fr/"
  repo = "https://github.com/vjeantet/hugo-theme-casper"

+++

# AIR theme for hugo

Air is a single-column theme for [Hugo](http://gohugo.io/).
Ported from [Casper theme for Ghost ](https://github.com/TryGhost/Casper), [vjeantet/hugo-theme-casper](https://github.com/vjeantet/hugo-theme-casper)

- blog demo : http://syui.github.io/hugo-theme-air

- blog source : https://github.com/syui/hugo-theme-air

![Hugo Air Theme screenshot](https://raw.githubusercontent.com/syui/hugo-theme-air/master/images/screen.gif)

```bash
$ git clone https://github.com/syui/hugo-theme-air

$ cd hugo-theme-air

$ cp config.toml.backup config.toml

$ hugo new post/foo.md

$ hugo server
---------------------------------
$ curl 127.0.0.1:1313/hugo-theme-air 
```

## Features

* [VincentGarreau/particles.js](https://github.com/VincentGarreau/particles.js/)
* Google Analytics (optional)
* Disqus ( can disable comments by content)
* Share buttons on Facebook, Twitter, Google (can disable share by content)
* Big cover image (optional)
* Custom cover by content (optional)
* Tagging
* Pagination
* Menu

# Theme usage and asumptions
* All blog posts are in the ```post``` folder (```content/post```)
* The homepage displays a paginated list of contents from the post Section (other contents may be added to main menu, see bellow)

# Installation

## Installing this theme

    mkdir themes
    cd themes
    git clone https://github.com/syui/hugo-theme-air

## Build your website with this theme

    hugo server -t hugo-theme-air

# Configuration

**config.toml**

``` bash
$ cat config.toml.backup

$ cp config.toml.backup config.toml
```

Example : [config.toml](https://github.com/syui/hugo-theme-air/blob/master/config.toml.backup)

## Multiple authors configuration

In addition to providing data for a single author as shown in the example above, multiple authors
can be configured via data/authors/\*.(yml, toml, json) entries. If the key provided in
.Site.Params.author matched a data/authors/\* entry, it will be used as the default. Overrides
per page can be done by a simple author = other_author_key entry in the front matter. For those
pages where you want to omit the author block completely, a .Params.noauthor entry is also
available.

``` bash
$ hugo new post/foo.md

$ cat content/post/foo.md
```

Example author definition file:


``` yml
name: John Doe
bio: The most uninteresting man in the world.
location: Normal, IL
website: http://example.com

```

Example override author per page file:
``` toml
+++
author = ""
date = "2014-07-11T10:54:24+02:00"
title = ""
...
+++

Contents here

```

## Metadata on each content file, example

``` toml
+++
author = ""
date = "2014-07-11T10:54:24+02:00"
draft = false
title = "dotScale 2014 as a sketch"
slug = "dotscale-2014-as-a-sketch"
tags = ["event","dotScale","sketchnote"]
image = "images/2014/Jul/titledotscale.png"
comments = true     # set false to hide Disqus comments
share = true        # set false to share buttons
menu = ""           # set "main" to add this content to the main menu
+++

Contents here
```

## Create new content based with default metadata from this theme
You can easyly create a new content with all metadatas used by this theme, using this command 
```
hugo new -t hugo-theme-air post/my-post.md
```
## Hosting for GitHub Pages

```bash
# build
$ hugo 

$ cd public

# preview
$ jekyll server
$ rm -rf _site

# make repository
$ git init
$ git remote add origin $url
$ git add .
$ git commit -m "first commit"
$ git push -u origin master

# push branch:gh-pages
$ git checkout -b gh-pages
$ git commit -m "open pages"
$ git push -u origin gh-pages

# open
$ curl user.github.io/repository
```

## Costomize for particles.js

> static/js/particles.js

key | option type / notes | example
----|---------|------|------
`particles.number.value` | number | `40`
`particles.number.density.enable` | boolean | `true` / `false` 
`particles.number.density.value_area` | number | `800`
`particles.color.value` | HEX (string) <br /> RGB (object) <br /> HSL (object) <br /> array selection (HEX) <br /> random (string) | `"#b61924"` <br /> `{r:182, g:25, b:36}` <br />  `{h:356, s:76, l:41}` <br /> `["#b61924", "#333333", "999999"]` <br /> `"random"`
`particles.number.density.value_area` | number | `800`
`particles.shape.type` | string <br /> array selection | `"circle"` <br /> `"edge"` <br /> `"triangle"` <br /> `"polygon"` <br /> `"star"` <br /> `"image"` <br /> `["circle", "triangle", "image"]`
`particles.shape.stroke.width` | number | `2`
`particles.shape.stroke.color` | HEX (string) | `"#222222"`
`particles.shape.polygon.nb_slides` | number | `5`
`particles.shape.image.src` | path link <br /> svg / png / gif / jpg | `"assets/img/yop.svg"` <br /> `"http://mywebsite.com/assets/img/yop.png"`
`particles.shape.image.width` | number <br />(for aspect ratio) | `100`
`particles.shape.image.height` | number <br />(for aspect ratio) | `100`
`particles.opacity.value` | number (0 to 1) | `0.75`
`particles.opacity.random` | boolean | `true` / `false` 
`particles.opacity.anim.enable` | boolean | `true` / `false` 
`particles.opacity.anim.speed` | number | `3`
`particles.opacity.anim.opacity_min` | number (0 to 1) | `0.25`
`particles.opacity.anim.sync` | boolean | `true` / `false`
`particles.size.value` | number | `20`
`particles.size.random` | boolean | `true` / `false` 
`particles.size.anim.enable` | boolean | `true` / `false` 
`particles.size.anim.speed` | number | `3`
`particles.size.anim.size_min` | number | `0.25`
`particles.size.anim.sync` | boolean | `true` / `false`
`particles.line_linked.enable` | boolean | `true` / `false`
`particles.line_linked.distance` | number | `150`
`particles.line_linked.color` | HEX (string) | `#ffffff`
`particles.line_linked.opacity` | number (0 to 1) | `0.5`
`particles.line_linked.width` | number | `1.5`
`particles.move.enable` | boolean | `true` / `false`
`particles.move.speed` | number | `4`
`particles.move.direction` | string | `"none"` <br /> `"top"` <br /> `"top-right"` <br /> `"right"` <br /> `"bottom-right"` <br /> `"bottom"` <br /> `"bottom-left"` <br /> `"left"` <br /> `"top-left"`
`particles.move.random` | boolean | `true` / `false`
`particles.move.straight` | boolean | `true` / `false`
`particles.move.out_mode` | string <br /> (out of canvas) | `"out"` <br /> `"bounce"`
`particles.move.bounce` | boolean <br /> (between particles) | `true` / `false`
`particles.move.attract.enable` | boolean | `true` / `false`
`particles.move.attract.rotateX` | number | `3000`
`particles.move.attract.rotateY` | number | `1500`
`interactivity.detect_on` | string | `"canvas", "window"`
`interactivity.events.onhover.enable` | boolean | `true` / `false`
`interactivity.events.onhover.mode` | string <br /> array selection | `"grab"` <br /> `"bubble"` <br /> `"repulse"` <br /> `["grab", "bubble"]`
`interactivity.events.onclick.enable` | boolean | `true` / `false`
`interactivity.events.onclick.mode` | string <br /> array selection | `"push"` <br /> `"remove"` <br /> `"bubble"` <br /> `"repulse"` <br /> `["push", "repulse"]`
`interactivity.events.resize` | boolean | `true` / `false`
`interactivity.events.modes.grab.distance` | number | `100`
`interactivity.events.modes.grab.line_linked.opacity` | number (0 to 1) | `0.75`
`interactivity.events.modes.bubble.distance` | number | `100`
`interactivity.events.modes.bubble.size` | number | `40`
`interactivity.events.modes.bubble.duration` | number <br /> (second) | `0.4`
`interactivity.events.modes.repulse.distance` | number | `200`
`interactivity.events.modes.repulse.duration` | number <br /> (second) | `1.2`
`interactivity.events.modes.push.particles_nb` | number | `4`
`interactivity.events.modes.push.particles_nb` | number | `4`
`retina_detect` | boolean | `true` / `false`

[VincentGarreau/particles.js](https://raw.githubusercontent.com/VincentGarreau/particles.js/master/README.md)

### color value black

``` json
    "color": {
-      "value": "#fff"
+      "value": "#000"
```

### move speed 3

``` json
    "move": {
-      "speed": 1,
-      "speed": 3,
```

### shape type circle

``` json
    "shape": {
-      "type": "edge",
+      "type": "circle",
```

### onclick mode remove

```
      "onclick": {
-        "mode": "push"
+        "mode": "remove"
```

# Contact me

:beetle: open an issue in github

:bird: [https://twitter.com/syui__](https://twitter.com/syui__)
