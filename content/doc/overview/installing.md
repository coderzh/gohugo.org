---
aliases:
- /doc/installing/
date: 2013-07-01
lastmod: 2016-01-04
linktitle: Installing Hugo
menu:
  main:
    parent: getting started
next: /doc/overview/usage
prev: /doc/overview/quickstart
title: 安装 Hugo
weight: 20
translated: true
---

Hugo 是用 [Go][] 语言写的，支持多个平台。

最新的 release 版本可以在 [Hugo Releases](https://github.com/spf13/hugo/releases) 找到。 我们提供预构建好的二进制包括
<i class="fa fa-windows"></i>&nbsp;Windows,
<i class="fa fa-linux"></i>&nbsp;Linux,
<i class="fa freebsd-19px"></i>&nbsp;FreeBSD
和 <i class="fa fa-apple"></i>&nbsp;OS&nbsp;X (Darwin)
for x64, i386 和 ARM architectures.

你可以使用 Go 编译器工具链源码编译 Hugo，比如在其他的操作系统如 DragonFly BSD, OpenBSD, Plan&nbsp;9 和 Solaris 。
在 http://golang.org/doc/install/source 查看完整的操作系统和编译架构的支持列表。


## 安装 Hugo （二进制）

安装过程非常简单。只需要下载合适你系统版本的 [Hugo 二进制](https://github.com/spf13/hugo/releases) 。
下载完毕后它可以在任何地方运行。你并不需要把它安装到一个全局的地方。这适用于共享一台主机和系统并且没有特别权限的账号的情况。

更理想的是，为了更方便的使用，你应该把它安装到你的 `PATH` 环境变量所在的位置。

在 OS&nbsp;X，如果你有 [Homebrew](http://brew.sh/) ，安装过程就更简单了：只需要运行 `brew install hugo` 。

### 安装 Pygments （可选）

Hugo 有一个 *可选的* 关于源代码高亮（Pygments）的额外依赖。

如果你想要使用 [highlight shortcode](/doc/extras/highlighting/) 源代码高亮，
你必须安装基于 Python 的 Pygments 程序。安装过程见 [Pygments home page](http://pygments.org/) 。

## 升级 Hugo

升级 Hugo 只需要简单的下载和替换之前你放在 `PATH` 路径的二进制文件。

## 源码安装

### 下载并源码编译的必备工具

* [Git](http://git-scm.com/)
* [Mercurial](http://mercurial.selenic.com/)
* [Go][] 1.3+ (Go 1.4+ on Windows, 见 Go [Issue #8090](https://code.google.com/p/go/issues/detail?id=8090))

### 直接从 GitHub 获取

    $ export GOPATH=$HOME/go
    $ go get -v github.com/spf13/hugo

`go get` 将会获取 Hugo 以及所有依赖的库到你的 `$GOPATH/src` 目录，同时编译所有代码生成最终的 `hugo` （或 `hugo.exe`） 二进制文件，这就全部准备好了。

你可以使用 `-u` 参数执行 `go get` 用来更新 Hugo 的所有依赖。

    $ go get -u -v github.com/spf13/hugo

## 贡献 Hugo

请见： [贡献指引](/doc/community/contributing/) 。

[Go]: http://golang.org/
