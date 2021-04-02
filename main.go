package main

import (
        // "log"
        "time"
        "fmt"
				// "sync"
)




func main() {
        // start := time.Now()

        cfg, err := OpenFile()
				if err != nil {
                ErrorLog.Println(err)
        }
        dd := NewDispatcher(cfg.Concurrency).Start()

				DefaultClient = ConfigClient(cfg.Timeout * time.Second)
				DefaultRequest = ConfigRequest(cfg.Method, cfg.Url, cfg.Body)

				fmt.Println(cfg.Requests)
        for i := 0; i < cfg.Requests; i++ {
                dd.Submit(Job{
												ID: i,
												StatusCode:  cfg.StatusCode,
										})
        }
        // end := time.Now()
        // log.Print(end.Sub(start).Seconds())
				fmt.Println(SumLatency)
				fmt.Println(HttpError)
}
