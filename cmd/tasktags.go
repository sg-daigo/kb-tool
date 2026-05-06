/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"text/tabwriter"

	"github.com/sg-daigo/kanboard-go"
	"github.com/spf13/cobra"
)

var taskID int

// tasktagsCmd represents the tasktag command
var tasktagsCmd = &cobra.Command{
	Use:   "tasktags",
	Short: "Get task tags",
	Long:  `Get task tags`,
	RunE: func(cmd *cobra.Command, args []string) error {
		request := kanboard.NewTaskTagsRequest(taskID)
		res, err := kanboard.SendRequestRaw[kanboard.TaskTagsParams](request, opts.KbConf, opts.Logger)
		if err != nil {
			return fmt.Errorf("send request failed: %w", err)
		}

		if opts.ToJson {
			fmt.Println(string(res.Result))
		} else {
			var result kanboard.TaskTagsResult
			err = json.Unmarshal(res.Result, &result)
			if err != nil {
				return fmt.Errorf("failed to unmarshal result: %w", err)
			}

			// tagID を取り出してソートする
			var keys []string
			for key := range result {
				keys = append(keys, key)
			}
			slices.Sort(keys)

			// 表示
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tName")
			for _, key := range keys {
				fmt.Fprintf(w, "%s\t%s\n", key, result[key])
			}
			w.Flush()
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(tasktagsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tasktagCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tasktagCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	tasktagsCmd.Flags().IntVarP(&taskID, "task", "t", 0, "Task ID")

}
