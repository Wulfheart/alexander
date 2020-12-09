package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/wulfheart/godip-influence/defaultInfluences"
	"github.com/zond/godip/variants"
	"wulfheartalexander/common"

	"github.com/spf13/cobra"
)


// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new game",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		v, found := variants.Variants[variant]
		if !found {
			panic(fmt.Errorf("variant %q not found", variant))
		}
		s, err := v.Start()
		if err != nil {
			panic(err)
			return
		}
		res, err := json.Marshal(common.NewResponseDTOfromState(s, defaultInfluences.ConvertToInfluence(defaultInfluences.Classical), v))
		if err != nil {
			panic(err)
		}
		fmt.Println(string(res))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&variant, "variant", "v", "", "variant to be used (required)")
	initCmd.MarkFlagRequired("variant")
}
