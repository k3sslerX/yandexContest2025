package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	const Max = 50
	var Windows [Max * Max][Max * Max]int
	var n, m, x, y, enabledLights, result int
	var tmp string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	parts := strings.Fields(str)
	n, _ = strconv.Atoi(parts[0])
	m, _ = strconv.Atoi(parts[1])
	x, _ = strconv.Atoi(parts[2])
	y, _ = strconv.Atoi(parts[3])
	for i := 0; i < n*x; i++ {
		scanner.Scan()
		tmp = scanner.Text()
		tmpRune := []rune(tmp)
		for j := 0; j < m*y; j++ {
			if tmpRune[j] == 'X' {
				Windows[i][j] = 1
			} else {
				Windows[i][j] = 0
			}
		}
	}
	for i := 0; i < n*x; i += x {
		for j := 0; j < m*y; j += y {
			for p := i; p < i+x; p++ {
				for q := j; q < j+y; q++ {
					enabledLights += Windows[p][q]
				}
			}
			if float64(enabledLights) >= (float64(x)*float64(y))/2 {
				result++
			}
			enabledLights = 0
		}
		enabledLights = 0
	}
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString(strconv.Itoa(result))
	_ = writer.Flush()
}
