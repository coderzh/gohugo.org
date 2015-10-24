+++
screenshot = "/images/angels-ladder.screenshot.png"
thumbnail = "/images/angels-ladder.tn.png"
demo = "/theme/angels-ladder/"
source = "https://github.com/tanksuzuki/angels-ladder.git"
name = "Angel's Ladder"
license = "MIT"
licenselink = "https://github.com/tanksuzuki/angels-ladder/blob/master/LICENSE.md"
description = "Angel's Ladder is a simple blog theme for Hugo."
homepage = "http://tanksuzuki.com/"
tags = ["blog"]
features = ["blog"]
min_version = 0.13

[author]
    name = "tanksuzuki"
    homepage = "http://tanksuzuki.com/"

+++

# Angel's Ladder

## Overview

Angel's Ladder is a simple blog theme for Hugo.

* Simple and clean design
* Responsive design
* Pagination
* Tagging
* Disqus
* Source code highlighting
* Google Analytics

## Screenshot

All of the display [here](https://github.com/tanksuzuki/angels-ladder/tree/master/images/tn_full.png).

![](https://raw.githubusercontent.com/tanksuzuki/angels-ladder/master/images/readme.png?token=ALXg6utVKJxUhSjlO_-voAEgA75oYTseks5VXtd1wA%3D%3D)  


## Installation

Clone this repository to your hugo theme directory.

```
cd themes
git clone https://github.com/tanksuzuki/angels-ladder
hugo server -t angels-ladder -D -w
```


## Configuration

To take full advantage of the features in this theme, you can add variables to your site config file.

The following is the example configuration.

```
baseurl = "http://tanksuzuki.com/"
languageCode = "ja"
title = "TANKSUZUKI.COM"
disqusShortname = "tanksuzuki"

[Params]
subtitle = "I would like to be a layer 3 switch."
facebook = "https://facebook.com/foobar"
twitter = "https://twitter.com/foobar"
github = "https://github.com/foobar"
profile = "/images/profile.png"
copyright = "Written by Asuka Suzuki"
analytics = "UA-XXXXXXXX-X"
```

Details of each parameter are as follows.

| Parameter | Required | Comment |
| :--- | :--- | :--- |
| baseurl | yes | Enter the title of your site. |
| languageCode | yes | Enter the language code of HTML. Example: en, ja. |
| title | yes | Enter the title of your site. |
| disqusShortname | no | Enter the short name of the disqus. If you do not enter, disqus section is hidden. |
| subtitle | no | Enter the subtitle of your site. If you do not enter, subtitle is hidden. |
| facebook | no | Enter the URL of Facebook. If you do not enter, the link is hidden. |
| twitter | no | Enter the URL of Twitter. If you do not enter, the link is hidden. |
| github | no | Enter the URL of Github. If you do not enter, the link is hidden. |
| profile | no | Enter the path to the profile image. If you do not enter, profile section will be hidden. |
| copyright | no | Enter the copyright notice. If you do not enter, copyright display is hidden. |
| analytics | no | Enter the tracking ID of Google analytics. If you do not enter, the analysis will be skipped. |


## Style Customization

The default theme color is `#29abe2`.
Please change the favorite color.

For original styles, please edit the `layouts/static/custom.css`.


## License

Open sourced under the [MIT license](https://github.com/tanksuzuki/angels-ladder/blob/master/LICENSE.md).


## Author

Asuka Suzuki


## Contact

Please contact me via [email](https://github.com/tanksuzuki) / [Twitter](https://twitter.com/tanksuzuki) :smile:
