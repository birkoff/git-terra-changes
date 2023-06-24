package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Define Flags
	diff := flag.String("diff", "diff.txt", "Git diff file")
	mappings := flag.String("mappings", "gtcModToCompMap.json", "Json file containing the mappings Module to Component")
	logfile := flag.String("log", "git-terra-changes.log", "Log putput file")
	out := flag.String("out", "git-terra-changes-components.txt", "Output Fie")
	liveDir := flag.String("live_dir", "infrastructure/live", "Terragrunt/Terraform Live directory")
	modulesDir := flag.String("modules_dir", "infrastructure/modules", "Terraform Modules directory")

	flag.Parse()

	// Set log output to stdout and file
	logFile, err := os.OpenFile(*logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
	log.SetPrefix("[INFO] ")
	log.Println("===> GIT TERRAFORM & TERRAGRUNT CHANGES <===")
	log.SetPrefix("[DEBUG] ")
	log.Println("Processing Git changes File with the following parameters.")
	log.Println("diff:", *diff)
	log.Println("mappings:", *mappings)
	log.Println("logfile:", *logfile)
	log.Println("out:", *out)
	log.Println("liveDir:", *liveDir)
	log.Println("modulesDir:", *modulesDir)

	// Open the mappings file
	mappings_, mappingsErr := readJSONFile(*mappings)
	if mappingsErr != nil {
		log.SetPrefix("[ERROR] ")
		log.Fatal(mappingsErr)
	}

	// Open the git-changes-file
	gitChangesFile, err := os.Open(*diff)
	if err != nil {
		log.SetPrefix("[ERROR] ")
		log.Fatal(err)
	}
	defer gitChangesFile.Close()

	scanner := bufio.NewScanner(gitChangesFile)

	components := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		log.SetPrefix("[DEBUG] ")
		log.Println("===> Reading file line: ", line)
		setComponent(line, &components, mappings_, *liveDir, *modulesDir)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.SetPrefix("[ERROR] ")
		log.Fatal(err)
	}

	// Write the components to a file
	writeToFile(components, *out)
}

func writeToFile(components []string, out string) {
	log.SetPrefix("[INFO] ")
	log.Println("Writing the list of components to a file...")

	f, err := os.Create(out)
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

func setComponent(line string, components *[]string, mappings_ map[string]string, liveDir string, modulesDir string) {
	log.Println("Setting component for line: ", line)

	liverDirParts := strings.Split(liveDir, "/")

	//	Separate the line into its path by / add them to a list and print them
	var path []string = strings.Split(line, "/")

	if len(path) < len(liverDirParts) {
		log.SetPrefix("[WARNING] ")
		log.Println("The path is not valid, skipping: ", line)
		log.SetPrefix("[INFO] ") // Reset the prefix
		return
	}

	// concat path[0] and path[1] to get the first 2 elements of the path
	linePath := ""
	for i := 0; i < len(liverDirParts); i++ {
		if i > 0 {
			linePath += "/"
		}
		linePath += path[i]
	}

	log.Println("linePath == liveDir: ", linePath, liveDir)
	if linePath == liveDir {
		// Store in components the last item of the path array only if is not there already
		if !strings.Contains(strings.Join(*components, " "), path[len(path)-1]) {
			log.Println("the Path is a component: ", line)
			*components = append(*components, path[len(path)-1])
		}
	} else if linePath == modulesDir {
		module := path[len(path)-1]

		// Store in components the last item of the path array only if is not there already
		if _, ok := mappings_[module]; ok {
			component := mappings_[module]
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
