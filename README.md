# uBlacklist subscription compilation

> uBlacklist 订阅地址合集

[![Update uBlacklist subscription weekly](https://github.com/eallion/uBlacklist-subscription-compilation/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/eallion/uBlacklist-subscription-compilation/actions/workflows/go.yml)
[![GitHub issues](https://img.shields.io/github/issues/eallion/uBlacklist-subscription-compilation?logo=GitHub&color=4ec100&style=flat)](https://github.com/eallion/uBlacklist-subscription-compilation/issues/new/choose) ![Visitors](https://visitor-badge.laobi.icu/badge?page_id=eallion.uBlacklist-subscription-compilation)

> 订阅源地址：`https://git.io/ublacklist`  

这是一个 uBlacklist 插件的订阅地址合集，搜集了网上大部分的订阅地址合并成一个。  
通过 [Github Actions](https://github.com/eallion/uBlacklist-subscription-compilation/actions/workflows/go.yml) 每周自动更新一次。

### 功能

#### 🧱 **屏蔽中文内容农场！**

【什么是[内容农场](https://zh.wikipedia.org/wiki/%E5%85%A7%E5%AE%B9%E8%BE%B2%E5%A0%B4)？】

通过匹配域名的方式，屏蔽搜索引擎的搜索结果。  
主要是屏蔽中文 SEO 垃圾站，如机器翻译、AI 生成的内容……  

#### 🔍️ 支持的搜索引擎（由 [uBlacklist](https://github.com/iorate/ublacklist) 插件决定）

此扩展支持以下搜索引擎。

<!-- prettier-ignore-start -->

|              | 网页               | 图片               | 视频               | 新闻               |
| ------------ | ------------------ | ------------------ | ------------------ | ------------------ |
| Google       | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| Bing         | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| Brave        | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| DuckDuckGo   | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| Ecosia       | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| Kagi         | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| SearXNG      | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| Startpage    | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| Yahoo! JAPAN | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |                    |
| Yandex       | :heavy_check_mark: |                    | :heavy_check_mark: |                    |

<!-- prettier-ignore-end -->

没有计划提供额外的内置搜索引擎支持。您可以使用 [SERPINFO](https://ublacklist.github.io/docs/serpinfo) 自行添加对任何搜索引擎的支持。

### 用法

#### ⬇️ 下载 uBlacklist 浏览器插件

- Chrome: [Chrome Web Store](https://chrome.google.com/webstore/detail/ublacklist/pncfbmialoiaghdehhbnbhkkgmjanfhe)
- Firefox: [Firefox Add-ons](https://addons.mozilla.org/en-US/firefox/addon/ublacklist/)
- Safari (macOS / iOS): [App Store](https://apps.apple.com/us/app/ublacklist-for-safari/id1547912640)
- Edge*: [Chrome Web Store](https://chrome.google.com/webstore/detail/ublacklist/pncfbmialoiaghdehhbnbhkkgmjanfhe)

#### 🔧 设置插件

##### 1. 设置 - 订阅

> 订阅黑名单列表：

添加订阅：

- 订阅源地址：`https://git.io/ublacklist`  

<blockquote>
<details>
    <summary>【👉点击展示】订阅链接内容</summary>
<br />

`https://git.io/ublacklist` 的完整内容如下：

```bash
# curl -I https://git.io/ublacklist

# HTTP/1.1 301 Moved Permanently
# cache-control: public, max-age=31536000, immutable
# x-lru-cache: HIT
location: https://raw.githubusercontent.com/eallion/uBlacklist-subscription-compilation/main/uBlacklist.txt
# content-length: 0
# date: 
# x-github-backend: Kubernetes
# x-github-request-id: 
```

复制订阅长链接：

```bash
https://raw.githubusercontent.com/eallion/uBlacklist-subscription-compilation/main/uBlacklist.txt
```

</details>

</blockquote>

如下图所示：（先点`添加订阅`）

![](tools/x2yWi62OWl.png)

##### 2. 设置 - 常规：（选填）

> 在谷歌的搜索结果中将不会显示以下网站：

```plain
*://*/so.php
*://*/so.php?s=*
*://*/cha.php?s=*
*://*/list.php?s=*
*://*/?s=*
*://*/so/*
```

屏蔽 AI 内容农场：

> 来自：https://github.com/laylavish/uBlockOrigin-HUGE-AI-Blocklist?tab=readme-ov-file#ublacklist

### 添加/删除 域名

##### 1. 插件设置

少量域名，请在自己浏览器插件的 `常规` 设置中添加即可，在行首添加 `@` 符号可取消屏蔽：

```
# 加入屏蔽
*://*.baidu.com/*

# 取消屏蔽
@*://*.baidu.com/*
```

##### 2. 提交 URL

如有需求添加、删除域名，请至 [eallion/uBlacklist-Subscription](https://github.com/eallion/uBlacklist-Subscription) 提交 PR。  

如果有什么了不得的特殊域名，可以加入到 [reverse_url.txt](https://github.com/eallion/uBlacklist-subscription-compilation/blob/main/reverse_url.txt) 中，相当于白名单。

### 感谢名单

##### uBlacklist

- [https://github.com/iorate/uBlacklist](https://github.com/iorate/uBlacklist)

##### Contribution

- https://github.com/eallion/uBlacklist-Subscription.git

##### Subscription

> 添加订阅源的时候，各订阅源均采用 Mit License 或无 License。如果有侵权行为，我会第一时间删除。

- https://github.com/arosh/ublacklist-stackoverflow-translation.git
- https://github.com/cobaltdisco/Google-Chinese-Results-Blocklist.git
- https://github.com/dallaslu/penzai-list.git
- https://github.com/gyli/Blocklist.git
- https://github.com/h-matsuo/uBlacklist-subscription-for-developer.git
- https://github.com/littleserendipity/uBlacklist-Subscription.git
- https://github.com/liubiantao/uBlacklist-Websites.git
- https://github.com/nonPointer/uBlacklist-Subscription.git
- https://github.com/Paxxs/Google-Blocklist
- https://github.com/xNathan/uBlacklist-subscription.git
- https://github.com/YeSilin/uBlacklist.git
- https://github.com/youzeliang/uBlacklist-Subscription
- https://github.com/zweie/some-rules-for-ublacklist
- https://github.com/laylavish/uBlockOrigin-HUGE-AI-Blocklist

### 推荐

推荐另一个插件：<https://github.com/danny0838/content-farm-terminator>

uBlacklist 插件对百度无效。  
如果需要百度屏蔽插件请用：<https://github.com/zhangolve/search-engine-filter>  

### [LICENSE](https://github.com/me-shaon/GLWTPL)

```txt
GLWT（Good Luck With That，祝你好运）公共许可证
            版权所有© 除作者外的所有人

任何人都被允许复制、分发、修改、合并、销售、出版、再授权
或任何其它行为，但风险自负。

作者对这个项目中的代码的行为一无所知。
代码处于可用或不可用状态，没有第三种可能


                祝你好运公共许可证
            复制、分发和修改的条款和条件

  0. 只要你永远不要留下任何可以追踪到原作者的线索，
你就可以随心所欲地做任何事，因此，不能因此责怪或追究
原作者的责任。

在任何情况下，作者均不对因使用或与本软件有关的合同诉讼、
侵权或其他方式产生的任何索赔、损害或其他责任负责。

自求多福吧。
```
