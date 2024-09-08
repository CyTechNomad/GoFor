package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type AuthType string
const (
    AuthBearer AuthType = "bearer"
    AuthBasic  AuthType = "basic"
)

type Auth struct {
	Type     AuthType
	Token    string
	Username string
	Password string
}

var auth Auth
var headers []string
var timeout int

// RootCmd is the base command for GoFor CLI
var RootCmd = &cobra.Command{
	Use:   "gofor",
	Short: "GoFor is a simple HTTP client written in Go",
	Long: `GoFor is a lightweight and efficient HTTP client.
It allows you to send HTTP requests like GET, POST, PUT, and DELETE.`,
	// This is the function that runs if no subcommand is provided
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to GoFor! Use --help to see available commands.")
	},
}

// Execute runs the root command and all its subcommands
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you can define global flags and configuration settings
	RootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.gofor.yaml)")

	RootCmd.PersistentFlags().StringArrayVarP(&headers, "header", "H", []string{}, "HTTP headers to include in the request")
	RootCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 30, "Timeout for the request in seconds")
	RootCmd.PersistentFlags().StringVarP(&auth.Token, "token", "T", "", "Bearer token for authentication")
}

func splitHeader(header string) []string {
	split := strings.SplitN(header, ":", 2)
	if len(split) != 2 {
		fmt.Println("Invalid header:", header)
		return nil
	}
	return split
}

func ApplyAuthentication(req *http.Request) {
	if auth.Token != "" {
		req.Header.Add("Authorization", "Bearer "+auth.Token)
	} else if auth.Username != "" && auth.Password != "" {
		req.SetBasicAuth(auth.Username, auth.Password)
	}
}
