---
author: Michael Henderson
date: 2015-02-22
lastmod: 2016-01-04
linktitle: Installing on Mac
menu:
  main:
    parent: tutorials
next: /doc/tutorials/installing-on-windows
prev: /doc/tutorials/github-pages-blog
title: Installing on a Mac
toc: true
weight: 10
translated: true
---

# 在 Mac 上安装 Hugo

本篇教程的目标是详细讲解在 Mac 电脑上安装 Hugo 的方法。

## 假设

1. 你知道如何打开一个终端窗口。
2. 你运行的是一个现代64位的 Mac。
3. 你将把 `~/Sites` 作为网站的起点。

## 选择方法

在你的 Mac 电脑上安装 Hugo 有三种方法：使用 `brew` 工具， 从发行版或者源码安装。
其中并没有“最好”的方法。你应该根据你的实际情况选择最佳的方法。

各个方法都有自己的优点和缺点。

1. `Brew` 是维护起来最简单轻便的方法。缺点也不严重。默认的安装包将会是
   最近的发布版本，因此在下一个版本发布前，将无法提供缺陷修复的版本（除非
   你使用 `--HEAD` 选项来安装）。`brew` 上的发行版可能会滞后几天，因为它需要
   和其他组配合。尽管如此，如果你希望使用一个稳健、被广泛使用的源码，我还是推荐
   使用 `brew`。它不仅有效，而且非常容易升级。

2. 下载 tarball 并安装也非常的简单。你需要有少量的命令行技能。
   升级也很容易。你仅需要在新的二进制上重复相同的步骤就行。这可以赋予你一定的灵活性，在电脑上
   拥有不同的版本。如果你想使用 `brew`， 那么选择使用二进制安装是个好主意。 

3. 从源码编译是使用最多的方法。这种方法的好处是你不需要等待一个添加了新特性或者修复缺陷的新发布版本。
   而缺点是需要多花一些时间进行配置。它的配置并不多，但比另两种方法多一些。

既然这篇教程是“初学者”指南，我将会详细讲解前两种方法，并快速阐述第三种方法。

## Brew

### 第一步：如果你还没安装 `brew` 的话， 安装它

访问 `brew` 的网站，http://brew.sh/， 并按照那里的指导安装。最重要的一步是：

```
ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

在我运行这条命令时，遇到了文件夹权限的问题。Google 上的搜索结果引导我浏览一些页面，通过更新 `/usr/local` 文件夹权限的方法解决问题。看起来挺可怕的， 但从那之后就一直没问题。

### 第二步：运行 `brew` 命令来安装 `hugo`

如果你想使用最新的可能有缺陷的版本，把安装命令 `brew install hugo` 替换为 `brew install hugo --HEAD`。

```
$ brew install hugo
==> Downloading https://homebrew.bintray.com/bottles/hugo-0.13_1.yosemite.bottle.tar.gz
######################################################################## 100.0%
==> Pouring hugo-0.13_1.yosemite.bottle.tar.gz
🍺  /usr/local/Cellar/hugo/0.13_1: 4 files,  14M
```

`Brew` 应该已经更新了路径来包含 Hugo。可以通过打开一个新的终端窗口，并运行下面的几条命令来确认：

```
$ # show the location of the hugo executable
$ which hugo
/usr/local/bin/hugo

$ # show the installed version
$ ls -l $( which hugo )
lrwxr-xr-x  1 mdhender admin  30 Mar 28 22:19 /usr/local/bin/hugo -> ../Cellar/hugo/0.13_1/bin/hugo

$ # verify that hugo runs correctly
$ hugo version
Hugo Static Site Generator v0.13 BuildDate: 2015-03-09T21:34:47-05:00
```

### 第三步：完成

你已经安装了 Hugo。现在你需要设置你的网站。阅读[快速入门](/doc/overview/quickstart/)，
搜索剩下的文档，如果你还有问题，[尽管提问!](http://discuss.gohugo.io/ "论坛")

## 从 tarball 安装

### 第一步：确定安装位置

当使用 tarball 安装时，你需要决定是否把二进制文件安装在 `/usr/local/bin` 还是你的主目录。它有三个预装位置：

1. 安装到 `/usr/local/bin` ，这样系统中的所有用户都可以使用。这是个好主意，因为这个位置是一般执行文件的标准安装位置。缺点是你可能需要提升权限才能把软件安装到这个地方。同时，如果你的系统有多个用户，他们将使用同一个版本的软件。有时如果你想要尝试一个新的发行版时，这可能会是个问题。

2. 安装到 `~/bin` 这样只有你可以执行它。这是个好主意，因为易于安装，易于维护，而且不需要提升权限。缺点是只有你可以运行 Hugo。如果你的网站有其他的用户，他们将不得不维护他们自己的那一份软件。这样人们可能会使用不同的版本。当然，这确实会让你体验不同的发行版变得更加方便。

3. 安装到你的 `sites` 文件夹。如果你只构建一个网站的话，这也是个不错的选择。这样可以让所有的东西都在一个地方。如果你想要尝试新的发行版，你仅仅需要把整个网站复制一份，更新其中的 Hugo 可执行文件，这样就可以了。

所有这三个地方都没有问题。我将描述第二种方法，这很大程度上是因为我喜欢这种方式。

### 第二步：下载 tarball

1. 在浏览器中打开 <https://github.com/spf13/hugo/releases>。

2. 滚动页面，在当前的发布版本中，寻找写着“最新发行版本”的绿色标签。

3. 下载 Mac 平台的当前的 tarball。名字应该形如 `hugo_X.YY_darwin_amd64.zip`，其中 `X.YY` 是发行版号。

4. 默认情况下，这个 tarball 会保存在 `~/Downloads` 文件夹内。如果你选择使用一个不同的位置，你将需要按照下面的步骤来修改。

### 第三步：验证下载

确认下载过程中，tarball 没有损坏：

```
$ tar tvf ~/Downloads/hugo_0.13_darwin_amd64.zip
-rwxrwxrwx  0 0      0           0 Feb 22 04:02 hugo_0.13_darwin_amd64/hugo_0.13_darwin_amd64
-rwxrwxrwx  0 0      0           0 Feb 22 03:24 hugo_0.13_darwin_amd64/README.md
-rwxrwxrwx  0 0      0           0 Jan 30 18:48 hugo_0.13_darwin_amd64/LICENSE.md
```

`.md` 文件是文档。其他的文件是可执行文件。

### 第四步：安装到你的 bin 文件夹

```
$ # create the directory if needed
$ mkdir -p ~/bin

$ # make it the working directory
$ cd ~/bin

$ # extract the tarball
$ unzip ~/Downloads/hugo_0.13_darwin_amd64.zip
Archive:  hugo_0.13_darwin_amd64.zip
  inflating: hugo_0.13_darwin_amd64/hugo_0.13_darwin_amd64
  inflating: hugo_0.13_darwin_amd64/README.md
  inflating: hugo_0.13_darwin_amd64/LICENSE.md

$ ls -l
total 7704
lrwxr-xr-x  1 mdhender  staff       22 Sep 29 13:34 hugo -> hugo_0.12_darwin_amd/hugo_0.12_darwin_amd64
drwxr-xr-x@ 1 mdhender  staff      102 Sep  1 14:17 hugo_0.12_darwin_amd64
drwxrwxr-x@ 5 mdhender  staff      170 Mar 28 22:46 hugo_0.13_darwin_amd64
-rw-r-----@ 1 mdhender  staff  3942651 Mar 28 22:45 hugo_0.13_darwin_amd64.zip

$ ls -l hugo_0.13_darwin_amd64
total 27560
-rw-r--r--@ 1 mdhender  staff      2707 Jan 30 18:48 LICENSE.md
-rw-r--r--@ 1 mdhender  staff      6748 Feb 22 03:24 README.md
-rwxr-xr-x@ 1 mdhender  staff  14095060 Feb 22 04:02 hugo_0.13_darwin_amd64
```

我已经运行了 Hugo v0.12，这样你就可以知道要如何设置它。对于 v0.13 版本，在我们完成后也是一样的。

```
$ # create the link to the real executable
$ rm -f hugo
$ ln -s hugo_0.13_darwin_amd64/hugo_0.13_darwin_amd64 hugo
$ ls -l
total 7704
lrwxr-xr-x  1 mdhender  staff       22 Mar 28 22:49 hugo -> hugo_0.13_darwin_amd/hugo_0.12_darwin_amd64
drwxr-xr-x@ 1 mdhender  staff      102 Sep  1 14:17 hugo_0.12_darwin_amd64
drwxrwxr-x@ 5 mdhender  staff      170 Mar 28 22:46 hugo_0.13_darwin_amd64

$ # verify that it runs
$ ./hugo version
Hugo Static Site Generator v0.13 BuildDate: 2015-02-22T04:02:30-06:00
```

你可能需要在 `PATH` 变量中添加你的 bin 文件夹路径。`which` 命令将会为我们做这个检查。如果它可以找到 `Hugo`，就会打印出它的完整路径。否则，它什么都不会打印出来。

```
$ # check if hugo is in the path
$ which hugo
/Users/mdhender/bin/hugo
```

如果 `hugo` 不在你的 `PATH` 中，可以通过更新 `~/.bash_profile` 文件来添加。首先，启动一个编辑器：

```
$ nano ~/.bash_profile
```

添加一行来更新你的 `PATH` 变量：

```
export PATH=$PATH:$HOME/bin
```

然后通过点击 Control-X 来保存文件，然后点击 Y 来保存文件并返回到提示符处。

关闭终端，然后打开一个新的终端，以让修改后的 profile 文件生效。再次运行 `which hugo` 命令来验证。

### 第五步：完成

你已经安装好了 Hugo。现在你需要设置你的网站。阅读
[快速入门](/doc/overview/quickstart/)，搜索剩余的文档，如果你依然有问题
[尽管提问!](http://discuss.gohugo.io/ "论坛")

## 从源代码编译安装

如果你想要自己编译 Hugo，你将需要 [Go](http://golang.org)，它也可以从 Homebrew 中安装：`brew install go`。

### 第一步：获取源代码

如果你想要编译特定的版本，访问 <https://github.com/spf13/hugo/releases> 并下载你选择版本的源代码。
如果你想要编译最新修改过的 Hugo（可能包含缺陷），可以克隆 Hugo 的代码仓库：

```
git clone https://github.com/spf13/hugo
```

### 第二步：编译

在你的工作目录中创建包含源码的文件夹，然后获取 Hugo 的依赖：

```
mkdir -p src/github.com/spf13
ln -sf $(pwd) src/github.com/spf13/hugo

# set the build path for Go
export GOPATH=$(pwd)

go get
```

这将会获取依赖包的最新完整版本，因此如果 Hugo 编译失败，可能是因为某个依赖
的作者引入了一个破坏性的改动。

然后编译：

```
go build -o hugo main.go
```

然后把 `hugo` 可执行文件放到 `$PATH` 中的某处。

### 第三步：完成

你大概知道这之后你要做什么。
