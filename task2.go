package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func decompressVector(vector []int, n, a, b int) []int {
	decompressedVector := make([]int, n)
	if a >= b {
		return decompressedVector
	}
	for i := 0; i < n; i++ {
		//if i%2 == 1 {
		//	vector[i]++
		//}
		decompressedVector[i] = ((vector[i] * (b - a)) / 255) + a
	}
	//Min := decompressedVector[0]
	//Max := decompressedVector[0]
	//minIndex := 0
	//maxIndex := n - 1
	//for i := 1; i < n; i++ {
	//	if decompressedVector[i] < Min {
	//		Min = decompressedVector[i]
	//		minIndex = i
	//	} else if decompressedVector[i] > Max {
	//		Max = decompressedVector[i]
	//		maxIndex = i
	//	}
	//}
	//decompressedVector[minIndex] = a
	//decompressedVector[maxIndex] = b
	return decompressedVector
}

func main() {
	var n, a, b int
	var result int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	parts := strings.Fields(str)
	n, _ = strconv.Atoi(parts[0])
	scanner.Scan()
	str = scanner.Text()
	parts = strings.Fields(str)
	vectorSrc := make([]int, n)
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(parts[i])
		vectorSrc[i] = num
	}
	scanner.Scan()
	str = scanner.Text()
	parts = strings.Fields(str)
	vectorDocCompressed := make([]int, n)
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(parts[i])
		vectorDocCompressed[i] = num
	}
	scanner.Scan()
	str = scanner.Text()
	parts = strings.Fields(str)
	a, _ = strconv.Atoi(parts[0])
	b, _ = strconv.Atoi(parts[1])
	vectorDoc := decompressVector(vectorDocCompressed, n, a, b)
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		//fmt.Printf("%d * %d = %d\n", vectorSrc[i], vectorDoc[i], vectorSrc[i]*vectorDoc[i])
		result += vectorSrc[i] * vectorDoc[i]
	}
	_, _ = writer.WriteString(strconv.Itoa(result))
	_ = writer.Flush()
}
