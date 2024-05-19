package handlers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger-labs/ccapi/bucket"
	"github.com/hyperledger-labs/ccapi/chaincode"
)

const (
	TEMP_DIR = "/tmp"
)

type CreateExamRequest struct {
	PatientHash      string                `form:"patientHash" binding:"required"`
	DoctorHash       string                `form:"doctorHash" binding:"required"`
	Timestamp        string                `form:"timestamp" binding:"required"`
	Name             string                `form:"name" binding:"required"`
	UrlExamDocument  *multipart.FileHeader `form:"urlExamDocument" binding:"required"`
	DiagnosisHash    string                `form:"diagnosisHash" binding:"required"`
	TreatmentHash    string                `form:"treatmentHash" binding:"required"`
}

func CreateExam(c *gin.Context) {
	var req CreateExamRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("failed to bind form: %w", err).Error(),
		})
		return
	}

	// Ensure that the file is present and has a valid filename
	if req.UrlExamDocument == nil || req.UrlExamDocument.Filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "UrlExamDocument file is required and must have a valid filename",
		})
		return
	}

	timestamp, err := time.Parse("2006-01-02T15:04:05.000Z", req.Timestamp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("failed to parse timestamp: %w", err).Error(),
		})
		return
	}

	// Ensure the directory exists
	if _, err := os.Stat(TEMP_DIR); os.IsNotExist(err) {
		if err := os.MkdirAll(TEMP_DIR, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Errorf("failed to create temp directory: %w", err).Error(),
			})
			return
		}
	}

	// Save UrlExamDocument file locally
	examDocPath := TEMP_DIR + "/" + req.UrlExamDocument.Filename
	if err := c.SaveUploadedFile(req.UrlExamDocument, examDocPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("failed to save UrlExamDocument file: %w", err).Error(),
		})
		return
	}
	defer os.Remove(examDocPath)

	// Upload UrlExamDocument file to Min.io
	minioClient, err := bucket.CreateClient(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("failed to create minio client: %w", err).Error(),
		})
		return
	}

	saveLocationExamDocument, err := bucket.UploadFile(c, minioClient, examDocPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("failed to upload UrlExamDocument file: %w", err).Error(),
		})
		return
	}

	// Hash received UrlExamDocument file
	openedFileExamDocument, err := req.UrlExamDocument.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("failed to open UrlExamDocument file: %w", err).Error(),
		})
		return
	}
	defer openedFileExamDocument.Close()

	fileBytesExamDocument, err := io.ReadAll(openedFileExamDocument)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("failed to read UrlExamDocument file: %w", err).Error(),
		})
		return
	}

	fileHashExamDocument := sha256.Sum256(fileBytesExamDocument)
	fileHashStringExamDocument := fmt.Sprintf("%x", fileHashExamDocument)

	// Create new asset and put in ledger
	txArgs := map[string]interface{}{
		"patientHash":      req.PatientHash,
		"doctorHash":       req.DoctorHash,
		"timestamp":        timestamp.Format("2006-01-02T15:04:05.000Z"),
		"name":             req.Name,
		"urlExamDocument":  saveLocationExamDocument,
		"diagnosisHash":    req.DiagnosisHash,
		"treatmentHash":    req.TreatmentHash,
		"fileHashStringExamDocument": fileHashStringExamDocument,
	}

	txArgsBytes, err := json.Marshal(txArgs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("failed to marshal tx args: %w", err).Error(),
		})
		return
	}

	argList := [][]byte{txArgsBytes}

	channelResponse, _, err := chaincode.Invoke(os.Getenv("CHANNEL"), os.Getenv("CCNAME"), "createExam", argList, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("failed to invoke chaincode: %w", err).Error(),
		})
		return
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(channelResponse.Payload, &payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("failed to unmarshal payload: %w", err).Error(),
		})
		return
	}

	c.JSON(http.StatusOK, payload)
}
