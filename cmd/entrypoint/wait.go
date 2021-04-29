package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"time"
)
var (
	timeoutSeconds       int
	requestTimeoutMillis int
	periodMillis         int
	url                  string

	waitCommand = &cobra.Command{
		Use:   "wait",
		Short: "wait unit sidecar ready",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := &http.Client{
				Timeout: time.Duration(requestTimeoutMillis) * time.Millisecond,
			}
			timeoutAt := time.Now().Add(time.Duration(timeoutSeconds) * time.Second)

			var err error
			for time.Now().Before(timeoutAt) {
				err = checkIfReady(client, url)
				if err == nil {
					log.Println("sidecar is ready")
					return nil
				}
				log.Println("sidecar is not ready")
				time.Sleep(time.Duration(periodMillis) * time.Millisecond)
			}
			return fmt.Errorf("sidecar is not ready in %d second(s)", timeoutSeconds)
		},
	}
)

func checkIfReady(client *http.Client, url string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != 200 {
		return fmt.Errorf("http status code not 200")
	}
	return nil
}

func Execute() error{
	return waitCommand.Execute()
}

func init() {
	waitCommand.PersistentFlags().IntVar(&timeoutSeconds, "timeoutSeconds", 60, "maximum number of seconds to wait for sidecar to be ready")
	waitCommand.PersistentFlags().IntVar(&periodMillis, "periodMillis", 500, "number of milliseconds to wait between attempts")
	waitCommand.PersistentFlags().IntVar(&requestTimeoutMillis, "requestTimeoutMillis", 500, "number of milliseconds to wait for response")
	waitCommand.PersistentFlags().StringVar(&url, "url", "http://localhost:8080/ready", "URL to use in requests")
}
