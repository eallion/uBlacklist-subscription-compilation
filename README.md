# uBlacklist-Websites
谷歌搜索结果中屏蔽机器翻译网站


## 介绍

事情发生于上周公司企业微信群里的一段对话：

<img src="https://github.com/liubiantao/uBlacklist-Websites/raw/master/imgs/chat1.png" alt="chat1" style="max-width:100%;width:300px" width="300">
<img src="https://github.com/liubiantao/uBlacklist-Websites/raw/master/imgs/chat2.png" alt="chat2" style="max-width:100%;width:300px" width="300">

今天搜索过程中又被垃圾站恶心到了，忍无可忍向谷歌点了举报，后来逛 V 站的时候发现了这个好方法，决定维护一个屏蔽列表，大家一起来跟流氓们做斗争。


## 用法

下载插件 [uBlacklist - Chrome 网上应用店](https://chrome.google.com/webstore/detail/ublacklist/pncfbmialoiaghdehhbnbhkkgmjanfhe)



在设置中添加 Subscription: https://raw.githubusercontent.com/liubiantao/uBlacklist-Websites/master/uBlacklist.txt

<img src="https://github.com/liubiantao/uBlacklist-Websites/raw/master/imgs/setting.png" alt="setting" style="max-width:100%;width:500px" width="500">


### 对比一下效果

之前：

<img src="https://github.com/liubiantao/uBlacklist-Websites/raw/master/imgs/before.png" alt="before" style="max-width:100%;width:500px" width="500">

之后：

<img src="https://github.com/liubiantao/uBlacklist-Websites/raw/master/imgs/after.png" alt="after" style="max-width:100%;width:500px" width="500">


## 如何贡献网址

方法一（还没实现）：

直接复制地址栏的地址，粘贴到这里 [addUrl.txt](https://github.com/liubiantao/uBlacklist-Websites/blob/master/addUrl.txt)，提交 PR。
之后，CI 会自动做去重以及格式化处理并更新 `uBlacklist.txt` 

方法二：

按格式修改 [uBlacklist.txt](https://github.com/liubiantao/uBlacklist-Websites/blob/master/uBlacklist.txt) 并提交 PR。

## 感谢

[google 搜索的时候建议屏蔽这些垃圾网站 - V2EX](https://www.v2ex.com/t/593519)