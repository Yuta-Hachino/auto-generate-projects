package storage

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
)

// createStorageClient は新しいStorageクライアントを作成します。
func CreateStorageClient(ctx context.Context) (*storage.Client, error) {
	return storage.NewClient(ctx)
}

// uploadFile は指定したバケットにテキストファイルをアップロードします。
func UploadFile(ctx context.Context, client *storage.Client, bucketName, objectName, content string) error {
	// バケットを指定
	bucket := client.Bucket(bucketName)

	// オブジェクトを作成
	obj := bucket.Object(objectName)
	w := obj.NewWriter(ctx)

	// ファイルの内容を書き込み
	if _, err := w.Write([]byte(content)); err != nil {
		w.Close()
		return err
	}

	// 書き込みを閉じてアップロードを確定
	if err := w.Close(); err != nil {
		return err
	}

	log.Printf("File uploaded to %s/%s\n", bucketName, objectName)
	return nil
}
