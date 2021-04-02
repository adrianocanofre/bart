package main


func NewDispatcher(num int) *disp {
        return &disp{
                Workers:  make([]*Worker, num),
                WorkChan: make(JobChannel),
                Queue:    make(JobQueue),
	      }
}


func (d *disp) Start() *disp {
        l := len(d.Workers)
        for i := 1; i <= l; i++ {
                wrk := NewWorker(i, make(JobChannel), d.Queue, make(chan struct{}))
                wrk.Start()
                d.Workers = append(d.Workers, wrk)
	      }
	      go d.process()
	      return d
}

func (d *disp) process() {
        for {
                select {
                        case job := <-d.WorkChan:
			                  jobChan := <-d.Queue
			                  jobChan <- job
		             }
        }
}

func (d *disp) Submit(job Job) {
        d.WorkChan <- job
}
