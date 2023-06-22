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
	//gitChangesFileName = "git_changes.txt"
	//mappingsFileName   = "gtcModToCompMap.json"
	logFileName        = "git-terra-changes.log"
	componentsFileName = "git-terra-changes-components.txt"
	liveComponentsPath = "infrastructure/live"
	modulesPath        = "infrastructure/modules"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./git-terra-changes <git_changes_file.txt> <mappings_file.json>")
		os.Exit(1)
	}

	gitChangesFileName := os.Args[1]
	mappingsFileName := os.Args[2]

	mappings, mappingsErr := readJSONFile(mappingsFileName)
	if mappingsErr != nil {
		log.SetPrefix("[ERROR] ")
		log.Fatal(mappingsErr)
	}

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

	// Open the git-changes-file
	gitChangesFile, err := os.Open(gitChangesFileName)
	if err != nil {
		log.SetPrefix("[ERROR] ")
		log.Fatal(err)
	}
	defer gitChangesFile.Close()

	scanner := bufio.NewScanner(gitChangesFile)

	components := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		log.Println("===> Reading file line: ", line)
		setComponent(line, &components, mappings)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.SetPrefix("[ERROR] ")
		log.Fatal(err)
	}

	// Write the components to a file
	writeToFile(components)
}

func writeToFile(components []string) {
	log.Println("Writing the list of components to a file...")

	f, err := os.Create(componentsFileName)
	if err != nil {
		log.SetPrefix("[ERROR] ")
		log.Fatal(err)
	}
	defer f.Close()

	if len(components) == 0 {
		log.SetPrefix("[WARNING] ")
		log.Println("No components found. Creating empty file.")
		fmt.Fprintln(f, "")
		return
	}

	log.Println("List of components: ")
	for _, component := range components {
		log.Println(component)
		fmt.Fprintln(f, component)
	}
}

func setComponent(line string, components *[]string, mappings map[string]string) {
	//	Separate the line into its path by / add them to a list and print them
	var path []string = strings.Split(line, "/")

	if len(path) < 2 {
		log.SetPrefix("[WARNING] ")
		log.Println("The path is not valid, skipping: ", line)
		log.SetPrefix("[INFO] ") // Reset the prefix
		return
	}
	// concat path[0] and path[1] to get the first 2 elements of the path
	linePath := path[0] + "/" + path[1]

	if linePath == liveComponentsPath {
		// Store in components the last item of the path array only if is not there already
		if !strings.Contains(strings.Join(*components, " "), path[len(path)-1]) {
			log.Println("the Path is a component: ", line)
			*components = append(*components, path[len(path)-1])
		}
	} else if linePath == modulesPath {
		module := path[len(path)-1]

		// Store in components the last item of the path array only if is not there already
		if _, ok := mappings[module]; ok {
			component := mappings[module]
			if !strings.Contains(strings.Join(*components, " "), component) {
				log.Println("The module path maps with a component: ", line, component)
				*components = append(*components, component)
			}
		}
	} else {
		log.Println("The path is not in the scope, skipping: ", line)
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
