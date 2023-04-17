package main

import (
	"easy-password-generator/pkg/generator"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "password-generator",
	Short: "A password generator CLI tool",
	Long:  `A CLI tool to generate secure, random passwords based on user-specified criteria.`,
	Run: func(cmd *cobra.Command, args []string) {
		password, err := generator.GeneratePassword(length, excludeUpper, excludeLower, excludeNumbers, excludeSpecial)
		if err != nil {
			fmt.Println("Error generating password:", err)
			return
		}
		fmt.Println("Generated password:    ", password)
	},
}
var (
	length                                                     int
	excludeUpper, excludeLower, excludeNumbers, excludeSpecial bool
)

func init() {
	rootCmd.Flags().IntVarP(&length, "length", "l", 10, "length of the password")
	rootCmd.Flags().BoolVarP(&excludeUpper, "upper", "u", false, "exclude uppercase letters")
	rootCmd.Flags().BoolVarP(&excludeLower, "lower", "L", false, "exclude lowercase letters")
	rootCmd.Flags().BoolVarP(&excludeNumbers, "numbers", "n", false, "exclude numbers")
	rootCmd.Flags().BoolVarP(&excludeSpecial, "special", "s", false, "exclude special characters")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
