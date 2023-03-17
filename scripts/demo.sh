
BENEFICIARY=$(accountsd keys show beneficiary -a)

accountsd tx accounts deploy simple-vesting "{\"beneficiary\":\"$BENEFICIARY\",\"startAfter\":\"10s\",\"duration\":\"100s\"}" --from funder --funds 100000000utest --broadcast-mode=block

accountsd tx accounts execute cosmos1yu6jl57lqju7qllneuf225qawsd3g7dvdqcqye examples.vesting.v1.MsgWithdrawUnlockedCoins '{}' --from beneficiary --broadcast-mode=block