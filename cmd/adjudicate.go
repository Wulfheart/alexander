package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/wulfheart/godip-influence/defaultInfluences"
	"github.com/zond/godip/variants"
	"go.uber.org/zap"
	"wulfheartalexander/common"
	"wulfheartalexander/logging"

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
		logging.Logger.Debug("Adjudication command called")
		v, found := variants.Variants[variant]
		if !found {
			logging.Logger.Error("Variant not found", zap.String("variant", variant))
		}
		logging.Logger.Debug("Variant found", zap.String("variant", variant))
		p := common.RequestDTO{}
		if err := json.Unmarshal([]byte(data), &p); err != nil {
			logging.Logger.Error("Error unmarshalling input data")
		}
		logging.Logger.Debug("Input data unmarshalled successfully")
		s := p.State(v)
		logging.Logger.Debug("Phase of game state created")
		if err := s.Next(); err != nil {
			logging.Logger.Error(err.Error())
		}
		logging.Logger.Debug("Phase of game adjudicated")
		res, err := json.Marshal(common.NewResponseDTOfromState(s, defaultInfluences.ConvertToInfluence(defaultInfluences.Classical), v))
		if err != nil {
			logging.Logger.Error(err.Error())
		}
		logging.Logger.Debug("Response JSON created")
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
