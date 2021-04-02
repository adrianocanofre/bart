package main

import (
        "os"
        "net/http"
        "time"
        "gopkg.in/yaml.v2"
        "bytes"
        "log"
)


func OpenFile() (*Config, error) {
        config := &Config{}

        file, err := os.Open("./load.yaml")
        if err != nil {
                return nil, err
        }
        defer file.Close()

        d := yaml.NewDecoder(file)

        if err := d.Decode(&config); err != nil {
                return nil, err
        }

        return config, nil
}

func ConfigClient(timeout time.Duration)(http.Client){
        tr := &http.Transport{
                MaxIdleConns:       100,
                MaxIdleConnsPerHost:  20,
                DisableKeepAlives: true,
        }

        client := http.Client{
                Timeout : timeout ,
                Transport: tr,
        }

        return client
}

func ConfigRequest(method string, url string, body string)(*http.Request){
        var jsonStr = []byte(body)

        request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
        request.Header.Set("Content-Type", "application/json")
        if err != nil {
                ErrorLog.Println(err)
        }
        return request
}


func init() {
        _, err := os.Stat("log")

        if os.IsNotExist(err) {
                errDir := os.MkdirAll("log", 0755)
                if errDir != nil {
                        log.Fatal(err)
                }
        }
        access, err := os.OpenFile("log/access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
        if err != nil {
                log.Fatal(err)
        }

        error, err := os.OpenFile("log/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
        if err != nil {
                log.Fatal(err)
        }

        AccessLog = log.New(access, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
        ErrorLog = log.New(error, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
