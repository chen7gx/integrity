package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/chen7gx/integrity/x/integrity/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

func CmdCreateHash() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-hash [details] [hash]",
		Short: "Broadcast message createHash",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDetails := args[0]
			argHash := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateHash(
				clientCtx.GetFromAddress().String(),
				argDetails,
				argHash,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
