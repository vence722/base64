package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: base64 infile outfile [linelength]. If line length is 0, then all the input will be in a line.")
		os.Exit(1)
	}

	infile := os.Args[1]
	outfile := os.Args[2]
	lineLen := 0
	if len(os.Args) >= 4 {
		lineLen, _ = strconv.Atoi(os.Args[3])
	}

	f, _ := os.Open(infile)
	defer f.Close()
	of, _ := os.Create(outfile)
	defer of.Close()
	data, _ := ioutil.ReadAll(f)

	buf := bytes.NewBuffer([]byte{})
	encoder := base64.NewEncoder(base64.StdEncoding, buf)
	defer encoder.Close()
	encoder.Write(data)

	result := insertLineDelimiter(buf.String(), "\r\n", lineLen)

	of.WriteString(result)
}

func insertLineDelimiter(src string, dlm string, lineLen int) string {
	lenSrc := len(src)
	lenDlm := len(dlm)
	lines := lenSrc / lineLen
	if lenSrc%lineLen == 0 {
		lines--
	}
	bufSize := lenSrc + lines*lenDlm
	buffer := make([]byte, bufSize, bufSize)

	i := 0
	j := 0
	for i < lenSrc {
		if i > 0 && i%lineLen == 0 {
			for k := 0; k < lenDlm; k++ {
				buffer[j+k] = byte(dlm[k])
			}
			j += lenDlm
		}
		buffer[j] = byte(src[i])
		i++
		j++
	}
	return string(buffer)
}
