+++
screenshot = "/images/tinyce.screenshot.png"
thumbnail = "/images/tinyce.tn.png"
demo = "/theme/tinyce/"
source = "https://github.com/roperzh/tinyce-hugo-theme.git"
name = "tinyce"
license = "MIT"
licenselink = "https://github.com/roperzh/tinyce-hugo-theme/blob/master/LICENSE.md"
description = "Tinyce is a general purpose hugo theme."
homepage = "http://roperzh.github.io/tinyce-hugo-theme/"
tags = ["discuss", "mobile", "responsive", "flat", "blog"]
features = ["discuss", "mobile", "responsive", "flat", "blog"]
min_version = 0.14

[author]
    name = "roperzh"
    homepage = "http://roperzh.com/"
    repo = "https://github.com/roperzh/tinyce-hugo-theme"

+++

Tinyce
===

Tinyce is a general-purpose [Hugo](http://hugo.spf13.com) theme.

![screenshot - 08092014 - 07 02 07 pm](https://cloud.githubusercontent.com/assets/4419992/3867787/120f0158-2011-11e4-9f50-0ccfa162cdd2.png)

You can check the live version [here](http://roperzh.github.io/tinyce-hugo-theme/).

Getting started
===

Once you have Hugo up and running, you must:

Create a Hugo project:

```bash
$ hugo new site my-site
```

And then install tinyce:

```bash
$ cd themes
$ git clone https://github.com/roperzh/tinyce-hugo-theme
```

*If you are using Hugo v0.11 or below, please read [this](https://github.com/roperzh/tinyce-hugo-theme/#usage-with-hugo-v011-and-below)*

Now you can set tinyce as the default theme for your site by adding this line
in the `config` file:

```toml
theme = "tinyce-hugo-theme"
```

Or you can start the server with:

```bash
hugo server --theme=tinyce-hugo-theme --buildDrafts --watch
```

Setting Variables
===

Tinyce needs a custom archetype in order to work properly, behind the covers
you will only need to define two new variables in the front matter section of
your contents.

- **color** : the color associated with the post, valid colors are: `red`,
`green`, `blue`, `yellow`, `orange` and `grey`.

- **icon** : a valid font-awesome icon name, without the `fa-` prefix. For example,
instead of `fa-heart`, you only need to put `heart`. [here](http://fortawesome.github.io/Font-Awesome/cheatsheet/)
is a cheatsheet of valid icons.

***Example***

```toml
+++
Categories = ["Development", "GoLang"]
Description = "My short description"
Tags = ["Development", "golang"]
color = "green"
icon = "heart"
title = "Go is cool"
date = 2014-08-09T05:04:40Z
+++

```

Usage with Hugo v0.11 and below
===

The latest version of the theme includes some features bundled with Hugo v0.12,
if you are using an old version, please install tinyce like this:

```bash
$ cd themes
$ git clone https://github.com/roperzh/tinyce-hugo-theme
$ cd tinyce-hugo-theme
$ git checkout 08a497
```

That's it!

Contributing
===

1. Fork it

2. Create your feature branch (`git checkout -b my-new-feature`)

3. Commit your changes (`git commit -am 'Add some feature'`)

4. Push to the branch (`git push origin my-new-feature`)

5. Create new Pull Request
