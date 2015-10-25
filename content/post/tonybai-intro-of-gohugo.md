---
author: "tonybai"
categories:
- 搭建指南
date: 2015-09-23
redirect: "http://tonybai.com/2015/09/23/intro-of-gohugo/"
tags:
- 入门
thumbnail: "/post-thumbnail/tonybaicom-hyde-theme-2.png"
title: "使用Hugo搭建静态站点"

---

虽然前一篇Blog宣称自己要用Markdown开始写Post，但实际操作起来还是发现了诸多不兼容问题(插件与主题间、插件与插件间的)，让编写和修改文章变得十分繁琐，于是我研究了一下静态Web站点生成工具Hugo。Hugo是由前Docker的重量级员工(2015年8月末从Docker离职)：Steve Francia实现的一个开源静态站点生成工具框架，类似于Jekyll、Octopress或Hexo，都是将特定格式(最常见的是Markdown格式)的文本文件转换为静态html文件而生成一个静态站点，多用于个人Blog站点、项目文档(Docker的官方manual Site就是用Hugo生成的)、初创公司站点等。这类工具越来越多的受到程序员等颇具“极客”精神的群体的欢迎，结合github.com等版本控制服务，采用具有简单语法格式但强大表达力的Markdown标记语言，人们可以在很短时间内就构建出一个符合自己需求的静态Web站点。在这些工具中，Hugo算是后起之秀了，它最大的优点就是Fast! 一个中等规模的站点在几分之一秒内就可以生成出来。其次是良好的跨平台特性、配置简单、使用方便等。这一切均源于其良好的基因：采用Go语言实现。Steve Francia除了Hugo平台自身外，还维护了一个Hugo Theme的Repo，这个Hugo主题库可以帮助Hugo使用者快速找到自己心仪的主题并快速搭建起静态站点。目前国内使用Hugo的人还不多，但感觉其趋势是在逐渐增多。这里写下这篇Post，也算是为大家入个门，引个路吧。

...
