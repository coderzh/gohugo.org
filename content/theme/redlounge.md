+++
screenshot = "/images/redlounge.screenshot.png"
thumbnail = "/images/redlounge.tn.png"
demo = "/theme/redlounge/"
source = "https://github.com/tmaiaroto/hugo-redlounge.git"
name = "Red Lounge"
description = "A clean, responsive, template with red accents."
license = "Apache 2.0"
licenselink = "https://github.com/tmaiaroto/hugo-redlounge/blob/master/LICENSE.md"
source_repo = "http://github.com/tmaiaroto/hugo-redlounge"
tags = ["redlounge", "red", "raleway", "libre baskerville", "blog", "gallery"]
features = ["blog", "gallery"]
min_version = 0.14

[author]
    name = "Tom Maiaroto"
    homepage = "http://www.shift8creative.com"

+++

Red Lounge
===========

This is an open-source Hugo theme designed by [@shift8creative](http://www.twitter.com/shift8creative) to be responsive and clean. It uses Pure.css and contains a few web fonts from Google as well as Font Awesome icons. 
So despite it being simple, it has a lot of flexibility in terms of typography and design elements. It's also quite configurable as it has a few variables to change the
appearance and features of the pages. By default, excess features (and JavaScript) is not included.

## Configuration and Options

### Sidebar Title & Tagline Params

The sidebar can be configured with your main site config using params. For example, in ```config.toml```

```
[params]
	sidebartitle = "My Site"
	sidebartagline = "Is super awesome"
    sidebarphoto = "/img/photo.png"
```

This three properites will not be shown if not set. You will want to keep these lines short since there's limited space. 
Alternatively you could add custom HTML using the ```sidebarheader.html``` partial and keep those values empty strings.

### Menus

There are a few menus this theme allows you to define (all optional) in your main config.

Main - This menu is pretty basic and goes on the left panel under the site title and description. It contains red markers to separate items and call attention to the 
fact that it is more important than the other lists/menus that you may have on your site. It is optional, but always shows a link to the home page.    
This would be a good place to link to your various sections.

```
[[menu.main]]
    name = "Blog"
    url = "/posts"
```


Social - This menu goes underneath the main menu and was originally designed to contain links out to social media accounts, RSS, etc. making use of Font Awesome.    
However, you can use it for whatever you want. Just keep in mind space is limited here on the left panel.

```
[[menu.social]]
	pre = "<i class='fa fa-twitter'></i>"
    url = "http://www.twitter.com/shift8creative"
    identifier = "twitter"
```

Footer - The footer menu might also contain links to social media accounts...It's up to you. It appears right above the copyright notice at the bottom of each page.    
This menu is simply plain text links centered and they are gray to match the footer. So there's less attention being drawn here. Perhaps good for notices, terms of service, etc.

```
[[menu.footer]]
    name = "Blog"
    url = "/posts"
```

### Categories

Some assembly required here.

```.Params.categories``` coming from front-matter are displayed as tags on list pages with boxes. By default they are all going to be gray with white text. The coloring is up to you.

Each label will have the following class: ```class="post-category post-category-{{ . | urlize }}"```

Note the name is going to be urlized. So for example: ```post-category-technology``` or ```post-category-golang``` and so on. This allows you to create your own CSS around the 
categories you end up defining. You can then set the background color to be something specific and then all instances of that category label will match.

You can easily include a categories CSS file, without modifying template partials, by using the site config params. Something like the following:

```
[params]
	categoriescss = "/css/my-categories.css"
```

This will be included in the header.html file before headend.html partial, so you can still include additional code in that partial afterward.

### Comments

Comments use Disqus, so the main config needs to define ```disqusShortname``` like normal. However, each page can disable comments from appearing with front matter. Simply set 
```nocomment = true``` and they will be hidden.

### Lightbox

Lightbox is included with the theme but won't be available for use (not even linked in the HTML) unless you enable it. This way it stays out of the way and saves on bandwidth. 
Should you decide you want to use it, simply add to your front-matter: ```lightbox = true``` and then in your markdown you'll need add links with a ```data-lightbox``` attribute. 
Markdown wants to add titles when you add quotes so the syntax is a little weird. Alternatively, you can use HTML (which is likely easier in this case). So the following should 
use images in your ```static``` directory:

```
<a href="/image.jpg" title="" data-lightbox="set1" data-title="This is my caption"><img src="/image-thumbnail.jpg" alt=""></a>
``` 

Also note that Lightbox requires jQuery, so turning this on for a page or archetype will also link jQuery in the head section of your pages from Google's CDN.

### Hiding & Showing Things

Comments can be hidden on a per page basis with ```nocomment = true``` but there's also some other things that can be hidden. Sometimes simply by not defining them, other times 
by explicitly setting variables. They are as follows:

 - ```nodate = true``` Hides the date on a page
 - ```noauthor = true``` Hides the author (which may simply be defined per page, but also could be set higher up in the config, so this overrides)
 - ```noread = true``` Hides the read time displayed on list pages
 - ```nopaging = true``` Hides the next/prev links that navigate through pages
 - ```notoc = true``` Hides the table of contents from a page
 - ```totop = true``` Shows a "To Top" link, fixed in the top right of the page after scrolling beyond 1000px, that takes the user back to the top of the page (not shown by default)
 - ```socialsharing = true``` Shows social media sharing buttons from a partial template (which can be overwritten)

### Other Params

 Aside from the above variables that hide/show things, there are a few other variables that this theme will look for and use. These can be placed in any front matter. In some cases 
 you'll want to define these in the archetypes so you don't need to keep defining the values. These are all optional of course.

 - ```ogimage``` OpenGraph image tag
 - ```ogtitle``` OpenGraph title tag (if not set, will use ```.Title```)
 - ```ogdescription``` OpenGraph description tag (if not set, will use ```.Description``` - also used the same way for meta description)
 - ```banner``` A banner image for a page that sits above the title and stretches across the width of the page
 - ```bannerheight``` An optional height for the banner image wrapper to help cut off larger images (specified as just a number, none by default)
 - ```bannerfill``` If set true, this will attempt to fill the available width of the content area with the banner image by applying a 100% width to the image style
 - ```bannerinline``` Use this instead to place a banner image after the title, author, tags, etc. which floats to the right of the content (does not use the fill/height params)
 - ```thumbnail``` Set this to an image path if you want to show a thumbnail with pages in the list view; originally intended for author avatars, any small image could work (50x50 pixels)
 - ```author``` The author name
 - ```authorlink``` A link for the author
 - ```authorlinktarget``` The link target, ie. "_blank" (useful if the link goes to an external site or something)
 - ```authortwitter``` The full URL to the author's Twitter profile (opens in a new window)
 - ```authorlinkedin``` The full URL to the author's LinkedIn profile (opens in a new window)
 - ```authorfacebook``` The full URL to the author's Facebook page/profile (opens in a new window)
 - ```authorgoogleplus``` The full URL to the author's Google+ page/profile (opens in a new window)

### Template Partials

There are a few partials being used so that key areas can be easily overwritten.

 - ```doctype.html``` This contains the HTML document declaration and you may wish to change it from the default value of english for language, add namespaces, etc.
 - ```header.html``` This contains the ```<head>``` portion of the page
 - ```meta.html``` This contains some basic meta tags, feel free to modify as needed (within header.html)
 - ```og.html``` OpenGraph only meta tags (within header.html)
 - ```headend.html``` This easily provides the ability to add custom style sheets and JavaScript right before ```</head>``` to override styles, etc. (within header.html)
 - ```authorsocial.html``` If an author is set this template partial will be used to optionally show links to their social media profiles if also set (within single.html)
 - ```socialsharing.html``` Allows you to change what's displayed when ```socialsharing = true```
 - ```sidebar.html``` The sidebar which is already pretty customizable with site params and menus, but can also easily be changed if need be
 - ```sidebarheader.html``` Placed above the h1 and h2 elements in the sidebar (which will appear if your site config was set), allowing for further sidebar header adjustments
 - ```singletop.html``` Placed at the top of the page for single.html (but under the banner)
 - ```listtop.html``` Placed at the top of the page for list.html
 - ```footer.html``` The footer
 - ```bodyend.html``` Right before ```</body>``` (within list.html, single.html, and index.html - be sure to include it when/if overwriting with your own index.html)
