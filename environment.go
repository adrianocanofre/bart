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
        Url string `yaml:"url"`
        Method string `yaml:"method"`
        Timeout time.Duration `yaml:"timeout"`
        StatusCode int `yaml:"statusCode"`
        Output string `yaml:"output"`
        Body string `yaml:"body"`
        PathLog string `yaml:"pathLog"`
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
