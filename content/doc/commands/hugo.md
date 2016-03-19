---
date: 2015-11-25T23:10:39+01:00
doc:
- commands
slug: hugo
title: hugo
url: /doc/commands/hugo
translated: true
---

## hugo

使用 hugo 构建你的网站

### 大纲


hugo 是主要的命令，用来构建你的网站。

hugo 是一个速度快同时扩展性很好的静态网站生成器，
由 spf13 和他的朋友们用 Go 语言实现。

完整的文档请见：http://gohugo.io/。

```
hugo
```

### 选项

```
  -b, --baseURL="": 主地址，比如： http://spf13.com/
  -D, --buildDrafts[=false]: 包含标记为 draft 的内容。
  -F, --buildFuture[=false]: 包含发布日期为未来的内容。
      --cacheDir="": 缓存的目录路径。默认为： $TMPDIR/hugo_cache/
      --canonifyURLs[=false]: 如果设置成 true，所有相对路径会和 baseURL 拼接到一起
      --config="": 配置文件 （默认是 config.yaml|json|toml）
  -d, --destination="": 目标文件写到哪里
      --disableRSS[=false]: 不要生成 RSS 文件
      --disableSitemap[=false]: 不要生成 Sitemap 文件
      --editor="": 如果提供了的话，会用指定的编辑器打开新的文章
      --ignoreCache[=false]: 忽略从缓存中读取，但是还是会写缓存
      --log[=false]: 开启日志
      --logFile="": 日志路径 （如果设置了，日志会自动打开）
      --noTimes[=false]: 不同步文件的修改时间
      --pluralizeListTitles[=true]: 列表的标题自动使用复数形式
      --preserveTaxonomyNames[=false]: 维持原样的 taxonomy 名称 ("Gérard Depardieu" vs "gerard-depardieu")
  -s, --source="": 读文件的路径
      --stepAnalysis[=false]: 显示程序不同步骤时的内存和耗时信息
  -t, --theme="": 主题名 （位于 in /doc/themes/THEMENAME/）
      --uglyURLs[=false]: 如果 true，使用 /filename.html 代替 /filename/
  -v, --verbose[=false]: 显示输出的详细信息
      --verboseLog[=false]: 详细信息写日志
  -w, --watch[=false]: 监视文件系统变化并重新生成
```

### 同时

* [hugo benchmark](/doc/commands/hugo_benchmark/)	 - hugo 的性能测试
* [hugo check](/doc/commands/hugo_check/)	 - 检查文章内容
* [hugo config](/doc/commands/hugo_config/)	 - 输出网站的配置信息
* [hugo convert](/doc/commands/hugo_convert/)	 - 转换文章内容到不同的格式
* [hugo gen](/doc/commands/hugo_gen/)	 - 一些有用的生成器集合
* [hugo import](/doc/commands/hugo_import/)	 - 从其他的网站导入到 hugo
* [hugo list](/doc/commands/hugo_list/)	 - 显示文章内容的列表
* [hugo new](/doc/commands/hugo_new/)	 - 生成新文章内容
* [hugo server](/doc/commands/hugo_server/)	 - 一个高性能的 Web 服务器
* [hugo undraft](/doc/commands/hugo_undraft/)	 - 把文章的 draft 状态由 'True' 变成 'False'
* [hugo version](/doc/commands/hugo_version/)	 - 输出 Hugo 的版本信息

