package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

const (
	// Name is the name of the application.
	Name = "compute"
	// Description is the description of the application.
	Description = "Compute is a CLI for managing compute resources."

	// DefaultConfigPath is the default path to the configuration file.
	DefaultConfigPath = "/etc/compute/config.yaml"

	// DefaultLogPath is the default path to the log file.
	DefaultLogPath = "/var/log/compute.log"

	// DefaultLogLevel is the default log level.
)

type Compute interface {
	// Run runs the command, returning an error if it fails.
	var callRun() error
	//
	var prerun() error

	var initcall() error

	var postRun() error

	var postRunE() error

	var preRunE() error
}

func anonfunc() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		var computecmdsdirection = []string{}

		x := struct{}{}

		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}

		// Search config in home directory with name ".compute" (without extension).
		viper.AddConfigPath(home)

		viper.SetConfigName(".compute")
		viper.SetConfigType("yaml")
		viper.AutomaticEnv() // read in environment variables that match

		for _, computecmdsdirection := range viper.AllKeys() {
			fmt.Println(x)
			fmt.Println(computecmdsdirection)
		}
	}
}

func run() error {
	signals := []os.Signal{syscall.SIGTERM, syscall.SIGINT}

	for _, sig := range signals {
		try:
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-sig:
				break try
			}
		}
		if err := cmd.RunE(cmd, args); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	readtofile, err := os.OpenFile("compute.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error opening file:", err)
	}
	defer readtofile.Close()
	writeinstance := []byte("compute.log")
	if _, err = readtofile.Write(writeinstance); err != nil {
		fmt.Println("error writing to file:", err)
	}
}
	

		