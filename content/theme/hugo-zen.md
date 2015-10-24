+++
screenshot = "/images/hugo-zen.screenshot.png"
thumbnail = "/images/hugo-zen.tn.png"
demo = "/theme/hugo-zen/"
source = "https://github.com/rakuishi/hugo-zen.git"
name = "Hugo Zen"
license = "MIT"
licenselink = "https://github.com/rakuishi/hugo-zen/blob/master/LICENSE.md"
description = "Hugo Zen is a minimal hugo theme."
tags = ["blog"]
features = ["blog", "themes"]
min_version = 0.13

[author]
    name = "OCHIISHI Koichiro"
    homepage = "http://rakuishi.com/"

+++

# Hugo Zen

Hugo Zen is a minimal hugo theme with [Skeleton](https://github.com/dhg/Skeleton/) and has ~100 lines of custom CSS.

![screenshot](/images/screenshot.png)

## Installation & Usage

Clone this repository to your hugo theme directory.

	$ git clone https://github.com/rakuishi/hugo-zen.git themes/hugo-zen
	$ hugo server --theme=hugo-zen --buildDrafts --watch

## Configuration

In this theme you can add variables to your site config file. The following is the example config:

	baseurl = "http://rakuishi.com/"
	languageCode = "ja"
	title = "rakuishi.com"
	author = "rakuishi"

	[params]
	  logo      = "/images/logo.jpg"
	  copyright = "rakuishi All rights reserved."
	  twitter   = "https://twitter.com/rakuishi07"
	  facebook  = "https://www.facebook.com/ochiishikoichiro"
	  github    = "https://github.com/rakuishi/"
	  email     = "rakuishi@gmail.com"
	  analytics = "UA-12345678-9"

## License

MIT License

## Author

[OCHIISHI Koichiro](https://github.com/rakuishi)
