package repository

import (
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

func SavePhoto(photoId int, file multipart.File) error {
	//создает папку для статики если ее нет
	if _, err := os.Stat("/static"); os.IsNotExist(err) {
		err := os.Mkdir("/static", 0755)
		if err != nil {
			return err
		}
	}

	name := strconv.Itoa(photoId)
	//создает файл для сохранения фотки
	tmpfile, err := os.Create("./static/" + name + ".png")
	defer tmpfile.Close()
	if err != nil {
		return err
	}

	//сохраянет фотку в файл
	_, err = io.Copy(tmpfile, file)
	if err != nil {
		return err
	}

	return nil
}

func GetPhoto(photoId int) (*os.File, error) {
	name := strconv.Itoa(photoId)
	path := "./static" + name + ".png"

	img, err := os.Open(path)
	defer img.Close()
	if err != nil {
		return nil, err
	}

	return img, nil
}
