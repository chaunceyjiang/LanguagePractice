package main
/* 
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomNo int
}
type Result struct {
	id          int
	sumOfDigits int
	job         Job
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	// n:=rand.Int63n(2)
	var n int64 = 2
	time.Sleep(time.Duration(n * (1 << 30)))
	return sum
}
func worker(wg *sync.WaitGroup, id int) {
	for job := range jobs { //从channel中读取数据，若cannel中没有数据，则阻塞在这
		output := Result{
			id:          id,
			job:         job,
			sumOfDigits: digits(job.randomNo), //工作
		}
		results <- output //输出
	}
	// channel被关闭，woker任务结束
	wg.Done()
}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomNo := rand.Intn(999)
		job := Job{
			id:       i,
			randomNo: randomNo,
		}
		jobs <- job
	}
	close(jobs)
}
func createWorkerPoll(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}
	wg.Wait()
	close(results)
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d ,input random no %d ,sum digits %d Worker id %d\n",
			result.job.id, result.job.randomNo, result.sumOfDigits, result.id)
	}
	done <- true
}
func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs) //产生任务

	done := make(chan bool)
	go result(done) //取结果

	noOfWorkers := 10
	createWorkerPoll(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "second")
}
 */