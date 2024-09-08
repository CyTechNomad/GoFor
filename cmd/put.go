package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var putData string
var putContentType string

var putCmd = &cobra.Command{
	Use:   "put [url]",
	Short: "Sends a PUT request to the specified URL with a request body",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		makePutRequest(url, putData, putContentType)
	},
}

func init() {
	RootCmd.AddCommand(putCmd)

	putCmd.Flags().StringVarP(&putData, "data", "d", "", "Data to include in the PUT request body")
	putCmd.Flags().StringVarP(&putContentType, "content-type", "c", "application/json", "Content-Type header for the PUT request")
}

func makePutRequest(url, data, contentType string) {
	client := &http.Client{
		Timeout: 30 * time.Second, // Adjust the timeout as needed
	}

	body := bytes.NewBuffer([]byte(data))

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		fmt.Printf("Error creating PUT request: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", contentType)

	for _, header := range headers {
		keyValue := splitHeader(header)
		if keyValue != nil {
			req.Header.Add(keyValue[0], keyValue[1])
		}
	}

	ApplyAuthentication(req)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making PUT request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}
	fmt.Println("Response Body:")
	fmt.Println(string(bodyResp))
}
