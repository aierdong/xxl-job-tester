package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"strings"
)

type XxlJobRequest struct {
	JobId          int    `json:"jobId"`
	JobHandler     string `json:"executorHandler"`
	Params         string `json:"executorParams"`
	BlockStrategy  string `json:"executorBlockStrategy"`
	Timeout        int    `json:"executorTimeout"`
	LogId          int    `json:"logId"`
	LogDateTime    int    `json:"logDateTime"`
	GlueType       string `json:"glueType"`
	BroadcastIndex int    `json:"broadcastIndex"`
	BroadcastTotal int    `json:"broadcastTotal"`
}

var (
	protocol string
	host     string
	port     int
)

func main() {
	request := XxlJobRequest{}
	var rootCmd = &cobra.Command{
		Use:   "xxl-job-tester",
		Short: "A client to simulate XXL-Job Admin call",
		Long:  `This tool allows you to simulate XXL-Job Admin calls to your executor.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return callXxlJobExecutor(&request)
		},
	}

	rootCmd.Flags().StringVarP(&protocol, "protocol", "", "http", "The protocol of your executor (http or https)")
	rootCmd.Flags().StringVarP(&host, "host", "", "localhost", "The host of your executor")
	rootCmd.Flags().IntVarP(&port, "port", "p", 8080, "The port of your executor")
	rootCmd.Flags().StringVarP(&request.JobHandler, "job_handler", "j", "", "The job handler")
	rootCmd.Flags().StringVarP(&request.Params, "params", "P", "", "The parameters")
	rootCmd.Flags().StringVarP(&request.BlockStrategy, "block_strategy", "b", "SERIAL_EXECUTION", "The block strategy")
	rootCmd.Flags().IntVarP(&request.Timeout, "timeout", "", 0, "The timeout, seconds")
	rootCmd.Flags().IntVarP(&request.LogId, "log_id", "", 1, "The log id")
	rootCmd.Flags().IntVarP(&request.LogDateTime, "log_date_time", "", 0, "The log date time")
	rootCmd.Flags().StringVarP(&request.GlueType, "glue_type", "", "BEAN", "The glue type")
	rootCmd.Flags().IntVarP(&request.BroadcastIndex, "broadcast_index", "", 0, "The broadcast index")
	rootCmd.Flags().IntVarP(&request.BroadcastTotal, "broadcast_total", "", 1, "The broadcast total")

	if err := rootCmd.MarkFlagRequired("job_handler"); err != nil {
		log.Fatal(err)
	}

	protocol = strings.ToLower(protocol)
	if protocol != "http" && protocol != "https" {
		log.Fatal("Invalid protocol")
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func callXxlJobExecutor(x *XxlJobRequest) error {
	url := fmt.Sprintf("%s://%s:%d/run", protocol, host, port)

	jsonBody, err := json.Marshal(x)
	if err != nil {
		fmt.Println("Error marshaling request:", err)
		return err
	}

	fmtJson, _ := json.MarshalIndent(x, "", "  ")
	fmt.Println("URL:", url)
	fmt.Println("Request:", string(fmtJson))
	fmt.Println("================================")

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}

	defer func() { _ = resp.Body.Close() }()

	fmt.Println("Response:")
	fmt.Println("Status  :", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}
	fmt.Println("Body    :", string(body))

	return nil
}
