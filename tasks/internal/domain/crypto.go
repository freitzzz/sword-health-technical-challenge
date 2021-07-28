package domain

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
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

	return hex.EncodeToString(encryptedSummary)
}

func decryptSummary(block cipher.Block, summary string) string {

	decodedDecryptedSummary, _ := hex.DecodeString(summary)

	decryptedSummary := make([]byte, len(decodedDecryptedSummary))

	block.Decrypt(decryptedSummary, decodedDecryptedSummary)

	return strings.TrimSpace(string(decryptedSummary))

}
