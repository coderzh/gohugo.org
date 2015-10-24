+++
screenshot = "/images/grid-side.screenshot.png"
thumbnail = "/images/grid-side.tn.png"
demo = "/theme/grid-side/"
source = "https://github.com/chipsenkbeil/grid-side.git"
name = "GridSide"
license = "MIT"
licenselink = "https://github.com/chipsenkbeil/grid-side/blob/master/LICENSE.md"
description = "Personal portfolio, blog, and gallery for use by the Hugo generator."
homepage = "https://github.com/chipsenkbeil/grid-side"
tags = ["blog", "portfolio", "projects"]
features = ["blog", "portfolio", "gallery"]
min_version = 0.14

[author]
    name = "Chip Senkbeil"
    homepage = "http://chipsenkbeil.com/"


+++

GridSide Theme
==============

The GridSide theme is a multi-page portfolio and blog utilizing the
[Materialize][materialize] frontend framework. Comments can be added using
Disqus. 

The theme contains a main page with a single grid of images representing
different sections of the website. Sections can be _post_, _gallery_, or
_project_ for various rendering.

- Current Materialize version is `0.97.0`.
- Current Font Awesome version is `4.4.0`.
- Current SideComments.js version is `0.0.3`.
- Current Masonry.js version is `3.3.1`.
- Current imagesLoaded.js version is `3.1.8`.
- Current Modernizr.js version is `2.8.3`.
- Current Highlight.js version is `8.7` and contains all 130 languages.
- Current lazysizes.js version is `1.2.0`.
- Current ls.noscript.js version is `1.2.0`.
- Current lightbox.js version is `2.8.1`.
- Current infinitescroll.js version is `2.1.0`.

![GridSide Theme Screenshot](https://raw.githubusercontent.com/chipsenkbeil/grid-side/master/images/screenshot.png)

Contents
--------

- [Installation](#installation)
- [Getting Started](#getting-started)
    - [The Config File](#the-config-file) 
    - [Setting the homepage header](#setting-the-homepage-header)
    - [Setting the homepage footer](#setting-the-homepage-footer)
    - [Adding homepage cells](#adding-homepage-cells)
    - [Specifying the main menu](#specifying-the-main-menu)
    - [Creating a post](#creating-a-post)
    - [Creating a project page](#creating-a-project-page)
    - [Creating a gallery image](#creating-a-gallery-image)
    - [Adding a custom post section](#adding-a-custom-post-section)
    - [Adding a custom project section](#adding-a-custom-project-section)
    - [Adding a custom gallery section](#adding-a-custom-gallery-section)
    - [Nearly Finished](#nearly-finished)
- [Contributing](#contributing)
- [License](#license)

Installation
------------

Inside the folder of your Hugo site run:

    $ mkdir themes
    $ cd themes
    $ git clone https://github.com/chipsenkbeil/grid-side

For more information read the official [setup guide][setup_guide] of Hugo.

Getting Started
---------------

### The Config File ###

Take a look inside the [`exampleSite`][exampleSite] folder of this theme.
You'll find a file called [`config.toml`][config.toml]. To use it, copy the
[`config.toml`][config.toml] in the root folder of your Hugo site. The config
file contains detailed explanation of each available property. Feel free
to customize this theme as you like.

### Setting the header ###

The header of the homepage serves as the main attraction for visitors. You
can provide your name, a description, and your email address on top of an
image.

```toml
[Params.Header]

    name = "Your name"
    description = "Your description"
    email = "your@email.com"
    image = "path/to/your/header/image"
```

Each field in the header is optional, meaning that you can choose to not
provide a name, description, or email. There are additional options you can use
in the header section as well! For more information, check out the
example [`config.toml`][config.toml].

### Setting the homepage footer ###

The footer of the homepage serves to provide contact links and other
miscellaneous connections from your main website. Each entry is composed of a
font awesome icon and a url. The icon is specified via `icon_class` and
represents the specific font awesome icon. E.g. `twitter` becomes
`fa fa-twitter`. The url is specified via `icon_link`.

```toml
[Params.Footer]

    [[Params.Footer.List]]

        icon_class = "twitter"
        icon_link = "https://www.twitter.com/..."

    [[Params.Footer.List]]

        icon_class = "github"
        icon_link = "https://www.github.com/..."
```

Each field in the header is optional, meaning that you can choose to not
provide a name, description, or email. There are additional options you can use
in the header section as well! For more information, check out the
example [`config.toml`][config.toml].

### Adding homepage cells ###

The other main aspect of the homepage is the homepage cells, or the grid of
images below your header. Each cell contains an image depicting its content
along with a title that is visible upon hovering over it. Each cell acts as a
hyperlink to other content on your site.

```toml
[Params.Cells]

    [[Params.Cells.List]]

        name = "Section name"
        link = "/some/path/on/your/site"
        image = "/some/image/file"
```

There are additional options you can use for each cell as well! For more
information, check out the example [`config.toml`][config.toml].

### Specifying the main menu ###

The main menu used on all list templates is specified via the menu name,
"Main". The fastest way to set your menu is to specify the `SectionPagesMenu`
option at the root of your config.

```toml
SectionPagesMenu = "Main"
```

### Creating a post ###

Each post in a post section should have the following front matter:

```toml
+++
title = "Post title"
date = "YYYY-MM-DD"
tags = [ "tag1", "tag2", ... ]
categories = [ "category1", ... ]
+++
```

Additionally, you can specify an image via the front matter `image = "url"`.

### Creating a project page ###

Each project in a project section should have the following front matter:

```toml
+++
title = "Project title"
tags = [ "tag1", "tag2", ... ]
categories = [ "category1", ... ]
+++
```

Furthermore, an image should be provided via `image = "url"` for more visual
effect. If one is not provided, a placeholder will be used.

Additionally, you can provide videos as an alternative to an image, which will
be displayed using the HTML5 video tag. Each of the following is optional:

```toml
+++
video_mp4 = "/path/to/mp4"
video_webm = "/path/to/webm"
video_ogv = "/path/to/ogv"
video_3gp = "/path/to/3gp"
video_fallback = "/path/to/image"
+++
```

The video fallback option is used both as the poster of the loading HTML5 video
and as the filler image if HTML5 video is not supported by the browser.

### Creating a gallery image ###

Each gallery image needs to be specified through the front matter of an
individual content item.

```toml
+++
title = "Image title used in lightbox"
date = "YYYY-MM-DD"
image = "/path/to/image"
+++
```

### Adding a custom post section ###

By default, the theme provides a custom view of `post/`. If you would like
to have a different section name than post, you can specify the section by
creating the following:

```
For layouts/custom_post_section/single.html:

{{ partial "post/single.html" . }}

```

```
For layouts/section/custom_post_section.html:

{{ partial "post/list.html" . }}

```

### Adding a custom project section ###

By default, the theme provides a custom view of `project/`. If you would like
to have a different section name than project, you can specify the section by
creating the following:

```
For layouts/custom_project_section/single.html:

{{ partial "project/single.html" . }}

```

```
For layouts/section/custom_project_section.html:

{{ partial "project/list.html" . }}

```

### Adding a custom gallery section ###

By default, the theme provides a custom view of `gallery/`. If you would like
to have a different section name than gallery, you can specify the section by
creating the following:

```
For layouts/custom_gallery_section/single.html:

{{ partial "gallery/single.html" . }}

```

```
For layouts/section/custom_gallery_section.html:

{{ partial "gallery/list.html" . }}

```

### Nearly Finished ###

In order to see your site in action, run Hugo's built-in local server. 

    $ hugo server -w

Now enter `localhost:1313` in the address bar of your browser.

Contributing
------------

Report any bugs using the [issue tracker][issue_tracker]. Submit your own bug
fixes or feature additions via a [pull request][pull_request].

License
-------

This theme is released under the MIT License. For more information read the
[license][license].

[materialize]: http://www.materializecss.com/
[setup_guide]: http://gohugo.io/overview/installing/
[exampleSite]: https://github.com/chipsenkbeil/grid-side/tree/master/exampleSite
[config.toml]: https://github.com/chipsenkbeil/grid-side/blob/master/exampleSite/config.toml
[issue_tracker]: https://github.com/chipsenkbeil/grid-side/issues
[pull_request]: https://github.com/chipsenkbeil/grid-side/pulls
[license]: https://github.com/chipsenkbeil/grid-side/blob/master/LICENSE
