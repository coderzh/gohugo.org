---
categories:
- 功能改进
date: 2015-11-21T17:48:04+08:00
description: "hugo 自动化部署脚本 deploy.py"
keywords:
- Hugo
- 自动化
- 部署
title: "Hugo 自动化部署脚本 deploy.py"
url: "/2015/11/21/hugo-deploy-script/"
author: CoderZh

---

之前我写了一个脚本用来自动部署我的 Hugo 博客，今天闲来无事完善了一下这个脚本，使这个脚本更加通用一些。

脚本路径：[https://github.com/coderzh/coderzh-hugo-blog/blob/master/deploy.py](https://github.com/coderzh/coderzh-hugo-blog/blob/master/deploy.py)

<!--more-->

### 原理

`deploy.py` 会自动执行 hugo 命令生成静态站点，然后将生成的文件拷贝到上层的一个目录里，然后，在那个目录里将文件 push 到你指定的 Git Repository 里。

### 使用方法

1. 将 `deploy.py` 放到你的 Hugo 站点目录。（和 config.yaml 等文件放一起）
1. 编辑 `deploy.py` 文件，修改你要部署到的 Git Repository：

    ```python
    GIT_REPO = [
        # [别名,   分支名,     Git Repo 路径]
        ['origin',  'gh-pages', 'git@github.com:coderzh/hugo-blog-deployed.git'],
        ['gitcafe', 'gh-pages', 'git@gitcafe.com:coderzh/coderzh-hugo-blog.git'],
    ]
    # 部署到哪里，相对上一级目录。比如下面的配置，会部署到 ../gh-pages 目录里
    DEPLOY_DIR = 'gh-pages'
    ```

1. 如果你的网站需要指定皮肤，需要在 config 文件中指定 `theme` 。因为我的脚本在生成静态文件时并不会指定皮肤。

    ```yaml
    theme: "rapid"
    ```

1. 第一次执行，使用 `first` 参数，它会做一些初始化的操作。并使用 `-t` 表示只是测试一下，并不会真的 push 。

    ```
    python deploy.py first -t
    ```

1. 中间可能需要输入密码，如果是自动化部署，可在 Git Repo 里添加一个没有密码的 SSH Key 。
1. 如果一切正常，切换到 `DEPLOY_DIR` 目录，`git log` 看看 commit 记录是否正常。如果一切也如你所愿。则可以把 `-t` 参数去掉重新执行一遍，执行真的 push 操作：

    ```
    python deploy.py first
    ```

1. 执行完成后，应该已经将生成的静态页面自动 push 到了你指定的 `GIT_REPO` 里。
1. 之后如需再次手工部署，只需要使用 `manual` 参数，速度会快很多：

    ```
    python deploy.py manual
    ```

1. 如果你想通过 `webhook` 来自动部署，使用 `auto` 参数，这样在执行 deploy.py 时，会使用 Git 自动更新你当前的 Hugo 站点目录 ，然后部署：

    ```
    python deploy.py auto
    ```

That's all, 祝你好运！
    
