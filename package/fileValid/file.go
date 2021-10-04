package imgvalid

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	errmsg "trungpham/gowebbasic/package/logMessages"
)

var appErr = errmsg.NewFileMsg()

func CheckImage(file *multipart.FileHeader) (bool, string, error) {

	if file.Size > 5242880 {
		return false, appErr.FileOver5MB, nil
	}

	src, err := file.Open()

	if err != nil {
		return false, appErr.FileNotOpen, err
	}

	buff := make([]byte, 512)
	_, err = src.Read(buff)

	if err != nil {
		return false, appErr.ReadBuffFail, err
	}

	filetype := http.DetectContentType(buff)

	switch filetype {
	case "image/jpeg", "image/jpg", "image/gif", "image/png":

	default:
		return false, appErr.NotImageFile, nil
	}
	defer src.Close()

	return true, filetype, nil
}

func CopyFile(file *multipart.FileHeader, dstFileName string, path string) error {
	localFile, err := os.Create(path + dstFileName)

	if err != nil {
		return err
	}

	defer localFile.Close()

	rootFile, err := file.Open()

	if err != nil {
		return err
	}

	if _, err := io.Copy(localFile, rootFile); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		return err
	}

	return nil
}
