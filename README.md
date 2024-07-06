# go-toolbox

类似Java Hutool的工具包

所有方法都和 Hutool 包保持一致

注：

- 并不保证所有方法都有
- 常用方法将会持续添加进去
- 使用上尽量和 Java Hutool 保持一致
- 尽量不引入其他过多依赖

欢迎各位大佬加入项目

# coll

集合类工具

# id

> id类工具

- RandomUUID 随机获取UUID
- SimpleUUID 去掉了横线的UUID
- ObjectID 通过时间戳生成的MongoID
- GetSnowFlakeNextId 获取雪花ID
- GetSnowFlakeNextIdStr 获取雪花ID字符串

# number

> 数值类工具

- Min 返回整数中的最小值
- Max 返回整数中的最大值

# str

> 字符串类工具

- IsEmpty 检测是否是空串
- IsNotEmpty 检测是否是空串
- Replace 替换字符串
- ReplaceAll 替换字符串
- IndexOf 返回字符在原始字符串的下标
- SubString 切割原始字符，返回子串 范围：[s,e)
- ParseStr 解析字符串
- FillBefore 将已有字符串【向前】填充为规定长度，如果已有字符串超过这个长度则返回这个字符串
- FillAfter 将已有字符串【向后】填充为规定长度，如果已有字符串超过这个长度则返回这个字符串
- Fill 将已有字符串填充为规定长度，如果已有字符串超过这个长度则返回这个字符串
- Similar 计算两个字符串的相似度
- SimilarScale 计算两个字符串的相似度百分比
- IsAlphaNumericOrChinese 检查字符是否为字母、数字或汉字
- Format 通过map中的参数 格式化字符串