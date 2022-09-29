package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "personal-finance",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:
			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		RunRest()
	},
}

func init() {
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(makeMigrationCmd)
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.env.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("env")
	}

	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var migrateUpCmd = &cobra.Command{
	Use:   "migrate:up",
	Short: "Migrate up migration file",
	Long:  `Migrate up migration file on folder migration/sql`,
	Run: func(cmd *cobra.Command, args []string) {
		mgr := WiringMigration()
		if err := mgr.Up(); err != nil {
			fmt.Printf("Migrate up error: %v\n", err.Error())
		}
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "migrate:down",
	Short: "Migrate down migration file",
	Long:  `Migrate down migration file on folder migration/sql`,
	Run: func(cmd *cobra.Command, args []string) {
		mgr := WiringMigration()
		if err := mgr.Down(); err != nil {
			fmt.Printf("Migrate down error: %s", err.Error())
		}
	},
}

var makeMigrationCmd = &cobra.Command{
	Use:   "make:migration [name] [ext]",
	Short: "Make migration file",
	Long:  `Make migration file on folder migration/sql`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		mgr := WiringMigration()
		_ = mgr.Create(args[0], args[1])
	},
}
