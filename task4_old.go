package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var m, n int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	parts := strings.Fields(str)
	m, _ = strconv.Atoi(parts[0])
	n, _ = strconv.Atoi(parts[1])
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		scanner.Scan()
		str := scanner.Text()
		for j := 0; j < n; j++ {
			if str[j] == 'o' {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = 1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

		}
	}
	//writer := bufio.NewWriter(os.Stdout)
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		_, _ = writer.WriteString(strconv.Itoa(matrix[i][j]))
	//	}
	//	_, _ = writer.WriteString("\n")
	//}
	//_ = writer.Flush()
}
