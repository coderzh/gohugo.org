---
date: 2013-07-01
linktitle: Configuration
menu:
  doc:
    name: 配置文件
    parent: getting started
next: /doc/overview/source-directory
notoc: true
prev: /doc/overview/usage
title: 配置 Hugo
weight: 40
---

通常的使用情况下，一个网站并不需要一个配置文件，因为它的目录结构和模板就提供了主要的配置。

Hugo 需要在源目录查找一个 `config.toml` 的配置文件。如果这个文件不存在，将会查找 `config.yaml`，然后是 `config.json` 。

这个配置文件是一个整站的配置。它给 Hugo 提供了如何构建站点的方式，比如全局的参数和菜单。

## 示例

下面是一个典型的 yaml 格式的配置文件的示例：

    ---
    baseurl: "http://yoursite.example.com/"
    ...

下面是一个 toml 格式的带了一些默认值的配置文件。
在 `[params]` 下面的值将会构成模板里的 `.Site.Params` 变量：

    contentdir = "content"
    layoutdir = "layouts"
    publishdir = "public"
    builddrafts = false
    baseurl = "http://yoursite.example.com/"
    canonifyurls = true
    
    [taxonomies]
      category = "categories"
      tag = "tags"
    
    [params]
      description = "Tesla's Awesome Hugo Site"
      author = "Nikola Tesla"

这是一个 yaml 格式的配置文件，设置了一些更多的选项：

    ---
    baseurl: "http://yoursite.example.com/"
    title: "Yoyodyne Widget Blogging"
    footnotereturnlinkcontents: "↩"
    permalinks:
      post: /:year/:month/:title/
    params:
      Subtitle: "Spinning the cogs in the widgets"
      AuthorName: "John Doe"
      GitHubUser: "spf13"
      ListOfFoo:
        - "foo1"
        - "foo2"
      SidebarRecentLimit: 5
    ...

## 配置变量

下面是 Hugo 定义好的变量列表，以及他们的默认值，你可以设置他们：

    ---
    archetypedir:               "archetype"
    # 根路径，比如： http://spf13.com/
    baseURL:                    ""
    # 包含标记了 draft 的内容
    buildDrafts:                false
    # 包含 publishdate 是 future 的内容
    buildFuture:                false
    # 让所有相对路径的 URL 相对于你的内容的根路径。注意这个选项不会影响绝对路径。
    relativeURLs:               false
    canonifyURLs:               false
    # 配置文件 (默认是 path/config.yaml|json|toml)
    config:                     "config.toml"
    contentdir:                 "content"
    dataDir:                    "data"
    defaultExtension:           "html"
    defaultLayout:              "post"
    # 生成文件目标位置
    destination:                ""
    disableLiveReload:          false
    # 不生成 RSS 文件
    disableRSS:                 false
    # 不生成 Sitemap 文件
    disableSitemap:             false
    # 如果提供了，生成新文章时将使用这个编辑器
    editor:                     ""
    footnoteAnchorPrefix:       ""
    footnoteReturnLinkContents: ""
    languageCode:               ""
    layoutdir:                  "layouts"
    # 开启日志
    log:                        false
    # 日志路径（如果设置了，默认开启日志）
    logFile:                    ""
    # "yaml", "toml", "json"
    metaDataFormat:             "toml"
    newContentEditor:           ""
    # 不要同步文件的修改时间
    noTimes:                    false
    paginate:                   10
    paginatePath:               "page"
    permalinks:
    # Pluralize titles in lists using inflect(不太明白)
    pluralizeListTitles:         true
    publishdir:                 "public"
    # 代码高亮的样式
    pygmentsStyle:              "monokai"
    # ture: 使用 pygments-css，false: 直接使用 color-codes（不使用 css ）
    pygmentsUseClasses:         false
    # 默认的站点地图
    sitemap:
    # 源文件在文件系统的路径
    source:                     ""
    staticdir:                  "static"
    # 显示程序在不同步骤的内存和耗时的分析
    stepAnalysis:               false
    # 使用哪个皮肤 (路径在 /themes/THEMENAME/)
    theme:                      ""
    title:                      ""
    # 如果是 true ，使用 /filename.html 代替 /filename/
    uglyURLs:                   false
    # 详细的输出
    verbose:                    false
    # 输出详细的日志
    verboseLog:                 false
    # 监控文件变化并重新生成
    watch:                      false
    ---




## 设置 Blackfriday 渲染

[Blackfriday](https://github.com/russross/blackfriday) 是 Hugo 中使用的 [Markdown](http://daringfireball.net/projects/markdown/) 渲染引擎。 Hugo 中的 Blackfriday 的默认设置已经很健全，可以适用于大多数的情况。

但是同时 Hugo 还提供了一些选项 --- 正如下面的表格里的所示的， 对应于 Blackfriday 源码里的选项。  ([html.go](https://github.com/russross/blackfriday/blob/master/html.go) 和 [markdown.go](https://github.com/russross/blackfriday/blob/master/markdown.go)):

<table class="table table-bordered">
<thead>
<tr>
<th>Flag</th><th>Default</th><th>Blackfriday flag</th>
</tr>
</thead>

<tbody>
<tr>
<td><code><strong>angledQuotes</strong></code></td>
<td><code>false</code></td>
<td><code>HTML_SMARTYPANTS_ANGLED_QUOTES</code></td>
</tr>
<tr>
<td class="purpose-title">Purpose:</td>
<td class="purpose-description" colspan="2">Enable/Disable smart angled double quotes.<br>
<small><strong>Example:</strong>&nbsp;<code>"Hugo"</code> renders to «Hugo» instead of “Hugo”.</small></td>
</tr>

<tr>
<td><code><strong>fractions</strong></code></td>
<td><code>true</code></td>
<td><code>HTML_SMARTYPANTS_FRACTIONS</code></td>
</tr>
<tr>
<td class="purpose-title">Purpose:</td>
<td class="purpose-description" colspan="2">Enable/Disable smart fractions.<br>
<small><strong>Example:</strong>&nbsp;<code>5/12</code> renders to <sup>5</sup>&frasl;<sub>12</sub> (<code>&lt;sup&gt;5&lt;/sup&gt;&amp;frasl;&lt;sub&gt;12&lt;/sub&gt;</code>)<br>
<strong>Caveat:</strong> Even with <code>fractions = false</code>,
Blackfriday would still convert 1/2, 1/4 and 3/4 to ½&nbsp;(<code>&amp;frac12;</code>),
¼&nbsp;(<code>&amp;frac14;</code>) and ¾&nbsp;(<code>&amp;frac34;</code>) respectively,
but only these three.</small></td>
</tr>

<tr>
<td><code><strong>plainIdAnchors</strong></code></td>
<td><code>false</code></td>
<td><code>FootnoteAnchorPrefix</code> and <code>HeaderIDSuffix</code></td>
</tr>
<tr>
<td class="purpose-title">Purpose:</td>
<td class="purpose-description" colspan="2">If <code>true</code>, then header and footnote IDs are generated without the document ID.<br>
<small><strong>Example:</strong>&nbsp;<code>#my-header</code> instead of <code>#my-header:bec3ed8ba720b9073ab75abcf3ba5d97</code>.</small></td>
</tr>

<tr>
<td><code><strong>extensions</strong></code></td>
<td><code>[]</code></td>
<td><code>EXTENSION_*</code></td>
</tr>
<tr>
<td class="purpose-title">Purpose:</td>
<td class="purpose-description" colspan="2">Use non-default additional extensions.<br>
<small><strong>Example:</strong>&nbsp;Add <code>"hardLineBreak"</code> to use <code>EXTENSION_HARD_LINE_BREAK</code>.</small></td>
</tr>

<tr>
<td><code><strong>extensionsmask</strong></code></td>
<td><code>[]</code></td>
<td><code>EXTENSION_*</code></td>
</tr>
<tr>
<td class="purpose-title">Purpose:</td>
<td class="purpose-description" colspan="2">Extensions in this option won't be loaded.<br>
<small><strong>Example:</strong>&nbsp;Add <code>"autoHeaderIds"</code> to disable <code>EXTENSION_AUTO_HEADER_IDS</code>.</small></td>
</tr>
</tbody>
</table>


**注意** 这些选项必须在 `blackfriday` 项下面，可以 **同时用在 site 和 page 级别** 。如果在 page 级别设置了，将会覆盖 site 上的设置。比如：

<table class="table">
<thead>
<tr>
<th>TOML</th><th>YAML</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td style="width: 50%;"><pre><code>[blackfriday]
  angledQuotes = true
  fractions = false
  plainIdAnchors = true
  extensions = ["hardLineBreak"]
</code></pre></td>
<td><pre><code>blackfriday:
  angledQuotes: true
  fractions: false
  plainIdAnchors: true
  extensions:
    - hardLineBreak
</code></pre></td>
</tr>
</tbody>
</table>

## 说明

修改配置对于 [LiveReload](/doc/extras/livereload/) 来说是没反应的。

当你修改配置时，请重启 `hugo server --watch` 。 