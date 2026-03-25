package frontend_app

 import (
 	"encoding/json"
 	"fmt"
 	"net/http"
 	"os"
 	"strings"
 	"time"
 )

 // APIError represents a structured error response from the backend API.
 type APIError struct {
 	Code    int    `json:"code"`
 	Message string `json:"message"`
 }

 // handleError logs the error and sends a standardized error response to the client.
 func handleError(w http.ResponseWriter, err error, statusCode int, customMessage string) {
 	// Log the error for debugging purposes. Include the custom message if provided.
 	logMessage := fmt.Sprintf("Error: %s", err.Error())
 	if customMessage != "" {
 		logMessage = fmt.Sprintf("%s - %s", logMessage, customMessage)
 	}
 	fmt.Fprintf(os.Stderr, "%s\n", logMessage)

 	// Construct the API error response.
 	apiError := APIError{
 		Code:    statusCode,
 		Message: customMessage,
 	}

 	// Marshal the error response to JSON.
 	jsonResponse, err := json.Marshal(apiError)
 	if err != nil {
 		// If we fail to marshal the error, log this error and send a generic 500.
 		fmt.Fprintf(os.Stderr, "Error marshaling error response: %s\n", err.Error())
 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
 		return
 	}

 	// Set the response headers.
 	w.Header().Set("Content-Type", "application/json")
 	w.WriteHeader(statusCode)

 	// Write the JSON response to the client.
 	_, err = w.Write(jsonResponse)
 	if err != nil {
 		fmt.Fprintf(os.Stderr, "Error writing error response: %s\n", err.Error())
 	}
 }

 // isValidEmail checks if the email address is valid based on a simple regex.
 func isValidEmail(email string) bool {
 	if len(email) < 3 || len(email) > 254 {
 		return false
 	}
 	if !strings.Contains(email, "@") {
 		return false
 	}
 	parts := strings.Split(email, "@")
 	if len(parts) != 2 {
 		return false
 	}
 	if len(parts[0]) == 0 || len(parts[1]) == 0 {
 		return false
 	}
 	if !strings.Contains(parts[1], ".") {
 		return false
 	}
 	return true
 }

 // formatTime formats a time.Time object into a human-readable string.
 func formatTime(t time.Time) string {
 	return t.Format(time.RFC3339) // Use a standard format
 }

 // getEnvVar retrieves an environment variable and returns a default value if it's not set.
 func getEnvVar(key, defaultValue string) string {
 	value := os.Getenv(key)
 	if value == "" {
 		return defaultValue
 	}
 	return value
 }