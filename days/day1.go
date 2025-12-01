package days

import (
	"bufio"
	"os"
	"strconv"
)

func Day1() (int, error) {
	f, err := os.Open(".data/day1.txt")
	if err != nil {
		return -1, err
	}

	r := bufio.NewReader(f)

	pos := 50
	count := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		lr := line[0]
		amount, err := strconv.Atoi(line[1 : len(line)-1])
		if err != nil {
			return -1, err
		}
		switch lr {
		case 'L':
			pos -= amount
		case 'R':
			pos += amount
		}

		pos = (pos + 100) % 100

		if pos == 0 {
			count++
		}
	}

	defer f.Close()

	return count, nil
}
