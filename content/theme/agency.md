+++
screenshot = "/images/agency.screenshot.png"
thumbnail = "/images/agency.tn.png"
demo = "/theme/agency/"
source = "https://github.com/digitalcraftsman/hugo-agency-theme.git"
name = "Agency"
license = "Apache License 2.0"
licenselink = "//github.com/digitalcraftsman/hugo-agency-theme/blob/master/LICENSE"
description = "A one page portfolio for companies and freelancers."
homepage = "//github.com/digitalcraftsman/hugo-agency-theme"
tags = ["portfolio", "company", "contact form", "clients", "team"]
features = ["portfolio with modals", "contact form", "timeline"]
min_version = 0.14

[author]
  name = "digitalcraftsman"
  homepage = "//github.com/digitalcraftsman"

# If porting an existing theme
[original]
  name = "David Miller"
  homepage = "//www.ironsummitmedia.com/"
  repo = "//github.com/IronSummitMedia/startbootstrap-agency"

+++

# Agency Theme

Agency Theme is a one page portfolio for companies and freelancers based on the [original Bootstrap theme](//github.com/IronSummitMedia/startbootstrap-agency) by [David Miller](//github.com/davidtmiller). This Hugo theme features several content sections, a responsive portfolio grid with hover effects, full page portfolio item modals, a timeline, and a contact form.

![Hugo Agency Theme screenshot](https://raw.githubusercontent.com/digitalcraftsman/hugo-agency-theme/master/images/screenshot.png)


## Contents

- [Installation](#installation)
- [Getting started](#getting-started)
    - [The config file](#the-config-file)
    - [Change the hero background](#change-the-hero-background)
    - [Present your skills](#present-your-skills) 
    - [Create your portfolio](#create-your-portfolio)
    - [Show what happened](#show-what-happened)
    - [Introduce your team](#introduce-your-team)
    - [List your clients](#list-your-clients)
    - [Make the contact form working](#make-the-contact-form-working)
    - [Nearly finished](#nearly-finished)
- [Contributing](#contributing)
- [License](#license)
- [Annotations](#annotations)


## Installation

Inside the folder of your Hugo site run:

    $ mkdir themes
    $ cd themes
    $ git clone https://github.com/digitalcraftsman/hugo-agency-theme

For more information read the official [setup guide](//gohugo.io/overview/installing/) of Hugo.


## Getting started

After installing the Agency Theme successfully it requires a just a few more steps to get your site running.


### The config file

Take a look inside the [`exampleSite`](//github.com/digitalcraftsman/hugo-agency-theme/tree/master/exampleSite) folder of this theme. You'll find a file called [`config.toml`](//github.com/digitalcraftsman/hugo-agency-theme/blob/master/exampleSite/config.toml). To use it, copy the [`config.toml`](//github.com/digitalcraftsman/hugo-agency-theme/blob/master/exampleSite/config.toml) in the root folder of your Hugo site. Feel free to change the strings in this theme.


### Change the hero background

The hero acts as an eye-catcher for your site. So consider to give him a nice background. You just need to replace the [`header-bg.jpg`](//github.com/digitalcraftsman/hugo-agency-theme/blob/master/static/img/header-bg.jpg) at [`static/img`](//github.com/digitalcraftsman/hugo-agency-theme/tree/master/static/img) with your own background image. But it's important that you keep the original filename.


### Present your skills

This section should show your capabilities and skills. You can change the services at `[params.services.list]` in the [`config.toml`](//github.com/digitalcraftsman/hugo-agency-theme/blob/master/exampleSite/config.toml).

All icons are part of Fontawesome's icon font. Look at the website of [Fontawesome](//fortawesome.github.io/Font-Awesome/icons/) for more icons. The icons are represented by their corresponding CSS class of Fontawesome. A skill is defined like this example:

```toml
[[params.services.list]]
      icon = "fa-shopping-cart"
      title = "E-Commerce"
      description = "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Minima maxime quam architecto quo inventore harum ex magni, dicta impedit."
```


### Create your portfolio

Beside the [`config.toml`](//github.com/digitalcraftsman/hugo-agency-theme/blob/master/exampleSite/config.toml), there is under `data` another subfolder called [`projects`](//github.com/digitalcraftsman/hugo-agency-theme/tree/master/exampleSite/data/projects) which hosts the files that will appear as your projects in the portfolio section. Such a project file might look like [this one](//github.com/digitalcraftsman/hugo-agency-theme/blob/master/exampleSite/data/projects/2014-07-05-project-1.yaml) written in YAML:

```yaml
modalID: 1
title: Round Icons
subtitle: Lorem ipsum dolor sit amet consectetur.
date: 2014-07-05
img: roundicons.png
preview: roundicons-preview.png
client: Start Bootstrap
clientLink: "#"
category: Graphic Design
description: Use this area to describe your project. Lorem ipsum dolor sit amet, consectetur adipisicing elit. Est blanditiis dolorem culpa incidunt minus dignissimos deserunt repellat aperiam quasi sunt officia expedita beatae cupiditate, maiores repudiandae, nostrum, reiciendis facere nemo! <br><br>**Want these icons in this portfolio item sample?** You can download 60 of them for free, courtesy of [RoundIcons.com](//getdpd.com/cart/hoplink/18076?referrer=bvbo4kax5k8ogc), or you can purchase the 1500 icon set [here](//getdpd.com/cart/hoplink/18076?referrer=bvbo4kax5k8ogc).
```

Copy [`projects`](//github.com/digitalcraftsman/hugo-agency-theme/tree/master/exampleSite/data/projects) inside the `data` folder in the **root** directory of your site. Let's make some changes.

Pay attention to the `modalID`. It must be a unique integer and be incremented with each new project you want to add to the portfolio. Otherwise, the corresponding modal can't be rendered.

Furthermore, you can use Markdown syntax for URLs like here `[text](//url.to/source)` in the description.

To give your projects an image, save those under [`static/img/portfolio`](github.com/digitalcraftsman/hugo-agency-theme/tree/master/static/img/portfolio). Don't forget to set the appropriate **filename** under `img` in your project.


### Show what happened

This theme features a timeline for important events in your company or your career too. You can add a new event by copying the following snippet to the `[params.about]` section in the [`config.toml`](//github.com/digitalcraftsman/hugo-agency-theme/blob/master/exampleSite/config.toml).

```toml
[[params.about.events]]
      img = "1.jpg"
      date = "2009-2011"
      title = "Our Humble Beginnings"
      description = "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Sunt ut voluptatum eius sapiente, totam reiciendis temporibus qui quibusdam, recusandae sit vero unde, sed, incidunt et ea quo dolore laudantium consectetur!"
```

The image set under `img` needs to be stored at [`static/img/about`](//github.com/digitalcraftsman/hugo-agency-theme/tree/master/static/img/about). The events will be listed from the top to the bottom.


### Introduce your team

Let the visitors or potential clients know who you are. To add a team member paste the code below into the [`config.toml`](//github.com/digitalcraftsman/hugo-agency-theme/blob/master/exampleSite/config.toml). The `img` field refers to the shown image. Paste those of you or your colleages into [`static/img/team`](//github.com/digitalcraftsman/hugo-agency-theme/tree/master/static/img/team) 

```toml
[[params.team.members]]
    img = "1.jpg"
    name = "Kay Garland"
    position = "Lead Designer"
    social = [
        ["fa-twitter", "#"],
        ["fa-facebook", "#"],
        ["fa-linkedin", "#"]
    ]
```

As you can see there's an option to link individual social networks. The first index of the array represents the icon (or CSS class) of [Fontawesome](//fortawesome.github.io/Font-Awesome/icons/). The last index is simply the link to the social network profiles.


### List your clients

You can also show some of your clients. To do so, paste the client's logos into [`static/img/logos`](//github.com/digitalcraftsman/hugo-agency-theme/tree/master/static/img/logos) and add the example below to the [`config.toml`](//github.com/digitalcraftsman/hugo-agency-theme/blob/master/exampleSite/config.toml).

```toml
[[params.clients]]
    logo = "designmodo.jpg"
    link = "#"
```

***The logos require a dimension of 200 x 50 pixels.***


### Make the contact form working

Since this page will be static, you can use [formspree.io](//formspree.io/) as proxy to send the actual email. Each month, visitors can send you up to one thousand emails without incurring extra charges. Begin the setup by following the steps below:

1. Enter your email address under 'email' in the [`config.toml`](//github.com/digitalcraftsman/hugo-agency-theme/blob/master/exampleSite/config.toml)
2. Upload the generated site to your server
3. Send a dummy email yourself to confirm your account
4. Click the confirm link in the email from [formspree.io](//formspree.io/)
5. You're done. Happy mailing!


### Nearly finished

In order to see your site in action, run Hugo's built-in local server. 

    $ hugo server -w

Now enter `localhost:1313` in the address bar of your browser.


## Contributing

Did you found a bug or got an idea for a new feature? Feel free to use the [issue tracker](//github.com/digitalcraftsman/hugo-agency-theme/issues) to let me know. Or make directly a [pull request](//github.com/digitalcraftsman/hugo-agency-theme/pulls).


## License

This theme is released under the Apache License 2.0 For more information read the [License](//github.com/digitalcraftsman/hugo-agency-theme/blob/master/LICENSE).


## Annotations

Thanks to [Steve Francia](//github.com/spf13) for creating Hugo and the awesome community around the project.
