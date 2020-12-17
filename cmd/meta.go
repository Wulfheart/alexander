/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zond/godip/variants"
	"wulfheartalexander/common"
	"wulfheartalexander/logging"
)

// metaCmd represents the meta command
var metaCmd = &cobra.Command{
	Use:   "meta",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		v, found := variants.Variants[variant]
		if !found {
			logging.Logger.Error(fmt.Sprintf("variant %q not found", variant))
		}
		res, err := json.Marshal(common.CreateMetaDtoFromVariant(v))
		if err != nil {
			logging.Logger.Error(err.Error())
		}
		fmt.Println(string(res))


	},
}

func init() {
	rootCmd.AddCommand(metaCmd)

	metaCmd.Flags().StringVarP(&variant, "variant", "v", "", "variant to be used (required)")
	metaCmd.MarkFlagRequired("variant")
}
