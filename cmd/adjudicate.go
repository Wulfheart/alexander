package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/wulfheart/godip-influence/defaultInfluences"
	"github.com/zond/godip/variants"
	"wulfheartalexander/common"

	"github.com/spf13/cobra"
)

var (
	data string
)
// adjudicateCmd represents the adjudicate command
var adjudicateCmd = &cobra.Command{
	Use:   "adjudicate",
	Short: "Adjudicate an existing game",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		v, found := variants.Variants[variant]
		if !found {
			panic(fmt.Errorf("variant %q not found", variant))
		}
		p := common.RequestDTO{}
		if err := json.Unmarshal([]byte(data), &p); err != nil {
			panic(err)
		}
		s := p.State(v)
		if err := s.Next(); err != nil {
			panic(err)
		}
		res, err := json.Marshal(common.NewResponseDTOfromState(s, defaultInfluences.ConvertToInfluence(defaultInfluences.Classical), v))
		if err != nil {
			panic(err)
		}
		fmt.Println(string(res))
	},
}

func init() {
	rootCmd.AddCommand(adjudicateCmd)
	adjudicateCmd.Flags().StringVar(&data, "data", "", "data from previous phases (required)")
	adjudicateCmd.MarkFlagRequired("data")
	adjudicateCmd.Flags().StringVarP(&variant, "variant", "v", "", "variant to be used (required)")
	adjudicateCmd.MarkFlagRequired("variant")
}
