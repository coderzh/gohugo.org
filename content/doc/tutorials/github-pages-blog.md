---
aliases:
- /tutorials/github_pages_blog/
author: Spencer Lyon
date: 2014-03-21
lastmod: 2016-01-14
linktitle: Hosting on GitHub
menu:
  main:
    parent: tutorials
next: /doc/tutorials/installing-on-mac
prev: /doc/tutorials/creating-a-new-theme
title: Hosting on GitHub Pages
toc: true
weight: 10
translated: true
---

> *This tutorial was contributed by [Spencer Lyon](http://sglyon.com/).*

## 简介

很多 Hugo 用户表示希望看到关于如何把 Hugo 生成的博客托管在 GitHub Pages 上的教程。这就是本篇教程的目的。我们仅仅需要读者已经正确安装了 Hugo，并且能适应 git 和 GitHub。

在本篇教程中，我将向你讲解我在创建 [http://spencerlyon2.github.io/hugo_gh_blog](http://spencerlyon2.github.io/hugo_gh_blog) 这一示例博客过程中的主要步骤。这个博客的源代码位于 [GitHub](https://github.com/spencerlyon2/hugo_gh_blog)。希望读者朋友可以下载这个示例仓库代码并跟着我一起操作。

### 找个地方放你的文件

因为我们的目标是使用 GitHub Pages 服务来托管网站，那么我们自然就需要在一个 GitHub 仓库中托管页面的内容。因此，第一步要么在 GitHub 上创建一个新的仓库，要么在网站内容所在的仓库中新建一个文件夹。因此，我创建了仓库 [spencerlyon2/hugo_gh_blog](https://github.com/spencerlyon2/hugo_gh_blog)。

## 创建博客

### 编写 `config.yaml` 文件

创建一个新的 Hugo 站点的第一步是[编写配置文件](/doc/overview/configuration/)。这个配置文件很重要的原因至少有以下两个：（1）这是全站设置的地方，以及（2）这个配置文件在一定程度上规定了 Hugo 生成网站的方式。比如我创建了一个 `config.yaml`，里面有如下内容：

    ---
    contentdir: "content"
    layoutdir: "layouts"
    publishdir: "public"
    indexes:
      category: "categories"
    baseurl: "http://spencerlyon2.github.io/hugo_gh_blog"
    title: "Hugo Blog Template for GitHub Pages"
    canonifyurls: true
    ...

> **警告：** 在本篇教程完成后，Hugo 的 `canonifyurls: true` 先前默认设置已经改为 `false` 了。
> 如果你在本教程中使用的是 Spencer 的 https://github.com/spencerlyon2/hugo_gh_blog, 
> **请确保在你的 `config.yaml` 中手动添加 `canonifyurls: true`**
> 否则你*将会*遇到诸如 CSS 无法载入等问题。

> 更多信息详见[“额外部分：URLs 页面”中的“标准化: 警告”](/doc/extras/urls/)。

### 定义网站结构

Hugo 假设你以有意义的方式来组织网站的内容，并使用相同的结构来渲染网站。注意我们在配置文件中有这么一行 `contentdir: "content"`。这表明网站的所有实际内容都应该放在一个叫做 `content` 的文件夹中。Hugo 把 `content` 文件夹中所有文件夹都当做子集。比如在我们的例子中，仅需要一个子集：放置我们博客文章的地方。那么我们创建两个新的文件夹：

```
▾ <root>/
    ▾ content/
        ▾ posts/
```

### 创建 HTML 模板

下一步是要定义这个新网站的界面和外观。因为 Hugo 将使用用户（你）编写的 HTML 模板来生成网站，因此这一步是非常主观的。我将仅仅展示一个可用于生成博客的主题。我决定以一个叫做 [Lanyou](http://lanyon.getpoole.com/) 的 Jekyll 主题为基础来构建这个示例项目。Lanyou 主题由纯 CSS 构成，在本示例仓库的 `/static/css` 文件夹中有一个轻微修改过的版本。如果你正在跟着本教程学习，那么你需要从示例仓库中抓取 `static` 文件夹，并把它与你刚创建的 `content` 文件夹放在一起。

因为完整组成一个完全的网站需要太多的文件，我无法在这里详细地讲解每个文件。然而，我将会在把这些所说的都完成后，展示出来文件夹的目录结构：

```
▾ <root>/
    ▾ content/
        ▾ posts/
            <blog posts>.md
    ▾ static/
        ▾ css/
            lanyon.css
            poole.css
    ▾ layouts/
        ▾ partials/
            <templates to be used in other files>.html
        ▾ posts/
            li.html
            single.html
            summary.html
        ▾ indexes/
            category.html
            indexes.html
            posts.html
        index.html
    README.md
```

示例仓库中的每个文件都有良好的注释，用于从整体上描述文件的作用，同时也解释了文件中所有的主要部件。如果你对使用/不使用 Hugo 进行网站开发不熟悉，我建议你在这些文件中搜寻一番，找到 Hugo 模板工作的原理以及把网站维系在一起的方式。

### 添加一些内容

创建博客的最后一步是添加一些实际的博文。为此，简单的为每篇新博文创建一个 Markdown 文件（加上后缀 `.md`）。在每篇博文的顶部，你应该添加一个元数据部分，来告诉 Hugo 有关博文的一些信息（见 [docs](/doc/content/front-matter/))。比如，看下示例仓库中的 `/content/posts/newest.md` 文件顶部的元数据部分：

    ---
    title: "Just another sample post"
    date: "2014-03-29"
    description: "This should be a more useful description"
    categories:
        - "hugo"
        - "fun"
        - "test"
    ---

这部分的关键在于强制的 `title` 和 `date` 信息，也包括可选项 `description` 和 `categories` 。这些元素中的每一个都在 `/layouts` 文件夹中所有模板上使用，也会把网站上其他页面的博文信息提供给 Hugo。

## 配置 `git` 工作流 

一旦网站建立起来并可以工作，我们需要把它推到 GitHub 代码仓库上的正确分支，这样网站就可以被 GitHub Pages 服务托管。有很多方法可以达到这个目的。这里我会向你展示我当前用于管理托管在 GitHub Pages 服务上的网站的工作流。

GitHub Pages 将会把任何代码仓库的网站托管起来，只要它有一个叫做 `gh-pages` 的分支，而且在这个分支的根部有一个有效的 `index.html`。典型的工作流大概是将网站的内容放在仓库的 `master` 分支，而把生成的网站放在 `gh-pages` 分支。这可以很好的分割输入和输出，但做起来却有些枯燥。作为变通方案，我们将使用 `git subtree` 命令集来设置 `public` 文件夹（或者你在 `config.yaml` 中设置的任意 `publishdir`）为 `gh-pages` 分支的根的镜像。这将可以使我们可以在 `master` 分支做所有的事，运行Hugo让网站输出到 `public` 文件夹，然后直接把这个文件夹直接推到 GitHub Pages 服务的正确位置来托管网站。

为了让它正确的工作起来，我们将在命令行中执行一系列的命令。我将把它们一起放在这里，这样便于复制和粘贴，并使用注释来解释每行的作用。注意这些命令需要在 `<root>` 文件夹中运行（也是 Hugo 项目中的 `content` 和 `layout` 文件夹所在的地方）。同时注意你需要修改那些有 GitHub 地址的示例仓库的命令，这样它们会指向你的仓库。

    # Create a new orphand branch (no commit history) named gh-pages
    git checkout --orphan gh-pages

    # Unstage all files
    git rm --cached $(git ls-files)

    # Grab one file from the master branch so we can make a commit
    git checkout master README.md

    # Add and commit that file
    git add .
    git commit -m "INIT: initial commit on gh-pages branch"

    # Push to remote gh-pages branch
    git push origin gh-pages

    # Return to master branch
    git checkout master

    # Remove the public folder to make room for the gh-pages subtree
    rm -rf public

    # Add the gh-pages branch of the repository. It will look like a folder named public
    git subtree add --prefix=public git@github.com:spencerlyon2/hugo_gh_blog.git gh-pages --squash

    # Pull down the file we just committed. This helps avoid merge conflicts
    git subtree pull --prefix=public git@github.com:spencerlyon2/hugo_gh_blog.git gh-pages

    # Run hugo. Generated site will be placed in public directory (or omit -t ThemeName if you're not using a theme)
    hugo -t ThemeName


    # Add everything
    git add -A

    # Commit and push to master
    git commit -m "Updating site" && git push origin master

    # Push the public subtree to the gh-pages branch
    git subtree push --prefix=public git@github.com:spencerlyon2/hugo_gh_blog.git gh-pages

在执行完这些命令后，等待 GitHub 服务器的更新，我们刚才创建的网站 [http://spencerlyon2.github.io/hugo_gh_blog](http://spencerlyon2.github.io/hugo_gh_blog) 就可以访问了。

### `deploy.sh`

现在，当你要在博客中添加新的博文时，你将需要依照下列步骤执行：

* 在 `content/posts` 文件夹中创建新的 Markdown 博文
* 通过以服务器模式 `hugo server --watch` 运行 Hugo 来预览你的博文
* 不以服务器模式运行 Hugo，这样生成的 url 对于网站而言是正确的
* 在 `master` 分支添加并注释新的博文
* 推送到 `master` 分支
* 推送共有的 subtree 到远程 `gh-pages` 分支

在前面列表中的前两步是你在编写时预览内容的一个简单方便的方法。这是一种动态的，非常合理的方法。然而，所有剩下的步骤在你每次向网站添加新内容时都是一样的。为了让这个重复的过程更加简单，我利用了 [Chimer Arts & Maker Space](https://github.com/chimera/chimeraarts.org) 网站的代码仓库中的一个脚本，在 [Hugo Showcase](/showcase) 中高亮显示。这段脚本在 `deploy.sh` 文件中，内容如下：

**注意：**

第一个命令 `hugo` 假定你以所有的默认设置来运行。

为了使用某个主题，确保使用 `-t ThemeName` 来指定它（或者在配置文件中包含主题）。

    hugo -t ThemeName

为了构建所有的草稿博文 *（如果你仅有草稿，将不会生成网站）*

    hugo --buildDrafts

**Deploy.sh:**

    #!/bin/bash
    
    echo -e "\033[0;32mDeploying updates to GitHub...\033[0m"
    
    # Build the project.
    hugo
    
    # Add changes to git.
    git add -A
    
    # Commit changes.
    msg="rebuilding site `date`"
    if [ $# -eq 1 ]
      then msg="$1"
    fi
    git commit -m "$msg"
    
    # Push source and build repos.
    git push origin master
    git subtree push --prefix=public git@github.com:spencerlyon2/hugo_gh_blog.git gh-pages

现在我可以把工作流列表中的最后四项改为一个命令 `bash deploy.sh`。这个脚本可以接收一个可选参数，作为 git 提交改动的信息。如果你想包含一个自定义的提交信息，在使用 bash 运行脚本时，把它放在引号内并放在脚本的后面：`bash deploy.sh "<my commit msg>"`。如果你选择不指定的话，脚本会使用当前时间生成一个提交信息。

## 配置 `git` 工作流的另一种方法
上面的方法使用 `git subtree` 来部署 Git 的 `gh-pages` 分支。这个方法非常有效，但有一个缺点：它需要将生成的内容提交到源分支。

另外还有一种方法：

1. 在 `master` 分支设置你的 Hugo 站点
2. 创建独立的 `gh-pages` 分支。（详见[这里](https://help.github.com/articles/creating-project-pages-manually/)。)
3. 按照下面的指导操作。

这样，假设你已经设置好了 `gh-pages`，在 `master` 中也有一些你希望发表的提交的内容：

```
# 在代码树的根目录获取部署脚本，并让它可执行。
wget https://github.com/X1011/git-directory-deploy/raw/master/deploy.sh && chmod +x deploy.sh

# 为了让它构建在 "dist" 之外的文件夹中，可以参见 deploy.sh 中的选项。
# 把网站构建在 /dist 文件夹中。
hugo -d dist

# 运行上面安装的 deploy.sh 脚本。
./deploy.sh
```

这个脚本会使用上一次的提交信息作为 `gh-pages` 分支的提交信息。

关于部署脚本的更多信息，可以看下这个 [README](https//github.com/X1011/git-directory-deploy)。

## 托管个人/组织的页面

正如 [GitHub 这篇文章中](https://help.github.com/articles/user-organization-and-project-pages/)说提到的，除了项目页面，你也可以托管用户/组织的页面。主要的不同之处在于：

> - 你必须使用 `username.github.io` 这样的命名形式。
> - `master` 分支的内容将会用于构建和发布到你的 GitHub Pages 站点。

这种情况下就简单多了：我们将创建两个不同的仓库，一个用于托管 Hugo 的内容，其中有一个 git 子模块的 `public` 文件夹。

一步一步来：

1. 在 GitHub 上创建 `<your-project>-hugo` 代码仓库（它将用于托管 Hugo 的内容）
2. 在 GitHub 上创建 `<username>.github.io` 代码仓库（它将用于托管 `public` 文件夹：静态的网站）
3. `git clone <<your-project>-hugo-url> && cd <your-project>-hugo`
4. 让你的网站在本地生效（`hugo server --watch -t <yourtheme>`）
5. 当你对结果满意时，<kbd>Ctrl</kbd>+<kbd>C</kbd> (关闭服务)并 `rm -rf public`（不用着急，使用 `hugo -t <yourtheme>` 总会重新生成这个文件夹的）
6. `git submodule add git@github.com:<username>/<username>.github.io.git public`
7. 基本上完成：添加 `deploy.sh` 脚本来帮你（让它可执行：`chmod +x deploy.sh`）:

```
#!/bin/bash

echo -e "\033[0;32mDeploying updates to GitHub...\033[0m"

# Build the project.
hugo # if using a theme, replace by `hugo -t <yourtheme>`

# Go To Public folder
cd public
# Add changes to git.
git add -A

# Commit changes.
msg="rebuilding site `date`"
if [ $# -eq 1 ]
  then msg="$1"
fi
git commit -m "$msg"

# Push source and build repos.
git push origin master

# Come Back
cd ..
```
8. `./deploy.sh "Your optional commit message" `把变动推送到 `<username>.github.io`（小心，你也许想要提交变化到 `<you-project>-hugo` 仓库）。

就是这样！你的个人页面运行在 [http://username.github.io/](http://username.github.io/)（在10分钟延迟之后）。

## 结论

希望这篇教程可以帮助你开启建立并运行网站的步伐！如果你有任何问题，欢迎通过[论坛](/doc/community/mailing-list/)与社区联系。
