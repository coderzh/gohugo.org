---
aliases:
- /doc/source-directory/
date: 2013-07-01
lastmod: 2015-02-09
menu:
  main:
    parent: getting started
next: /doc/content/organization
notoc: true
prev: /doc/overview/configuration
title: Source Organization
weight: 50
translated: true
---

Hugo 使用一个单一的目录用来生成完整的站点。

源目录的顶层一般有以下一些元素：

    ▸ archetypes/
    ▸ content/
    ▸ data/
    ▸ layouts/
    ▸ static/
    ▸ themes/
      config.toml

学习更多这些不同的目录和他们的目的：

* [config](/doc/overview/configuration/)
* [data](/doc/extras/datafiles/)
* [archetypes](/doc/content/archetypes/)
* [content](/doc/content/organization/)
* [layouts](/doc/layout/overview/)
* [static](/doc/themes/creation#toc_4)
* [themes](/doc/themes/overview/)


## 例子

一个目录的例子：

    .
    ├── config.toml
    ├── archetypes
    |   └── default.md
    ├── content
    |   ├── post
    |   |   ├── firstpost.md
    |   |   └── secondpost.md
    |   └── quote
    |   |   ├── first.md
    |   |   └── second.md
    ├── data
    ├── layouts
    |   ├── _default
    |   |   ├── single.html
    |   |   └── list.html
    |   ├── partials
    |   |   ├── header.html
    |   |   └── footer.html
    |   ├── taxonomies
    |   |   ├── category.html
    |   |   ├── post.html
    |   |   ├── quote.html
    |   |   └── tag.html
    |   ├── post
    |   |   ├── li.html
    |   |   ├── single.html
    |   |   └── summary.html
    |   ├── quote
    |   |   ├── li.html
    |   |   ├── single.html
    |   |   └── summary.html
    |   ├── shortcodes
    |   |   ├── img.html
    |   |   ├── vimeo.html
    |   |   └── youtube.html
    |   ├── index.html
    |   └── sitemap.xml
    ├── themes
    |   ├── hyde
    |   └── doc
    └── static
        ├── css
        └── js

这个目录结构告诉了我们一些关于这个站点的信息：

1. 这个网站打算有两个不同类型的文章内容：*posts* 和 *quotes* 。
2. 它将在文章内容里应用两个不同的分类：*categories* 和 *tags* 。
3. 它将在三种不同的视图里呈现这些内容：列表、摘要和完整的内容。
