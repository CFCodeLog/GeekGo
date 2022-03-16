package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/hhxsv5/go-redis-memory-analysis"
)

var client redis.UniversalClient
var ctx context.Context

const (
	ip   string = "127.0.0.1"
	port uint16 = 6379
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%v:%v", ip, port),
		Password:     "",
		DB:           0,
		PoolSize:     128,
		MinIdleConns: 100,
		MaxRetries:   5,
	})

	ctx = context.Background()
}

func main() {
	reset()
	analysis("before_all")
	write(10000, "len10_10k", generateValue(10))
	analysis("after_len10_10k")

	reset()
	write(50000, "len10_50k", generateValue(10))
	analysis("after_len10_50k")

	reset()
	write(500000, "len10_500k", generateValue(10))
	analysis("after_len10_500k")
	// write(10000, "len1000_10k", generateValue(1000))
	// write(50000, "len1000_50k", generateValue(1000))
	// write(500000, "len1000_500k", generateValue(1000))

	// write(10000, "len5000_10k", generateValue(5000))
	// write(50000, "len5000_50k", generateValue(5000))
	reset()
	write(10000, "len5000_10k", generateValue(5000))
	analysis("after_len5000_10k")
	reset()
	write(50000, "len5000_50k", generateValue(5000))
	analysis("after_len5000_50k")
	reset()
	write(500000, "len5000_500k", generateValue(5000))
	analysis("after_len5000_500k")

	reset()
}
func reset() {
	client.FlushAll(ctx)
}
func write(num int, key, value string) {
	for i := 0; i < num; i++ {
		k := fmt.Sprintf("%s:%v", key, i)
		cmd := client.Set(ctx, k, value, -1)
		err := cmd.Err()
		if err != nil {
			fmt.Println(cmd.String())
		}
	}
}

func analysis(name string) {
	analysis, err := gorma.NewAnalysisConnection(ip, port, "")
	if err != nil {
		fmt.Println("something wrong:", err)
		return
	}
	defer analysis.Close()

	analysis.Start([]string{":"})

	err = analysis.SaveReports("./reports" + name)
	if err == nil {
		fmt.Println("done")
	} else {
		fmt.Println("error:", err)
	}
}

func generateValue(size int) string {
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		arr[i] = 'a'
	}
	return string(arr)
}
