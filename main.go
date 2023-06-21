package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	gitChangesFileName = "git_changes.txt"
	logFileName        = "log_file.txt"
	componentsFileName = "component_list.txt"
	mappingsFileName   = "mappings.json"
	liveComponentsPath = "infrastructure/live"
	modulesPath        = "infrastructure/modules"
)

func main() {
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
	log.SetPrefix("[INFO] ")
	log.Println("Processing Git Changes File.")
	log.Println("Generating list of components to be deployed.")

	mappings, err := readJSONFile(mappingsFileName)
	if err != nil {
		log.SetPrefix("[ERROR] ")
		log.Fatal(err)
	}

	// Open the git-changes-file
	gitChangesFile, err := os.Open(gitChangesFileName)
	if err != nil {
		log.SetPrefix("[ERROR] ")
		log.Fatal(err)
	}
	defer gitChangesFile.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(gitChangesFile)

	components := []string{}

	// Iterate over each line and print it
	for scanner.Scan() {
		line := scanner.Text()
		log.Println("Reading file line: ", line)

		//	Separate the line into its path by / add them to a list and print them
		var path []string = strings.Split(line, "/")

		// concat path[0] and path[1] to get the first 2 elements of the path
		linePath := path[0] + "/" + path[1]

		if linePath == liveComponentsPath {
			// Store in components the last item of the path array only if is not there already
			if !strings.Contains(strings.Join(components, " "), path[len(path)-1]) {
				components = append(components, path[len(path)-1])
			}
		}

		if linePath == modulesPath {
			module := path[len(path)-1]

			// Store in components the last item of the path array only if is not there already
			if _, ok := mappings[module]; ok {
				component := mappings[module]
				if !strings.Contains(strings.Join(components, " "), component) {
					components = append(components, component)
				}
			}
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.SetPrefix("[ERROR] ")
		log.Fatal(err)
	}

	// Write the components to a file
	log.Println("Writing the list of components to a file...")

	f, err := os.Create(componentsFileName)
	if err != nil {
		log.SetPrefix("[ERROR] ")
		log.Fatal(err)
	}
	defer f.Close()

	log.Println("List of components: ")
	for _, component := range components {
		log.Println(component)
		fmt.Fprintln(f, component)
	}
}

func readJSONFile(filename string) (map[string]string, error) {
	// Read the JSON file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a map to store the JSON data
	mappings := make(map[string]string)
	err = json.NewDecoder(file).Decode(&mappings)

	if err != nil {
		return nil, err
	}

	return mappings, nil
}
