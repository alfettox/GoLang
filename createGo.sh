#!/bin/bash

# Define the content of the Go file
GO_CODE='

package main

import(
    "fmt"
)

func main() {
    
    
    }

    
'

# Create the Go file
echo "$GO_CODE" > new_go_file.go

# Create a shell script to compile and run the Go file
echo '#!/bin/bash
go run new_go_file.go
'
