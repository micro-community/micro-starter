package workerpool

type TaskFunction = func() (interface{}, error)
type Task struct {
	f TaskFunction
}

// a worker group that runs a number of task concurrently
type Pool struct {
	Tasks       []*Task
	concurrency int
	tasksChan   chan *Task
}

func NewTask(f TaskFunction) *Task {
	return &Task{f: f}
}

// execute task function and add its return value to channel
func (t *Task) Run(resultsChan chan<- interface{}) {
	result, err := t.f()
	if err != nil {
		resultsChan <- err
	} else {
		resultsChan <- result
	}
}

// initializes a new pool with the given tasks at the given concurrency
func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *Task, 1),
	}
}

// loop for taskChan and execute every task
func (p *Pool) execute(results chan<- interface{}) {
	for task := range p.tasksChan {
		task.Run(results)
	}
}

// runs all work within the pool and blocks until it's finished
func (p *Pool) Run() *[]interface{} {
	taskCount := len(p.Tasks)
	// pass each task result to this channel
	taskResultsChan := make(chan interface{}, taskCount)
	// store each task result in this array
	var resultsArray []interface{}

	for i := 0; i < p.concurrency; i++ {
		go p.execute(taskResultsChan)
	}

	// dispatch task
	for _, task := range p.Tasks {
		p.tasksChan <- task
	}
	close(p.tasksChan)

	for i := 0; i < taskCount; i++ {
		resultsArray = append(resultsArray, <-taskResultsChan)
	}

	return &resultsArray
}

type CallbackFunc = func(taskResult interface{}) (interface{}, error)

// runs all work within the pool, and execute callback function on each task result
func (p *Pool) RunWith(callback CallbackFunc) *[]interface{} {
	taskCount := len(p.Tasks)
	taskResultsChan := make(chan interface{}, taskCount)

	var resultsArray []interface{}

	for i := 0; i < p.concurrency; i++ {
		go p.execute(taskResultsChan)
	}

	// dispatch task
	for _, task := range p.Tasks {
		p.tasksChan <- task
	}
	close(p.tasksChan)

	for i := 0; i < taskCount; i++ {
		value, err := callback(<-taskResultsChan)
		if err != nil {
			resultsArray = append(resultsArray, err)
		} else {
			resultsArray = append(resultsArray, value)
		}
	}

	return &resultsArray
}
