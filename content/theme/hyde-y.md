+++
screenshot = "/images/hyde-y.screenshot.png"
thumbnail = "/images/hyde-y.tn.png"
demo = "/theme/hyde-y/"
source = "https://github.com/enten/hyde-y.git"
name = "Hyde-Y"
license = "MIT"
licenselink = "https://github.com/enten/hyde-y/LICENSE"
description = "An elegant open source and mobile first theme, extended with new integrations and layout improvements"
homepage = "https://github.com/enten/hyde-y"
tags = ["blog", "technical", "personal"]
features = ["blog", "technical", "personal"]
min_version = 0.14

[author]
    name = "Steven Enten"
    homepage = "http://enten.fr"

[original]
    author =  "Andrei Mihu"
    homepage = "https://github.com/zyro/hyde-x"
    repo = "https://github.com/zyro/hyde-x"

+++

Hyde-Y
======

<small>Forked from [Hyde-X](https://github.com/zyro/hyde-x)</small>

You can find a live site using this theme [here](http://enten.github.io/hugo-boilerplate/)
and the corresponding source code [here](https://github.com/enten/hugo-boilerplate).

## Screenshot

![preview](https://raw.githubusercontent.com/enten/hyde-y/master/images/screenshot.png)

## Installation

```
$ cd your_site_repo/
$ mkdir themes
$ cd themes
$ git clone https://github.com/enten/hyde-y
```

See the [official Hugo themes documentation](http://gohugo.io/themes/installing) for more info.

## Usage

This theme expects a relatively standard Hugo blog/personal site layout:
```
.
└── content
    ├── post
    |   ├── post1.md
    |   └── post2.md
    ├── code
    |   ├── project1.md
    |   ├── project2.md
    ├── license.md        // this is used in the sidebar footer link
    └── other_page.md
```

Just run `hugo --theme=hyde-y` to generate your site!

## Configuration

### Hugo

An example of what your site's `config.toml` could look like. All theme-specific parameters are under `[params]` and standard Hugo parameters are used where possible.

``` toml
# hostname (and path) to the root eg. http://spf13.com/
baseurl = "http://www.example.com"

# Site title
title = "sitename"

# Copyright
copyright = "(c) 2015 yourname."

# Language
languageCode = "en-EN"

# Metadata format
# "yaml", "toml", "json"
metaDataFormat = "yaml"

# Theme to use (located in /themes/THEMENAME/)
theme = "hyde-y"

# Pagination
paginate = 10
paginatePath = "page"

# Enable Disqus integration
disqusShortname = "your_disqus_shortname"

[permalinks]
    post = "/:year/:month/:day/:slug/"
    code = "/:slug/"

[taxonomies]
    tag = "tags"
    topic = "topics"

[author]
    name = "yourname"
    email = "yourname@example.com"

#
# All parameters below here are optional and can be mixed and matched.
#
[params]
    # You can use markdown here.
    brand = "foobar"
    topline = "few words about your site"
    footline = "code with <i class='fa fa-heart'></i>"

    # Sidebar position
    # false, true, "left", "right"
    sidebar = "left"

    # Text for the top menu link, which goes the root URL for the site.
    # Default (if omitted) is "Home".
    home = "home"

    # Select a syntax highight.
    # Check the static/css/highlight directory for options.
    highlight = "default"

    # Google Analytics.
    googleAnalytics = "Your Google Analytics tracking code"

    # Sidebar social links.
    github = "enten/hugo-boilerplate" # Your Github profile ID
    bitbucket = "" # Your Bitbucket profile ID
    linkedin = "" # Your LinkedIn profile ID (from public URL)
    googleplus = "" # Your Google+ profile ID
    facebook = "" # Your Facebook profile ID
    twitter = "" # Your Twitter profile ID
    youtube = ""  # Your Youtube channel ID
    flattr = ""  # populate with your flattr uid

[blackfriday]
    angledQuotes = true
    fractions = false
    hrefTargetBlank = false
    latexDashes = true
    plainIdAnchors = true
    extensions = []
    extensionmask = []

```

### Menu

Create `data/Menu.toml` to configure the sidebar navigation links. Example below.

```toml
[about]
    Name = "About"
    URL = "/about"

[posts]
    Name = "Posts"
    Title = "Show list of posts"
    URL = "/post"

[tags]
    Name = "Tags"
    Title = "Show list of tags"
    URL = "/tags"
```

### Foot menu

Create `data/FootMenu.toml` to configure the footer navigation links. Example below.

```toml
[license]
    Name = "license"
    URL = "/license"
```

## Built-in colour themes

You can rebuild the stylesheet. To do it you need [npm](http://npmjs.com/) and run the task `build:css`.

```bash
$ vi scss/_00-config.less
# edit configuration

$ npm install
$ npm run build:css

> hyde-y@0.0.4 build:css /home/steven/code/hyde-y
> grunt build:css

Running "less:development" (less) task
File static/css/style.css created

Running "cssmin:target" (cssmin) task
>> 1 file created. 27.04 kB → 14.38 kB

Done, without errors.
```

The task `watch` allows to rebuild the stylesheet when a change is spotted on `scss/*.less` files.

## Tips

* If you've added `theme = "hyde-y"` to your `config.toml`, you don't need to keep using the `--theme=hyde-y` flag!
* Although all of the syntax highlight CSS files under the theme's `static/css/highlight` are bundled with the site, only the one you choose will be included in the page and delivered to the browser.
* Change the favicon by providing your own as `static/favicon.png` (and `static/touch-icon-144-precomposed.png` for Apple devices) in your site directory.
* Hugo makes it easy to override theme layout and behaviour, read about it [here](http://gohugo.io/themes/customizing).
* Pagination is set to 10 items by default, change it by updating `paginate = 10` in your `config.toml`.

## Changes and enhancements from the original theme

* Highly customizable: every layout blocks are pieces of HTML code individually stored.
* Consistent hierarchy for the partials and templates files to assist the overrides of any layout block.
* Adjustable sidebar in config file.
* Project layout (badges, github ribbon) for posts stored in `content/code/` folder.
* CSS built with [KNACSS](http://knacss.com/) micro-framework.
* Client-side syntax highlighting through [highlight.js](https://highlightjs.org/), sane fallback if disabled or no JS - infinitely more flexible than the standard Hugo highlighting.
* Disqus integration: comment counts listed under blog entry names in post list, comments displayed at the bottom of each post.
* Google Analytics integration.
* Google Authorship metadata.
* Paginated blog listing.
* [FontAwesome](http://fortawesome.github.io/Font-Awesome) social links.
* ...many other small layout tweaks!

## Attribution

Obviously largely a port of the awesome [Hyde](https://github.com/poole/hyde) and [Hyde-X](https://github.com/zyro/hyde-x) themes.

## Questions, ideas, bugs, pull requests?

All feedback is welcome! Head over to the [issue tracker](https://github.com/enten/hyde-y/issues).

## License

Open sourced under the [MIT license](https://github.com/enten/hyde-y/blob/master/LICENSE).
