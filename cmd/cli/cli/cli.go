package cli

import (
	"github.com/blushft/strana/app"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "strana",
	Short: "Your way analytics",
	RunE:  runServer,
}

func init() {
	flgs := rootCmd.PersistentFlags()

	flgs.StringP("config", "c", "./config.yaml", "Config file")
	flgs.Bool("debug", false, "Enable debugging")
	flgs.StringP("server.host", "l", "", "Server listener host")
	flgs.StringP("server.port", "p", "8863", "Server listening port")

	viper.BindPFlags(flgs)

	cobra.OnInitialize(setupConfig)

	rootCmd.AddCommand(configCommand())
}

func setupConfig() {
	viper.SetDefault("config", "./config.yaml")

	viper.SetEnvPrefix("strana")
	viper.AutomaticEnv()
}

func configure() (*config.Config, error) {
	logger.Log().Info("getting configuration")

	v := viper.GetViper()
	cf := viper.GetString("config")

	viper.SetConfigFile(cf)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	conf, err := config.NewConfig(v)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func Execute() error {
	return rootCmd.Execute()
}

func runServer(cmd *cobra.Command, args []string) error {
	logger.Log().Info("starting strana...")
	conf, err := configure()
	if err != nil {
		return err
	}

	a, err := app.New(*conf)
	if err != nil {
		return err
	}

	return a.Start()
}
