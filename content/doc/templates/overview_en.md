---
aliases:
- /doc/templates/
- /layout/templates/
- /layout/overview/
date: 2013-07-01
lastmod: 2015-05-22
linktitle: Overview
menu:
  main:
    parent: layout
next: /doc/templates/go-templates
prev: /doc/themes/creation
title: Hugo Templates
toc: true
weight: 10
---

Hugo uses the excellent Go html/template library for its template engine.
It is an extremely lightweight engine that provides a very small amount of
logic. In our experience it is just the right amount of logic to be able
to create a good static website.

While Hugo has a number of different template roles, most complete
websites can be built using just a small number of template files.
Please don’t be afraid of the variety of different template roles. They
enable Hugo to build very complicated sites. Most sites will only
need to create a [/layouts/\_default/single.html](/doc/templates/content/) & [/layouts/\_default/list.html](/doc/templates/list/)

If you are new to Go's templates, the [Go Template Primer](/layout/go-templates/)
is a great place to start.

If you are familiar with Go’s templates, Hugo provides some [additional
template functions](/doc/templates/functions/) and [variables](/doc/templates/variables/) you will want to be familiar
with.

## Primary Template roles

There are 3 primary kinds of templates that Hugo works with.

### [Single](/doc/templates/content/)
Render a single piece of content

### [List](/doc/templates/list/)
Page that list multiple pieces of content

### [Homepage](/doc/templates/homepage/)
The homepage of your site

## Supporting Template Roles (optional)

Hugo also has additional kinds of templates all of which are optional

### [Partial Templates](/doc/templates/partials/)
Common page parts to be included in the above mentioned templates

### [Content Views](/doc/templates/views/)
Different ways of rendering a (single) content type

### [Taxonomy Terms](/doc/templates/terms/)
A list of the terms used for a specific taxonomy, e.g. a Tag cloud

## Other Templates (generally unnecessary)

### [RSS](/doc/templates/rss/)
Used to render all rss documents

### [Sitemap](/doc/templates/sitemap/)
Used to render the XML sitemap

### [404](/doc/templates/404/)
This template will create a 404.html page used when hosting on GitHub Pages


