package multi_process

/**
1. 新到来的Job会被写入JobQueue通道中
排程中心工作：
1. 排程中心会不断地从JobQueue中读取新的Job
2. 有新的Job时会从WorkerPool中读取一个JobChannel
3. 将新Job写入JobChannel
Worker工作
1. 不断地将当前的JobChannel注册到WorkerPool
2. 当JobChannel中有可读的Job时，处理该Job
*/
import (
	"fmt"
	"reflect"
	"time"
)

type Task struct {
	Num int
}
type Job struct {
	Task Task
}

var (
	MaxWorker = 5
	JobQueue  chan Job
)

type Worker struct {
	id         int
	WorkerPool chan chan Job // 工作者池，通道的通道，每一个元素都是一个job通道
	JobChannel chan Job      // 工作通道，每个元素是一个job
	exit       chan bool
}

func NewWorker(workerPool chan chan Job, id int) Worker {
	fmt.Printf("a new worker(%d)\n", id)
	return Worker{
		id:         id,
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		exit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			// 将当前的任务队列注册到工作池
			w.WorkerPool <- w.JobChannel
			fmt.Println("register private JobChannel to public WorkerPool", w)
			select {
			case job := <-w.JobChannel:
				fmt.Println("get a job from private w.JobChannel")
				fmt.Println(job)
				time.Sleep(5 * time.Second)
			case <-w.exit:
				fmt.Println("worker exit", w)
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.exit <- true
	}()
}

// 排程中心
type Scheduler struct {
	WorkerPool chan chan Job // 和workerPool用的是同一个
	MaxWorkers int
	Workers    []*Worker
}

// 创建排程中心
func NewScheduler(maxWorkers int) *Scheduler {
	pool := make(chan chan Job, maxWorkers)
	return &Scheduler{
		WorkerPool: pool,
		MaxWorkers: maxWorkers,
	}
}

// 工作池初始化
func (s *Scheduler) Create() {
	workers := make([]*Worker, s.MaxWorkers)
	for i := 0; i < s.MaxWorkers; i++ {
		worker := NewWorker(s.WorkerPool, i)
		worker.Start()
		workers[i] = &worker
	}
	s.Workers = workers
	go s.schedule()
}

// 工作池的关闭
func (s *Scheduler) Shutdown() {
	workers := s.Workers
	for _, w := range workers {
		w.Stop()
	}
	time.Sleep(time.Second)
	close(s.WorkerPool)
}

// 排程
func (s *Scheduler) schedule() {
	for {
		select {
		case job := <-JobQueue:
			fmt.Println("get a job from JobQueue")
			go func(job Job) {
				//	从workerPool获取jobChannel
				jobChannel := <-s.WorkerPool
				fmt.Println("get a private jobChannel from public s.WorkerPool", reflect.TypeOf(jobChannel))
				jobChannel <- job
				fmt.Println("worker's private jobChannel add one job")
			}(job)
		}
	}
}

func WorkPool() {
	JobQueue = make(chan Job, 5)
	scheduler := NewScheduler(MaxWorker)
	scheduler.Create()
	time.Sleep(1 * time.Second)
	go createJobQueue()
	time.Sleep(100 * time.Second)
	scheduler.Shutdown()
	time.Sleep(10 * time.Second)

}

func createJobQueue() {
	for i := 0; i < 30; i++ {
		task := Task{Num: i}
		job := Job{Task: task}
		JobQueue <- job
		fmt.Println("JobQueue add one job")
		time.Sleep(time.Second)
	}
}
