package domain

import (
	"crypto/aes"
	"strings"
	"testing"
)

func TestEncryptSummaryReturnsEncryptedStringOfSummaryBasedOnProvidedCipherBlock(t *testing.T) {

	block, _ := aes.NewCipher([]byte(strings.Repeat("x", 32)))

	summary := strings.Repeat("a", aes.BlockSize)

	encryptedSummary := encryptSummary(block, summary)

	expectedecryptedSummaryBytes := make([]byte, len(summary))

	block.Encrypt(expectedecryptedSummaryBytes, []byte(summary))

	expectedecryptedSummary := string(expectedecryptedSummaryBytes)

	if encryptedSummary != expectedecryptedSummary {
		t.Fatalf("Function output is different than the expected encrypted summary: \nOutput: %s\nExpected: %s", encryptedSummary, expectedecryptedSummary)
	}

}

func TestEncryptSummaryReturnsEncryptedStringOfSummaryWithPaddingSpacesIfItsSizeIsShorterThanCipherBlockSize(t *testing.T) {

	block, _ := aes.NewCipher([]byte(strings.Repeat("x", 32)))

	summary := strings.Repeat("a", aes.BlockSize-1)

	encryptedSummary := encryptSummary(block, summary)

	expectedecryptedSummaryBytes := make([]byte, aes.BlockSize)

	block.Encrypt(expectedecryptedSummaryBytes, []byte(summary+" "))

	expectedecryptedSummary := string(expectedecryptedSummaryBytes)

	if encryptedSummary != expectedecryptedSummary {
		t.Fatalf("Function output is different than the expected encrypted summary: \nOutput: %s\nExpected: %s", encryptedSummary, expectedecryptedSummary)
	}

}

func TestDecryptSummaryReturnsDecryptedStringOfSummaryWithoutPaddingSpacesIfItsSizeIsShorterThanCipherBlockSize(t *testing.T) {

	block, _ := aes.NewCipher([]byte(strings.Repeat("x", 32)))

	summary := strings.Repeat("a", aes.BlockSize-1)

	encryptedSummary := encryptSummary(block, summary)

	decryptedSummary := decryptSummary(block, encryptedSummary)

	if summary != decryptedSummary {
		t.Fatalf("Function output is different than the expected decrypted summary: \nOutput: %s\nExpected: %s", decryptedSummary, summary)
	}

}
