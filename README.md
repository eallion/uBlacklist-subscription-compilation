# uBlacklist subscription compilation
> uBlacklist 订阅地址合集

![Update uBlacklist subcription weekly](https://github.com/eallion/uBlacklist-subscription-compilation/workflows/Update%20uBlacklist%20subcription%20weekly/badge.svg?branch=master)

这是一个 uBlacklist 订阅地址合集，搜集了网上大部分的订阅地址合并成一个。通过 Github Actions 每周自动更新一次。

# 功能

通过匹配域名的方式，屏蔽搜索引擎的搜索结果。主要是屏蔽中文SEO垃圾站。  

支持的网站：（由 uBlacklist 插件决定）
- [Google](https://www.google.com/ncr?gws_rd=ssl)
- [DuckDuckGo](https://duckduckgo.com/)
- [Startpage](https://www.startpage.com/)

# 用法：

### 一、下载 uBlacklist 浏览器插件：
- Chrome： [https://chrome.google.com/webstore/detail/ublacklist/pncfbmialoiaghdehhbnbhkkgmjanfhe](https://chrome.google.com/webstore/detail/ublacklist/pncfbmialoiaghdehhbnbhkkgmjanfhe)
- Firefox： [https://addons.mozilla.org/en-US/firefox/addon/ublacklist](https://addons.mozilla.org/en-US/firefox/addon/ublacklist/)

### 二、设置插件

#### 1、设置 - 一般：

>以下网站将不会显示在 Google 的搜索结果中：

```
*://*/so.php
*://*/so.php?s=*
*://*/cha.php?s=*
*://*/list.php?s=*
*://*/?s=*
*://*/so/*
```
如图：

![](https://cdn.jsdelivr.net/gh/eallion/uBlacklist-subscription-compilation@master/tools/PKO0JoJBpU.png)


#### 2、设置 - 订阅：

> 订阅黑名单列表:

添加订阅：  
将 [该列表地址](https://raw.githubusercontent.com/eallion/uBlacklist-subscription-compilation/master/uBlacklist.txt) 添加到 Subscription 订阅中。

`https://raw.githubusercontent.com/eallion/uBlacklist-subscription-compilation/master/uBlacklist.txt`

![](https://cdn.jsdelivr.net/gh/eallion/uBlacklist-subscription-compilation@master/tools/xN5B7whxYr.png)

# 感谢名单：

uBlacklist: 
- [https://github.com/iorate/uBlacklist](https://github.com/iorate/uBlacklist)

Subscription:

- [https://github.com/iorate/ublacklist-example-subscription](https://github.com/iorate/ublacklist-example-subscription)

- [https://github.com/YeSilin/uBlacklist](https://github.com/YeSilin/uBlacklist)
- [https://github.com/cobaltdisco/Google-Chinese-Results-Blocklist](https://github.com/cobaltdisco/Google-Chinese-Results-Blocklist)
- [https://github.com/gyli/Blocklist](https://github.com/gyli/Blocklist)
- [https://github.com/h-matsuo/uBlacklist-subscription-for-developer](https://github.com/h-matsuo/uBlacklist-subscription-for-developer)
- [https://github.com/littleserendipity/uBlacklist-Subscription](https://github.com/littleserendipity/uBlacklist-Subscription)
- [https://github.com/liubiantao/uBlacklist-Websites](https://github.com/liubiantao/uBlacklist-Websites)
- [https://github.com/nonPointer/uBlacklist-Subscription](https://github.com/nonPointer/uBlacklist-Subscription)
- [https://github.com/scyrte/uBlacklist-Subscription](https://github.com/scyrte/uBlacklist-Subscription)
- [https://github.com/xNathan/uBlacklist-subscription](https://github.com/xNathan/uBlacklist-subscription)
- [https://github.com/yangjingchuang/ublacklist-subscription](https://github.com/yangjingchuang/ublacklist-subscription)


# LICENSE
```
GLWT（祝你好运）公共许可证
版权所有（C）每个人，除了作者

任何人都被允许复制、分发、修改、合并、销售、出版、再授权或
任何其它操作，但风险自负。

作者对这个项目中的代码一无所知。
代码处于可用或不可用状态，没有第三种情况。


                祝你好运公共许可证
            复制、分发和修改的条款和条件

0 ：在不导致作者被指责或承担责任的情况下，你可以做任何你想
要做的事情。

无论是在合同行为、侵权行为或其它因使用本软件产生的情形，作
者不对任何索赔、损害承担责任。

祝你好运及一帆风顺。
```
