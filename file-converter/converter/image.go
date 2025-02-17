package converter

import (
	"file-converter/utils"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func checkImageReadable(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("erreur ouverture du fichier : %v", err)
	}
	defer file.Close()

	_, format, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("erreur décodage image : %v", err)
	}
	fmt.Println("✅ Image lue avec succès. Format détecté :", format)
	return nil
}

func ConvertPNGToJPEG(inputPath string) (string, error) {
	fmt.Println("🔍 Vérification de l'image PNG :", inputPath)

	err := checkImageReadable(inputPath)
	if err != nil {
		fmt.Println("❌ Erreur lecture PNG :", err)
		return "", err
	}

	outputPath := utils.GenerateNewFileName(inputPath, ".jpg")
	fmt.Println("📂 Chemin de sortie :", outputPath)

	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("❌ Erreur : Impossible d'ouvrir l'image PNG :", err)
		return "", err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("❌ Erreur : Impossible de décoder l'image PNG :", err)
		return "", err
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("❌ Erreur : Impossible de créer le fichier JPEG :", err)
		return "", err
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 90})
	if err != nil {
		fmt.Println("❌ Erreur : Échec d'encodage JPEG :", err)
		return "", err
	}

	fmt.Println("✅ Conversion réussie :", outputPath)
	return outputPath, nil
}

func ConvertJPEGToPNG(inputPath string) (string, error) {
	fmt.Println("🔍 Vérification de l'image JPEG :", inputPath)

	err := checkImageReadable(inputPath)
	if err != nil {
		fmt.Println("❌ Erreur lecture JPEG :", err)
		return "", err
	}

	outputPath := utils.GenerateNewFileName(inputPath, ".png")
	fmt.Println("📂 Chemin de sortie :", outputPath)

	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("❌ Erreur : Impossible d'ouvrir l'image JPEG :", err)
		return "", err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("❌ Erreur : Impossible de décoder l'image JPEG :", err)
		return "", err
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("❌ Erreur : Impossible de créer le fichier PNG :", err)
		return "", err
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	if err != nil {
		fmt.Println("❌ Erreur : Échec d'encodage PNG :", err)
		return "", err
	}

	fmt.Println("✅ Conversion réussie :", outputPath)
	return outputPath, nil
}
