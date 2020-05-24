---
title: 学习 Vim 命令总结
date: "2019-09-28T20:54:03.284Z"
description: 学习 Vim 命令总结
---

> 可以使用 vscode-vim 扩展，但是要注意一些ctrl+字母的快捷键会无效，必须去掉冲突的快捷键

`esc` 回到普通模式

`i` 普通模式进入插入模式

`:` 进入命令模式

`:wa` 保存全部文件

`dd` 删除一整行

`u` 撤销

`r` 替换

`x` 删除一个字符

`cw` 修改从光标到一个单词结尾的字符

`%` 匹配括号移动

`*` 移动到匹配的下个单词

`#` 移动到匹配的上个单词

`gg` 到第一行

`.` 重复上次命令

`G` 最后一行（大写的G就是按 shift+g）

`数字+G` 到第几行，然后按 `` 可以跳回去

`ma` 标记为 a, 然后按 `a 跳过去

`ye` 复制单词

`0y$` 从行头复制到行尾

`y2/foo` 复制 2 个foo之间的字符

`gU` 全变大写

`gu` 全变小写

`^` 本行第一个字符

`g_` 本行最后一个字符

`fa` 到下一个为 a 的字符处

`t,` 到下一个为 , 的符号处

`dt"` 删除到 " 前的字符

`vi"` 选择 " 里面的字符

`va"` 选择包括 " 的字符

`J` 连成一行

`<<` 左缩进

`>>` 右缩进

`==` 自动缩进

## 用正则替换文本

`%s/old/new/g` 替换所有匹配的字符

`10,$s/old/new` 替换10行到最后的字符

`10,20s/old/new` 替换10到20行到最后的字符

## 录制宏

`qa` 开始录制到a

`yp` 复制粘贴

`ctrl+a` 递增 1（在 vscode-vim 中无效）

`q` 停止录制

`@a` 递增写入

`10@a` 重复写10次

`10@@` 递增写 10 次