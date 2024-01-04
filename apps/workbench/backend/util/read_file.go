package util

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/Boostport/mjml-go"
	"github.com/deamgo/workbench/initialize"
	"os"
)

// PathRole Structure maps database tables
type PathRole struct {
	Role   string
	Path   string
	Method string
}

func GetRoles() []PathRole {

	// Open CSV file
	file, err := os.Open(initialize.GetConfig().PermissionConfig.Path)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return nil
	}
	defer file.Close()

	// create CSV reader
	reader := csv.NewReader(file)

	// read table header
	header, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV header:", err)
		return nil
	}

	// Create a map to store column indexes
	headerMap := make(map[string]int)
	for i, col := range header {
		headerMap[col] = i
	}

	// read data
	var userRoles []PathRole
	for {
		row, err := reader.Read()
		if err != nil {
			break // File read end
		}

		// parse data
		role := PathRole{
			Role:   row[headerMap["Role"]],
			Path:   row[headerMap["Path"]],
			Method: row[headerMap["Method"]],
		}

		userRoles = append(userRoles, role)
	}

	return userRoles
}

func ParseMJMLFile(filePath string, ctx context.Context) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	html, err := mjml.ToHTML(ctx, string(file), mjml.WithMinify(true))
	if err != nil {
		return "", err
	}
	return html, nil
}
