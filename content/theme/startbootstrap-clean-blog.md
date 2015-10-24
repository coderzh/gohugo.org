+++
screenshot = "/images/startbootstrap-clean-blog.screenshot.png"
thumbnail = "/images/startbootstrap-clean-blog.tn.png"
demo = "/theme/startbootstrap-clean-blog/"
source = "https://github.com/humboldtux/startbootstrap-clean-blog"
name = "Startbootstrap Clean Blog"
license = "Apache 2.0"
licenselink = "https://github.com/humboldtux/startbootstrap-clean-blog/blob/master/LICENSE.md"
description = "Startbootstrap Clean Blog"
homepage = "https://github.com/humboldtux/startbootstrap-clean-blog"
tags = ["blog", "bootstrap"]
features = ["blog", "bootstrap"]
min_version = 0.14

[author]
  name = "Benoît Benedetti"
  homepage = "http://www.humboldtux.com"

# original theme
[original]
  name = "Startbootstrap Clean Blog"
  homepage = "http://startbootstrap.com/template-overviews/clean-blog/"
  repo = "https://github.com/IronSummitMedia/startbootstrap-clean-blog"

+++

## About

Hugo theme based on [Start Bootstrap Clean Blog](http://startbootstrap.com/template-overviews/clean-blog/).

![Screenshot](https://raw.githubusercontent.com/humboldtux/startbootstrap-clean-blog/master/images/tn.png)

## Setup

### As is

You can use the site as is, it will work. You can for example use it with [HugoBasicExample](https://github.com/spf13/HugoBasicExample).

#### full features

If you want to have all the features and pages of the original theme, you will need a similar *config.toml*:

``` toml
baseurl = "http://humboldtux.github.io/sbcb-demo"
title = "Start Bootstrap Clean Blog"
canonifyurls = true
paginate = 10
theme = "startbootstrap-clean-blog"
languageCode = "en-us"
copyright = "Code released under the Apache 2.0 license."

[author]
  name = "Benoît Benedetti"
    
[params]
  DateForm = "Mon, Jan 2, 2006"
  Description = "Your site description"
  Author = "Benoît Benedetti"
  email = "benoit.benedetti@gmail.com"
  Ganalytics = "XX-00000000-0"
              
[[params.social]]
  title = "twitter"
  url = "https://twitter.com/humboldtux"
[[params.social]]
  title = "github"
  url = "https://github.com/humboldtux"
[[params.social]]
  title = "facebook"
  url = "https://www.facebook.com/FACEBOOKHANDLE"
                                      
[[menu.main]]
  name = "home"
  url = "/"
  weight = -200
[[menu.main]]
  name = "Archives"
  url = "/post/"
  weight = -180
```
                                                  
You can find a [Demo site](http://humboldtux.github.io/sbcb-demo/), and have a look at the source on the
[repo site](https://github.com/humboldtux/sbcb-demo) for configuration parameters.

### Custom

You can have a look at my [personal website](http://www.humboldtux.net), where i customize it and use it in a different way than the original theme.

## TODO

* multi-language support:
  * encoding
  * https://github.com/spf13/hugo/tree/master/examples/multilingual
  * http://discuss.gohugo.io/t/multiple-language-websites/537
  * http://discuss.gohugo.io/t/hugo-multisite-workflow/103
* "Edit on github" feature http://discuss.gohugo.io/t/how-to-implement-edit-this-page-on-github/1166
* Display archive posts per year/month
* dynamic configuration of email for JS contact form
* Advanced pagination, make pagination button look like older posts button
* pre and next links in post https://github.com/antzucaro/az.com/blob/master/layouts/_default/single.html
* Suggested links at the bottom of post
* use post title as post image name
* mailing list subscription form
* Google search form
 * http://discuss.gohugo.io/t/how-are-you-implementing-site-search/986
* multi author support, authors list page, author page
* use {{ .Content }} for 404.html
* deploy workflow
  * Makefile
  * http://discuss.gohugo.io/t/deployment-workflow/90
  * http://discuss.gohugo.io/t/my-deployment-process/807
  * http://discuss.gohugo.io/t/automated-hugo-deployments-anywhere-using-wercker/539
* updatedate post metadata for updated post
* comment system support: disqus, Discourse, Kaiju
* Syntax highlighting
* Less support
* post series : http://discuss.gohugo.io/t/creating-post-series/1115
* Add menu entry for taxonomies
* DRY template
