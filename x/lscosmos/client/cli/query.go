package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/persistenceOne/pstake-native/v2/x/lscosmos/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group lscosmos queries under a subcommand
	cmd := &cobra.Command{
		Use:                        queryRoute,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdQueryAllState(),
		CmdQueryParams(),
		CmdQueryHostChainParams(),
		CmdQueryDelegationState(),
		CmdQueryAllowListedValidators(),
		CmdQueryCValue(),
		CmdQueryModuleState(),
		CmdQueryIBCTransientStore(),
		CmdQueryUnclaimed(),
		CmdQueryFailedUnbondings(),
		CmdQueryPendingUnbondings(),
		CmdQueryUnbondingEpoch(),
		CmdQueryHostAccountUndelegation(),
		CmdQueryDelegatorUnbodingEpochEntry(),
		CmdQueryHostAccounts(),
		CmdQueryDepositModuleAccount(),
		CmdDelegatorUnbondingEpochEntries(),
	)

	return cmd
}

// CmdQueryHostChainParams implements the host chain params query command
func CmdQueryHostChainParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "host-chain-params",
		Short: "shows host chain parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.HostChainParams(context.Background(), &types.QueryHostChainParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryAllState implements the all state query command
func CmdQueryAllState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-state",
		Short: "shows genesis state",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AllState(context.Background(), &types.QueryAllStateRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryDelegationState implments the delegation state query command
func CmdQueryDelegationState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegation-state",
		Short: "shows delegation state",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.DelegationState(context.Background(), &types.QueryDelegationStateRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryAllowListedValidators implements allow listed validators query command.
func CmdQueryAllowListedValidators() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allow-listed-validators",
		Short: "shows allow listed validators",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AllowListedValidators(context.Background(), &types.QueryAllowListedValidatorsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryCValue implements the c value query command
func CmdQueryCValue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "c-value",
		Short: "shows current c-value of the protocol",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.CValue(context.Background(), &types.QueryCValueRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryModuleState implements the module state query command
func CmdQueryModuleState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "module-state",
		Short: "shows current module state",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ModuleState(context.Background(), &types.QueryModuleStateRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryIBCTransientStore implements the IBC transient store query command
func CmdQueryIBCTransientStore() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ibc-transient-store",
		Short: "shows amount in ibc-transient-store",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.IBCTransientStore(context.Background(), &types.QueryIBCTransientStoreRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryUnclaimed implements the unclaimed unbondings query command
func CmdQueryUnclaimed() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unclaimed [delegator-address]",
		Args:  cobra.ExactArgs(1),
		Short: "shows unclaimed amounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			delegatorAddress, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.Unclaimed(context.Background(), &types.QueryUnclaimedRequest{DelegatorAddress: delegatorAddress.String()})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryFailedUnbondings implements the failed unbondings query command
func CmdQueryFailedUnbondings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "failed-unbondings [delegator-address]",
		Args:  cobra.ExactArgs(1),
		Short: "shows failed unbondings request",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			_, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.FailedUnbondings(context.Background(), &types.QueryFailedUnbondingsRequest{DelegatorAddress: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryPendingUnbondings implements the pending unbondings query command
func CmdQueryPendingUnbondings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pending-unbondings [delegator-address]",
		Args:  cobra.ExactArgs(1),
		Short: "shows pending unbondings",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			_, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.PendingUnbondings(context.Background(), &types.QueryPendingUnbondingsRequest{DelegatorAddress: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryUnbondingEpoch implements the unbonding epoch query command
func CmdQueryUnbondingEpoch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unbonding-epoch [epoch-number]",
		Args:  cobra.ExactArgs(1),
		Short: "Shows unbonding epoch details",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			epochNumber, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}

			res, err := queryClient.UnbondingEpochCValue(context.Background(), &types.QueryUnbondingEpochCValueRequest{EpochNumber: epochNumber})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHostAccountUndelegation implements the host account undelegation query command
func CmdQueryHostAccountUndelegation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "host-account-undelegation [epoch-number]",
		Args:  cobra.ExactArgs(1),
		Short: "Shows host account undelegation for the given epoch number",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			epochNumber, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}

			res, err := queryClient.HostAccountUndelegation(context.Background(), &types.QueryHostAccountUndelegationRequest{EpochNumber: epochNumber})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryDelegatorUnbodingEpochEntry implements the delegation unbonding epoch entry query command
func CmdQueryDelegatorUnbodingEpochEntry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegator-unbonding-epoch-entry [delegator-address] [epoch-number]",
		Args:  cobra.ExactArgs(2),
		Short: "Shows host account undelegation for the given epoch number",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			delegatorAddress, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			epochNumber, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			res, err := queryClient.DelegatorUnbondingEpochEntry(context.Background(), &types.QueryDelegatorUnbondingEpochEntryRequest{DelegatorAddress: delegatorAddress.String(), EpochNumber: epochNumber})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHostAccounts implements the host accounts query command
func CmdQueryHostAccounts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "host-accounts",
		Short: "shows the host accounts (delegation and reward) ica owner id",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.HostAccounts(context.Background(), &types.QueryHostAccountsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryDepositModuleAccount implements the deposit module account query command
func CmdQueryDepositModuleAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit-module-account",
		Short: "shows the balance of deposit module account",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.DepositModuleAccount(context.Background(), &types.QueryDepositModuleAccountRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdDelegatorUnbondingEpochEntries implements the delegator unbonding epoch entries query command
func CmdDelegatorUnbondingEpochEntries() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegator-unbonding-epoch-entries [delegator-address]",
		Args:  cobra.ExactArgs(1),
		Short: "Shows host account undelegations for all the epoch number the entry has been made.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			delegatorAddress, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.DelegatorUnbondingEpochEntries(context.Background(), &types.QueryAllDelegatorUnbondingEpochEntriesRequest{DelegatorAddress: delegatorAddress.String()})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
