package cmd

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/sg-daigo/kanboard-go"
	"github.com/spf13/cobra"
)

type GlobalOption struct {
	KbConf  kanboard.ServerConfig
	Logger  *slog.Logger
	IsDebug bool
	ToJson  bool
}

var opts GlobalOption

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kb-tools",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		level := new(slog.LevelVar)
		if opts.IsDebug {
			level.Set(slog.LevelDebug)
		} else {
			level.Set(slog.LevelInfo)
		}
		handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
		opts.Logger = slog.New(handler)
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kb-tools.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	opts.KbConf.Token = os.Getenv("KB_TOKEN")
	rootCmd.PersistentFlags().StringVarP(&opts.KbConf.Server, "server", "s", "http://localhost", "Server url")
	rootCmd.PersistentFlags().BoolVarP(&opts.IsDebug, "debug", "d", false, "Toggle debug mode")
	rootCmd.PersistentFlags().BoolVarP(&opts.ToJson, "json", "j", false, "Show JSON")
}

func Println(w io.Writer, a ...any) {
	_, _ = fmt.Fprintln(w, a...)
}

func Printf(w io.Writer, format string, a ...any) {
	_, _ = fmt.Fprintf(w, format, a...)
}
