package domain

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
	"strings"

	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/logging"
)

const (
	taskSummaryAESKeyEnvKey = "task_summary_aes_key"
)

func createTaskSummaryCipher() cipher.Block {

	key := os.Getenv(taskSummaryAESKeyEnvKey)

	block, cerr := aes.NewCipher([]byte(key))

	if cerr != nil {
		logging.LogError("Failed to create AES Cipher for task summary encryption")
		logging.LogError(cerr.Error())

		panic("Cannot proceed without cipher for encrpyting task summary")
	}

	return block

}

func encryptSummary(block cipher.Block, summary string) string {

	summaryToEncrypt := summary

	summaryLen := len(summaryToEncrypt)

	blockSize := block.BlockSize()

	if summaryLen < blockSize {
		summaryToEncrypt = summaryToEncrypt + strings.Repeat(" ", blockSize-summaryLen)
	}

	encryptedSummary := make([]byte, len(summaryToEncrypt))

	block.Encrypt(encryptedSummary, []byte(summaryToEncrypt))

	return string(encryptedSummary)
}

func decryptSummary(block cipher.Block, summary string) string {

	decryptedSummary := make([]byte, len(summary))

	block.Decrypt(decryptedSummary, []byte(summary))

	return strings.TrimSpace(string(decryptedSummary))

}
