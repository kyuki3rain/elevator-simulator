/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/kyuki3rain/elevator-simulator/field"
	"github.com/spf13/cobra"
)

type Options struct {
	endTime           int
	rate              float64
	floorNumber       int
	elevatorNumber    int
	elevatorMaxPeople int
	sleep             int
}

var (
	o = &Options{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "elevator-simulator",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		f := field.New(1, o.endTime, o.rate, o.floorNumber, o.elevatorNumber, o.elevatorMaxPeople, o.sleep)

		f.Loop(o.sleep)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.elevator-simulator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().IntVarP(&o.endTime, "time", "t", 100, "endTime option")
	rootCmd.Flags().Float64VarP(&o.rate, "rate", "r", 1.0, "createHumanRate option")
	rootCmd.Flags().IntVarP(&o.floorNumber, "floor", "f", 10, "floorNumber option")
	rootCmd.Flags().IntVarP(&o.elevatorNumber, "elev", "e", 5, "elevatorNumber option")
	rootCmd.Flags().IntVarP(&o.elevatorMaxPeople, "max", "m", 10, "elevatorMaxPeople option")
	rootCmd.Flags().IntVarP(&o.sleep, "sleep", "s", 100, "sleep option")
}
