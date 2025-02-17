package api

import (
	"file-converter/converter"
	"file-converter/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
)

func ConvertTextHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fichier non trouvé"})
		return
	}

	tempPath := "temp/" + file.Filename
	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Échec de la sauvegarde"})
		return
	}

	ext := utils.GetFileExtension(file.Filename)
	var outputData []byte
	var outputFile string

	// Conversion CSV → JSON
	if ext == ".csv" {
		outputData, err = converter.CSVToJSON(tempPath)
		outputFile = utils.GenerateNewFileName(tempPath, ".json")
	} else if ext == ".json" {
		outputData, err = converter.JSONToCSV(tempPath)
		outputFile = utils.GenerateNewFileName(tempPath, ".csv")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format non supporté"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur de conversion"})
		return
	}

	err = utils.WriteFile(outputFile, outputData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'enregistrement"})
		return
	}

	c.File(outputFile)
}

func ConvertImageHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fichier non trouvé"})
		return
	}

	tempPath := "temp/" + file.Filename
	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Échec de la sauvegarde"})
		return
	}

	ext := utils.GetFileExtension(file.Filename)
	var outputFile string

	if ext == ".png" {
		outputFile, err = converter.ConvertPNGToJPEG(tempPath)
	} else if ext == ".jpg" || ext == ".jpeg" {
		outputFile, err = converter.ConvertJPEGToPNG(tempPath)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format non supporté"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur de conversion"})
		return
	}

	c.File(outputFile)
}


func CompressHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fichier non trouvé"})
		return
	}

	tempPath := "temp/" + file.Filename
	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Échec de la sauvegarde"})
		return
	}

	info, err := os.Stat(tempPath)
	if err != nil {
		fmt.Println("❌ Erreur : Fichier temporaire introuvable")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur interne"})
		return
	}
	fmt.Println("📂 Fichier à compresser :", tempPath, "| Taille :", info.Size(), "octets")

	outputFile := "temp/" + utils.GenerateNewFileName(file.Filename, ".zip")

	err = converter.CompressToZIP([]string{tempPath}, outputFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur de compression"})
		return
	}

	info, err = os.Stat(outputFile)
	if err != nil {
		fmt.Println("❌ Erreur : Fichier ZIP introuvable")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur interne"})
		return
	}
	fmt.Println("✅ Fichier ZIP créé :", outputFile, "| Taille :", info.Size(), "octets")

	c.File(outputFile)
}
