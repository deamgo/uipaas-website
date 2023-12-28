package util

import (
	"encoding/csv"
	"fmt"
	"os"
)

// PathRole Structure maps database tables
type PathRole struct {
	Role   string
	Path   string
	Method string
}

func GetRoles() []PathRole {

	// Gets the directory where the executable file is located
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Open CSV file
	file, err := os.Open(rootDir + "/auth/permission/permissions.csv")
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
