+++
screenshot = "/images/tachyons.screenshot.png"
thumbnail = "/images/tachyons.tn.png"
demo = "/theme/tachyons/"
source = "https://github.com/marloncabrera/tachyons.git"
author = "Marlon Cabrera Oliveira"
description = "Hugo theme using tachyons.io css3 framework & geomicons.com icons"
license = "MIT"
name = "tachyons"
source_repo = "https://github.com/marloncabrera/tachyons.git"
tags = ["blog", "personal"]

+++

Tachyons
==========

Tachyons is a responsive & fast [Hugo](http://gohugo.io) theme made with the amazing [TACHYONS](http://www.tachyons.io) CSS framework & [GEOMICONS](http://geomicons.com).

Please see a demo site [here](http://marloncabrera.github.io/tachyons) and another example [here](http://nolram.com).

![Tachyons screenshot](https://raw.githubusercontent.com/marloncabrera/tachyons/master/images/tn.png)

## Contents

- [Usage](#usage)
- [Options](#options)
  - [Header](#header)
  - [Footer](#footer)
  - [Icon Reference Table](#icon-reference-table)
- [Author](#author)
- [Credits](#credits)
- [License](#license)

## Usage

From [Hugo Quickstart](http://http://gohugo.io/overview/quickstart/):

* **Install Hugo**

Go to [Hugo Releases](https://github.com/spf13/hugo/releases) and download the appropriate version for your OS and architecture.
Save the main executable as *hugo* (or *hugo.exe* on Windows) somewhere in your PATH as we will be using it in the next step.
More complete instructions are available at [Installing Hugo](http://gohugo.io/overview/installing/).

* **Create a new site & install Tachyon Theme**

```
$ hugo new site /path/to/site
$ cd /path/to/site
$ mkdir themes
$ cd themes
$ git clone https://github.com/marloncabrera/tachyons.git
```

See [Options](#options) below to customize your blog.

* **Create Some Content**

```
$ hugo new post/first.md
```

* **Run Hugo and see the results**

```
$ hugo server --theme=tachyons --buildDrafts --watch
```

Open your browser at [http://localhost:1313](http://localhost:1313) to see the results.

## Options

Customize your blog editing the header & footer under themes/tachyons/layouts/partials/ directory.


### Header

Replace "Your_blog_name_here" to change your new Blog name.
If you need to create more sections, you have to add the section name to the nav menu like the example below, then pick up a specific icon from [Icon Reference Table](#icon-reference-table).


```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    ...
    ...
  </head>
  <body>
      <header class="tc pal bb b-near-white">
    <h1 class="f6 thin i link-child"><a href="/">Your_blog_name_here</a></h1>
        <nav>
          <ul class="list pln">
            <li class="dib prm"><a href="/"><span class="js-geomicon geomicon" data-icon="home"></span>Home</a></li>
            <li class="dib prm"><a href="/post"><span class="js-geomicon geomicon" data-icon="compose"></span>Blog</a></li>
            <li class="dib prm"><a href="/about"><span class="js-geomicon geomicon" data-icon="user"></span>About</a></li>
          </ul>
        </nav>
        </header>
```

### Footer

On footer.html you can customize or remove the links to Github or Twitter and include Google Analytics javascript code between the ```<script> ...</script>``` tags along with Geomicons portion.


```html
footer class="tc center bt b--near-white pvx phm pax-m phxl-l pvx-l">
<small class="f5 mw7 db center phm lh-copy">
<nav>
<ul class="list pln">
<li class="dib prm"><a href="https://www.github.com"><span class="js-geomicon geomicon blue" data-icon="github"></span></a></li>
<li class="dib prm"><a href="https://www.twitter.com"><span class="js-geomicon geomicon blue" data-icon="twitter"></span></a></li>
</ul>
</nav>
Copyright-  2015
</small>
</footer>
<script src="http://d2v52k3cl9vedd.cloudfront.net/geomicons/0.2.0/geomicons.min.js.gz"></script>
<script>
var icons = document.querySelectorAll('.geomicon');
Geomicons.inject(icons);
</script>
</body>
</html>
```


### Icon Reference Table



Icon    | ID
--------|--------
![bookmark](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/bookmark.svg) | bookmark
![calendar](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/calendar.svg) | calendar
![camera](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/camera.svg) | camera
![chat](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/chat.svg) | chat
![check](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/check.svg) | check
![chevron-down](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/chevronDown.svg) | chevron-down
![chevron-left](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/chevronLeft.svg) | chevron-left
![chevron-right](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/chevronRight.svg) | chevron-right
![chevron-up](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/chevronUp.svg) | chevron-up
![clock](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/clock.svg) | clock
![close](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/close.svg) | close
![cloud](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/cloud.svg) | cloud
![cog](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/cog.svg) | cog
![compose](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/compose.svg) | compose
![dribbble](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/dribbble.svg) | dribbble
![expand](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/expand.svg) | expand
![external](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/external.svg) | external
![facebook](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/facebook.svg) | facebook
![file](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/file.svg) | file
![folder](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/folder.svg) | folder
![geolocation](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/geolocation.svg) | geolocation
![github](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/github.svg) | github
![grid](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/grid.svg) | grid
![heart](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/heart.svg) | heart
![home](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/home.svg) | home
![info](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/info.svg) | info
![link](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/link.svg) | link
![list](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/list.svg) | list
![lock](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/lock.svg) | lock
![mail](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/mail.svg) | mail
![music-note](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/musicNote.svg) | music-note
![next](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/next.svg) | next
![no](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/no.svg) | no
![pause](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/pause.svg) | pause
![picture](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/picture.svg) | picture
![pin](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/pin.svg) | pin
![play](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/play.svg) | play
![previous](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/previous.svg) | previous
![refresh](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/refresh.svg) | refresh
![repost](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/repost.svg) | repost
![search](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/search.svg) | search
![shopping-cart](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/shoppingCart.svg) | shopping-cart
![skull](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/skull.svg) | skull
![speaker-volume](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/speakerVolume.svg) | speaker-volume
![speaker](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/speaker.svg) | speaker
![star](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/star.svg) | star
![tag](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/tag.svg) | tag
![trash](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/trash.svg) | trash
![triangle-down](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/triangleDown.svg) | triangle-down
![triangle-left](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/triangleLeft.svg) | triangle-left
![triangle-right](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/triangleRight.svg) | triangle-right
![triangle-up](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/triangleUp.svg) | triangle-up
![twitter](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/twitter.svg) | twitter
![user](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/user.svg) | user
![video](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/video.svg) | video
![warning](https://cdn.rawgit.com/jxnblk/geomicons-open/master/dist/icons/warning.svg) | warning


## Author
**Marlon Cabrera Oliveira**
- <http://nolram.com>
- <http://github.com/marloncabrera>
- <http://twitter.com/marloncabrera>

## Credits
**Steve Francia**
- <http://gohugo.io>
- <https://github.com/spf13>
- <https://twitter.com/spf13>

**Adam Morse**
- <http://tachyons.io>
- <http://github.com/mrmrs>
- <http://twitter.com/mrmrs_>

**Brent Jackson**
- <http://geomicons.com>
- <http://github.com/jxnblk>
- <http://twitter.com/jxnblk>


##License

Open sourced under the [MIT license](license.md).
