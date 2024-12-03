package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/klaus112/advent-of-code-2023/common"
)

type converter struct {
	sourceRangeStart      int
	destinationRangeStart int
	rangeLength           int
}

type almanac struct {
	Seeds                 []int
	SeedToSoil            []converter
	SoilToFertilizer      []converter
	FertilizerToWater     []converter
	WaterToLight          []converter
	LightToTemperature    []converter
	TemperatureToHumidity []converter
	HumidityToLocation    []converter
}

// correct answer part 1: 107430936
// seems like this will take to long for part 2 so i should probably do something like this:
// https://www.reddit.com/r/advent-of-code-2023/comments/18b560a/comment/kc4dbhq/

func main() {
	f, err := os.Open(common.InputFile)
	if err != nil {
		fmt.Println("could not open file", err)

		return
	}
	defer f.Close()

	almanac, err := parseAlmanac(f)
	if err != nil {
		fmt.Println(err)

		return
	}

	locations := almanac.getLocationsForSeeds()
	fmt.Printf("locations: %v\n", locations)

	lowestValue := math.MaxInt64
	for _, location := range locations {
		if location < lowestValue {
			lowestValue = location
		}
	}

	fmt.Println("Lowest Value is: ", lowestValue)
}

func (a almanac) getLocationsForSeeds() []int {
	res := make([]int, 0, len(a.Seeds))
	for i := range a.Seeds {
		res = append(res, a.getLocationForSeed(a.Seeds[i]))
	}

	return res
}

func (a almanac) getLocationForSeed(seed int) int {
	soil := convertersToOutput(seed, a.SeedToSoil)
	fertilizer := convertersToOutput(soil, a.SoilToFertilizer)
	water := convertersToOutput(fertilizer, a.FertilizerToWater)
	light := convertersToOutput(water, a.WaterToLight)
	temperature := convertersToOutput(light, a.LightToTemperature)
	humidity := convertersToOutput(temperature, a.TemperatureToHumidity)
	return convertersToOutput(humidity, a.HumidityToLocation)
}

func convertersToOutput(input int, converters []converter) int {
	for _, converter := range converters {
		if res := converter.toOutput(input); res > 0 {
			return res
		}
	}

	return input // no mappings found return it unmapped
}

// toOutput converts the input to the respective output value.
// Will return 0 if the value was not found.
func (c converter) toOutput(in int) int {
	// not in range
	if in < c.sourceRangeStart || c.sourceRangeStart+c.rangeLength-1 < in {
		return 0
	}

	for i := 0; i < c.rangeLength; i++ {
		if (c.sourceRangeStart + i) == in {
			return c.destinationRangeStart + i
		}
	}

	return 0
}

// parseAlmanac parses the input file to our expected structure.
func parseAlmanac(in io.Reader) (almanac, error) {
	scanner := bufio.NewScanner(in)

	res := almanac{
		SeedToSoil:            make([]converter, 0),
		SoilToFertilizer:      make([]converter, 0),
		FertilizerToWater:     make([]converter, 0),
		WaterToLight:          make([]converter, 0),
		LightToTemperature:    make([]converter, 0),
		TemperatureToHumidity: make([]converter, 0),
		HumidityToLocation:    make([]converter, 0),
	}

	part := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			// skip empty lines
			continue
		}
		switch {
		case strings.HasPrefix(line, "seeds:"):
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				return res, fmt.Errorf("seeds line not correctly formatted: '%s'", line)
			}
			seeds, err := getValues(common.RemoveEmpty(parts[1]))
			if err != nil {
				return almanac{}, err
			}
			res.Seeds = seeds
		case strings.HasPrefix(line, "seed-to-soil"):
			part = 1
		case strings.HasPrefix(line, "soil-to-fertilizer"):
			part = 2
		case strings.HasPrefix(line, "fertilizer-to-water"):
			part = 3
		case strings.HasPrefix(line, "water-to-light"):
			part = 4
		case strings.HasPrefix(line, "light-to-temperature"):
			part = 5
		case strings.HasPrefix(line, "temperature-to-humidity"):
			part = 6
		case strings.HasPrefix(line, "humidity-to-location"):
			part = 7
		default:
			err := assignInputValues(&res, line, part)
			if err != nil {
				return res, err
			}
		}
	}

	return res, nil
}

func assignInputValues(in *almanac, line string, part int) error {
	conv, err := getConverterValue(line)
	if err != nil {
		return fmt.Errorf("line to converter mapping failed: %w", err)
	}

	switch part {
	case 1:
		in.SeedToSoil = append(in.SeedToSoil, conv)
	case 2:
		in.SoilToFertilizer = append(in.SoilToFertilizer, conv)
	case 3:
		in.FertilizerToWater = append(in.FertilizerToWater, conv)
	case 4:
		in.WaterToLight = append(in.WaterToLight, conv)
	case 5:
		in.LightToTemperature = append(in.LightToTemperature, conv)
	case 6:
		in.TemperatureToHumidity = append(in.TemperatureToHumidity, conv)
	case 7:
		in.HumidityToLocation = append(in.HumidityToLocation, conv)
	default:
		return fmt.Errorf("unexpected part number: %d", part)
	}

	return nil
}

func getConverterValue(line string) (converter, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return converter{}, fmt.Errorf("getConverterValue: line does not consist of 3 parts: '%s'", line)
	}

	destinationStartRange, err := strconv.Atoi(parts[0])
	if err != nil {
		return converter{}, fmt.Errorf("getConverterValue: %w", err)
	}

	sourceStartRange, err := strconv.Atoi(parts[1])
	if err != nil {
		return converter{}, fmt.Errorf("getConverterValue: %w", err)
	}

	rangeLength, err := strconv.Atoi(parts[2])
	if err != nil {
		return converter{}, fmt.Errorf("getConverterValue: %w", err)
	}

	return converter{
		sourceRangeStart:      sourceStartRange,
		destinationRangeStart: destinationStartRange,
		rangeLength:           rangeLength,
	}, nil
}

func getValues(line string) ([]int, error) {
	parts := strings.Split(line, " ")
	res := make([]int, 0, len(parts))
	for _, val := range parts {
		converted, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("convertion failed in line '%s': %w", line, err)
		}
		res = append(res, converted)
	}

	return res, nil
}
