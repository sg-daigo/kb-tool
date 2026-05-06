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

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Get all projects",
	Long:  `Get all projects.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		request := kanboard.NewRequest[struct{}]("getAllProjects", nil)
		res, err := kanboard.SendRequestRaw[struct{}](request, opts.KbConf, opts.Logger)
		if err != nil {
			return fmt.Errorf("send request failed: %w", err)
		}

		if opts.ToJson == true {
			fmt.Println(string(res.Result))
		} else {
			var projects []kanboard.ProjectResult
			err := json.Unmarshal(res.Result, &projects)
			if err != nil {
				return fmt.Errorf("json unmarshal faild: %w", err)
			}

			slices.SortFunc(projects, func(a, b kanboard.ProjectResult) int {
				return int(a.ID) - int(b.ID)
			})

			// 表示
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tName")
			for _, project := range projects {
				fmt.Fprintf(w, "%d\t%s\n", project.ID, project.Name)
			}
			w.Flush()
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
