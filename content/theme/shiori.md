+++
screenshot = "/images/shiori.screenshot.png"
thumbnail = "/images/shiori.tn.png"
demo = "/theme/shiori/"
source = "https://github.com/chibicode/hugo-theme-shiori.git"
name = "Shiori"
license = "MIT"
licenselink = "https://github.com/chibicode/hugo-theme-shiori/blob/master/LICENSE.md"
description = "A Bootstrap-based Hugo theme."
homepage = "http://github.com/chibicode/hugo-theme-shiori"
tags = ["blog", "personal"]
features = ["blog"]
min_version= 0.13

[author]
    name =  "Shusaku Uesugi"
    homepage = "http://chibicode.com/"

[original]
    author = "Elle Kasai"
    homepage = "http://ellekasai.github.io/shiori/"
    repo = "https://github.com/ellekasai/shiori"

+++

# Shiori

Partial port of [Shiori](http://github.com/ellekasai/shiori) theme to [Hugo](http://gohugo.io).

![](images/snapshot.png)

## Changed from the Original

- To customize the theme, modify Sass files under [static_src](static_src) and run `gulp build`.
- Removed "you can customize this section" messages.
- Removed prev/post navigation on each post (can't be implemented until [this feature](https://github.com/spf13/hugo/issues/319#issuecomment-77797461) is implemented.

## Example Config

```toml
...

title = "Shiori"
author = "Shu Uesugi"

[params]
    author_name = "Shu Uesugi"
    twitter_username = "chibicode"
    gravatar_id = "d747ad92018acd87a77899704fc6693d"
```

## License

[MIT License](LICENSE.md)
