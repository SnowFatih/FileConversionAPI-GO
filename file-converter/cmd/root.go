package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "file-converter",
	Short: "Outil CLI pour convertir et compresser des fichiers",
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Erreur :", err)
		os.Exit(1)
	}
}
