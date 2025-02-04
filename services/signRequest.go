package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func SignAWSRequest(tiploc string) map[string]string {
	// Configuration
	awsRegion := os.Getenv("AWS_REGION")
	service := os.Getenv("SERVICE")
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	apiKey := os.Getenv("API_KEY")

	// Request details
	httpMethod := "GET"
	apiHostName := os.Getenv("API_HOST_NAME")
	apiEndpoint := fmt.Sprintf("/services/%v", tiploc) // Ensure this matches exactly
	requestPayload := ""                               // Empty payload for GET requests
	queryString := ""                                  // No query parameters

	// Step 1: Generate timestamps
	now := time.Now().UTC()
	amzDate := now.Format("20060102T150405Z") // ISO 8601 Basic Format
	dateStamp := now.Format("20060102")       // YYYYMMDD

	// Step 2: Canonical Headers
	headers := map[string]string{
		"accept":                       "application/json",
		"access-control-allow-headers": "Content-Type,Access-Control-Allow-Origin,Authorization,X-Amz-Date,X-Api-Key,X-Amz-Security-Token",
		"access-control-allow-origin":  "*",
		"host":                         apiHostName,
		"x-amz-date":                   amzDate,
		"x-api-key":                    apiKey,
	}

	// Step 3: Construct Canonical Request
	signedHeaders := getSignedHeaders(headers)
	canonicalHeaders := getCanonicalHeaders(headers)
	payloadHash := hashSHA256(requestPayload)

	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		httpMethod,
		apiEndpoint,
		queryString,
		canonicalHeaders,
		signedHeaders,
		payloadHash,
	)

	// Step 4: Create the String-to-Sign
	credentialScope := fmt.Sprintf("%s/%s/%s/aws4_request", dateStamp, awsRegion, service)
	stringToSign := fmt.Sprintf("AWS4-HMAC-SHA256\n%s\n%s\n%s",
		amzDate,
		credentialScope,
		hashSHA256(canonicalRequest),
	)

	// Step 5: Calculate the Signature
	signingKey := getSignatureKey(secretKey, dateStamp, awsRegion, service)
	signature := hex.EncodeToString(hmacSHA256(stringToSign, signingKey))

	// Step 6: Construct Authorization Header
	authHeader := fmt.Sprintf("AWS4-HMAC-SHA256 Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		accessKey, credentialScope, signedHeaders, signature)

	// Step 7: Return Headers
	return map[string]string{
		"Authorization": authHeader,
		"x-amz-date":    amzDate,
		"x-api-key":     apiKey,
		"Host":          apiHostName,
		"Accept":        "application/json",
	}
}

// Helper: Hash input using SHA256 and return hex string
func hashSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// Helper: Generate HMAC-SHA256 signature
func hmacSHA256(data string, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return h.Sum(nil)
}

// Helper: Get AWS signing key
func getSignatureKey(secret, dateStamp, region, service string) []byte {
	kDate := hmacSHA256(dateStamp, []byte("AWS4"+secret))
	kRegion := hmacSHA256(region, kDate)
	kService := hmacSHA256(service, kRegion)
	kSigning := hmacSHA256("aws4_request", kService)
	return kSigning
}

// Helper: Get sorted signed headers string
func getSignedHeaders(headers map[string]string) string {
	var keys []string
	for key := range headers {
		keys = append(keys, strings.ToLower(key))
	}
	sort.Strings(keys)
	return strings.Join(keys, ";")
}

// Helper: Construct Canonical Headers string
func getCanonicalHeaders(headers map[string]string) string {
	var buf bytes.Buffer
	var keys []string
	for key := range headers {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		buf.WriteString(strings.ToLower(key) + ":" + strings.TrimSpace(headers[key]) + "\n")
	}
	return buf.String()
}
