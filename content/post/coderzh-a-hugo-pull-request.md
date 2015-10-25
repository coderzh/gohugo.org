---
categories:
- 功能改进
tags:
- PullRequest
keywords:
- Hugo
date: 2015-09-03T08:55:31+08:00
title: "给Hugo提交了一个PR"
author: CoderZh
thumbnail: http://7xlx3k.com1.z0.glb.clouddn.com/pullrequest.jpg
weight: -1

---

给Hugo提交了一个Pull Request，并且通过了，记录一下。PR的内容是在config里增加了disablePathToLower选项，可以配置是否需要阻止Hugo自动将路径转换为小写。在站点从旧的系统里迁移到Hugo中时会非常有用。

<!--more-->

在上一篇博客[使用hugo搭建个人博客站点](http://blog.coderzh.com/2015/08/29/hugo/)  里提到，Hugo强制将所有路径都变成了小写。小写路径虽好，但是如果之前你的网站路径不是全小写的，迁移过来路径发生变化，将会导致路径失效，是不可接受的。

上文提到，我硬改了Hugo代码，让它不要强制转换小写路径。这样的做法不够优雅，假如别人也和我有一样的需求怎么办？我们还是希望Hugo本身就能够提供这样的功能。于是，我决定把代码改的优雅一些，给Hugo作者提交一个Pull Request，将这个功能合并到Hugo代码里。

于是，我在config文件里增加了一个选项：

```python
# Do not make the url/path to lowercase
disablePathToLower: true 
```

 增加一个开关其实很简单，只需要在command/hugo.go里的LoadDefaultSettings函数里增加一行：

```go
viper.SetDefault("DisablePathToLower", false)
```

开关生效的实现也很简单，找到最终转换为小写的函数，加入disablePathToLower开关的判断：

```go
func MakePathSanitized(s string) string {
    if viper.GetBool("DisablePathToLower") {
        return MakePath(s)
    } else {
        return strings.ToLower(MakePath(s))
    }
}
```

之前的函数名叫MakePathToLower，由于我在内部加了个开关，于是名字不适用了，于是修改为了：MakePathSanitized。

修改完成之后，需要把测试案例跑一遍，保证你的修改不会破坏原来的逻辑。

```go
go test ./...
```

当然，也要保证编译是成功的：

```go
go build
```

在我修改的第一个版本时，我犯了一个错误，导致测试案例没有通过。原因是第一个版本我增加的配置选项是PathToLower，默认值是true。而有些测试案例是没有去设置PathToLower的，默认读到的会是false，从而导致执行失败的执行结果。所以，config里增加的选项，默认值尽量是false吧，也会好理解一些。于是就有了：disablePathToLower

提交Pull Request时，有几点需要注意的：

 1. 将自己的修改合并成一个commit再提交Pull Request。（这样在主Resp里更加规范和整洁）
 1. commit message里的描述尽量简洁清晰，如果有对应的issue ID，最好加上：See #1234 or Fixes #1234 之类的。这样可以自动关联起来。
 1. 不要在自己的master分支上修改提交PR，而是应该自己开一个单独分支，由该分支提交Pull Request到原Repo的master。

最后，我提交了这么一个Pull Request：[add configuration variable: "disablePathToLower" See #557](https://github.com/spf13/hugo/pull/1392)

不久后，Hugo的主要维护者之一通过了我的PR，很友善的将我蹩脚的中国式英语的commit message调整了一下合并到了Hugo的master：

![HugoPR](http://7xlx3k.com1.z0.glb.clouddn.com/HugoPR.png-w)

commit: [https://github.com/spf13/hugo/commit/52d94fa67578f6b63035e73b236ca8abd40d0006](https://github.com/spf13/hugo/commit/52d94fa67578f6b63035e73b236ca8abd40d0006)

这就是开源项目的好处，开源社区的魅力所在：**你可以使用它，修改它，贡献自己的代码，参与其中，让它变好，让所有人受益**。

