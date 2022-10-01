package main
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// Note: Ignoring potential erros from input.Err()
	fmt.Printf("%s\t%s\n", "Count", "line")
	var dupes = 0
	for line, n := range counts {
		if n > 1 {
			dupes++
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	fmt.Printf("%d\t%s\n", dupes, "Duplicates")
}