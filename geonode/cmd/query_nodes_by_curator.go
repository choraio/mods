package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/choraio/mods/geonode/types/v1"
)

// QueryNodesByCuratorCmd creates and returns the query nodes-by-curator command.
func QueryNodesByCuratorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodes-by-curator [curator]",
		Short: "query nodes by curator",
		Long:  "query nodes by curator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := getQueryClient(cmd)
			if err != nil {
				return err
			}

			req := v1.QueryNodesByCuratorRequest{
				Curator: args[0],
			}

			res, err := c.NodesByCurator(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
