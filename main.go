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
	size := len(fileSlices)
	epured := make([][]byte, 0, size)

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

func calcValue(ints [][]int, i int, j int) int {
	if i == 0 || j == 0 {
		return 1
	}
	top := ints[i-1][j]
	left := ints[i][j-1]
	topLeft := ints[i-1][j-1]
	if top <= left && top <= topLeft {
		return top + 1
	} else if left <= top && left <= topLeft {
		return left + 1
	}
	return topLeft + 1
}

func bsq(slices [][]byte) (x int, y int, size int) {
	ints := make([][]int, len(slices))
	x = 0
	y = 0
	size = 0

	ints[0] = make([]int, len(slices[0]))
	for i := 0; i < len(ints); i++ {
		ints[i] = make([]int, len(slices[i]))
		for j := 0; j < len(ints[i]); j++ {
			if slices[i][j] == '.' {
				ints[i][j] = calcValue(ints, i, j)
			} else {
				ints[i][j] = 0
			}
			if ints[i][j] > size {
				size = ints[i][j]
				x = i
				y = j
			}
		}
	}
	return x, y, size
}

func fillXSlice(slices [][]byte, x int, y int, size int) []string {
	repres := make([]string, len(slices))

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			slices[x-i][y-j] = 'x'
		}
	}
	for i := 0; i < len(slices); i++ {
		repres[i] = string(slices[i])
	}
	return repres
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
	x, y, size := bsq(slices)
	repres := fillXSlice(slices, x, y, size)
	for _, sli := range repres {
		fmt.Println(sli)
	}
}
