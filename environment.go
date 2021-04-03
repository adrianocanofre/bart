package main

import (
        "time"
        "net/http"
				"log"
)

var (
        DefaultRequest *http.Request
        DefaultClient http.Client
        SumLatency int64
        AccessLog  *log.Logger
        ErrorLog   *log.Logger
				HttpError  int
        cfg *Config
)

type Config struct {
        Requests int `yaml:"requests"`
        Concurrency int `yaml:"concurrency"`
        Output string `yaml:"output"`
        Body string `yaml:"body"`
        PathLog string `yaml:"pathLog"`
        Http   struct {
          Url string `yaml:"url"`
          Method string `yaml:"method"`
          StatusCode int `yaml:"statusCode"`
          Body string `yaml:"body"`
          Timeout time.Duration `yaml:"timeout"`
        } `yaml:"http"`
}

type Job struct {
		    ID          int
        StatusCode  int
        UpdatedAt  time.Time
}

type disp struct {
        Workers  []*Worker
        WorkChan JobChannel
        Queue    JobQueue
}


type JobChannel chan Job
type JobQueue chan chan Job

type Worker struct {
        ID      int
        JobChan JobChannel
        Queue   JobQueue
        Quit    chan struct{}
}
