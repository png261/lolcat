package lolcat

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strings"
)

func rgb(i int) (int, int, int) {
	var f = rand.Float64()
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func Print(str string) {
	for pos, char := range str {
		r, g, b := rgb(pos)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, char)
	}
}

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intented to work with pipes.")
		fmt.Println(`Example: echo "hello world"  | gololcat`)
	}

	reader := bufio.NewReader(os.Stdin)
	j := 0

	var str []string

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		str = append(str, string(input))
		j++
	}

	Print(strings.Join(str, ""))
}
