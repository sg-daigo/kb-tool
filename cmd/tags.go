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

// tagsCmd represents the tags command
var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Get all tags",
	Long:  `Get all tags.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		request := kanboard.NewRequest[struct{}]("getAllTags", nil)
		result, err := kanboard.SendRequestRaw[struct{}](request, opts.KbConf, opts.Logger)
		if err != nil {
			return fmt.Errorf("send request failed: %w", err)
		}

		if opts.ToJson {
			fmt.Println(string(result.Result))
		} else {
			var tags []kanboard.AllTagsResponse
			err = json.Unmarshal(result.Result, &tags)
			if err != nil {
				return fmt.Errorf("failed to unmarshal tags: %w", err)
			}

			slices.SortFunc(tags, func(a, b kanboard.AllTagsResponse) int {
				return int(a.ID) - int(b.ID)
			})

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			_, _ = fmt.Fprintln(w, "ID\tProject\tName")
			for _, tag := range tags {
				_, _ = fmt.Fprintf(w, "%d\t%d\t%s\n", tag.ID, tag.ProjectID, tag.Name)
			}
			_ = w.Flush()
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(tagsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tagsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tagsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
