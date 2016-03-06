---
date: 2014-03-10
lastmod: 2015-12-24
linktitle: Migrating from Jekyll
menu:
  main:
    parent: tutorials
next: /doc/tutorials/create-a-multilingual-site
prev: /doc/tutorials/mathjax
title: Migrate to Hugo from Jekyll
toc: true
weight: 10
translated: true
---

**说明** Hugo 0.15 支持 `hugo import jekyll` 命令，详见：[import from Jekyll](/doc/commands/hugo_import_jekyll/).
## 移动静态内容到 `static` 文件夹
Jekyll 有一条规则，任何不以 `_` 开始的文件夹都会被原样复制到 `_site` 文件夹中。Hugo 把所有的静态内容放在 `static` 文件夹中。因此你需要把它们都移动到那里去。
使用 Jekyll，看起来这样的路径

    ▾ <root>/
        ▾ images/
            logo.png

应该变为

    ▾ <root>/
        ▾ static/
            ▾ images/
                logo.png

另外，你也想要让那些放在根目录下的文件（比如 `CNAME`）也移动到 `static` 中。

## 创建 Hugo 配置文件
Hugo 可以 JSON、YAML 或者 TOML 格式读入你的配置。Hugo 也支持自定义参数的配置。详细内容可以参考 [Hugo 配置文档](/doc/overview/configuration/)。

## 把发布文件夹配置为 `_site`
Jekyll 的默认发布文件夹是 `_site`，而 Hugo 的是 `public`。如果，你和我一样，把 [`_site` 文件夹映射为 `gh-pages` 分支的 git 子模块](http://blog.blindgaenger.net/generate_github_pages_in_a_submodule.html)，下面两种方法，你需要做其中之一：

1. 修改指向 `gh-pages` 的子模块，由 `_site` 指向 public （推荐）。

        git submodule deinit _site
        git rm _site
        git submodule add -b gh-pages git@github.com:your-username/your-repo.git public

2. 或者，修改 Hugo 的配置，使用 `_site`，而不是 `public`。

        {
            ..
            "publishdir": "_site",
            ..
        }

## 把 Jekyll 模板转换为 Hugo 的模板
主要工作就在这里。文档是你的朋友。如果你需要重新记起如何构建博客的话，应该参考 [Jekyll 模板文档](http://jekyllrb.com/docs/templates/) 并参考 [Hugo 模板](/doc/layout/templates/) 来学习 Hugo 的方式。

作为一个参考数据，为 [heyitsalex.net](http://heyitsalex.net/) 转换模板花费我不超过几个小时的时间。

## 把 Jekyll 插件转换为 Hugo 的短代码 
Jekyll 有[插件](http://jekyllrb.com/docs/plugins/)；Hugo 有[短代码](/doc/shortcodes/)。转换是非常琐碎的。

### 实现
作为一个例子，我使用一个自定义的 [`image_tag`](https://github.com/alexandre-normand/alexandre-normand/blob/74bb12036a71334fdb7dba84e073382fc06908ec/_plugins/image_tag.rb) 插件，在运行 Jekyll 时生成带标题的图片。当我阅读短代码的时候，我发现 Hugo 有一个内建的好用的短代码，可以完成同样的功能。

Jekyll 的插件：

```ruby
    module Jekyll
      class ImageTag < Liquid::Tag
        @url = nil
        @caption = nil
        @class = nil
        @link = nil
        // Patterns
        IMAGE_URL_WITH_CLASS_AND_CAPTION =
        IMAGE_URL_WITH_CLASS_AND_CAPTION_AND_LINK = /(\w+)(\s+)((https?:\/\/|\/)(\S+))(\s+)"(.*?)"(\s+)->((https?:\/\/|\/)(\S+))(\s*)/i
        IMAGE_URL_WITH_CAPTION = /((https?:\/\/|\/)(\S+))(\s+)"(.*?)"/i
        IMAGE_URL_WITH_CLASS = /(\w+)(\s+)((https?:\/\/|\/)(\S+))/i
        IMAGE_URL = /((https?:\/\/|\/)(\S+))/i
        def initialize(tag_name, markup, tokens)
          super
          if markup =~ IMAGE_URL_WITH_CLASS_AND_CAPTION_AND_LINK
            @class   = $1
            @url     = $3
            @caption = $7
            @link = $9
          elsif markup =~ IMAGE_URL_WITH_CLASS_AND_CAPTION
            @class   = $1
            @url     = $3
            @caption = $7
          elsif markup =~ IMAGE_URL_WITH_CAPTION
            @url     = $1
            @caption = $5
          elsif markup =~ IMAGE_URL_WITH_CLASS
            @class = $1
            @url   = $3
          elsif markup =~ IMAGE_URL
            @url = $1
          end
        end
        def render(context)
          if @class
            source = "<figure class='#{@class}'>"
          else
            source = "<figure>"
          end
          if @link
            source += "<a href=\"#{@link}\">"
          end
          source += "<img src=\"#{@url}\">"
          if @link
            source += "</a>"
          end
          source += "<figcaption>#{@caption}</figcaption>" if @caption
          source += "</figure>"
          source
        end
      end
    end
    Liquid::Template.register_tag('image', Jekyll::ImageTag)
```

写成 Hugo 的短代码是这种形式：

    <!-- image -->
    <figure {{ with .Get "class" }}class="{{.}}"{{ end }}>
        {{ with .Get "link"}}<a href="{{.}}">{{ end }}
            <img src="{{ .Get "src" }}" {{ if or (.Get "alt") (.Get "caption") }}alt="{{ with .Get "alt"}}{{.}}{{else}}{{ .Get "caption" }}{{ end }}"{{ end }} />
        {{ if .Get "link"}}</a>{{ end }}
        {{ if or (or (.Get "title") (.Get "caption")) (.Get "attr")}}
        <figcaption>{{ if isset .Params "title" }}
            {{ .Get "title" }}{{ end }}
            {{ if or (.Get "caption") (.Get "attr")}}<p>
            {{ .Get "caption" }}
            {{ with .Get "attrlink"}}<a href="{{.}}"> {{ end }}
                {{ .Get "attr" }}
            {{ if .Get "attrlink"}}</a> {{ end }}
            </p> {{ end }}
        </figcaption>
        {{ end }}
    </figure>
    <!-- image -->

### 使用
我简单的修改：

    {% image full http://farm5.staticflickr.com/4136/4829260124_57712e570a_o_d.jpg "One of my favorite touristy-type photos. I secretly waited for the good light while we were "having fun" and took this. Only regret: a stupid pole in the top-left corner of the frame I had to clumsily get rid of at post-processing." ->http://www.flickr.com/photos/alexnormand/4829260124/in/set-72157624547713078/ %}

成这个样子（这个例子使用了扩展版本的命名 `fig`，和内建的 `figure` 不同）：

    {{%/* fig class="full" src="http://farm5.staticflickr.com/4136/4829260124_57712e570a_o_d.jpg" title="One of my favorite touristy-type photos. I secretly waited for the good light while we were having fun and took this. Only regret: a stupid pole in the top-left corner of the frame I had to clumsily get rid of at post-processing." link="http://www.flickr.com/photos/alexnormand/4829260124/in/set-72157624547713078/" */%}}

这么做的好处是，命名参数的短代码，相对而言，更加易读。

## 收尾工作
### 修改内容
依据 Jekyll 生成的每篇博文的自定义数量，这个步骤将或多或少需要一些精力。除了 `hugo server --watch` 是你的朋友外，没有其他直接快速的规则。测试你的修改，并根据需要修改错误。

### 清除
这时你将需要移除 Jekyll 的配置。如果还有其他不再使用的东西，也删除掉。

## 一个 diff 形式的实际例子
[Hey, it's Alex](http://heyitsalex.net/) 从 Jekyll 移植到 Hugo 花费了不到一个_带孩子的爸爸的一天_。查看这个 [diff](https://github.com/alexandre-normand/alexandre-normand/compare/869d69435bd2665c3fbf5b5c78d4c22759d7613a...b7f6605b1265e83b4b81495423294208cc74d610)，你可以看到所有的修改（和所犯的错误）。
