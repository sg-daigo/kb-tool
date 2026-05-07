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

var ProjectID int

// projectcolumnsCmd represents the projectcolumns command
var projectcolumnsCmd = &cobra.Command{
	Use:   "projectcolumns",
	Short: "Get project columns",
	Long:  `Get project columns`,
	RunE: func(cmd *cobra.Command, args []string) error {
		request := kanboard.NewColumnsRequest(ProjectID)
		res, err := kanboard.SendRequestRaw[kanboard.ColumnsParams](request, opts.KbConf, opts.Logger)
		if err != nil {
			return fmt.Errorf("send request failed: %w", err)
		}

		if opts.ToJson {
			fmt.Println(string(res.Result))
		} else {
			var result []kanboard.ColumnsResult
			err = json.Unmarshal(res.Result, &result)
			if err != nil {
				return fmt.Errorf("failed to unmarshal result: %w", err)
			}

			slices.SortFunc(result, func(a, b kanboard.ColumnsResult) int {
				return int(a.ID) - int(b.ID)
			})

			// 表示
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			Println(w, "ID\tPosition\tTitle")
			for _, column := range result {
				Printf(w, "%d\t%s\t%s\n", column.ID, column.Position, column.Title)
			}
			_ = w.Flush()
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(projectcolumnsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectcolumnsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectcolumnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	projectcolumnsCmd.Flags().IntVarP(&ProjectID, "project", "p", 0, "Project ID")
}
