package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct { //定义结构体
	queue []string
	cond  *sync.Cond
}

func (q *Queue) Enqueue(item string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, item) //实际的入队操作。
	fmt.Printf("入队:[%s]\n", item)
	q.cond.Broadcast() //广播给消费者立刻消费掉。
	//q.cond.Signal() //也可以发送单个通知。
}
func (q *Queue) Dequeue() string { //收到广播后立刻消费。
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.queue) == 0 { //如果没有数据就继续等待。
		fmt.Println("队列空，等待中。")
		q.cond.Wait()
		return ""
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	fmt.Println("出队:", result)
	return result
}

func SyncGoSliceMain(iEnq int, iDnq int, iTime int) {
	q := Queue{ //实例化结构体
		queue: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	for i := 0; i < iEnq; i++ {
		func(i int) {
			go func() {
				for {
					Msg := fmt.Sprintf("EnqueueId : %06d", i)
					q.Enqueue(Msg) //入队
					time.Sleep(time.Second * 2)
				}
			}()
		}(i)
	}

	for i := 0; i < iDnq; i++ {
		go func() {
			for {
				q.Dequeue() //出队
				time.Sleep(time.Second)
			}
		}()
	}
	for i := 0; i < iTime; i++ {
		time.Sleep(time.Second)
	}

}
