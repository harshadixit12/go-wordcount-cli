/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"go-wordcount-cli/lib"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var warnColor = "\033[33m"
var errorColor = "\033[31m"
var linesShort,linesLong, bytesShort, bytesLong, charsShort,charsLong string = "-l", "--lines","-c","--bytes","-m","--chars";

var validArgs = [6] string {linesLong, linesShort, bytesLong, bytesShort, charsLong, charsShort}


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-wordcount-cli",
	Short: "print newline, word, and byte counts for each file",
	Long: `This is an attempt to implement wc tool from linux in go, as a learning exercise. The CLI will count the number of new linesm words and bytes in each 
	file given as input. Supported arguments: -l and --lines for lines, -c and --bytes for bytes, -m and --chars for characters.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var lines, _ = cmd.Flags().GetBool("lines");
		var bytes, _ = cmd.Flags().GetBool("bytes");
		var chars, _ = cmd.Flags().GetBool("chars");
		var words, _ = cmd.Flags().GetBool("words");

		var allFlags = lines && bytes && chars && words || (!lines && !bytes && !chars && !words);

		if (len(args) < 1) {
			fmt.Println(warnColor + "Please give a file as input.")
			return
		}

		data, err := os.ReadFile(args[0])

		if (err != nil) {
			fmt.Println(errorColor + "Invalid input:", err)
			return
		}

		if (len(args) > 1 && !util.ArrayContains(validArgs, args[1])) {
			fmt.Println(errorColor + "Invalid argument: ", args[1])
			return
		}

		var solution = "";

		if (bytes || allFlags) {
			solution += strconv.Itoa(len(data)) + " Bytes"
		}

		if (chars || allFlags) {
			solution += "\n" + strconv.Itoa(len(data)) + " Chars"
		}

		if (lines || allFlags) {
			solution += "\n" + strconv.Itoa(strings.Count(string(data), "\n")) +" Lines"
		}

		if (words || allFlags) {
			solution += "\n" + strconv.Itoa(len(strings.Fields(string(data)))) +" Words"
		}

		fmt.Println(solution)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-wordcount-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolP("lines", "l", false, "Number of lines in the input file.")
	rootCmd.PersistentFlags().BoolP("bytes", "c", false, "Number of bytes in the input file.")
	rootCmd.PersistentFlags().BoolP("chars", "m", false, "Number of chars in the input file.")
	rootCmd.PersistentFlags().BoolP("words", "w", false, "Number of words in the input file.")
}


