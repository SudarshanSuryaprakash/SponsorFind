# SponsorFind
This is a go program that finds whether a company passed in as a command line argument sponsors tier 2 skilled worker visas

To use,

Either simply run the executable with an argument like so: "./sponsor.exe <company name>"
  You will need to have tier2.csv in the same directory as the executable
  
Otherwise:

- Install Go
- In the directory where the sponsor.go file is present, ensure that tier2.csv is present too.
- execute the command "go run sponsor.go <comapny name>" OR build an executable with "go build sponsor.go" and run the executable with an argument like so: "./sponsor.exe <company name>"
  
  
