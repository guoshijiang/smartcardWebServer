package cmd

import (
	"github.com/spf13/cobra"
	"smartcardWebServer/pkg/cli"
)

var RootCmd = &cobra.Command{
	Use: "Smartcard App Service",
	Short: "Smartcard App cmd",
	Long: `Smartcard App Service is a tool to manage golang service`,
	Example:`
		smartcardWebServer 
	`,
}

func init(){
	RootCmd.AddCommand(cli.SmartCardDataCmd("smartcard"))
}