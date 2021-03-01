package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id      int
	randnum int
}

type Result struct {
	job         Job
	sumofdigits int
}

// 创建缓冲信道
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// 每个子任务，在每个woker中
func digits(number int) int {
	sum := 0
	num := number
	for num != 0 {
		digit := num % 10
		sum += digit
		num /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

// 创建了一个工作者（Worker），读取 jobs 信道的数据，根据当前的 job 和 digits 函数的返回值，
// 创建了一个 Result 结构体变量，然后将结果写入 results 缓冲信道。
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randnum)}
		results <- output
	}
	wg.Done()
}

// 创建了一个 Go 协程的工作池
func createWokerPool(numOfWokers int) {
	var wg sync.WaitGroup
	for i := 0; i < numOfWokers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	//所有协程完成执行之后，函数会关闭 results 信道。
	// 因为所有协程都已经执行完毕，于是不再需要向 results 信道写入数据了。
	close(results)
}

// 我们继续编写一个函数，把作业分配给工作者。
func allocate(numOfJobs int) {
	for i := 0; i < numOfJobs; i++ {
		randomnum := rand.Intn(999)
		job := Job{i, randomnum}
		jobs <- job
	}
	// 当写入所有的 job 时，它关闭了 jobs 信道。
	close(jobs)
}

// 读取 results 信道和打印输出的函数。
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randnum, result.sumofdigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	numOfJobs := 100
	go allocate(numOfJobs)

	// 创建了 done 信道，并将其传递给 result 协程。
	done := make(chan bool)
	go result(done)

	// 创建了一个有 10 个协程的工作池
	numOfWokers := 10
	createWokerPool(numOfWokers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
