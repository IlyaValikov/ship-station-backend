package repository

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

type MinioRepository interface{
    UploadServiceImage(userID, shipID uint64, imageBytes []byte, contentType string) (string, error)
    RemoveServiceImage(userID, shipID uint64) error
}

func (r *Repository) UploadServiceImage(shipID, userID uint, imageBytes []byte, contentType string) (string, error) {
    objectName := fmt.Sprintf("ships/%d/image", shipID)

	reader := io.NopCloser(bytes.NewReader(imageBytes))

	_, err := r.mc.PutObject(context.TODO(), "images-bucket", objectName, reader, int64(len(imageBytes)), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
        return "", errors.New("ошибка при добавлении изображения в минио бакет")
    }

	// Формирование URL изображения
	imageURL := fmt.Sprintf("http://localhost:9000/images-bucket/%s", objectName)

    return imageURL, nil
}

func (r *Repository) RemoveServiceImage(shipID, userID uint) error {
    objectName := fmt.Sprintf("ships/%d/image", shipID)
	err := r.mc.RemoveObject(context.TODO(), "images-bucket", objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return errors.New("не удалось удалить изображение из бакет")
	}

    if err := r.db.Table("ships").
	Where("ship_id = ?", shipID).
	Update("photo", nil).Error; 
	err != nil {
        return errors.New("ошибка при обновлении URL изображения в базе данных")
    }
	
    return nil
}
