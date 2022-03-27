### 1 总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用
### 2 实现一个从 socket connection 中解码出 goim 协议的解码器。

#### 答1 
···
粘包 怎么从一串字符里协议好获取的内容。

方式1: fix length
发送方，每次发送固定长度的数据，并且不超过缓冲区，接受方每次按固定长度区接受数据。例 MACU

方式2: delimiter based
发送方，在数据包添加特殊的分隔符，用来标记数据包边界.例 /n ,; 这些分隔符

方式3: length field based
发送方，在消息数据包头添加包长度信息 例：双方协定前2个字符是长度
···

#### 答2
见 main.go