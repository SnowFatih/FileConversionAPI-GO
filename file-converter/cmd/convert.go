package cmd

import (
	"fmt"
	"file-converter/converter"
	"file-converter/utils"
	"strings"

	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert [inputFile] [outputFile]",
	Short: "Convertir un fichier d'un format à un autre",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFile := args[1]
		var err error

		fmt.Println("📂 Conversion en cours :", inputFile, "→", outputFile)

		inputExt := strings.ToLower(utils.GetFileExtension(inputFile))
		outputExt := strings.ToLower(utils.GetFileExtension(outputFile))
		switch {
		// CSV → JSON
		case inputExt == ".csv" && outputExt == ".json":
			outputFile, err = converter.CSVToJSON(inputFile)

		// JSON → CSV
		case inputExt == ".json" && outputExt == ".csv":
			outputFile, err = converter.JSONToCSV(inputFile)

		// PNG → JPEG
		case inputExt == ".png" && outputExt == ".jpg":
			outputFile, err = converter.ConvertPNGToJPEG(inputFile)

		// JPEG → PNG
		case (inputExt == ".jpg" || inputExt == ".jpeg") && outputExt == ".png":
			outputFile, err = converter.ConvertJPEGToPNG(inputFile)

		// Compression ZIP
		case outputExt == ".zip":
			err = converter.CompressToZIP([]string{inputFile}, outputFile)

		default:
			fmt.Println("❌ Format non supporté :", inputExt, "→", outputExt)
			return
		}

		if err != nil {
			fmt.Println("❌ Erreur lors de la conversion :", err)
		} else {
			fmt.Println("✅ Conversion réussie :", outputFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
