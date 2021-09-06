package taskpool

type Task struct {
	F func(p ...interface{}) error
	Param interface{}
}

func (t *Task) Do() error{
	return t.F(t.Param)
}

type Pool struct {
	tasks chan *Task
	cap int
}

func NewTask(f func(params ...interface{}) error,p ...interface{}) *Task{
	return &Task{f,p}
}
//dasdasdasd
func NewPool(cap int) *Pool{
	return &Pool{
		make(chan *Task),
		cap,
	}
}
//dasdasdasd
func (p *Pool) Enqueue(t *Task){
	p.tasks <- t
}

func (p *Pool) Run(){
	//根据内存池里数量启动协程
	for i := 0; i < p.cap; i++{
		go p.do(i+1)
	}
}

func (p *Pool) Stop(){
	close(p.tasks)
}

func (p *Pool) do(id int) {
	//task不断的从dequeue队列中拿到任务
	for  {
		//如果拿到任务,则执行task任务
		if task,ok := <- p.tasks;ok{
			task.Do()
		}else {
			break
		}
	}
}