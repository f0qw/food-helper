package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A brief description of my app",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, Cobra!")
	},
}

var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "A brief description of sub command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, Sub Command!")
	},
}

var nestedCmd = &cobra.Command{
	Use:   "nested",
	Short: "A brief description of nested command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, Nested Command!")
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
	subCmd.AddCommand(nestedCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
