package cmd

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
    Use: "delete [url]",
    Short: "Sends a DELETE request to the specified URL",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        url := args[0]
        makeDeleteRequest(url)
    },
}

func init() {
    RootCmd.AddCommand(deleteCmd)
}

func makeDeleteRequest(url string) {
    client := &http.Client{
        Timeout: time.Duration(timeout) * time.Second,
    }

    req, err := http.NewRequest(http.MethodDelete, url, nil)
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
        fmt.Println("Error making DELETE request:", err)
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
