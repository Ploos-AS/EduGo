package workerpool

import (
	"context"
	"sync"
)

type Job struct {
	ID    int
	Value int
}

type Result struct {
	ID     int
	Square int
}

func Run(ctx context.Context, workers int, jobs []Job) []Result {
	if workers < 1 {
		workers = 1
	}
	in := make(chan Job)
	out := make(chan Result)

	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-in:
					if !ok {
						return
					}
					select {
					case out <- Result{ID: job.ID, Square: job.Value * job.Value}:
					case <-ctx.Done():
						return
					}
				}
			}
		}()
	}

	go func() {
		defer close(in)
		for _, job := range jobs {
			select {
			case in <- job:
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	results := make([]Result, 0, len(jobs))
	for result := range out {
		results = append(results, result)
	}
	return results
}
