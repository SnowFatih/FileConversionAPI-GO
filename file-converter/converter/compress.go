package converter

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CompressToZIP(files []string, outputPath string) error {
	fmt.Println("📦 Début de la compression :", outputPath)

	for _, filePath := range files {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Println("❌ Erreur : Fichier introuvable :", filePath)
			return err
		}
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("❌ Erreur : Impossible de créer le fichier ZIP :", err)
		return err
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	for _, filePath := range files {
		fmt.Println("➕ Ajout du fichier :", filePath)

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("❌ Erreur : Impossible d'ouvrir le fichier", filePath, err)
			return err
		}

		info, err := file.Stat()
		if err != nil {
			fmt.Println("❌ Erreur : Impossible de lire les informations du fichier", err)
			file.Close()
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			fmt.Println("❌ Erreur : Impossible de créer un header ZIP", err)
			file.Close()
			return err
		}

		header.Name = filepath.Base(filePath)
		header.Method = zip.Deflate 

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			fmt.Println("❌ Erreur : Impossible d'ajouter le fichier dans l'archive", err)
			file.Close()
			return err
		}

		_, err = io.Copy(writer, file)
		if err != nil {
			fmt.Println("❌ Erreur : Problème lors de la copie du fichier dans l'archive", err)
			file.Close()
			return err
		}

		fmt.Println("✅ Fichier ajouté :", filePath)
		file.Close()
	}

	fmt.Println("✅ Compression terminée :", outputPath)
	return nil
}
