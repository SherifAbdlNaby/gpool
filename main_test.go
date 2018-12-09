package main

import (
	"fmt"
	"pipeline/semp"
	"pipeline/work"
	"pipeline/workerpool"
	"sync"
	"testing"
)

var workersCountValues = []int{100000, 10}
var workAmountValues = []int{10000, 100000}

/*var workersCountValues = []int{1, 2, 3, 4, 10, 100, 1000, 10000, 100000}
var workAmountValues = []int{10, 100, 1000, 10000, 100000}
*/
/*func BenchmarkCallAndWait(b *testing.B) {

	for i := 0; i < 2; i++ {
		for _, workercount := range workersCountValues {
			var workerPool worker.ConcurrencyBounder
			var name string
			if i == 0 {
				name = "Workerpool"
			}
			if i == 1 {
				name = "Semaphore"
			}

			b.Run(fmt.Sprintf("[%s]W[%d]", name, workercount), func(b *testing.B) {
				if i == 0 {
					workerPool = workerpool.NewWorkerPool(workercount)
				}
				if i == 1 {
					workerPool = semp.NewSempWorker(int64(workercount))
				}

				workerPool.Start()

				b.ResetTimer()

				for i2 := 1; i2 < b.N; i2++ {
					x, _ := workerPool.Enqueue(&sr, 13376800)
					<-x
				}

				b.StopTimer()
				workerPool.Stop()
			})
		}
	}
}
*/
func BenchmarkBulkWait(b *testing.B) {
	for _, workercount := range workersCountValues {
		for _, work := range workAmountValues {
			for i := 0; i < 2; i++ {
				var workerPool worker.ConcurrencyBounder
				var name string
				if i == 0 {
					name = "Workerpool"
				}
				if i == 1 {
					name = "Semaphore"
				}

				b.Run(fmt.Sprintf("[%s]W[%d]J[%d]", name, workercount, work), func(b *testing.B) {

					b.ResetTimer()

					for i2 := 1; i2 < b.N; i2++ {
						if i == 0 {
							workerPool = workerpool.NewWorkerPool(workercount)
						}
						if i == 1 {
							workerPool = semp.NewSempWorker(int64(workercount))
						}
						workerPool.Start()
						b.StartTimer()
						wg := sync.WaitGroup{}
						wg.Add(work)
						for i3 := 0; i3 < work; i3++ {
							go func() {
								x, _ := workerPool.Enqueue(&sr, 13376800)
								<-x
								wg.Done()
							}()
						}
						wg.Wait()
						b.StopTimer()
						workerPool.Stop()
					}
				})
			}
		}
	}
}

/*func BenchmarkWorkersOverhead(b *testing.B) {

	for i := 0; i < 2; i++ {
		for _, workercount := range workersCountValues {
				var workerPool worker.ConcurrencyBounder
				var name string
				if i == 0 {
					name = "Workerpool"
				}
				if i == 1 {
					name = "Semaphore"
				}

				b.Run(fmt.Sprintf("[%s]W[%d]", name, workercount), func(b *testing.B) {


					b.ResetTimer()

					for i2 := 1; i2 < b.N; i2++ {
						if i == 0 {
							workerPool = workerpool.NewWorkerPool(workercount)
						}
						if i == 1 {
							workerPool = semp.NewSempWorker(int64(workercount))
						}
						workerPool.Start()
						workerPool.Stop()
					}

					b.StopTimer()
				})
		}
	}
}

*/
