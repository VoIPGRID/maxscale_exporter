package main

import (
  "fmt"
  "strings"
  "net/http"
  "log"
  "encoding/json"
  "io/ioutil"
  "strconv"
  "time"
  "os"
)

type Server struct {
  Server string
  Address string
  Port int
  Connections int
  Status string
}

type Service struct {
  Name string `json:"Service Name"`
  Router string `json:"Router Module"`
  Sessions int `json:"No. Sessions"`
}

type Status struct {
  Name string `json:"Variable_name"`
  Value int `json:"Value"`
}

type Variable struct {
  Name string `json:"Variable_name"`
  Value json.Number `json:"Value,num_integer"`
}

type Event struct {
  Duration string `json:"Duration"`
  Queued int  `json:"No. Events Queued"`
  Executed int  `json:"No. Events Executed"`
}

func maxscaleUrl(path string) string {
  return "http://" + os.Getenv("MAXSCALE_MAXINFO_JSON_LISTENER_TCP_ADDR") + path
}

func listenAddr() string {
  return os.Getenv("MAXSCALE_EXPORTER_LISTEN_ADDR")
}

func servers(w http.ResponseWriter, r *http.Request) {
  url := maxscaleUrl("/servers")
  resp, err := http.Get(url)

  defer resp.Body.Close()

  if err != nil {
    panic(err)
  }

  jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  var jsonData []Server

  err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)

  if err != nil {
    panic(err)
  }

  // fmt.Println(jsonData)
  fmt.Fprintf(w, "#TYPE maxscale_servers_connections gauge\n")
  fmt.Fprintf(w, "#HELP maxscale_servers_connections Server connections\n")
  for _,element := range jsonData {
    fmt.Fprintf(w, "maxscale_servers_connections{server=\"%s\", address=\"%s\", port=\"%d\"} %d\n", element.Server, element.Address, element.Port, element.Connections)
  }
  fmt.Fprintf(w, "\n")
}

func services(w http.ResponseWriter, r *http.Request) {
  url := maxscaleUrl("/services")
  resp, err := http.Get(url)

  defer resp.Body.Close()

  if err != nil {
    panic(err)
  }

  jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  var jsonData []Service

  err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)

  if err != nil {
    panic(err)
  }
  fmt.Fprintf(w, "#TYPE maxscale_services_sessions gauge\n")
  fmt.Fprintf(w, "#HELP maxscale_services_sessions Service Sessions\n")
  // fmt.Println(jsonData)
  for _,element := range jsonData {

    fmt.Fprintf(w, "maxscale_services_sessions{name=\"%s\", router=\"%s\"} %d\n", element.Name, element.Router, element.Sessions)


  }
  fmt.Fprintf(w, "\n")
}

func status(w http.ResponseWriter, r *http.Request) {
  url := maxscaleUrl("/status")
  resp, err := http.Get(url)

  defer resp.Body.Close()

  if err != nil {
    panic(err)
  }

  jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  var jsonData []Status

  err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)

  if err != nil {
    panic(err)
  }

  // fmt.Println(jsonData)

  for _,element := range jsonData {

    switch element.Name {
      case "Uptime",
      "Uptime_since_flush_status",
      "Threads_created",
      "Threadpool_threads",
      "Read_events",
      "Write_events",
      "Hangup_events",
      "Error_events",
      "Accept_events":
      fmt.Fprintf(w, "#TYPE maxscale_status_%s counter\n", strings.ToLower(element.Name))
      fmt.Fprintf(w, "#HELP maxscale_status_%s %s\n", strings.ToLower(element.Name), element.Name)
      fmt.Fprintf(w, "maxscale_status_%s %d\n\n", strings.ToLower(element.Name), element.Value)
      case "Threads_running",
      "Threads_connected",
      "Connections",
      "Client_connections",
      "Backend_connections",
      "Listeners",
      "Zombie_connections",
      "Internal_descriptors",
      "Event_queue_length",
      "Pending_events",
      "Max_event_queue_length",
      "Max_event_queue_time",
      "Max_event_execution_time":
      fmt.Fprintf(w, "#TYPE maxscale_status_%s gauge\n", strings.ToLower(element.Name))
      fmt.Fprintf(w, "#HELP maxscale_status_%s %s\n", strings.ToLower(element.Name), element.Name)
      fmt.Fprintf(w, "maxscale_status_%s %d\n\n", strings.ToLower(element.Name), element.Value)

    }
  }
}

func variables(w http.ResponseWriter, r *http.Request) {
  url := maxscaleUrl("/variables")
  resp, err := http.Get(url)

  defer resp.Body.Close()

  if err != nil {
    panic(err)
  }

  jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  var jsonData []Variable

  err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)

  if err != nil {
    panic(err)
  }

  // fmt.Println(jsonData)

  for _,element := range jsonData {

    switch element.Name {
    case "MAXSCALE_UPTIME":
      fmt.Fprintf(w, "#TYPE maxscale_variables_%s counter\n", strings.ToLower(element.Name))
      fmt.Fprintf(w, "#HELP maxscale_variables_%s %s\n", strings.ToLower(element.Name), element.Name)
      fmt.Fprintf(w, "maxscale_variables_%s %s\n\n", strings.ToLower(element.Name), element.Value)
      case "MAXSCALE_SESSIONS",
      "MAXSCALE_POLLSLEEP",
      "MAXSCALE_NBPOLLS",
      "MAXSCALE_THREADS":
      fmt.Fprintf(w, "#TYPE maxscale_variables_%s gauge\n", strings.ToLower(element.Name))
      fmt.Fprintf(w, "#HELP maxscale_variables_%s %s\n", strings.ToLower(element.Name), element.Name)
      fmt.Fprintf(w, "maxscale_variables_%s %s\n\n", strings.ToLower(element.Name), element.Value)

    }
  }
}

func events(w http.ResponseWriter, r *http.Request) {
  url := maxscaleUrl("/event/times")
  resp, err := http.Get(url)

  defer resp.Body.Close()

  if err != nil {
    panic(err)
  }

  jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  var jsonData []Event

  err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)

  if err != nil {
    panic(err)
  }
  fmt.Fprintf(w, "#TYPE maxscale_events_executed_seconds histogram\n")
  fmt.Fprintf(w, "#HELP maxscale_events_executed_seconds Events Executed\n")

  executedsum := float64(0)
  executedcount := 0
  executedtime := 0.1
  for _,element := range jsonData {
    executedcount += element.Executed
    executedsum = executedsum + (float64(element.Executed) * executedtime)
    executedtime += 0.1
    switch element.Duration {
    case "< 100ms":
      fmt.Fprintf(w, "maxscale_events_executed_seconds_bucket{le=\"0.100000\"} %d\n", executedcount)
    case "> 3000ms":
      fmt.Fprintf(w, "maxscale_events_executed_seconds_bucket{le=\"+Inf\"} %d\n", executedcount)
    default:
      durationf := strings.Split(element.Duration, " ")
      ad := strings.Trim(durationf[len(durationf)-1], "ms")
      duurr, _ := strconv.ParseFloat(ad, 64)
      hurrr := duurr / 1000
      fmt.Fprintf(w, "maxscale_events_executed_seconds_bucket{le=\"%f\"} %d\n", hurrr, executedcount)
    }
  }
  fmt.Fprintf(w, "maxscale_events_executed_seconds_sum %d\n", int(executedsum))
  fmt.Fprintf(w, "maxscale_events_executed_seconds_count %d\n\n", executedcount)

  fmt.Fprintf(w, "#TYPE maxscale_events_queued_seconds histogram\n")
  fmt.Fprintf(w, "#HELP maxscale_events_queued_seconds Events Queued\n")

  queuedsum := float64(0)
  queuedcount := 0
  queuedtime := 0.1
  for _,element := range jsonData {
    queuedcount += element.Queued
    queuedsum = queuedsum + (float64(element.Queued) * queuedtime)
    queuedtime += 0.1
    switch element.Duration {
    case "< 100ms":
      fmt.Fprintf(w, "maxscale_events_queued_seconds_bucket{le=\"0.100000\"} %d\n", queuedcount)
    case "> 3000ms":
      fmt.Fprintf(w, "maxscale_events_queued_seconds_bucket{le=\"+Inf\"} %d\n", queuedcount)
    default:
      durationf := strings.Split(element.Duration, " ")
      ad := strings.Trim(durationf[len(durationf)-1], "ms")
      duurr, _ := strconv.ParseFloat(ad, 64)
      hurrr := duurr / 1000
      fmt.Fprintf(w, "maxscale_events_queued_seconds_bucket{le=\"%f\"} %d\n", hurrr, queuedcount)
    }
  }

  fmt.Fprintf(w, "maxscale_events_queued_seconds_sum %d\n", int(queuedsum))
  fmt.Fprintf(w, "maxscale_events_queued_seconds_count %d\n", queuedcount)
}

func metrics(w http.ResponseWriter, r *http.Request) {
  t := time.Now()
  r.ParseForm()
  fmt.Printf("[%s] remote=%s protocol=%s useragent=%s contentlength=%d host=%s request_url=%s", t.Format(time.RFC3339), r.RemoteAddr, r.Proto, r.Header["User-Agent"][0], r.ContentLength, r.Host, r.RequestURI)
  fmt.Println(r.Form["url_long"])
  for k, v := range r.Form {
    fmt.Println("key:", k)
    fmt.Println("val:", strings.Join(v, ""))
  }
  servers(w, r)
  services(w, r)
  status(w, r)
  variables(w, r)
  events(w, r)
}



func main() {
  http.HandleFunc("/metrics", metrics)
  err := http.ListenAndServe(listenAddr(), nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
