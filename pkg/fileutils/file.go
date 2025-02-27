package fileutils

import (
	"bufio"
	"os"
	"path/filepath"
)

// ReadLinesInRange lit un fichier dans le répertoire "data" et retourne les lignes lues.
func ReadLinesInRange(fileName string) ([]string, error) {
	// Construire le chemin complet du fichier
	filePath := filepath.Join("data", fileName)

	// Ouvrir le fichier
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Lire les lignes du fichier
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Vérifier s'il y a eu une erreur de lecture pendant le scan
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Retourner les lignes lues
	return lines, nil
}
