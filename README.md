# GoHugo.org

Hugo 中文文档： [http://www.gohugo.org/doc/](http://www.gohugo.org/doc/)

## Hugo 介绍

[spf13](http://spf13.com) 对于 Hugo 的英文原版介绍地址在这，就不另行翻译了：[https://gohugo.io/overview/introduction/](https://gohugo.io/overview/introduction/)

原文大致意思是当前的静态网站生成工具对环境依赖过多，性能较差，于是使用 Go 语言写了一个静态网站生成器 Hugo。不仅解决了环境依赖、性能较差的问题，还有使用简单、部署方便等诸多优点，通过 LiveReload 实时刷新，极大的优化文章的写作体验。

## 关于翻译文档

本翻译文档目的是让 Hugo 在中国能得到更好的推广，让有需要和对 Hugo 感兴趣的人能从此文档中获得微薄帮助。

此中文文档翻译自官方最新的 `v0.15` 版本：

[https://github.com/spf13/hugo/tree/v0.15.docs](https://github.com/spf13/hugo/tree/v0.15.docs)

翻译文档最麻烦的是能否跟上官方文档的更新，于是我做了一个自动与官方最新版本生成 diff 的 Go 脚本，具体如何使用稍后放出。

此次翻译我也希望借助社区之力，如果你对此有兴趣，请前往本站的 GitHub **提交一个 Issue，标题注明希望翻译的文档名称** 即可。

Issue 提交地址： [https://github.com/coderzh/gohugo.org/issues](https://github.com/coderzh/gohugo.org/issues)

待翻译文档路径： [https://github.com/coderzh/gohugo.org/tree/master/content/doc](https://github.com/coderzh/gohugo.org/tree/master/content/doc)

然后 Fork 该 [Repo](https://github.com/coderzh/gohugo.org) ，通过 Pull Request 的方式提交过来便是。

## Hugo 能做什么

通过 Hugo 你可以快速搭建你的静态网站，比如博客系统、文档介绍、公司主页、产品介绍等等。相对于其他静态网站生成器来说，Hugo 具备如下特点：

  * 极快的页面编译生成速度。（ ~1&nbsp;ms 每页面）
  * 完全跨平台支持，可以运行在 <i class="fa fa-apple"></i>&nbsp;Mac OS&nbsp;X, <i class="fa fa-linux"></i>&nbsp;Linux, <i class="fa fa-windows"></i>&nbsp;Windows, 以及更多!
  * 安装方便 [Installation](/doc/overview/installing/)
  * 本地调试 [Usage](/doc/overview/usage/) 时通过 [LiveReload](/doc/extras/livereload/) 自动即时刷新页面。
  * 完全的皮肤支持。
  * 可以部署在任何的支持 HTTP 的服务器上。

## 从这里开始了解吧……

 * [安装 Hugo](/doc/overview/installing/)
 * [快速开始](/doc/overview/quickstart/)

对了，别忘记了给该 `Repo` 一个 `Star` 。

