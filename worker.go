package main

import (
        // "io/ioutil"
        // "strings"
        //"time"
        // "bytes"
        // "log"
)

func NewWorker(ID int, JobChan JobChannel, Queue JobQueue, Quit chan struct{}) *Worker {
        return &Worker{
                ID:      ID,
                JobChan: JobChan,
                Queue:   Queue,
                Quit:    Quit,
        }
}

func (wr *Worker) Start() {
        go func() {
        for {

                wr.Queue <- wr.JobChan
                select {
                        case job := <-wr.JobChan:
                              // log.Println("Worker", job.ID)
				                      callApi(job.ID, job.StatusCode)
                              // log.Println("Close", job.ID)
			                  case <-wr.Quit:

                              close(wr.JobChan)
                              return
                }
        }
        }()
}


func (wr *Worker) Stop() {
        close(wr.Quit)
}

func callApi(id, statusCode int) {
        HttpError +=1
        // start := time.Now()
        resp, _ := DefaultClient.Do(DefaultRequest)
        // finish := time.Since(start).Milliseconds()



        // if resp.StatusCode != statusCode{
        //
        //
        //         if resp.StatusCode >= 500 {
        //                 defer func() {
        //                     if err := recover(); err != nil {
        //                         // ErrorLog.Println(err)
        //                     }
        //                 }()
        //
        //         } else {
        //                 ErrorLog.Println(err)
        //         }
        //
        //
        // }else {
        //     SumLatency += finish
        // }

        // log.Printf("%d  %d:: ok", id, resp.StatusCode)
        // buf := new(bytes.Buffer)
        // buf.ReadFrom(resp.Body)
        // newBody := buf.String()

        AccessLog.Println(id)
        defer resp.Body.Close()
}
