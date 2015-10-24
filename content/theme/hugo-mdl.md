+++
screenshot = "/images/hugo-mdl.screenshot.png"
thumbnail = "/images/hugo-mdl.tn.png"
demo = "/theme/hugo-mdl/"
source = "https://github.com/jchatkinson/HugoMDL.git"
# theme.toml for Hugo-MDL
# See http://jatkinson.me for an example

name = "HugoMDL"
license = "Apache 2.0"
licenselink = "https://github.com/jchatkinson/HugoMDL/blob/master/LICENSE.md"
description = "a simple and responsive theme to show your bio, blog, and projects - based on Google's Material Design Lite library"
homepage = "http://jatkinson.me"
tags = ["blog", "projects","portfolio","Material Design Lite","MDL","Material Design"]
features = ["blog", "project showcase", "portfolio", "profile","responsive","Material Design"]
min_version = 0.14

[author]
  name = "Jeremy Atkinson"
  homepage = "http://jatkinson.me"
  email = "jeremy@jatkinson.me"
+++

# HugoMDL

HugoMDL is a multipage responsibe blog/portfolio theme for [Hugo](http://gohugo.io). The theme is based on the
[Material Design Lite](http://www.getmdl.io/) sample blog layout. 

At present it suits my needs, but feel free to help improve this theme by forking a copy and submitting a pull request. 

![HugoMDL Homepage Screenshot](https://raw.githubusercontent.com/jchatkinson/HugoMDL/master/images/screenshot.png)

![HugoMDL Projects Screenshot](https://raw.githubusercontent.com/jchatkinson/HugoMDL/master/images/projects.png)

![ HugoMDL Post Screenshot](https://raw.githubusercontent.com/jchatkinson/HugoMDL/master/images/post.png)

## Installation

Inside the folder of your Hugo site run:

    $ mkdir themes
    $ cd themes
    $ git clone https://github.com/jchatkinson/HugoMDL

For more information read the official [setup guide](//gohugo.io/overview/installing/) of Hugo.

(Note for myself)
To develop on c9, add the hugo binary to the workspace, then add it to the path using `export PATH=$PATH:/home/ubuntu/workspace/hugo`
To serve the site from c9, cd to site root and use `hugo server --bind="0.0.0.0" --port=8080  --watch --disableLiveReload` 

## Configure the Theme

There are a few different places you need to configure

### Site-wide configuration

Look at the `exampleSite\config.toml` file for the relevent site parameters you need to fill out

### Background Image

Simply replace `\static\images\background.jpg` with a background of your choice.

### About page

Look at the `exampleSite\about.md` file for an example 'About' page. 

### Posts

Posts use the following front-matter:

```
---
title: "First Post"
description: "first post with Hugo website engine"
date: "2015-08-18"
categories:
    - "post"
tags:
    - "meta"
    - "test"
cardthumbimage: "/images/default.jpg" #optional: default solid color if unset
cardheaderimage: "/images/default.jpg" #optional: default solid color if unset
cardbackground: "#263238" #optional: card background color; only shows when no image specified
#cardtitlecolor: "#fafafa" #optional: can be changed to make text visible over card image
"author":
    name: "Firstname Lastname"
    description: "Writer of stuff"
    website: "http://example.com/"
    email: "firstname@example.com"
    twitter: "https://twitter.com/"
    github: "https://github.com/"
    image: "/images/avatar-64x64.png"
---
```

`cardthumbimage` is used as the card media on the front page. If unset, a solid color will be displayed.
`cardheaderimage` is used as the card media on the post page. If unset, a solid color will be displayed.
`cardbackground` is used as the solid color if no image is specified. If unset, the default theme color is displayed.
`cardtitlecolor` is used to modify the default white title color (for instance, if you have a light colored background) 
`author` parameters are used to populate the card and post with author's data 


### Projects

Projects have similar frontmatter to posts. 

```
---
title: "My Project Title"
description: "Description of the sample project" #optional
cardthumbimage: "/images/default.jpg" #optional: default solid color if unset
cardheaderimage: "/images/default.jpg" #optional: default solid color if unset
cardbackground: "#263238" #optional: card background color; only shows when no image specified
#cardtitlecolor: "#fafafa" #optional: can be changed to make text visible over card image
repo: "http://github.com/" #optional: no icon appears if unset
web: "http://github.com/" #optional: no icon appears if unset
date: "2015-08-18"
categories:
    - "project"
tags:
    - "meta"
    - "project"
"author": # used to fill out the project page
    name: "Firstname Lastname"
    description: "Writer of stuff"
    website: "http://example.com/"
    email: "firstname@example.com"
    twitter: "https://twitter.com/"
    github: "https://github.com/"
    image: "/images/avatar-64x64.png"

---
```

## License

This theme is released under the Apache License 2.0. It uses content from Google's material design lite project, also released under Apache License 2.0.
