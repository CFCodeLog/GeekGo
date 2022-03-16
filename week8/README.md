1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
```
docker run --restart always --name=redis --publish=6379:6379 --hostname=redis --restart=on-failure --detach redis:latest
docker exec -it redis /bin/bash
redis-benchmark -t set,get -n 100000 -d 10
redis-benchmark -t set,get -n 100000 -d 20
redis-benchmark -t set,get -n 100000 -d 50
redis-benchmark -t set,get -n 100000 -d 100
redis-benchmark -t set,get -n 100000 -d 200
redis-benchmark -t set,get -n 100000 -d 1000
redis-benchmark -t set,get -n 100000 -d 5000
```

```
root@redis:/data# redis-benchmark -t set,get -n 100000 -d 10 -q
SET: 116279.07 requests per second, p50=0.215 msec
GET: 114547.53 requests per second, p50=0.215 msec

root@redis:/data# redis-benchmark -t set,get -n 100000 -d 20 -q
SET: 114678.90 requests per second, p50=0.215 msec
GET: 114416.48 requests per second, p50=0.215 msec

root@redis:/data# redis-benchmark -t set,get -n 100000 -d 30 -q
SET: 113122.17 requests per second, p50=0.223 msec
GET: 112485.94 requests per second, p50=0.223 msec

root@redis:/data# redis-benchmark -t set,get -n 100000 -d 40 -q
SET: 112107.62 requests per second, p50=0.223 msec
GET: 111607.14 requests per second, p50=0.223 msec

root@redis:/data# redis-benchmark -t set,get -n 100000 -d 50 -q
SET: 109289.62 requests per second, p50=0.223 msec
GET: 112612.61 requests per second, p50=0.223 msec

root@redis:/data# redis-benchmark -t set,get -n 100000 -d 100 -q
SET: 115074.80 requests per second, p50=0.215 msec
GET: 111111.12 requests per second, p50=0.223 msec

root@redis:/data# redis-benchmark -t set,get -n 100000 -d 200 -q
SET: 113250.28 requests per second, p50=0.223 msec
GET: 108108.11 requests per second, p50=0.231 msec

root@redis:/data# redis-benchmark -t set,get -n 100000 -d 500 -q
SET: 96805.42 requests per second, p50=0.239 msec
GET: 107758.62 requests per second, p50=0.231 msec

root@redis:/data# redis-benchmark -t set,get -n 100000 -d 1000 -q
SET: 115740.73 requests per second, p50=0.215 msec
GET: 117233.30 requests per second, p50=0.215 msec

root@redis:/data# redis-benchmark -t set,get -n 100000 -d 5000 -q
SET: 104275.29 requests per second, p50=0.239 msec
GET: 100704.94 requests per second, p50=0.247 msec
```


2.写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

```
Key	Count	Size	NeverExpire	AvgTtl(excluded never expire)	Size	Per key
len10_10k:*	10000	107.422 KB	10000	0	107.422	0.0107422
len10_50k:*	50000	537.109 KB	50000	0	537.109	0.01074218
len10_500k:*	500000	5.245 MB	500000	0	5370.88	0.01074176
len5000_10k:*	10000	654.297 KB	10000	0	654.297	0.0654297
len5000_50k:*	50000	3.195 MB	50000	0	3271.68	0.0654336
```
结论：相同长度的value写入次数越多，单key占用空间越小。