package domain

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
	"strings"
)

const (
	taskSummaryAESKeyEnvKey = "task_summary_aes_key"
)

func LoadTaskSummaryCipher() (cipher.Block, error) {

	key := os.Getenv(taskSummaryAESKeyEnvKey)

	return aes.NewCipher([]byte(key))

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
