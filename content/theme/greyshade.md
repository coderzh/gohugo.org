+++
screenshot = "/images/greyshade.screenshot.png"
thumbnail = "/images/greyshade.tn.png"
demo = "/theme/greyshade/"
source = "https://github.com/cxfksword/greyshade.git"
name = "Greyshade"
license = "MIT"
licenselink = "http://sm.mit-license.org/"
description = "Port of octopress's greyshade theme to Hugo"
homepage = "https://github.com/cxfksword/greyshade"
tags = ["blog", "octopress"]
features = ["blog"]
min_version = 0.14

[author]
  name = "cxfksword"
  homepage = "http://cxfksword.github.io"

# If porting an existing theme
[original]
  name = "Shashank Mehta"
  homepage = "http://github.com/shashankmehta/greyshade"
  repo = "https://github.com/shashankmehta/greyshade"

+++

Greyshade
============

[Greyshade](http://github.com/shashankmehta/greyshade) is a minimal, responsive theme ported from Octopress.

For demo, check out websites from the [wiki](https://github.com/shashankmehta/greyshade/wiki/Sites-using-Greyshade/)

![mobile view](/images/screenshot_mobile.png)

![screenshot](/images/screenshot.png)

### Features

1. Responsive design.
1. Fancybox integration
1. Gravatar support thanks to [Zhigang Fang](https://github.com/zhigang1992). Add your email id to `_config.yml` and your profile picture will appear in the left nav. If you don't want your email to be visible in the HTML code, you can provide the MD5 hashed value of your email as [expected](https://gravatar.com/site/implement/hash/) by Gravatar with an `email_md5` entry in `_config.yml`
1. [Schema.org](http://schema.org/) support thanks to [Nathan Shaughnessy](https://github.com/nathanshox)


### Setup

1. Fork or [download](https://github.com/cxfksword/greyshade/archive/master.zip) this theme repo
3. Add config to your site `config.toml` file

```
baseurl = "http://gohugo.io/"
languageCode = "en-us"
title = "your blog title"

[author]
name = "hugo"
# email will use for gravatar
email = ""

[taxonomies]
category = "categories"

[params]
# site description, will show under navigation
description = ""

# RSS / Email (optional) subscription links (change if using something like Feedburner)
subscribe_rss = "/index.xml"
subscribe_email = ""

# social links
facebook_user = ""
googleplus_user = ""
twitter_user = ""
github_user = ""
coderwall_user = ""
stackoverflow_user = ""
stackoverflow_user_id = ""
linkedin_user = ""
pinterest_user = ""
delicious_user = ""
pinboard_user = ""
quora_user = ""
instagram_user = ""
behance_user = ""
douban_user = ""

# share links
facebook_like = true
twitter_tweet_button = true
google_plus_one = "true"
addthis_profile_id = ""

# Disqus Comments
disqus_short_name = ""
disqus_show_comment_count = false

# google analytics
google_analytics_tracking_id = ""
```

### License
MIT: [http://sm.mit-license.org](http://sm.mit-license.org/)

### Author
**Shashank Mehtal**
- <https://github.com/shashankmehta>

### Ported By
**cxfksword**
- <https://github.com/cxfksword>
