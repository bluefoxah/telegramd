# mtproto协议生成工具

## 支持协议版本
> 当前支持的协议为layer71
[API scheme layer 71](https://github.com/telegramdesktop/tdesktop/blob/b0cc61c621c1643c1180921a83540c59be9642a3/Telegram/Resources/scheme.tl)

## 简单说明
proto生成的go代码文件太大，超过2.5MB，导致IDE语法分析器无法识别，拆分成多个proto文件

## TODO
- 重构代码生成器
- 增加mtproto的Encode和Decode各种出错条件的判断
- mtproto的Encode和Decode性能优化
