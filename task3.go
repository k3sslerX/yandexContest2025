package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type taxi struct {
	timestamp int
	id        int
	position  int
	meta      int
}
type order struct {
	timestamp int
	id        int
	position  int
	time      int
	meta      int
}

func main() {
	var n, l, s int
	var taxis []taxi
	var orders []order
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	parts := strings.Fields(str)
	n, _ = strconv.Atoi(parts[0])
	l, _ = strconv.Atoi(parts[1])
	s, _ = strconv.Atoi(parts[2])
	for i := 0; i < n; i++ {
		scanner.Scan()
		str := scanner.Text()
		parts := strings.Fields(str)
		if len(parts) == 4 && parts[0] == "TAXI" {
			timestamp, _ := strconv.Atoi(parts[1])
			id, _ := strconv.Atoi(parts[2])
			position, _ := strconv.Atoi(parts[3])
			taxis = append(taxis, taxi{
				timestamp: timestamp,
				id:        id,
				position:  position,
				meta:      i,
			})
		} else if len(parts) == 5 && parts[0] == "ORDER" {
			timestamp, _ := strconv.Atoi(parts[1])
			id, _ := strconv.Atoi(parts[2])
			position, _ := strconv.Atoi(parts[3])
			time, _ := strconv.Atoi(parts[4])
			orders = append(orders, order{
				timestamp: timestamp,
				id:        id,
				position:  position,
				time:      time,
				meta:      i,
			})
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(orders); i++ {
		accepted := 0
		acceptedStr := ""
		for j := 0; j < len(taxis); j++ {
			if taxis[j].meta < orders[i].meta {
				if orders[i].time == 0 {
					if orders[i].timestamp == taxis[j].timestamp && orders[i].position == taxis[j].position {
						if accepted < 5 {
							acceptedStr += strconv.Itoa(taxis[j].id)
							acceptedStr += " "
							accepted++
							continue
						}
					}
				} else {
					newPossiblePosition := (taxis[j].position + (orders[i].timestamp-taxis[j].timestamp)*s) % l
					maxPositionRange := orders[i].position - newPossiblePosition
					minPositionRange := orders[i].position - taxis[j].position
					if maxPositionRange < 0 {
						maxPositionRange += l
					}
					if minPositionRange < 0 {
						minPositionRange += l
					}
					if maxPositionRange <= orders[i].time*s && minPositionRange <= orders[i].time*s && accepted < 5 {
						acceptedStr += strconv.Itoa(taxis[j].id)
						acceptedStr += " "
						accepted++
					}
				}
			}
		}
		if accepted == 0 {
			acceptedStr = "-1 "
		}
		acceptedStr = strings.TrimSuffix(acceptedStr, " ")
		acceptedStr += "\n"
		_, _ = writer.WriteString(acceptedStr)
	}
	_ = writer.Flush()
}
