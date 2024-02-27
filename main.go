package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type City struct {
	min, max, mean float64
	count          int64
}

const MAX_CITIES int = 10000
const MEASUREMENTS_FILE_PATH string = "./data/measurements.txt"

var cities map[string]City = make(map[string]City, MAX_CITIES)
var keys []string = make([]string, 0, MAX_CITIES)

func main() {
	file, err := os.Open(MEASUREMENTS_FILE_PATH)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		data := strings.Split(scanner.Text(), ";")

		temperature, _ := strconv.ParseFloat(data[1], 64)
		city := data[0]
		if val, ok := cities[city]; ok {
			val.count += 1
			val.max = max(temperature, val.max)
			val.min = min(temperature, val.min)
			val.mean += temperature
			cities[city] = val
		} else {
			keys = append(keys, city)
			cities[city] = City{
				count: 1,
				max:   temperature,
				min:   temperature,
				mean:  temperature,
			}
		}
	}
	var result strings.Builder
	slices.Sort(keys)
	for _, key := range keys {
		city := cities[key]
		avg := city.mean / float64(city.count)
		result.WriteString(fmt.Sprintf("%s=%.1f/%.1f/%.1f, ", key, city.min, round(avg), city.max))
	}
	fmt.Printf("{%s}", result.String()[:result.Len()-2])

}

func round(number float64) float64 {
	return math.Ceil(number*10) / 10
}
