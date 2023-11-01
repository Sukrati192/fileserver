package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sukrati192/fileserver/contracts"
	"github.com/sukrati192/fileserver/services"
)

// GetMap godoc
// @Summary API to get the uploaded file
// @Description Checks if a file exists in IPFS storage with given id, then stream the file
// @Tags public
// @Produce octet-stream
// @Param fileID path string true "ID of the file"
// @Success 200 {} FileStream
// @Failure 500 {object} contracts.ErrorResponse
// @Router /file/{fileID} [get]
func GetFile(ctx echo.Context) error {
	cid := ctx.Param("fileID")
	fileName := fmt.Sprintf("%s.txt", cid)
	if err := services.IpfsShell.Get(cid, fileName); err != nil {
		fmt.Fprintf(os.Stderr, "error while getting file %s: %v\n", fileName, err)
		return ctx.JSON(http.StatusInternalServerError, contracts.ErrorResponse{Error: "Could not get file"})
	}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while opening file: %s\n", err)
		return ctx.JSON(http.StatusInternalServerError, contracts.ErrorResponse{Error: "Could not open file"})
	}
	defer os.Remove(fileName)
	var fileReader io.Reader = file
	return ctx.Stream(http.StatusOK, echo.MIMEOctetStream, fileReader)
}

// GetMap godoc
// @Summary API to upload a file
// @Description Uploads the file in IPFS storage and returns id, size and timestamp in response
// @Tags public
// @Accept mpfd
// @Produce json
// @Param file formData file true "File to be uploaded"
// @Success 200 {object} contracts.File
// @Failure 500 {object} contracts.ErrorResponse
// @Router /upload [post]
func UploadFile(ctx echo.Context) error {
	currTime := time.Now()
	file, fileHeader, err := ctx.Request().FormFile("file")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while getting file from request %v,%v: %v\n", file, fileHeader, err)
		return ctx.JSON(http.StatusInternalServerError, contracts.ErrorResponse{Error: "Could not upload file"})
	}
	defer file.Close()
	fileReader, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while reading file contents: %s\n", err)
		return ctx.JSON(http.StatusInternalServerError, contracts.ErrorResponse{Error: "Could not upload file"})
	}
	cid, err := services.IpfsShell.Add(bytes.NewReader(fileReader))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while adding file to ipfs: %s\n", err)
		return ctx.JSON(http.StatusInternalServerError, contracts.ErrorResponse{Error: "Could not upload file"})
	}
	fileSize := fmt.Sprint(fileHeader.Size, " bytes")
	fileResp := contracts.File{FileID: cid, Timestamp: currTime, Size: fileSize}
	return ctx.JSON(http.StatusCreated, fileResp)
}
