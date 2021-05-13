package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin) //輸入內容
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)             //去空白
	number, err := strconv.ParseFloat(input, 64) //轉型Float
	if err != nil {
		return 0, err
	}
	return number, nil
}
