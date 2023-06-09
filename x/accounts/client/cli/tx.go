package cli

import (
	"accounts/utils"
	"accounts/x/accounts/keeper"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	accountstypes "accounts/x/accounts/types"
)

const (
	FundsFlagName = "funds"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(schemas map[string]*keeper.Schema) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        accountstypes.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", accountstypes.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(GetDeployCmd(schemas), GetExecuteCmd(schemas))

	return cmd
}

func GetDeployCmd(schemas map[string]*keeper.Schema) *cobra.Command {

	cmd := &cobra.Command{
		Use:  "deploy [account-type] [init-msg-json]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			accountType := args[0]
			initMsgJSON := args[1]

			accountSchema, exists := schemas[accountType]
			if !exists {
				return fmt.Errorf("unkown account type %s, %#v", accountType, schemas)
			}
			msg, err := accountSchema.InitMsg.UnmarshalJSONString(initMsgJSON)
			if err != nil {
				return fmt.Errorf("encoding error: %s", err)
			}

			anyMsg, err := utils.MarshalAnyBytes(msg)
			if err != nil {
				return err
			}

			funds, err := maybeFunds(cmd.Flags())
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &accountstypes.MsgDeploy{
				Sender:      clientCtx.FromAddress.String(),
				Kind:        accountType,
				InitMessage: anyMsg,
				Funds:       funds,
			})
		},
	}

	cmd.Flags().String(FundsFlagName, "", "optional funds to send in deploy and execute [Coins string]")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func GetExecuteCmd(schemas map[string]*keeper.Schema) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "execute [contract-address] [type] [json]",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			accTypeResp, err := accountstypes.NewQueryClient(clientCtx).AccountKind(cmd.Context(), &accountstypes.QueryAccountKindRequest{Address: args[0]})
			if err != nil {
				return err
			}

			msgType := args[1]
			msgJSON := args[2]

			accountSchema, exists := schemas[accTypeResp.Kind]
			if !exists {
				return fmt.Errorf("unkown account type %s, got: %#v", accTypeResp.Kind, schemas)
			}

			msgSchema, exists := accountSchema.ExecuteMsgs[msgType]
			if !exists {
				return fmt.Errorf("unknown execute msg for account type %s", msgType)
			}

			protoMsg, err := msgSchema.UnmarshalJSONString(msgJSON)
			if err != nil {
				return fmt.Errorf("message construction: %w", err)
			}

			funds, err := maybeFunds(cmd.Flags())
			if err != nil {
				return err
			}

			anyMsg, err := utils.MarshalAnyBytes(protoMsg)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &accountstypes.MsgExecute{
				Sender:  clientCtx.FromAddress.String(),
				Address: args[0],
				Message: anyMsg,
				Funds:   funds,
			})
		},
	}

	cmd.Flags().String(FundsFlagName, "", "optional funds to send in deploy and execute [Coins string]")
	flags.AddTxFlagsToCmd(cmd)
	return cmd

}

func maybeFunds(flags *pflag.FlagSet) (sdk.Coins, error) {
	coinsStr, err := flags.GetString(FundsFlagName)
	if err != nil {
		return nil, err
	}

	if coinsStr == "" {
		return nil, nil
	}

	return sdk.ParseCoinsNormalized(coinsStr)
}
