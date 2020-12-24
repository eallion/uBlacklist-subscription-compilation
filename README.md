# uBlacklist subscription compilation
> uBlacklist 订阅地址合集

![Update uBlacklist subcription weekly](https://github.com/eallion/uBlacklist-subscription-compilation/workflows/Update%20uBlacklist%20subcription%20weekly/badge.svg?branch=main) [![](https://data.jsdelivr.com/v1/package/gh/eallion/uBlacklist-subscription-compilation/badge?style=rounded)](https://www.jsdelivr.com/package/gh/eallion/uBlacklist-subscription-compilation) ![Visitors](https://visitor-badge.laobi.icu/badge?page_id=eallion.uBlacklist-subscription-compilation)

这是一个 uBlacklist 订阅地址合集，搜集了网上大部分的订阅地址合并成一个。  
通过 [Github Actions](https://github.com/eallion/uBlacklist-subscription-compilation/actions) 每周自动更新一次。

本列表**比较激进**！会屏蔽少部分正规但低质量的中文网站。  
所以个人建议搜索中文资料请用百度，uBlacklist 插件对百度无效。

# 功能

通过匹配域名的方式，屏蔽搜索引擎的搜索结果。主要是屏蔽中文SEO垃圾站。  

支持的搜索引擎：（由 uBlacklist 插件决定）
- [Google](https://www.google.com/ncr?gws_rd=ssl)
- [DuckDuckGo](https://duckduckgo.com/)
- [Startpage](https://www.startpage.com/)

# 用法：

### 一、下载 uBlacklist 浏览器插件：
- Chrome： <https://chrome.google.com/webstore/detail/ublacklist/pncfbmialoiaghdehhbnbhkkgmjanfhe>
- Firefox： <https://addons.mozilla.org/en-US/firefox/addon/ublacklist/>
- Edge：<https://chrome.google.com/webstore/detail/ublacklist/pncfbmialoiaghdehhbnbhkkgmjanfhe>

### 二、设置插件

#### 设置 - 订阅：

> 订阅黑名单列表:

添加订阅：  
选择一个列表添加到 Subscription 订阅中：（如下图所示）  

- GitHub 列表 （大陆部分地区可能被墙或者速度慢）

`https://git.io/ublacklist`  
`https://git.io/ublocklist` # For BLM

- CDN 加速列表

`https://cdn.jsdelivr.net/gh/eallion/uBlacklist-subscription-compilation@main/uBlacklist.txt`
`https://cdn.jsdelivr.net/gh/eallion/uBlacklist-subscription-compilation@main/uBlocklist.txt` # For BLM

上面4个列表的内容一模一样，任意选择一个即可。

**能稳定访问 GitHub 的用户推荐使用：**  
```
https://git.io/ublacklist
```
> <https://git.io/ublacklist>

示例图：

![](https://cdn.jsdelivr.net/gh/eallion/uBlacklist-subscription-compilation@main/tools/x2yWi62OWl.png)

#### 设置 - 一般：（选填）

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

![](https://cdn.jsdelivr.net/gh/eallion/uBlacklist-subscription-compilation@main/tools/PKO0JoJBpU.png)

# 贡献

1. Fork 本仓库
2. 把需要禁用的域名按格式添加到 add_url.txt 里
3. 提交 PR

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
