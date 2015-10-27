---
categories:
- 搭建指南
keywords:
- Hugo
- 博客
- 搭建
- 部署
- 使用
- 教程
- 说明
- 迁移
tags:
- 入门
date: 2015-08-29T16:52:15+08:00
author: CoderZh
title: 使用hugo搭建个人博客站点
thumbnail: http://7xlx3k.com1.z0.glb.clouddn.com/HugoBlog.jpg-wt
weight: -100

---

Hugo是一个用Go语言编写的静态网站生成器，它使用起来非常简单，相对于Jekyll复杂的安装设置来说，Hugo仅需要一个二进制文件hugo(hugo.exe)即可轻松用于本地调试和生成静态页面。

<!--more-->

Hugo生成静态页面的效率很高，我的260多篇博客文章生成几乎是瞬间完成的，而之前用Jekyll需要等待10秒左右。

Hugo自带watch的调试模式，可以在我修改MarkDown文章之后切换到浏览器，页面会检测到更新并且自动刷新，呈现出最终效果，能极大的提高博客书写效率。再加上Hugo是使用Go语言编写，已经没有任何理由不使用Hugo来代替Jekyll作为我的个人博客站点生成器了。

###  静态网站生成器

什么是静态网站生成器？如果追溯到最早的网站形式，那时候的网页都是静态的，即一个内容不变的html文件放在服务器上，人们通过互联网访问浏览的都是这个一成不变的页面。后来，人们发现，需要和网页进行交互，能根据用户的输入动态呈现出相应的内容，这就是动态网站。那，为什么现在又回归使用静态网站呢？特别是对于博客网站这种特殊的形式。我的理解是：

 1. 访问速度提升。静态页面不需要像动态页面那样经常去查询数据库，而是直接将最终页面内容返回。
 1. 搜素引擎友好。便于搜索引擎索引，比如很多动态网站的页面地址是一样的，只是后面传入的参数不一样，容易让搜索引擎误认为是同一个页面。（虽然不会）
 1. 可以完全抛弃数据库，减少复杂度，将最复杂的一步交给静态网站生成器，自己只专注写作、生成、发布三个步骤。
 1. 博客文章可以以文本文件的方式（MarkDown）在本地维护管理，不需要像之前那样在网页的编辑器里提交到网站数据库。你可以方便的使用github管理你的博客文章，不会丢失，又能追溯到每一次的内容变更。

目前最流行的静态网站生成器是Jekyll，它是github创始人自己实现的一套ruby的静态网站生成器。一经推出，各个程序哥竞相效仿，一时间使用Jekyll搭建自己的博客变成了一件很高逼格的事情。

我也是跟风者之一，在Jekyll之前，我主要也是在博客园写文章，同时独立博客这块也尝试过自己实现的博客程序，用GAE写过，后来又用Tornado写过，都是动态网站。自己实现博客程序，总是在开始一段时间内很狂热，当一切实现完成，细节修缮好之后很快就失去了兴趣。因为你的注意力从只是写文章，经常转移到其他地方去。比如网站不好看，需要去改一改网站的风格样式，修复BUG，加一加功能等等。

所以，如果只是专注于写作，还是找一个稳定一些，提供大量现成皮肤，有稳定专业的组织维护更新的博客系统。使用Jekyll确实让我眼前一亮，原来博客程序可以这么玩。于是利用周末的时间，我将我之前所有的文章，包括博客园的文章，都迁移到了Jekyll上来。自从有了Jekyll，我终于可以开开心心的专注的写文章了，而且使用喜欢的MarkDown格式。：）

在使用了Jekyll一段时间后，它的问题也逐渐暴露出来：静态页面生成的效率不够高。因为我把以前的文章都导入了过来，一共有260来篇，每次编辑文章后，需要等等10秒，待它将所有页面检查并生成完成之后，才能看到最终的效果。这是我最不能忍受的一点。对于初学者，Jekyll还有很多问题，比如环境搭建非常复杂，导致使用Jekyll的人大都是一些喜欢折腾，不怕折腾的程序哥。

随着Go1.5版本的发布，让我意识到是时候好好玩一玩Go语言了。使用Go语言实现的静态网站生成器Hugo（雨果）立即吸引了我，它解决了我最大的痛点：生成的效率。文档、社区各方面的支持都不错，使用起来非常简单，各种皮肤直接套用，于是我又开始了Jekyll迁移到Hugo的漫漫长路。（一个周末的时间）

~~即使迁移到了Hugo，我还保留着原来的Jekyll博客，只是不会再更新了，用来怀念？还是哪天突然又跳回来也不一定。~~

~~我的Jeklly博客：[http://jekyll.coderzh.com](http://jekyll.coderzh.com)~~ Update(2015-09-20): 觉得没啥用了，还是干掉了

我的Hugo博客：[http://blog.coderzh.com/](http://blog.coderzh.com/)

### GitHub Pages

使用静态网站生成器生成好静态页面之后，需要把文件放到服务器上供别人浏览。比较传统的方式是租用VPS虚拟服务器，比如：linode、digitalocean。将生成好的静态页面手工上传到服务器上。如果你习惯这种方式部署，推荐你使用[digitalocean](https://www.digitalocean.com/?refcode=e131e2bba197)。

当然， 还有更好的方式，就是直接把网站托管到GitHub Pages。你只需要在GitHub上创建一个项目，然后将生成出来的静态页面文件push到这个项目的gh-pages分支，保证根目录有一个index.html文件即可。这样，一个免费、无限流量的博客系统就搭建完成了。同时，通过github你可以方便对博客文章进行管理和追踪。

### Hugo

前面的铺垫介绍的差不多了，该主角上场了。Hugo是什么？它主要做了什么？

 1. Hugo只有一个二进制文件（比如Windows里只是一个hugo.exe）
 1. Hugo可以将你写好的MarkDown格式的文章自动转换为静态的网页。
 1. Hugo内置web服务器，可以方便的用于本地调试。

### Hello Hugo

Hugo官方主页：[https://gohugo.io/](https://gohugo.io/)

Hugo的安装方式有两种，一种是直接下载编译好的Hugo二进制文件。如果只是使用Hugo推荐用这种方式。另一种方式是获取Hugo的源码，自己编译。由于各种不可预料的网络问题，第二种方式不是那么轻易能成功，虽然最后我还是折腾出来了。

Hugo二进制下载地址：[https://github.com/spf13/hugo/releases](https://github.com/spf13/hugo/releases)

下载下来后，只有一个叫hugo或者hugo.exe的程序，接下来开始生成自己的站点：

```bash
$ hugo new site mysite
```

然后hugo会自动生成这样一个目录结构：

```bash
  ▸ archetypes/
  ▸ content/
  ▸ layouts/
  ▸ static/
    config.toml
```

简要介绍一下，config.toml是网站的配置文件，这是一个TOML文件，全称是Tom's Obvious, Minimal Language，这是它的作者GitHub联合创始人Tom Preston-Werner 觉得YAML不够优雅，捣鼓出来的一个新格式。如果你不喜欢这种格式，你可以将config.toml替换为YAML格式的config.yaml，或者json格式的config.json。hugo都支持。

content目录里放的是你写的markdown文章，layouts目录里放的是网站的模板文件，static目录里放的是一些图片、css、js等资源。

进入生成的site目录：

```bash
$ cd mysite
```

创建一个页面：

```bash
$ hugo new about.md
```

如果是博客日志，最好将md文件放在content的post目录里。

```bash
$ hugo new post/first.md
```

执行完后，会在content/post目录自动生成一个MarkDown格式的first.md文件：

```
+++
date = "2015-01-08T08:36:54-07:00"
draft = true
title = "first"
 
+++
```

+++可以替换为Jekyll一样的\-\-\-，里面的内容是这篇文章的一些信息。下面就可以开始写你的文章内容，比如：

```
+++
date = "2015-01-08T08:36:54-07:00"
draft = true
title = "first"
 
+++

### Hello Hugo

 1. aaa
 1. bbb
 1. ccc

```

OK，刚才的about.md也有内容，该看看最后的效果了。然后你屁颠屁颠的使用hugo server启动，打开浏览器里一看，发现毛都没有！这是肿么了！

这是Hugo对初学者非常不友好的地方，默认生成的网站是没有任何皮肤模板的。为了看看第一个写的示例，还得去Github上把一个网页模板下载下来。如果你网络够好，网速够快，你可以在刚才的目录将Hugo官方的所有模板都下载下来：

```bash
$ git clone --recursive https://github.com/spf13/hugoThemes themes
```

我尝试过，也失败过，且从未成功一次性将所有的模板下载下来。所以，我们还是老老实实只下载其中一个模板来看看效果吧：

```bash
$ cd themes
$ git clone https://github.com/spf13/hyde.git
```

启动本地调试：

```bash
$ hugo server --theme=hyde --buildDrafts --watch
```

浏览器里打开：http://127.0.0.1:1313

![HugoFirstPost](http://7xlx3k.com1.z0.glb.clouddn.com/HugoFirstPost.png-w)

--watch或者-w 选项打开的话，将会监控到文章的改动从而自动去刷新浏览器，不需要自己手工去刷新浏览器，非常方便。

如果你看了上面的说明已经有冲动去试一试Hugo了，我的目的也算达到了，接下来你需要的只是查看官方的说明文档就够了，所以具体的一些设置我就不重复了。

官方文档：[https://gohugo.io/overview/introduction/](https://gohugo.io/overview/introduction/)

皮肤列表：[https://github.com/spf13/hugoThemes](https://github.com/spf13/hugoThemes)

常用文档：

  1. [Configuring Hugo](https://gohugo.io/overview/configuration/)
  1. [Front Matter](https://gohugo.io/content/front-matter/)
  1. [Menus](https://gohugo.io/extras/menus/)
  1. [Template Variables](https://gohugo.io/templates/variables/)
  1. [Hosting on GitHub Pages](https://gohugo.io/tutorials/github-pages-blog/)

### 遇到的问题

 1. 默认的ServerSide的代码着色会有问题，有些字的颜色会和背景色一样导致看不见。  
  解决方法：使用ClientSide的代码着色方案即可解决。（见：[Client-side Syntax Highlighting](https://gohugo.io/extras/highlighting/#client-side:c4210b265c792cac9a6cf6a5f53b671d)）
 1. URL全部被转成了小写，如果是旧博客迁移过来，将是无法接受的。  
  解决方法：~~我是直接改了Hugo的代码，将URL强制转换为小写那段逻辑去掉了，之后考虑在config里提供配置开关，然后给Hugo提一个PR。如果是Windows用户可以直接[https://github.com/coderzh/ConvertToHugo](https://github.com/coderzh/ConvertToHugo) 下载到我修改后的版本myhugo.exe。~~  
  Update(2015-09-03): 已经提交[PR](https://github.com/spf13/hugo/pull/1392)并[commit](https://github.com/spf13/hugo/commit/52d94fa67578f6b63035e73b236ca8abd40d0006)到Hugo，最新版本只需要在config里增加：  
  `disablePathToLower: true`
 1. 文章的内容里不能像Jekyll一样可以内嵌代码模板了。最终会生成哪些页面，有一套相对固定而复杂的规则，你会发现想创建一个自定义界面会非常的困难。  
  解决方法：无，看文档，了解它的规则。博客程序一般也不需要特别的自定义界面。Hugo本身已经支持了类似posts, tags, categories等内容聚合的页面，同时支持rss.xml，404.html等。如果你的博客程序复杂到需要其他的页面，好好想想是否必须吧。
 1. 如何将rss.xml替换为feed.xml？  
  解决方法：在config.yaml里加入：
 ```
 rssuri: "feed.xml"
 ```


### Jekyll迁移到Hugo

Jekyll的文章内容迁移到Hugo中，大部分内容是兼容的，但也有一些地方是不兼容的。主要有以下几个地方需要修改：

 1. Jekyll文章能从文件名里的日期部分读取到日期，并将剩余的部分当做的页面url的名称，比如：2015-08-29-first.md。而Hugo只认md文件里的date字段，url的名称如果用文件名的话将会使用完整的文件名（不会去除日期部分）。为了兼容，必须在md的Front Matter里写入url字段，用来说明该页面的相对url地址，从而保持兼容。
 1. tags, categories等字段必须用列表的方式，不像Jekyll中那样随意了。
 1. {% raw %} {% endraw %}将不需要了。
 1. {% highlight ruby %} 变成了{{\< highlight ruby >}} 。不过我还是推荐使用\`\`\` ruby \`\`\`形式，然后使用ClientSide的Highlight，这样两边都兼容。
 1. 需要将Jekyll里的public里的文件拷贝到Hugo的static目录里。
 1. Jekyll的文章必须放到Hugo的content/post目录里。
 1. Jekyll只需要push文章内容到github，服务器会自动生成静态页面。毕竟是github的亲儿子。而Hugo需要你将生成的public目录里的内容做为gh-pages分支push上去。具体的简便的方法见：[Hosting on GitHub Pages](https://gohugo.io/tutorials/github-pages-blog/)

假如你之前的博客和我一样是Jekyll的，可以尝试使用我写的[ConvertToHugo.py](https://github.com/coderzh/ConvertToHugo/blob/master/ConvertToHugo.py) 工具。这个转换工具逻辑相对比较简单和清晰，如果满足不了你的需求你也可以轻易在此基础上做些修改，如果我能收到PR当然是最好了。 

反正，我是完全使用CovertToHugo.py将我原来的Jekyll博客全部转过来了。而且，我找到了一个非常炫酷的主题，并且在此基础稍微修改下。

![HugoBlog](http://7xlx3k.com1.z0.glb.clouddn.com/HugoBlog.jpg)

Blog地址：[http://blog.coderzh.com](http://blog.coderzh.com)

Github: [https://github.com/coderzh/coderzh-hugo-blog](https://github.com/coderzh/coderzh-hugo-blog)

**Update(2015-10-08) 国庆在家给Hugo提了个PR，已经将Jekyll迁移功能集成到了Hugo里，pull最新的Hugo代码，编译Hugo，然后可以执行：**

```bash
$ hugo import jekyll YourJekyllDir TargetDir
```

将自动把你的Jekyll站点转换成Hugo的站点。欢迎试用并反馈。

### 关于部署

假设你需要部署在 `GitHub Pages` 上，首先在GitHub上创建一个Repository，命名为：`coderzh.github.io` （coderzh替换为你的github用户名）。

在站点根目录执行 `Hugo` 命令生成最终页面：

```bash
$ hugo --theme=hyde --baseUrl="http://coderzh.github.io/"
```

如果一切顺利，所有静态页面都会生成到 `public` 目录，将pubilc目录里所有文件 `push` 到刚创建的Repository的 `master` 分支。

```bash
$ cd public
$ git init
$ git remote add origin https://github.com/coderzh/coderzh.github.io.git
$ git add -A
$ git commit -m "first commit"
$ git push -u origin master
```

浏览器里访问：`http://coderzh.github.io/`

自动部署的脚本可以参考我的Repo里的`deploy.py`脚本：[https://github.com/coderzh/coderzh-hugo-blog](https://github.com/coderzh/coderzh-hugo-blog)

全自动化部署到GitHub和GitCafe，见我的另一篇博客：[通过webhook将Hugo自动部署至GitHub Pages和GitCafe Pages](http://blog.coderzh.com/2015/09/13/use-webhook-automated-deploy-hugo/)


