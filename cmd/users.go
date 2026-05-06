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

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Get all users",
	Long:  `Get all users.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		request := kanboard.NewRequest[struct{}]("getAllUsers", nil)
		req, err := kanboard.SendRequestRaw[struct{}](request, opts.KbConf, opts.Logger)
		if err != nil {
			return fmt.Errorf("send request failed: %w", err)
		}

		if opts.ToJson {
			fmt.Println(string(req.Result))
		} else {
			var users []kanboard.UserResult
			err := json.Unmarshal(req.Result, &users)
			if err != nil {
				return fmt.Errorf("failed to unmarshal result: %w", err)
			}
			slices.SortFunc(users, func(a, b kanboard.UserResult) int {
				return int(a.ID) - int(b.ID)
			})

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			Println(w, "ID\tName")
			for _, user := range users {
				Printf(w, "%d\t%s\n", user.ID, user.Username)
			}
			_ = w.Flush()
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(usersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
