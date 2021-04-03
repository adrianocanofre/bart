package main

import (
        // "log"
        "time"
        "fmt"
				// "sync"
)




func main() {
        // start := time.Now()

        dd := NewDispatcher(cfg.Concurrency).Start()

				DefaultClient = ConfigClient(cfg.Http.Timeout * time.Second)
				DefaultRequest = ConfigRequest(cfg.Http.Method, cfg.Http.Url, cfg.Http.Body)

				fmt.Println(cfg.Requests)
        for i := 0; i < cfg.Requests; i++ {
                dd.Submit(Job{
												ID: i,
												StatusCode:  cfg.Http.StatusCode,
										})
        }
        // end := time.Now()
        // log.Print(end.Sub(start).Seconds())
				fmt.Println(SumLatency)
				fmt.Println(HttpError)
}
