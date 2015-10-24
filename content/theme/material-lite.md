+++
screenshot = "/images/material-lite.screenshot.png"
thumbnail = "/images/material-lite.tn.png"
demo = "/theme/material-lite/"
source = "https://github.com/SamuelDebruyn/hugo-material-lite"
name = "Material Lite"
license = "Apache License 2.0"
licenselink = "https://github.com/SamuelDebruyn/hugo-material-lite/blob/master/LICENSE.md"
description = "A hugo theme based on Google's Material Design Lite"
homepage = "https://github.com/SamuelDebruyn/hugo-material-lite/"
tags = ["material", "minimalistic", "lite", "google"]
min_version = 0.14

[author]
  name = "Samuel Debruyn"
  homepage = "http://sa.muel.be"

# If porting an existing theme
[original]
  name = "Material Design Lite"
  homepage = "http://www.getmdl.io"
  repo = "https://github.com/google/material-design-lite"

+++

# Material Design Lite inspired theme for hugo

This is a theme for [hugo](http://gohugo.io), a static site generator. It is based on the [blog template](http://www.getmdl.io/templates/) of Google's [Material Design Lite](http://www.getmdl.io).

[![wercker status](https://app.wercker.com/status/242df4252d594182e3ccc5b3dac3205e/m/master "wercker status")](https://app.wercker.com/project/bykey/242df4252d594182e3ccc5b3dac3205e)

## License

Like the original template, this theme's license is Apache 2.0.

## Usage

### Content organization & features

This theme does currently **not support** lists, categories, tags...

This theme is optimized to use 2 content types: `post`s and `page`s. Pages are automatically displayed in the dropdown list beneath the author and posts are displayed on the homepage.

You can add a `description` and an `image` to a post or page using the front matter.

#### Some features

* posts can have featured images
* header contains tags for [Open Graph](http://ogp.me/) and [Twitter Cards](https://dev.twitter.com/cards/overview)
* links to profiles on Facebook, Twitter, Google Plus
* customizable share button
* pagination
* customizable background image/color
* customizable design (colors)
* responsive design
* code highlighting

### Configuration variables

These are the supported site parameters:

	TwitterUser = "YourTwitterUsername"
	FacebookUser = "YourFacebookUsername"
	GooglePlusUser = "YourGooglePlusUsername"
	copyright = "&copy; a short copyright statement"
	author = "John Doe"
	background = "a relative or absolute URL to a background image"
	bgcolor = "a background color; will be ignored if you use background"
	primary = "indigo"
	accent = "pink"

*Primary* and *accent* define the colors of the design. Check [the MDL customizer](http://www.getmdl.io/customize/index.html) to see the supported colors.

### Share button configuration

Create a file called *share.html* in *layouts/partials/* to override the configuration for the share button.

This is the default configuration. This button uses a plugin called *carrot* to share to different networks.  recommend keeping the `ui` part as is and only modifying the `networks` part. Check [this link](https://github.com/carrot/share-button/wiki/Configuration-Options) for detailed instructions.

	<script>
		// See https://github.com/carrot/share-button/wiki/Configuration-Options
		$(function() {
			var config =
				{
					ui: {
						flyout: "middle left",
						button_text: "<i class='material-icons'>share</i>"
					}
			};
			var share = new Share(".share", config);
		});
	</script>

### Header tags

If you want to add extra tags to the `<head>` of your page, you can create a file called *header-extra.html* in *layouts/partials/* and add your extra content to that file.

## Contributing

This theme is very basic and I welcome all contributions and feedback. Please create an issue or a pull request to contribute.

## Demo

### Code

https://github.com/SamuelDebruyn/hugo-theme-example

### Website

http://materialexample.sa.muel.be/
