package cli

import (
	"accounts/x/accounts/keeper"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/tx"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	accountstypes "accounts/x/accounts/types"
)

const (
	FundsFlagName = "flags"
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
				return fmt.Errorf("unkown account type %s", accountType)
			}
			msg, err := accountSchema.InitMsg.EncodeFromJSONStringToProto(initMsgJSON)
			if err != nil {
				return err
			}

			anyMsg, err := codectypes.NewAnyWithValue(msg)
			if err != nil {
				return err
			}

			funds, err := maybeFunds(cmd.Flags())
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &accountstypes.MsgDeploy{
				Sender:      clientCtx.From,
				Kind:        accountType,
				InitMessage: anyMsg,
				Funds:       funds,
			})
		},
	}

	cmd.Flags().String(FundsFlagName, "", "optional funds to send in deploy and execute [Coins string]")
	return cmd
}

func GetExecuteCmd(schemas map[string]*keeper.Schema) *cobra.Command {
	/*
		cmd := &cobra.Command{
			Use: "execute [contract-address] [type] [json]",
			Args: cobra.ExactArgs(3),
			RunE: func(cmd *cobra.Command, args []string) error {
				clientCtx, err := client.GetClientTxContext(cmd)
				if err != nil {
					return err
				}

				msgType := args[1]
				msgJSON := args[2]

				accountSchema, exists := schemas[accountType]
				if !exists {
					return fmt.Errorf("unkown account type %s", accountType)
				}

				funds, err := maybeFunds(cmd.Flags())
				if err != nil {
					return err
				}

				return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &accountstypes.MsgExecute{
					Sender:  clientCtx.From,
					Address: "",
					Message: nil,
					Funds:   nil,
				})
			},
		}

		cmd.Flags().String(FundsFlagName, "", "optional funds to send in deploy and execute [Coins string]")
	*/
	panic("impl")
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
