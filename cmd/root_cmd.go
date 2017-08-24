package main

import (
	"fmt"
	"os"

	"github.com/aditya-prasad/grpc-health"
	"github.com/spf13/cobra"
)

var (
	ServerURL   string
	ServiceName string
	Timeout     int
	Verbose     bool
)

var RootCmd = &cobra.Command{
	Use:   "grpc-health -u <server_url> -n <fully_qualified_service_name>",
	Short: "CLI tool to health check GRPC services",
	Long: `
GRPC Health Checker
-----------------------
CLI tool to health check GRPC services. The service is expected to implement the Health method as follows:

service SomeService {
  rpc Health (google.protobuf.Empty)  returns (google.protobuf.Empty) {}
  // other stuff
}
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if ServerURL == "" {
			return fmt.Errorf("server url not provided")
		}
		if ServiceName == "" {
			return fmt.Errorf("service name not provided")
		}
		grpcHealthChecker := grpc_health.NewGrpcHealthChecker()
		isHealthy := grpcHealthChecker.IsHealthy(ServerURL, ServiceName, Timeout, Verbose)
		if !isHealthy {
			os.Exit(1)
		}
		return nil
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().StringVarP(&ServerURL, "serverUrl", "u", "", "Server URL")
	RootCmd.Flags().StringVarP(&ServiceName, "serviceName", "n", "", "Fully qualified service name")
	RootCmd.Flags().IntVarP(&Timeout, "timeout", "t", 1, "Timeout for connecting in seconds")
	RootCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "Print logs")
}
