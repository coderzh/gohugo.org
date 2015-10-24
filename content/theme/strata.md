+++
screenshot = "/images/strata.screenshot.png"
thumbnail = "/images/strata.tn.png"
demo = "/theme/strata/"
source = "https://github.com/digitalcraftsman/hugo-strata-theme.git"
name = "Strata"
license = "Creative Commons Attribution 3.0 Unported"
licenselink = "//github.com/digitalcraftsman/hugo-strata-theme/blob/master/LICENSE.md"
description = "A responsive and minimal single-page portfolio."
homepage = "//github.com/digitalcraftsman/hugo-strata-theme"
tags = ["image gallery", "contact form", "google analytics", "customizable"]
features = ["", ""]
min_version = 0.14

[author]
  name = "digitalcraftsman"
  homepage = "//github.com/digitalcraftsman"

# If porting an existing theme
[original]
  name = "AJ"
  homepage = "//html5up.net/"
  repo = ""

+++

# Strata Theme

The Strata theme is a responsive and minimal one-page-portfolio based on the self-named template made by [HTML5UP](//html5up.net/). Noteworthy features of this Hugo theme are a custom about section, a portfolio with gallery for photographs or client works and a contact form.

![Screenshot](https://raw.githubusercontent.com/digitalcraftsman/hugo-strata-theme/dev/images/screenshot.png)


## Contents

- [Installation](#installation)
- [Getting started](#getting-started)
    - [The config file](#the-config-file) 
    - [Sidebar](#sidebar)
    - [Build up your portfolio](#build-up-your-portfolio)
    - [Make the contact form working](#make-the-contact-form-working)
    - [Nearly finished](#nearly-finished)
- [Contributing](#contributing)
- [License](#license)
- [Annotations](#annotations)


## Installation

Inside the folder of your Hugo site run:

    $ mkdir themes
    $ cd themes
    $ git clone git@github.com:digitalcraftsman/hugo-strata-theme.git

For more information read the official [setup guide](//gohugo.io/overview/installing/) of Hugo.


## Getting started

After installing the Strata Theme successfully it requires a just a few more steps to get your site finally running.


### The config file

Take a look inside the [`exampleSite`](https://github.com/digitalcraftsman/hugo-strata-theme/tree/dev/exampleSite) folder of this theme. You'll find a file called [`config.toml`](//github.com/digitalcraftsman/hugo-strata-theme/blob/dev/exampleSite/config.toml). To use it, copy the [`config.toml`](//github.com/digitalcraftsman/hugo-strata-theme/blob/dev/exampleSite/config.toml) in the root folder of your Hugo site. Feel free to customize this theme as you like.

### Sidebar

The sidebar provides a small overview of who you are. One of the first elements that will be spottet is the avatar in the sidebar. Replace it with a nice picture of you by either swapping the [`avatar.jpg`](https://github.com/digitalcraftsman/hugo-strata-theme/blob/dev/static/images/avatar.jpg) or by setting a new path to an image in [`config.toml`](//github.com/digitalcraftsman/hugo-strata-theme/blob/dev/exampleSite/config.toml):

```toml
[params.sidebar]
    avatar = "avatar.jpg"
```

The path is relative to [`./static/images`](https://github.com/digitalcraftsman/hugo-strata-theme/tree/dev/static/images).

Last but not least add a few words about you and your work.


### Build up your portfolio

As photograph or freelancer, your most important piece in the resume is the work you've done. Within the [`config.toml`](//github.com/digitalcraftsman/hugo-strata-theme/blob/dev/exampleSite/config.toml) add the following snippet to add a new item to the gallery:

```toml
[params.portfolio]

        # The images and thumbnails are stored under static/images
        # Create and change subfolders as you like
        [[params.portfolio.gallery]]
            image = "fulls/01.jpg"
            thumb = "thumbs/01.jpg"
            title = "Lorem ipsum dolor."
            description = "Lorem ipsum dolor sit amet."
```

### Make the contact form working

Since this page will be static, you can use [formspree.io](//formspree.io/) as proxy to send the actual email. Each month, visitors can send you up to one thousand emails without incurring extra charges. Begin the setup by following the steps below:

1. Enter your email address under 'email' in the [`config.toml`](//github.com/digitalcraftsman/hugo-strata-theme/blob/dev/exampleSite/config.toml)
2. Upload the generated site to your server
3. Send a dummy email yourself to confirm your account
4. Click the confirm link in the email from [formspree.io](//formspree.io/)
5. You're done. Happy mailing!


### Nearly finished

In order to see your site in action, run Hugo's built-in local server. 

    $ hugo server -w

Now enter [`localhost:1313`](//localhost:1313) in the address bar of your browser.


## Contributing

Did you found a bug or got an idea for a new feature? Feel free to use the [issue tracker](//github.com/digitalcraftsman/hugo-strata-theme/issues) to let me know. Or make directly a [pull request](//github.com/digitalcraftsman/hugo-strata-theme/pulls).


## License

This theme is released under the Creative Commons Attribution 3.0 Unported  License. For more information read the [License](//github.com/digitalcraftsman/hugo-strata-theme/blob/dev/LICENSE.md).


## Annotations

Thanks to [Steve Francia](//github.com/spf13) for creating [Hugo](//gohugo.io) and the awesome community around the project.
