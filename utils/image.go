package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type Images struct {
	Id        string
	Path      string
	CreatedAt int64
	UpdatedAt int64
}

func UploadImage(request *multipart.FileHeader) (Images, error) {
	img := Images{}
	file, _ := request.Open()

	tempFile, err := os.CreateTemp("assets/img/images", "image-*.jpg")
	if err != nil {
		fmt.Println("create Image", err.Error())
		return img, err
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("readIotil", err.Error())

		return img, err
	}
	tempFile.Write(fileBytes)

	fileName := tempFile.Name()
	newFileName := strings.Split(fileName, "/")
	fmt.Println(fileName, newFileName)
	img = Images{
		Id:        GenerateId(),
		Path:      newFileName[3],
		CreatedAt: CurrentMillis(),
		UpdatedAt: CurrentMillis(),
	}

	// image = service.ImageRepository.Create(ctx, tx, image)
	// imageResponses = append(imageResponses, utils.ToImageResponse(image))

	return img, nil
}
func GetImage() string {

	return ""
}
func RemoveImage() string {

	return ""
}
