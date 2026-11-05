package main

//go build -race main.go  编译后运行查看

import (
	"fmt"
	"sync"
	"time"
)

var m = make(map[int]int, 0)
var wg sync.WaitGroup

var mutex sync.Mutex

func test(num int) {
	mutex.Lock()
	var sum = 1

	for i := 1; i <= num; i++ {
		sum *= i
	}
	m[num] = sum
	fmt.Printf("key=%v value=%v\n", num, sum)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

func main() {
	for i := 0; i < 40; i++ {
		wg.Add(1)
		go test(i)
	}

	wg.Wait()

	/*
		key=0 value=1
		key=13 value=6227020800
		key=1 value=1
		key=2 value=2
		key=3 value=6
		key=7 value=5040
		key=8 value=40320
		key=10 value=3628800
		key=5 value=120
		key=11 value=39916800
		key=6 value=720
		key=12 value=479001600
		key=4 value=24
		key=22 value=-1250660718674968576
		key=19 value=121645100408832000
		key=21 value=-4249290049419214848
		key=17 value=355687428096000
		key=23 value=8128291617894825984
		key=9 value=362880
		key=20 value=2432902008176640000
		key=15 value=1307674368000
		key=16 value=20922789888000
		key=14 value=87178291200
		key=18 value=6402373705728000
		key=24 value=-7835185981329244160
		key=25 value=7034535277573963776
		key=26 value=-1569523520172457984
		key=35 value=6399018521010896896
		key=29 value=-7055958792655077376
		key=27 value=-5483646897237262336
		key=36 value=9003737871877668864
		key=37 value=1096907932701818880
		key=38 value=4789013295250014208
		key=39 value=2304077777655037952
		key=32 value=-6045878379276664832
		key=30 value=-8764578968847253504
		key=33 value=3400198294675128320
		key=34 value=4926277576697053184
		key=28 value=-5968160532966932480
		key=31 value=4999213071378415616
	*/
}
