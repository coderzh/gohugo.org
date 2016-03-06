---
author: Michael Henderson
date: 2015-03-30
lastmod: 2015-12-23
linktitle: Installing on Windows
menu:
  main:
    parent: tutorials
next: /doc/tutorials/mathjax
prev: /doc/tutorials/installing-on-mac
title: Installing on Windows
toc: true
weight: 10
translated: true
---

# 在 Windows 上安装 Hugo

本篇教程的目标是详细讲解在 Windows 电脑上安装 Hugo 的方法。

## 假设

1. 你知道如何打开一个命令提示窗口。
2. 你运行的是一个现代64位的 Windows。
3. 你的网站地址是 `example.com`。
4. 你将使用 `D:\Hugo\Sites` 作为网站的起点。
5. 你将使用 `D:\Hugo\bin` 存储可执行文件。

## 设置你的文件夹

你将需要一个存储 Hugo 可执行文件、博客内容（你创建的的那些文件），以及生成文件（Hugo 为你创建的 HTML）的地方。

1. 打开 Windows Explorer。
2. 创建一个新的文件夹，`D:\Hugo`。
3. 创建一个新的文件夹，`D:\Hugo\bin`。
4. 创建一个新的文件夹，`D:\Hugo\Sites`。

## 下载预先编译好的 Windows 版本的 Hugo 可执行文件

使用 go 编译 Hugo 的一个优势就是仅有一个二进制文件。你不需要运行安装程序来使用它。相反，你需要把这个二进制文件复制到你的硬盘上。我假设你将把它放在 `D:\Hugo\bin` 文件夹内。如果你选择放在其它的位置，你需要在下面的命令中替换为那些路径。

1. 在浏览器中打开 https://github.com/spf13/hugo/releases。
2. 当前的版本是 hugo_0.13_windows_amd64.zip。
3. 下载那个 ZIP 文件，并保存到 `D:\Hugo\bin` 文件夹中。
4. 在 Windows Explorer 中找到那个 ZIP 文件，并从中提取所有的文件。
5. 你应该可以看到一个 `hugo_0.13_windows_amd64.exe` 文件。
6. 把那个文件重命名为 `hugo.exe`。
7. 确保 `hugo.exe` 文件在 `D:\Hugo\bin` 文件夹。（有可能提取过程会把它放在一个子文件夹中。如果确实这样的话，使用 Windows Explorer 把它移动到 `D:\Hugo\bin`。）
8. 使用 `D:\Hugo\bin>set PATH=%PATH%;D:\Hugo\bin` 把 hugo.exe 可执行文件添加到你的 PATH路径中。

## 验证可执行文件

运行几个命令来验证可执行命令可以运行，然后构建一个示例网站作为起点。

1. 打开一个命令提示符窗口。
2. 在提示符中，输入 `hugo help` 并按下 Enter 键。你看到的输出应该以下面的文字开始：
    {{< nohighlight >}}A Fast and Flexible Static Site Generator built with love by spf13 and friends in Go. Complete documentation is available at http://gohugo.io
{{< /nohighlight >}}
如果你看到了，那么安装就完成了。如果你没有看到，仔细检查下你放置 `hugo.exe` 文件的路径，在 PATH 变量中输入了正确的路径。如果你依旧没有得到那个输出结果，在 Hugo 讨论列表中把你的命令和输出贴出来。
3. 在命令提示符中，跳转当前目录到 `Sites` 文件夹。
    {{< nohighlight >}}C:\Program Files> cd D:\Hugo\Sites
C:\Program Files> D:
D:\Hugo\Sites>
{{< /nohighlight >}}
4. 运行命令来生成一个新的网站。我使用 `example.com` 作为网站的名字。
    {{< nohighlight >}}D:\Hugo\Sites> hugo new site example.com
{{< /nohighlight >}}

5. 你现在应该拥有一个叫做 D:\Hugo\Sites\example.com 的文件夹。进入这个文件夹，并列出文件夹中的内容。你应该得到类似下面的输出内容：
    {{< nohighlight >}}D:\Hugo\Sites>cd example.com
D:\Hugo\Sites\example.com>dir
 Directory of D:\hugo\sites\example.com

04/13/2015  10:44 PM    <DIR>          .
04/13/2015  10:44 PM    <DIR>          ..
04/13/2015  10:44 PM    <DIR>          archetypes
04/13/2015  10:44 PM                83 config.toml
04/13/2015  10:44 PM    <DIR>          content
04/13/2015  10:44 PM    <DIR>          data
04/13/2015  10:44 PM    <DIR>          layouts
04/13/2015  10:44 PM    <DIR>          static
               1 File(s)             83 bytes
               7 Dir(s)   6,273,331,200 bytes free

{{< /nohighlight >}}

现在你就拥有了安装好的 Hugo 和工作的网站。你需要添加一个布局（或者主题），然后创建一些博客内容。前往 http://gohugo.io/overview/quickstart/ 了解相关的步骤。

## 故障排查

@dhersam 录制了一个常见问题的不错的视频，地址在 https://www.youtube.com/watch?v=c8fJIRNChmU
