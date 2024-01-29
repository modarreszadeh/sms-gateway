package queue

type T interface{}

type Queue struct {
	Queue     chan T
	QueueSize int
}

func New(queueSize int) *Queue {
	q := Queue{
		QueueSize: queueSize,
		Queue:     make(chan T, queueSize),
	}

	return &q
}

func (q *Queue) Enqueue(item T) {
	q.Queue <- item
}

func (q *Queue) DispatchProcess(process func(interface{})) {
	go func() {
		for {
			select {
			case item := <-q.Queue:
				go func() {
					process(item)
				}()
			}
		}
	}()
}
