+++
screenshot = "/images/hugo-uno.screenshot.png"
thumbnail = "/images/hugo-uno.tn.png"
demo = "/theme/hugo-uno/"
source = "https://github.com/SenjinDarashiva/hugo-uno.git"
name = "Hugo-Uno"
license = "Creative Commons Attribution 4.0 International"
licenselink = "https://creativecommons.org/licenses/by/4.0/"
description = "An elegant open source and mobile first theme"
homepage = "https://github.com/SenjinDarashiva/hugo-uno"
tags = ["Responsive", "Awesome fonts", "Chart.js", "Lightbox", "gallery"]
features = ["blog", ]

[author]
    name = "Fredrik Loch"
    homepage = "http://fredrikloch.me"

[original]
    author =  "daleanthony"
    homepage = "http://daleanthony.com/"
    repo = "https://github.com/daleanthony/Uno"

+++

hugo-uno
========

A responsive hugo theme with awesome font's, charts and light-box galleries, the theme is based on [Uno](https://github.com/daleanthony/Uno) for ghost.
A example site is available at [fredrikloch.me](http://fredrikloch.me)

A Swedish translation is available in the branch feature/swedish

## Usage
The following is a short tutorial on the usage of some features in the theme.
Configuration
-

To take full advantage of the features in this theme you can add variables to your site config file, the following is the example config from the example site:
```
languageCode = "en-us"
contentdir = "content"
publishdir = "public"
builddrafts = false
baseurl = "http://fredrikloch.me/"
canonifyurls = true
title = "Fredrik Loch"
author = "Fredrik Loch"
copyright = "This work is licensed under a Creative Commons Attribution-ShareAlike 4.0 International License."


[indexes]
   category = "categories"
   tag = "tags"
[Params]
  AuthorName = "Fredrik"
  github = "Senjindarashiva"
  bitbucket = "floch"
  flickr = "senjin"
  twitter = "senjindarshiva"
  email = "fredrik.loch@outlook.com"
  description = ""
  cv = "/pages/cv"
  muut = "fredrikloch"
  linkedin = "fredrikloch"
  cover = "/images/background-cover.jpg"
  logo = "/img/logo-1.jpg"
```

If you prefer to use discourse replace the "muut" line with the following(remember the trailing slash)

'''
  discourse = "http://discourse.yoursite.com/"
'''
Charts
-
To create charts I use [Chart.js](https://github.com/nnnick/Chart.js) which can be configured through basic js files. To add a chart to a post use the following short-code:
```
{{%/* chart id="basicChart" width=860 height=400 js="../../js/chartData.js" */%}}
```
Where the javascript file specified contains the data for the chart, a basic example could look like this:
```
$(function(){
  var chartData = {
      labels: ["Jekyll", "Hugo", "Wintersmith"],
      datasets: [
          {
              label: "Mean build time",
              fillColor: "#E1EDD7",
              strokeColor: "#E1EDD7",
              highlightFill: "#C1D8AB",
              highlightStroke: "#C1D8AB",
              data: [784, 100, 5255]
          }
      ]
  };
  var ctx = $("#chartData").get(0).getContext("2d");
  var myBarChart = new Chart(ctx).Bar(data1, {
      scaleBeginAtZero : true,
      responsive: true,
      maintainAspectRatio: false,
    }
  );
```
A running example can be found in my comparison between [Jekyll, Hugo and Winthersmith](http://fredrikloch.me/post/2014-08-12-Jekyll-and-its-alternatives-from-a-site-generation-point-of-view/)
Gallery
-
To add a gallery to the site we use basic html together with [lightGallery](http://sachinchoolur.github.io/lightGallery/index.html) to create a responsive light-box gallery.
```
<ul style="list-style: none;" id="lightGallery">
    <li data-src="pathToImg.jpg">
        <img src="pathToThumb.jpg"></img>
    </li>
    <li data-src="pathToImg.jpg">
        <img src="pathToThumb.jpg"></img>
    </li>
</ul>

<script src=../../js/lightGallery.min.js></script>
<script>
    $("#lightGallery").lightGallery();
</script>
```
A running example can be found in my short [review of hugo](http://fredrikloch.me/post/Lightbox image's and a short review of hugo/)
## Features

**Cover page**
The landing page for Hugo-Uno is a full screen 'cover' featuring your avatar, blog title, mini-bio and cover image.

**Built with SASS, using BEM**
If you know HTML and CSS making modifications to the theme should be super simple.

**Responsive**
Hugo-Uno looks great on all devices, even those weird phablets that nobody buys.

**Moot comments**
Moot integration allows users to comment on your posts.

**Font-awesome icons**
For more information on available icons: [font-awesome](http://fortawesome.github.io/Font-Awesome/)

**No-JS fallback**
While JS is widely used, some themes and websites don't provide fallback for when no JS is available (I'm looking at you [Squarespace](http://blog.squarespace.com/)). If for some weird reason a visitor has JS disabled your blog will still be usable.

## License
[Creative Commons Attribution 4.0 International](http://creativecommons.org/licenses/by/4.0/)

## Development

In order to develop or make changes to the theme you will need to have the sass compiler and bourbon both installed.

To check installation run the following commands from a terminal and you should see the `> cli output` but your version numbers may vary.

** SASS **
```bash

sass -v
> Sass 3.3.4 (Maptastic Maple)
```
If for some reason SASS isn't installed follow the instructions from the [Sass install page](http://sass-lang.com/install)

** Bourbon **
```bash

bourbon help
> Bourbon 3.1.8
```
If Bourbon isn't installed follow the installation instructions on the [Bourbon website](http://bourbon.io)

Once installation is verified we will need to go mount the bourbon mixins into the `scss` folder.

From the project root run `bourbon install` with the correct path
```bash
bourbon install --path assets/scss
> bourbon files installed to assets/scss/bourbon/
```

Now that we have the bourbon mixins inside of the `scss` src folder we can now use the sass cli command to watch the scss files for changes and recompile them.

```bash
sass --watch assets/scss:assets/css
>>>> Sass is watching for changes. Press Ctrl-C to stop.
```

To minify the css files use the following command in the assets folder

```bash
curl -X POST -s --data-urlencode 'input@css/uno.css' http://cssminifier.com/raw > css/uno.min.css
```
