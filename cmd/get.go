// cmd/get.go
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// getCmd represents the 'get' command
var getCmd = &cobra.Command{
    Use:   "get [url]",
    Short: "Sends a GET request to the specified URL",
    Long: `Sends a GET request to the specified URL and returns the response.

Example:
  gofor get https://jsonplaceholder.typicode.com/todos/1`,
    Args: cobra.MinimumNArgs(1), // Requires at least 1 argument (the URL)
    Run: func(cmd *cobra.Command, args []string) {
        url := args[0]
        makeGetRequest(url)
    },
}

func init() {
    // Register the 'get' subcommand with the root command
    RootCmd.AddCommand(getCmd)
}

// makeGetRequest sends a GET request to the specified URL
func makeGetRequest(url string) {
    client := &http.Client{
        Timeout: time.Duration(timeout) * time.Second,
    }
    
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    for _, header := range headers {
        keyValue := splitHeader(header)
        if keyValue != nil {
            req.Header.Add(keyValue[0], keyValue[1])
        }
    }

    ApplyAuthentication(req)

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error making Get request:", err)
        return
    }
    defer resp.Body.Close()

    fmt.Println("Status Code:", resp.StatusCode)

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    fmt.Println(string(body))
}



