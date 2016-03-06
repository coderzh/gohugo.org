---
author: Spencer Lyon
date: 2014-03-20
lastmod: 2015-05-22
menu:
  main:
    parent: tutorials
next: /doc/tutorials/migrate-from-jekyll
prev: /doc/tutorials/installing-on-windows
title: MathJax Support
toc: true
weight: 10
translated: true
---

## 什么是 MathJax?

[MathJax](http://www.mathjax.org/) 是一个 JavaScript 库，可以让你使用 LaTeX 风格的语法，在网页的 HTML （或者 Markdown）的源码中显示数学表达式。因为它是一个纯 Javascript 的库，在 Hugo 中是可以直接使用的，但也会有一些奇怪的问题，下面会详细讨论。

本篇教程并非介绍在网站上实际使用 MathJax 来渲染输入的数学表达式。相反，本篇教程收集了让 MathJax 在 Hugo 构建的网站上工作的一些的提示和技巧。

## 启用 MathJax

第一步是在你想显示数学表达式排版的页面时启用 MathJax。这么做的方法有很多种（喜欢折腾的读者可以参考 MathJax 文档中的 [载入和配置](http://docs.mathjax.org/en/latest/configuration.html) 部分，来获得更多包含 MathJax 的方法），但最简单的方法是使用安全的 MathJax CDN，可以在页面的源码中插入下面的 HTML 代码：

    <script type="text/javascript"
      src="https://cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML">
    </script>

保证所有页面都包含这段代码的方法是把它放在 `layouts/partials/` 文件夹中的一个模板中。比如，我把这段代码放在了我的模板 `footer.html` 的底部，因为我知道 footer 将会包含在我网站中的每个页面中。

### 选项和特性

MathJax 是一个拥有许多特性的稳定开源库。我鼓励有兴趣的读者看一下 [MathJax 文档](http://docs.mathjax.org/en/latest/index.html)，尤其是 [基本用法](http://docs.mathjax.org/en/latest/index.html#basic-usage) 和 [MathJax 配置选项](http://docs.mathjax.org/en/latest/index.html#mathjax-configuration-options)的部分。

## 与 Markdown 的问题

在启用 MathJax 后，任何在特定标记（见文档）中输入的数学表达式都将被处理，并排版显示在网页上。然而，和 Markdown 一起使用就会带来一个问题， 在 Markdown 中被下划线（`_`）包围的文字表示 `强调` 的意思，而 LaTeX (MathJax) 使用下划线来创建一个下标。下划线的这种 “双重标准” 会导致一些意料外的情况发生。

### 解决方法

解决这个问题有几种方法。其中之一是通过输入 `\_` 而不是 `_`，在数学表达式的代码的每个下划线都添加转义符号。如果你输入的表达式全是下标，这么做会变得非常冗余乏味。

另一种方法是把 MathJax 代码认为是逐字的代码，并不处理它。有一种这么做的方法是把数学表达式放到一个 `<div>` `</div>` 块中。Markdown 将会忽略这些部分，而且它们会被直接传入 MathJax 中并被正确处理。这用于显示有样式的数学公式是最好的，但对于内嵌的数学表达式来说，嵌入在 `<div>` 中是不可取的方法。Markdown 处理内嵌文字的语法是把它们包含在倒引号（`` ` ``）中。然而，你可能已经注意到了，在倒引号中包含的文字和标准的文字的渲染方式是不同的（在本网站上这些元素以红色高亮）。为了解决这个问题，我们可以创建一条新的 CSS，在所有的内嵌逐字文字包括 MathJax 代码上应用标准样式。下面我将展示完成这个目的的 HTML 和 CSS 代码（注意这个方法修改自[这篇博客文章](http://doswa.com/2011/07/20/mathjax-in-markdown.html)---所有荣誉归原作者所有）。

    <script type="text/x-mathjax-config">
    MathJax.Hub.Config({
      tex2jax: {
        inlineMath: [['$','$'], ['\\(','\\)']],
        displayMath: [['$$','$$'], ['\[','\]']],
        processEscapes: true,
        processEnvironments: true,
        skipTags: ['script', 'noscript', 'style', 'textarea', 'pre'],
        TeX: { equationNumbers: { autoNumber: "AMS" },
             extensions: ["AMSmath.js", "AMSsymbols.js"] }
      }
    });
    </script>

    <script type="text/x-mathjax-config">
      MathJax.Hub.Queue(function() {
        // Fix <code> tags after MathJax finishes running. This is a
        // hack to overcome a shortcoming of Markdown. Discussion at
        // https://github.com/mojombo/jekyll/issues/199
        var all = MathJax.Hub.getAllJax(), i;
        for(i = 0; i < all.length; i += 1) {
            all[i].SourceElement().parentNode.className += ' has-jax';
        }
    });
    </script>

和前面一样，这段内容应该包含在每个使用 MathJax 的页面的 HTML 源码中。下面的代码段包含了用于让逐字 MathJax 块使用与页面主体一样的字体样式来渲染的 CSS 样式。


    code.has-jax {font: inherit;
                  font-size: 100%;
                  background: inherit;
                  border: inherit;
                  color: #515151;}

在 CSS 代码中，注意这一行 `color: #515151;`。`#515151` 在我的 CSS 中是分配给 `body` 类的 `color` 元素的值。为了让公式与网页的主体适配，这个值应该和主体的颜色值相同。

### 使用

有了这样的配置，在 Hugo 生成的页面上自然使用 MathJax 就没什么问题了。为了包含内嵌数学表达式，仅需要把 LaTeX 代码放入 `` `$ TeX Code $` `` 或 `` `\( TeX Code \)` `` 中。为了包含有显示样式的数学表达式，仅需要把 LaTeX 代码放入 `<div>$$TeX Code$$</div>` 中。所有的数学表达式都将会在 Hugo 生成的网页中用合适的样式展示出来！
