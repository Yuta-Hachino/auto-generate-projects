package generate

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/Yuta-Hachino/auto-generate-projects/service/chatgpt"
	"github.com/Yuta-Hachino/auto-generate-projects/service/storage"
)

const firstSubPrompt = "をCloudRunサービスとしてデプロイしたい。terraformで書くとどうなる？必要なファイルのリストをカンマ区切り文字列にして教えて？その一行以外の説明などはいらない。"

const createFileTextSubFirstPrompt = "というプロンプトに対して貴方は"
const createFileTextSubSecondPrompt = "というファイル名のリストを返した。その中で"
const createFileTextSubThirdPrompt = "のコード（内容のテキスト）を提示して？回答をそのままファイルに張り付けるので余計な説明文などはいらない。"

func NewProject(sourcePrompt string) error {
	ctx := context.Background()
	// GCSクライアントを作成
	client, err := storage.CreateStorageClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// アップロード先のバケットとファイル名
	bucketName := "auto-generate-project-storage"

	firstPrompt := sourcePrompt + firstSubPrompt
	targetFilesCsv, err := chatgpt.SendMessage(firstPrompt)
	if err != nil {
		return errors.New("prompt or LLM APIs error.")
	}

	targetFiles := strings.Split(targetFilesCsv, ",")
	if len(targetFiles) == 0 {
		return errors.New("not found target files.")
	}

	createFileTextPrompt := firstPrompt + createFileTextSubFirstPrompt + targetFilesCsv + createFileTextSubSecondPrompt
	for _, targetFile := range targetFiles {
		fileContent, err := chatgpt.SendMessage(createFileTextPrompt + targetFile + createFileTextSubThirdPrompt)
		if err != nil {
			return errors.New("prompt or LLM APIs error." + targetFile)
		}

		// ファイルをアップロード
		if err := storage.UploadFile(ctx, client, bucketName, targetFile, fileContent); err != nil {
			log.Fatalf("Failed to upload file: %v", err)
			return errors.New("Failed to upload file")
		}
	}

	return nil
}
