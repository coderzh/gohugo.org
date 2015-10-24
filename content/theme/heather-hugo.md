+++
screenshot = "/images/heather-hugo.screenshot.png"
thumbnail = "/images/heather-hugo.tn.png"
demo = "/theme/heather-hugo/"
source = "https://github.com/hbpasti/heather-hugo.git"
name = "Heather Hugo"
license = "MIT"
licenselink = "//opensource.org/licenses/MIT"
description = "A Hyperminimal J̶e̶k̶y̶l̶l̶ Hugo Theme"
homepage = "//hbpasti.github.io/heather-hugo"
tags = ["blog"]
features = ["blog"]

[author]
  name = "hbpasti"
  homepage = "//hbpasti.github.io"

[original]
  name = "Heather"
  homepage = "//jxnblk.github.io/Heather/"
  repo = "https://github.com/jxnblk/Heather/"

+++

# Heather for Hugo

A Hyperminimal J̶e̶k̶y̶l̶l̶ Hugo Theme

http://hbpasti.github.io/heather-hugo/

## About Heather

[Heather](//jxnblk.com/Heather/) is hyperminimal theme
for [Jekyll](//jekyllrb.com) created by Brent Jackson
([jxnblk](//jxnblk.com/)).

## About J̶e̶k̶y̶l̶l̶ Hugo

[Hugo](//gohugo.io) is a (very) fast static site generator written in Go.

## Get Started

Once you have Hugo set up, create your blog with `hugo new` and then
add the heather-hugo theme to the `themes` directory.

You can just download and extract it there or add it as a submodule
with:

    git submodule add https://github.com/hbpasti/heather-hugo themes/heather-hugo

Then edit your blog's config file to use heather-hugo theme:

- `config.toml`

    ``` toml
    theme = "heather-hugo"
    ```

- `config.yaml`

    ``` yaml
    theme: "heather-hugo"
    ```

A sample YAML config file would look like this:

    title: "Heather"
    baseurl: "http://localhost:1313"
    languageCode: en-us
    theme: heather-hugo

    permalinks:
      post: /:year/:month/:title/

    taxonomies:
      tags: ["meta", "theme", "blog"]

    params:
      description: A Hyperminimal J̶e̶k̶y̶l̶l̶ Hugo Theme
      author: 
        name: "Hbpasti"
        email: "your@email.com"
    ...

---

MIT License
http://opensource.org/licenses/MIT
