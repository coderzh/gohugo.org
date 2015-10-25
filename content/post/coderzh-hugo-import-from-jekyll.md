---
categories:
- 功能改进
tags:
- Jekyll
- 迁移
date: 2015-10-11T10:06:11+08:00
description: "Jekyll一键导入到Hugo的方法"
keywords:
- hugo
- import
- migrate
- jekyll
- 导入
- 迁移
author: CoderZh
title: "使用Hugo Import一键迁移Jekyll"
thumbnail: http://7xlx3k.com1.z0.glb.clouddn.com/HugoImportJekyll.png-wt
weight: -1

---

![HugoImportJekyll](http://7xlx3k.com1.z0.glb.clouddn.com/HugoImportJekyll.png-wt)

<!--more-->

国庆长假期间，给Hugo提了几个PR，其中最主要的一个是给Hugo增加了内置的Jekyll迁移工具。这样之前的Jekyll用户可以通过一个简单的命令就可以将网站转化成Hugo站点：

```bash
$ hugo import jekyll your-jekyll-dir target-dir
```

得益于Hugo代码可读性非常好，原来计划需要几天来完成的功能1天左右的时间就完成了。完整的PR过程请见：[https://github.com/spf13/hugo/pull/1469](https://github.com/spf13/hugo/pull/1469)

### Hugo Import

主要实现了将原来的Jeklly网站一键转换为Hugo网站，具体实现了如下功能：

 1. 生成新的Hugo站点结构。(hugo new site)
 2. 读取Jeklly的_config.yml，相应内容转化到Hugo的config.yaml。
 3. 转换所有MarkDown的文章。（具体转换规则见下文）
 4. 拷贝Jekyll其他目录及文件到Hugo的static目录。

@bep在他的OS X上做了一些测试：

```bash
$ hugo import jekyll qrush.github.com qr2
  Importing...
  Congratulations! 72 posts imported!
  Now, start Hugo by yourself:
  $ git clone https://github.com/spf13/herring-cove.git qr2/themes/herring-cove
  $ cd qr2
  $ hugo server -w --theme=herring-cove
```

### MarkDown文章转换规则

Hugo和Jekyll都是使用MarkDown来写文章的，文章内容基本上是兼容的，除了一些FrontMatter的细节和局部细节。Hugo Import Jekyll主要的转换规则如下：

 1. 保持原Jekyll文章目录结构及文件名不变，将文章拷贝到了content/post目录下。
 2. 尽量保持原Jekyll文章链接地址不变，在文章的FrontMatter自动填入url字段。
 3. date字段转换成Hugo要求的time.RFC3339格式。
 4. draft字段自动生成。
 5. 删除layout字段。
 6. category字段转换成categories字段。
 7. excerpt_separator智能替换。(比如： \<\!\-\-more\-\-\>)
 8. 删除`{% raw %}{% endraw %}`标签。
 9. 替换`{% highlight %}{% endhighlight %}`标签为 \{\{\< highlight \>\}\}\{\{\< / highlight \>\}\}

### 代码

import.go: 

[https://github.com/spf13/hugo/blob/master/commands/import.go](https://github.com/spf13/hugo/blob/master/commands/import.go)

import_test.go: 

[https://github.com/spf13/hugo/blob/master/commands/import_test.go](https://github.com/spf13/hugo/blob/master/commands/import_test.go)

我fork的Hugo Repo（我的最新改动都在这里）：

[https://github.com/coderzh/hugo](https://github.com/coderzh/hugo)

### 最后

这个功能将会在下个正式版推出，现在需要使用的同学自行pull最新的Hugo代码编译吧：

```bash
$ go get -u -v github.com/spf13/hugo
```

欢迎试用，并且提出意见或建议！谢谢！
