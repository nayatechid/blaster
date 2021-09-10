## how to run 
go run *.go --subject "Email Subject" --receivers receivers.csv --template template.html

## how to build
go build *.go

## pre-requisites
require minimum go version 1.16 

# TODO
1. [x] send arguments from os.Args
2. [x] mail destinations using csv file
3. [x] mail template using html file
4. [ ] configuration using yaml
