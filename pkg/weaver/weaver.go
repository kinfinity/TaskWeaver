package weaver

import (
	"io"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	cli_runtime "taskweaver/pkg/weaver/cli-runtime"
	weaverCommand "taskweaver/pkg/weaver/cmd"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/*
Weaver Application

-	initialize client runtime
*/
var weaverInstance *weaverProcess
var once sync.Once

type weaverProcess struct {
	logger *log.Logger
}

func GetWeaverInstance() *weaverProcess {
	weaverInstance = &weaverProcess{
		logger: log.New(os.Stdout, "[WEAVER]: ", 0),
	}
	return weaverInstance
}

// get weaver logger
func (wp *weaverProcess) Logger() *log.Logger { return wp.logger }

func (wp *weaverProcess) Run() (err error) {
	/*
		Handle Root Command & Configs
	*/
	var (
		configFileName string = "/config.yaml"
		configFilePath string
	)

	var rootCmd = &cobra.Command{
		Use:   "weaver",
		Short: "Task Weaver Client - Weaver",
		Run: func(cmd *cobra.Command, args []string) {
			/*
				Initialize Runtime & Handle Commands
			*/
			initWeaver(wp.logger, cmd, configFilePath, configFileName)
		},
	}
	// Set up a persistent flag for the configFilePath
	rootCmd.PersistentFlags().StringVarP(&configFilePath, "config", "c", "", "config file (default is ./config.yaml)")

	// Bind the viper configuration to the configFilePath flag1
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	// Execute the Cobra command
	if err := rootCmd.Execute(); err != nil {
		wp.logger.Panic("fatal error config file: ", err.Error())
		os.Exit(1)
	}
	return
}

func initWeaver(logger *log.Logger, cmd *cobra.Command, configFilePath string, configFileName string) (err error) {
	// Setup Client Config
	configDir := setupConfig(logger, configFileName)

	viper.SetConfigType("yaml")
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(configDir)
	logger.Println("Default Config = ", configDir+configFileName)

	// Check config via flag
	if cmd.PersistentFlags().Changed("config") {
		// Configure Viper from --config
		viper.AddConfigPath(configFilePath) // --config overwrites  config in the .weaver directory
		logger.Println("Overwrite Config File =", configFilePath+configFileName)
		if _, err := os.Stat(configFilePath + configFileName); os.IsNotExist(err) {
			logger.Println("New Config File Not Found: ", err)
		}
	}

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		logger.Panic("fatal error config file:", err.Error())
	}
	logger.Println("Log File =", viper.GetString("weaver_config.log-file"))

	/*
		- initialize client
	*/
	cli_options := cli_runtime.ClientOptions{
		ConfigFile: "", // Override if specified with --config flag
		Name:       viper.GetString("weaver_config.client-name"),
		Version:    "",
		LogFile:    viper.GetString("weaver_config.log-file"),
	}
	// Open or create a log file
	logFile, err := os.OpenFile(cli_options.LogFile, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	// Create a logger
	logger = log.New(
		io.MultiWriter(
			os.Stdout,
			cli_runtime.NewFileWriter(logFile),
		),
		"["+strings.ToUpper(cli_options.Name)+"]: ",
		0,
	)
	cli_options.SetLogger(logger)

	client, err := cli_runtime.NewCli(&cli_options)
	if err != nil {
		cli_runtime.HandleError(err)
	}

	/*
		- call default weaver command
		- if no arguments, show help page
	*/
	command := weaverCommand.NewWeaverCommand()
	if err := client.RunCommand(command); err != nil {
		// Handle error
	}
	return
}

func setupConfig(logger *log.Logger, configName string) (configDir string) {
	// Check config in home
	homeDir, _ := os.UserHomeDir()
	configDir = path.Join(homeDir, ".weaver/config")
	configFilePath := path.Join(configDir, "/", configName)
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// Initialize default config
		logger.Printf("Creating new config at %s\n", configFilePath)
		err := cli_runtime.InitializeConfig(logger, configDir, configFilePath)
		if err != nil {
			logger.Fatalln(err)
		}
	} else if err != nil {
		panic(err)
	}
	return
}
