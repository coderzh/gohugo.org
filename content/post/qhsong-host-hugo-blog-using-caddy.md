---
categories:
- 搭建指南
keywords:
- LetsEncrypt
- Caddy
- Blog
tags:
- 入门
date: 2017-04-30T20:55:09+08:00 
author: qhsong
title: 使用Caddy和Hugo构建个人博客
thumbnail: http://7xlx3k.com1.z0.glb.clouddn.com/sqhme.png-wt
redirect: https://sqh.me/tech/host-hugo-blog-using-caddy/
weight: -100

---

## 介绍

[Hugo](http://gohugo.io/)是一个强大的静态博客生成器，由spf13使用Golang开发。在性能和生成文章的速度上都不错。目前的版本是0.20.6，依然没有迈出1.0的大版本号。一直在用Hugo进行博客的写作。（半年多没有更新博客了，逃

[Caddy](https://caddyserver.com/)同样也是一个由Golang开发的HTTP server，其最大的亮点就是内建了HTTPS和HTTP2的支持，同时自动能够向Let's Encrypt申请证书。同时，支持Middleware的功能，使用Golang撰写的Middleware能够支持各种各样的功能，同时任何人都可以撰写自己的Middleware，个性化这个HTTP Server。  <!--more-->


## 使用背景
我自己的博客项目是托管在Github上，每一次撰写完成后，需要Push到Github上，然后再到Host上Pull下来，然后本地Hugo自动更新。之前的证书呢是定期去Let's Encrypt申请，一到证书过期日子就非常的麻烦。
在学长的介绍下，认识了Caddy。Caddy有Hugo和Git插件，通过各种组合实现Hugo的自动更新和Github的webhook自动拉取，这样只要将博客文章push到了Github之后，剩下的事情就靠Caddy的中间件来完成，极大的自动化了这个过程。

## 过程

### 安装Caddy和Hugo
Caddy和Hugo的安装参见对应的安装文档，因为都是Golang编译出来的，所以只要把二进制Down下来就能够运行。安装Caddy的时候，一定要选择http.git和http.hugo这两个Middleware，否则之后的功能都不能实现。  
Hugo安装后，一定要放在`$PATH`中供Caddy调用。

### 设置Caddy配置文件
Caddy的配置文件比较简单，下面是我使用的一个配置文件
```
https://sqh.me {
    tls you@example.com
    gzip
    root /home/qhsong/sqh.me/public
    git github.com/qhsong/sqh.me {
        path /home/qhsong/sqh.me
        then hugo --destination=/home/qhsong/sqh.me/public
        hook /webhook GitHubSecretKey
        hook_type github
        clone_args --recursive
        pull_args --recurse-submodules
    }
    hugo
}
```
这个配置文件里面，tls设置了给Let's Encrypt证书申请的Email，root设置Hugo生成后的静态网站网址。
git 配置里面，填入了我的github地址。同时，指定了Clone的路径Path，使用then语句来指定克隆完成后做的事情。 
最后，在hook和hook_type中分别webhook的路径和type，这个一定要写，否则webhook不能正常调用。更详细的参数说明，详见[http.git](https://caddyserver.com/docs/http.git)。

### 设置Github的Webhook
Github的WebHook设置比较简单，设置好你的SecretKey，同时选择Content-Type为`application/json`，这样Caddy才能正常解析。

### 设置Systemd自动开机启动
Caddy同时提供了各种各样的启动脚本，见[caddy/dist/init](https://github.com/mholt/caddy/tree/master/dist/init)。选择一个适合你的系统的，然后复制到对应目录下，并进行修改就能使用。
考虑使用Supervisord，Systemd的配置有点繁琐。。

### Done
至此就可以完整的搞定这一切了


## 结束
Golang催生了很多有意思的东西，貌似Caddy已经商业化了。Golang社区也在蓬勃发展当中,还是有一些有意思的事情等待我们去发现的~ 虽然现在还在写C++。。。
