package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/choraio/mods/content/types/v1"
)

// QueryContentsCmd creates and returns the query contents command.
func QueryContentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contents",
		Short: "query all contents",
		Long:  "query all contents",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := getQueryClient(cmd)
			if err != nil {
				return err
			}

			req := v1.QueryContentsRequest{}

			res, err := c.Contents(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
