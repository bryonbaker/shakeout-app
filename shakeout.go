package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
    "io/ioutil"
    "path/filepath"
)

const logFilePath = "/data/access-log.log"

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Set the content type to HTML
    w.Header().Set("Content-Type", "text/html")

    // Write the response
    fmt.Fprintln(w, "<html><body><h1>Hello request received.</h1>")

    // Log the date and time of the request
    logRequest()

    // Display the contents of the access log
    displayLog(w)

    fmt.Fprintln(w, "</body></html>")
}

func logRequest() {
    // Create the /data directory if it doesn't exist
    err := os.MkdirAll(filepath.Dir(logFilePath), os.ModePerm)
    if err != nil {
        fmt.Printf("Error creating log directory: %s\n", err)
        return
    }

    // Open the log file in append mode, create it if it doesn't exist
    file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Printf("Error opening log file: %s\n", err)
        return
    }
    defer file.Close()

    // Get the pod name from the environment variable
    podName := os.Getenv("POD_NAME")
    if podName == "" {
        podName = "unknown"
    }

    // Write the current date and time to the log file
    logEntry := fmt.Sprintf("Request received at: %s. Processed by pod: %s\n", time.Now().Format("2006-01-02T15:04:05.000Z07:00"), podName)
    if _, err := file.WriteString(logEntry); err != nil {
        fmt.Printf("Error writing to log file: %s\n", err)
    }
}

func displayLog(w http.ResponseWriter) {
    // Read the log file
    data, err := ioutil.ReadFile(logFilePath)
    if err != nil {
        if os.IsNotExist(err) {
            // If the log file does not exist, display a message
            fmt.Fprintln(w, "<p>No log entries yet.</p>")
        } else {
            // If there is an error reading the log file, display an error message
            fmt.Fprintf(w, "<p>Error reading log file: %s</p>", err)
        }
        return
    }

    // Display the contents of the log file
    fmt.Fprintln(w, "<pre>"+string(data)+"</pre>")
}

func main() {
    http.HandleFunc("/hello", helloHandler)

    // Listen on port 9000
    err := http.ListenAndServe(":9000", nil)
    if err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}
