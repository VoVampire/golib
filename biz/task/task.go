package task

type Task struct {
}

func doTask(n int) {
	ch := make(chan Task, 3)
	for i := 0; i < n; i++ {
		go worker(ch)
	}

	tasks := tasks()
	for i := range tasks {
		ch <- tasks[i]
	}

	<-make(chan bool)
}

func worker(ch chan Task) {
	for {
		task := <-ch
		process(task)
	}
}

func process(Task) {

}

func tasks() []Task {
	return nil
}
