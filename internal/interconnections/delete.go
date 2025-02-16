package interconnections

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var connectionID string

	deleteConnectionCmd := &cobra.Command{
		Use:   `delete -i <connection_id> `,
		Short: "Deletes a interconnection.",
		Long:  "Deletes the specified interconnection.",
		Example: `  # Deletes the specified interconnection:
  metal interconnections delete -i 7ec86e23-8dcf-48ed-bd9b-c25c20958277`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			_, _, err := c.Service.DeleteInterconnection(context.Background(), connectionID).Execute()
			if err != nil {
				return err
			}
			fmt.Println("connection deletion initiated. Please check 'metal interconnections get -i", connectionID, "' for status")

			return nil
		},
	}

	deleteConnectionCmd.Flags().StringVarP(&connectionID, "id", "i", "", "The UUID of the interconnection.")
	_ = deleteConnectionCmd.MarkFlagRequired("id")

	return deleteConnectionCmd
}
