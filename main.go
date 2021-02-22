package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func getFileContent(filename string) []byte {
	st, err := os.Stat(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", filename, "does not exist")
		return nil
	}
	if st.Mode().IsRegular() == false {
		fmt.Fprintln(os.Stderr, "Error:", filename, "is not a regular file")
		return nil
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: could not read from", filename, "file")
		return nil
	}
	return data
}

func epurSlices(fileSlices [][]byte) [][]byte {
	epured := [][]byte{}

	for _, sli := range fileSlices {
		if sli != nil && len(sli) > 0 {
			epured = append(epured, sli)
		}
	}
	return epured
}

func getHeightFromSlices(fileSlices [][]byte) int {
	txt := string(fileSlices[0])
	value, err := strconv.Atoi(txt)
	if err != nil {
		return -1
	}
	return value
}

func verifySlicesLength(fileSlices [][]byte) bool {
	size := len(fileSlices[0])

	for i := 1; i < len(fileSlices); i++ {
		if len(fileSlices[i]) != size {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Error: Invalid number of arguments,", len(os.Args), "provided, 1 required")
		os.Exit(84)
	}
	data := getFileContent(os.Args[1])
	if data == nil {
		os.Exit(84)
	}
	slices := bytes.Split(data, []byte{byte('\n')})
	slices = epurSlices(slices)
	height := getHeightFromSlices(slices)
	if height == -1 {
		fmt.Fprintln(os.Stderr, "Error: invalid syntax in the map's first line")
		os.Exit(84)
	}
	if height != len(slices)-1 {
		fmt.Fprintln(os.Stderr, "Error: different sizes between file and first line token")
		os.Exit(84)
	}
	slices = slices[1:]
	if verifySlicesLength(slices) == false {
		fmt.Fprintln(os.Stderr, "Error: lines are not all equal in size")
		os.Exit(84)
	}
}
