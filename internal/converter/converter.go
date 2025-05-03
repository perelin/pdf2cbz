package converter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const tempDirName = ".pdf2cbz_tmp_files"

// ConvertPDFToCBZ converts a PDF file to CBZ format
func ConvertPDFToCBZ(pdfPath string) error {
	// Validate PDF file
	if err := validatePDF(pdfPath); err != nil {
		return fmt.Errorf("validatePDF: %w", err)
	}

	// Create temporary directory in the same directory as the input file
	tempDir := filepath.Join(filepath.Dir(pdfPath), tempDirName)
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return fmt.Errorf("createTempDir: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Extract images
	fmt.Println("Extracting pages as images...")
	if err := extractImages(pdfPath, tempDir); err != nil {
		return fmt.Errorf("extractImages: %w", err)
	}

	// Create CBZ file
	cbzPath := strings.TrimSuffix(pdfPath, filepath.Ext(pdfPath)) + ".cbz"
	fmt.Println("Creating CBZ archive...")
	if err := createCBZ(tempDir, cbzPath); err != nil {
		return fmt.Errorf("createCBZ: %w", err)
	}

	return nil
}

func validatePDF(pdfPath string) error {
	file, err := os.Open(pdfPath)
	if err != nil {
		return fmt.Errorf("failed to open PDF file: %w", err)
	}
	defer file.Close()

	// Basic validation - check if file exists and is readable
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	if info.Size() == 0 {
		return fmt.Errorf("PDF file is empty")
	}

	return nil
}

func getBaseFilename(path string) string {
	base := filepath.Base(path)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

func extractImages(pdfPath, tempDir string) error {
	cmd := exec.Command("pdftoppm",
		"-png",
		"-f", "1",
		"-progress",
		pdfPath,
		filepath.Join(tempDir, "page"))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pdftoppm failed: %w", err)
	}

	return nil
}

func createCBZ(tempDir, cbzPath string) error {
	// Create a zip file of the temporary directory
	cmd := exec.Command("zip", "-r", cbzPath, tempDir)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("zip failed: %w", err)
	}

	return nil
}
