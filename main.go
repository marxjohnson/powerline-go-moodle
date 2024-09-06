package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"fmt"
)

type PowerlineSegment struct {
	Name       string `json:"name"`
	Content    string `json:"content"`
	Foreground int    `json:"foreground"`
	Background int    `json:"background"`
}

func findTopLevelVersionPHP(startDir string) string {
	var data map[string]interface{}
	var prefix, pattern string

	currentDir := startDir
	for {
		if currentDir == filepath.Dir(currentDir) { // Stop at the root directory
			break
		}

		composerFile := filepath.Join(currentDir, "composer.json")
		if _, err := os.Stat(composerFile); err == nil {
			// File exists, load composer.json
			fileContent, err := ioutil.ReadFile(composerFile)
			if err != nil {
				break
			}
			err = json.Unmarshal(fileContent, &data)
			if err != nil {
				break
			}

			if name, ok := data["name"].(string); ok {
				if name == "moodle/moodle" || name == "totara/totara_meta" {
					break
				}
			}
		}
		currentDir = filepath.Dir(currentDir)
	}

	// Check for version.php
	versionFile := filepath.Join(currentDir, "version.php")
	if _, err := os.Stat(versionFile); err == nil {
		fileContent, err := ioutil.ReadFile(versionFile)
		if err != nil {
			return ""
		}
		content := string(fileContent)

		// Regex search for Moodle or Totara version
		if data["name"] == "moodle/moodle" {
			prefix = "M"
			pattern = `\$release\s*=\s*'([^ ]+) `
		} else {
			prefix = "T"
			pattern = `\$TOTARA->version\s*=\s*'([^']+)'`
		}
		re := regexp.MustCompile(pattern)
		match := re.FindStringSubmatch(content)
		if len(match) > 1 {
			return prefix + match[1]
		}
	}
	return ""
}

func main() {
	startDir, _ := os.Getwd() // Get current working directory

	releaseVersion := findTopLevelVersionPHP(startDir)

	if releaseVersion == "" {
		fmt.Println("{}")
	}

	segment := PowerlineSegment{
		Name: "moodle",
		Content: releaseVersion,
		Foreground: 15,
		Background: 166,
	}

	output, err := json.Marshal([1]PowerlineSegment{segment})

	if err != nil {
		fmt.Println("{}")
		os.Exit(1)
	}

	// Output the JSON string
	fmt.Println(string(output))
}
