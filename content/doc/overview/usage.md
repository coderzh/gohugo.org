---
aliases:
- /doc/usage/
date: 2013-07-01
lastmod: 2015-12-23
menu:
  main:
    parent: getting started
next: /doc/overview/configuration
notoc: true
prev: /doc/overview/installing
title: Using Hugo
weight: 30
translated: true
---

确保 `Hugo` 在你的 `PATH` 路径下或者提供一个它的路径。

{{< nohighlight >}}$ hugo help

Hugo is a Fast and Flexible Static Site Generator built with love by spf13 and friends in Go.

Complete documentation is available at http://gohugo.io

Usage:
  hugo [flags]
  hugo [command]

Available Commands:
  server          Hugo runs its own webserver to render the files
  version         Print the version number of Hugo
  config          Print the site configuration
  check           Check content in the source directory
  benchmark       Benchmark hugo by building a site a number of times
  new             Create new content for your site
  undraft         Undraft changes the content's draft status from 'True' to 'False'
  genautocomplete Generate shell autocompletion script for Hugo
  gendoc          Generate Markdown documentation for the Hugo CLI.
  help            Help about any command

Flags:
  -b, --baseURL="": hostname (and path) to the root, e.g. http://spf13.com/
  -D, --buildDrafts=false: include content marked as draft
  -F, --buildFuture=false: include content with publishdate in the future
      --cacheDir="": filesystem path to cache directory. Defaults: $TMPDIR/hugo_cache/
      --config="": config file (default is path/config.yaml|json|toml)
  -d, --destination="": filesystem path to write files to
      --disableRSS=false: Do not build RSS files
      --disableSitemap=false: Do not build Sitemap file
      --editor="": edit new content with this editor, if provided
  -h, --help=false: help for hugo
      --ignoreCache=false: Ignores the cache directory for reading but still writes to it
      --log=false: Enable Logging
      --logFile="": Log File path (if set, logging enabled automatically)
      --noTimes=false: Don't sync modification time of files
      --pluralizeListTitles=true: Pluralize titles in lists using inflect
  -s, --source="": filesystem path to read files relative from
      --stepAnalysis=false: display memory and timing of different steps of the program
  -t, --theme="": theme to use (located in /doc/themes/THEMENAME/)
      --uglyURLs=false: if true, use /filename.html instead of /filename/
  -v, --verbose=false: verbose output
      --verboseLog=false: verbose logging
  -w, --watch=false: watch filesystem for changes and recreate as needed


Additional help topics:
 hugo convert         Convert will modify your content to different formats hugo list            Listing out various types of content

Use "hugo help [command]" for more information about a command.
{{< /nohighlight >}}

## 一般用法

最常见的用法是在你的当前站点目录里执行 `hugo` ：

{{< nohighlight >}}$ hugo
0 draft content
0 future content
99 pages created
0 paginator pages created
16 tags created
0 groups created
in 120 ms
{{< /nohighlight >}}

你的站点将会生成到 `public` 目录，用来部署到你的 Web 服务器。


## 开发你的站点时获得实时的反馈

如果你想立即查看到新的内容改动，告诉 Hugo 让它监控这些改动。
Hugo 将会监控文件系统的变化，在文件保存时会立即重新构建你的站点：

{{< nohighlight >}}$ hugo -s ~/Code/hugo/docs
0 draft content
0 future content
99 pages created
0 paginator pages created
16 tags created
0 groups created
in 120 ms
Watching for changes in /Users/spf13/Code/hugo/docs/content
Press Ctrl+C to stop
{{< /nohighlight >}}

Hugo 甚至可以运行一个 Server ，并且同时创建一个预览的站点！
Hugo 实现了 [LiveReload](/doc/extras/livereload/) 技术用来自动重新加载打开的页面，只要是支持 JavaScript 的浏览器，包括移动设备。
这是最简单也是最常见的开发 Hugo 站点的方法：

{{< nohighlight >}}$ hugo server -ws ~/Code/hugo/docs
0 draft content
0 future content
99 pages created
0 paginator pages created
16 tags created
0 groups created
in 120 ms
Watching for changes in /Users/spf13/Code/hugo/docs/content
Serving pages from /Users/spf13/Code/hugo/docs/public
Web Server is available at http://localhost:1313/
Press Ctrl+C to stop
{{< /nohighlight >}}


## 部署你的站点

本地开发并执行完 `hugo server` 之后，你需要最后执行一次 `hugo` 命令 **不带 `server` 参数** ，
然后你就可以通过拷贝 `public` 目录（通过 FTP 、 SFTP 、 WebDAV 、 Rsync 、 git push 等等）到你的生产环境的 Web 服务器来 **部署你的站点**。

由于 Hugo 生成的是一个静态站点，你的站点可以托管到任何地方，包括 [Heroku][] 、 [GoDaddy][] 、 [DreamHost][] 、 [GitHub Pages][] 、 [Amazon S3][] 和 [CloudFront][] ，
或者任何便宜甚至免费的静态网站托管服务。

[Apache][], [nginx][], [IIS][]...  任何 Web 服务器都支持！

[Apache]: http://httpd.apache.org/ "Apache HTTP Server"
[nginx]: http://nginx.org/
[IIS]: http://www.iis.net/
[Heroku]: https://www.heroku.com/
[GoDaddy]: https://www.godaddy.com/
[DreamHost]: http://www.dreamhost.com/
[GitHub Pages]: https://pages.github.com/
[Amazon S3]: http://aws.amazon.com/s3/
[CloudFront]: http://aws.amazon.com/cloudfront/ "Amazon CloudFront"


### 关于部署的说明

运行 `hugo` 命令 *并不会* 删除之前生成的文件。这意味着你必须在运行 `hugo` 命令之前删除你的 `public/` 目录（或者你通过 `-d`/`--destination` 指定的目录）。
不然的话，你可能有运行遗留在生成目录的错误文件的风险（比如草稿或者未来的文章）。

一个简单的处理方法是使用不同的目录用于开发和部署环境。

启动一个会生成草稿内容（有助于编辑）的 Server ，指定一个不同的目标目录： `dev/` 目录

{{< nohighlight >}}$ hugo server -wDs ~/Code/hugo/docs -d dev
{{< /nohighlight >}}

当内容准备好发布时，使用默认的 `public/` 目录：

{{< nohighlight >}}$ hugo -s ~/Code/hugo/docs
{{< /nohighlight >}}

这将防止你不小心发布了还未准备好的内容。

### 或者, 直接用 Hugo 作为你的 Web 服务器！

是的，没错！由于 Hugo 运行的非常快，不仅在生成站点时，*也* 在运行 Web 服务时（得益于 Go 语言关于并发和多线程的良好设计）。
有些用户事实上选择使用 Hugo 本身作为 Web 服务器**用于他的生产环境服务器**。

不再需要其他的 Web 服务器了 （Apache 、 nginx 、 IIS ……）。

下面是执行的命令：

{{< nohighlight >}}$ hugo server --baseURL=http://yoursite.org/ \
              --port=80 \
              --appendPort=false \
              --bind=87.245.198.50
{{< /nohighlight >}}

注意 `bind` 选项，指定了服务器绑定到哪个网络接口（默认是 `127.0.0.1` ，适用于大多数开发的时候）。
有些服务器，比如 Amazon WS ，使用网络地址转换（NAT）导致你有时很难找到实际使用的 IP 地址。
使用 `--bind=0.0.0.0` 将可以绑定到所有接口。

通过这种方法，实际上你可以只部署你的源文件了，因为 Hugo 会在你的服务器上及时的生成 Web 站点。

你可以添加 `--disableLiveReload=true` 可选选项，如果你不希望将 LiveReload 相关的 JavaScript 代码添加到你的页面。

是不是很有趣？这里有一些 Hugo 用户贡献的不错的指引：

* [hugo, syncthing](http://fredix.ovh/2014/10/hugo-syncthing/) (French) by Frédéric Logier (@fredix)
