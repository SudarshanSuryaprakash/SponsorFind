package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
	"strings"
)

func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func main() {
    records := readCsvFile("./tier2.csv")
	var result []string
	args := os.Args[1:]
	query := strings.Join(args, " ")
	
	if len(query) == 0 {
		fmt.Println("Please provide an argument")
		return
	}
	
	var entry []string
	for _, s := range records {
		entry = append(entry, "\n" + s[0] + ", " + s[1] + ", " + s[2] + ", " + s[3] + ", " + s[4])
	}	

	for _, res := range entry {
		if strings.Contains(strings.ToLower(res), strings.ToLower(query)) {
			// result = res	
			result = append(result, res)
		} 
	} 

	if len(result) == 0 {
		fmt.Printf("%s does not sponsor tier 2 work visas", query)
		return
	}

	if len(result) > 1 {
		fmt.Printf("%s sponsors work visas!", query)
		for _, res := range result {
			fmt.Println(res)
		}
	} else {
		fmt.Println(result[0])
	}
}