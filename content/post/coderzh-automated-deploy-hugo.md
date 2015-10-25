---
categories:
- 部署上线
tags:
keywords:
- GitCafe
keywords:
- GitCafe
- 部署
date: 2015-09-13T08:48:52+08:00
description: "通过webhook将Hugo自动部署至GitHub Pages和GitCafe Pages"
author: CoderZh
title: "通过webhook将Hugo自动部署至GitHub Pages和GitCafe Pages"
thumbnail: http://7xlx3k.com1.z0.glb.clouddn.com/AutoDeploy.png-wt
weight: -1

---

本文的主要内容如标题所示，通过webhook将Hugo自动部署至GitHub Pages和GitCafe Pages。如果你正好有这个需求，看这篇文章正好，可以节省你不少时间。如果不是，了解一下也无妨。

<!--more-->

首先，必须解释一下，为什么需要自动部署，以及为什么需要需要同时部署到GitHub Pages和GitCafe Pages。

### 为什么要自动部署

使用Hugo生成的静态页面是在public文件夹里，部署的时候需要把public文件夹里的内容push到GitHub的gh-pages分支里。每次写完文章，除了push markdown格式的文章，还需要单独push生成的public文件夹里的东西，步骤稍显麻烦。

之前参照了官方的做法，使用subtree来push public，步骤简化不少。然而，这还是不够简单。因为每次修改文章之后，必须依赖一个脚本才能正确提交和部署。假如你在手机里浏览时，发现一个错别字，顺手就在GitHub的Web界面就把错别字改了，然而这样并不会重新生成静态页面和部署。有些不方便。

Hugo生成静态页面和部署的过程应该让机器自动来完成。**写作应该是一个相对单纯的事情，使用Hugo的人应该更专注于写作。**

### 为什么需要同时部署到GitHub Pages和GitCafe Pages

大公司很喜欢的一个词：容灾。GitHub出现不可访问的事情在国内也是常有的，而GitCafe作为国内的代码托管厂商，是否是一个稳定的存在也不好说。所以，将网站同时部署到这两个上面。通过DnsPod里CNAME设置线路“国内”和“国外”，不仅起到了任何一个挂掉，另一个可以继续工作的目的，还起到了CDN就近访问的作用。

使用GitCafe还有另外一个原因。GitHub Pages拒绝了一切百度的爬虫，所以，百度无法索引到GitHub Pages的网页。对于国内的搜索市场来说，百度的份额还是比较大的，虽然我认为看我的博客的人都不应该使用百度，但现实总是残酷的。如果希望网站被百度收录，就必须放到GitHub以外的地方。GitCafe就是一种比较好的选择。

[http://gitcafe.com](http://gitcafe.com)

如果，你只是希望使用官方的Hugo自动化部署到GitHub Pages，下面的内容你可以不用看了。你可以直接使用Wercker的服务来自动部署。

文档见：[
http://gohugo.io/tutorials/automated-deployments/](http://gohugo.io/tutorials/automated-deployments/)

由于Wercker还不支持GitCafe的部署，以及我需要使用特定的修改版本的Hugo来生成静态网页，并且希望这些步骤比较可控，所以，还是自己来折腾整个过程吧。

### webhook

webhook是GitHub上提供的Git的一种Hook机制，当代码发生变化时，比如代码被Push到GitHub的Repo时，GitHub会自动请求一个你指定的网页，并且把变更相关的参数都传递过来。入口在Repo的Settings - webhooks & services

![webhook](http://7xlx3k.com1.z0.glb.clouddn.com/WebHook.png-w)

说明文档：[https://developer.github.com/webhooks/](https://developer.github.com/webhooks/)

借助webhook的机制，我们就可实现当有新的文章Push之后，自动通知远程的一台机器执行一个脚本，脚本的内容就是生成静态页面和Push部署到最终的服务器。

webhook的Server接收webhook通知，然后执行一个脚本。这样的需求太普遍了，以至于完全不需要自己来实现。在GitHub里搜webhook可以搜出来很多。我主要挑选了Go语言的版本。主要有两个：

 * [https://github.com/qiniu/webhook](https://github.com/qiniu/webhook)
 * [https://github.com/adnanh/webhook](https://github.com/adnanh/webhook)

第一个是七牛写的，代码很简单，用法也很简单。开始打算用七牛的版本。最后调试的时候发现json解析失败，完全不可用啊！有点坑爹。于是换成了第二个，这个Repo有200多个Star。还是靠谱很多，最后部署，调试，非常顺利。

用法也很简单，首先安装webhook：

```
$ go get github.com/adnanh/webhook
```

写一个配置文件hooks.json，里面指定需要执行的脚本：

```
[
  {
    "id": "redeploy-webhook",
    "execute-command": "/var/scripts/redeploy.sh",
    "command-working-directory": "/var/webhook"
  }
]
```

指定端口启动：

```
$ /path/to/webhook -hooks hooks.json -port=9876 -verbose
```

然后它将接受webhook地址：（把它填到GitHub里的webhook里）

```
http://yourserver:9876/hooks/redeploy-webhook
```

### 自动部署

![AutoDeploy](http://7xlx3k.com1.z0.glb.clouddn.com/AutoDeploy.png-w)

大致的流程如上图。上图的DigitalOcean是一台VPS服务器，我用了很长时间了，速度和稳定性都不错。需要的同学使用这个链接购买，可以获得10美元的优惠：[https://www.digitalocean.com/?refcode=e131e2bba197](https://www.digitalocean.com/?refcode=e131e2bba197)

整个流程中，复杂度主要是在DigitalOcean的VPS上部署服务和脚本。

部署的脚本可以在我的GitHub上看到：[https://github.com/coderzh/coderzh-hugo-blog](https://github.com/coderzh/coderzh-hugo-blog) 

需要的同学可以参考下，代码见：[deploy.py](https://github.com/coderzh/coderzh-hugo-blog/blob/master/deploy.py)  

deploy.py放到你的主工程，也就是你写markdown的Repo下。比如：/var/coderzh-hugo-blog/下

adnanh-webhook的配置文件：

```
[
  {
    "id": "hugo-deploy",
    "execute-command": "/var/webhook/hugo-deploy.sh"
  }
]
```

hugo-deploy.sh里执行deploy.py：

```
#!/bin/bash
 
python /var/coderzh-hugo-blog/deploy.py --auto
```

剩下的是怎么在DigitalOcean的VPS上把这套东西部署起来。我使用nginx + supervisor搭建webhook的Server。

关于nginx和supervisor可以参考之前的一篇文章：[http://blog.coderzh.com/2014/05/19/digitalocean/](http://blog.coderzh.com/2014/05/19/digitalocean/)

supervisor的配置如下：

```
[program:webhook]
command=/root/gocode/bin/webhook -hooks /var/webhook/hooks.json -verbose -port=9876
user=root
directory=/var/webhook
autorestart=true
redirect_stderr=true
environment=HOME="/root",USER="root"
```

关于VPS上SSH Key的设置，见：[https://help.github.com/articles/generating-ssh-keys/](https://help.github.com/articles/generating-ssh-keys/)  为了自动部署方便，可以不设置密码。

当然，还有个大前提，在VPS上安装最新版本的golang。推荐使用gvm来安装。（记得安装1.5之前必须先把1.4先装上）

golang 安装：

```
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
source ~/.bashrc
gvm version
gvm install go1.4
gvm install go1.5.1
gvm use go1.5.1 --default
go version
```

最后，Good Luck！

