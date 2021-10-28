/*
官方文档地址：https://redis.uptrace.dev
仓库地址：github.com/go-redis/redis
参考文档：https://www.liwenzhou.com/posts/Go/go_redis/
*/

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
)

var rdb *redis.Client

// 初始化连接
func initRedisClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.100.61:6379",
		Password: "yuedong",
		DB:       0,
		PoolSize: 100,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = rdb.Ping(ctx).Result()
	return
}

func V8Example() {
	ctx := context.Background()
	err := rdb.Set(ctx, "key", "test_value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key:", val)
	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2:", val2)
	}
}

func redisExample() {
	ctx := context.Background()
	err := rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}
	val, err := rdb.Get(ctx, "score").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("score:", val)
	val2, err := rdb.Get(ctx, "name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
		return
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}

func redisExample2() {
	ctx := context.Background()
	zsetKey := "language_rank"
	// languages := []*redis.Z{
	// 	{Score: 90.0, Member: "Golang"},
	// 	{Score: 98.0, Member: "Java"},
	// 	{Score: 95.0, Member: "python"},
	// 	{Score: 97.2, Member: "JavaScript"},
	// 	{Score: 99.0, Member: "C/C++"},
	// }
	// zadd
	// num, err := rdb.ZAdd(ctx, zsetKey, languages...).Result()
	// if err != nil {
	// 	fmt.Printf("zadd failed, err:%\n", err)
	// 	return
	// }
	// fmt.Printf("zadd %d succ.\n", num)
	// 把Golang的分数加10
	// newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	// if err != nil {
	// 	fmt.Printf("zincrby failed, err:%v\n", err)
	// 	return
	// }
	// fmt.Printf("Golang's score is %f now.\n", newScore)
	// 取分值最高的三个
	fmt.Println("====================================")
	ret, err := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
	// 取95~100分的
	fmt.Println("=======取95~100分=========")
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

}

func main() {
	if err := initRedisClient(); err != nil {
		return
	}
	// V8Example()
	// redisExample()
	redisExample2()
}
