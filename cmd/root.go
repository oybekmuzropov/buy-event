package cmd

import (
	"fmt"
	"github.com/buy_event/config"
	"github.com/buy_event/service"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"log"
	"os"

	_ "github.com/lib/pq"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "buy_event",
	Short: "Task for creating command",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var userService *service.UserService
var purchaseService *service.PurchaseService
var logService *service.LogService
var cfg config.Config

func init() {
	cfg = config.Load()

	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)
	db, err := sqlx.Connect("postgres", psqlString)

	if err != nil {
		log.Printf("error while connecting database: %v", err)
		return
	}
	userService = service.NewUserService(db)
	purchaseService = service.NewPurchaseService(db)
	logService = service.NewLogService(db)

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.buy_event.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".buy_event")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
