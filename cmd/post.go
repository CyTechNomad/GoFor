package cmd

import (
    "fmt"
    "io"
    "net/http"
    "strings"
    "time"
   
    "github.com/spf13/cobra"
)

var postData string
var contentType string

var postCmd = &cobra.Command{
    Use: "post [url]",
    Short: "Sends a POST request to the specified URL with a request body",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        url := args[0]
        makePostRequest(url, postData, contentType)
    },
}

func init() {
    postCmd.Flags().StringVarP(&postData, "data", "d", "", "Data to send in the request body")
    postCmd.Flags().StringVarP(&contentType, "content-type", "c", "text/plain", "Content type of the request body")
    RootCmd.AddCommand(postCmd)
}

func makePostRequest(url, data, contentType string) {
    client := &http.Client{
        Timeout: time.Duration(timeout) * time.Second,
    }

    req, err := http.NewRequest(http.MethodPost, url, nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    req.Header.Add("Content-Type", contentType)
    req.Body = io.NopCloser(strings.NewReader(data))

    for _, header := range headers {
        keyValue := splitHeader(header)
        if keyValue != nil {
            req.Header.Add(keyValue[0], keyValue[1])
        }
    }

    ApplyAuthentication(req)

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error making POST request:", err)
        return
    }
    defer resp.Body.Close()

    fmt.Println("Status Code:", resp.StatusCode)

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    fmt.Println("Response Body:", string(body))
}
