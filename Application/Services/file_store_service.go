package Services

import (
	"io"
	"os"
	"path/filepath"
)

// FileStorageService implements file storage operations
type FileStorageService struct {
	ImageUploadPath string
	FileStoreFolder string
}

// NewFileStorageService creates a new instance of FileStorageService
func NewFileStorageService() *FileStorageService {
	return &FileStorageService{
		ImageUploadPath: filepath.Join(".", "ImageStorage", "images"), // Set your desired upload path
		FileStoreFolder: "ImageStorage",
	}
}

// GetFileUrl returns the URL of a stored file
func (fs *FileStorageService) GetFileUrl(fileName string) string {
	return filepath.Join("/", fs.FileStoreFolder, "images", fileName)
}

// SaveFile saves a file to the specified path
func (fs *FileStorageService) SaveFile(mediaBinaryStream io.Reader, fileName string) error {
	filePath := filepath.Join(fs.ImageUploadPath, fileName)
	output, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, mediaBinaryStream)
	return err
}

// DeleteFile deletes a file from storage
func (fs *FileStorageService) DeleteFile(fileName string) error {
	filePath := filepath.Join(fs.ImageUploadPath, fileName)
	err := os.Remove(filePath)
	return err
}

// HandleFileUpload handles file upload using Gin-Gonic
//func HandleFileUpload(c *gin.Context) {
//	file, err := c.FormFile("file")
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
//		return
//	}
//
//	// Initialize File Storage Service
//	fileStorageService := NewFileStorageService()
//
//	// Save file to storage
//	if err := fileStorageService.SaveFile(file, file.Filename); err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
//		return
//	}
//
//	// Return file URL
//	fileURL := fileStorageService.GetFileUrl(file.Filename)
//	c.JSON(http.StatusOK, gin.H{"fileUrl": fileURL})
//}
