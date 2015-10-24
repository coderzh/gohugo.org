+++
screenshot = "/images/slim.screenshot.png"
thumbnail = "/images/slim.tn.png"
demo = "/theme/slim/"
source = "https://github.com/zhe/hugo-theme-slim.git"
name = "Slim"
license = "MIT"
licenselink = "https://github.com/zhe/hugo-theme-slim/blob/master/LICENSE.md"
description = "A slim minimal theme for Hugo."
homepage = "https://github.com/zhe/hugo-theme-slim"
tags = ["minimal", "clean", "neat", "black white", "blog"]
features = ["blog"]
min_version = 0.13

[author]
    name = "Zhe Zhang"
    homepage = "http://zhezhang.co"

+++

# Slim

Slim is a minimal, clean and beautiful theme for [Hugo](http://gohugo.io/).

![Slim screenshot](https://raw.githubusercontent.com/zhe/hugo-theme-slim/master/images/screenshot.png)

[Demo](http://zhezhang.co).

## Installation

```
mkdir themes
cd themes
git clone https://github.com/zhe/hugo-theme-slim slim
```

See the [official docs](http://gohugo.io/themes/installing) for more information.

## Configuration
You could add `params` into your site's `config.toml` file:

```
[params]
  Subtitle = "Your site's subtitle/tagline"
  GithubID = "Your Github ID"
  TwitterID = "Your Twitter ID"
  AnalyticsID = "Your Google Analytics tracking code"
  Summary = true  # takes true or false
  Content = false  # takes true or false
  # if both are set to true, summary is shown.
```

if you use `config.yaml`, it could look like:

```
params:
  Subtitle: "Your site's subtitle/tagline"
  GithubID: "Your Github ID"
  TwitterID: "Your Twitter ID"
  AnalyticsID: "Your Google Analytics tracking code"
  Summary: true # takes true or false
  Content: false # takes true or false
  # if both are set to true, summary is shown
```

## Build your site

```
hugo server -t slim
```


## License

Open sourced under [MIT license](https://github.com/zhe/hugo-theme-slim/blob/master/LICENSE.md).
