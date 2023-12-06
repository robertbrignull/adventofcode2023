package day5

import (
	"fmt"
	"robertbrignull/adventofcode2023/shared"
	"strconv"
	"strings"
)

type Seeds []int

type RangeMapEntry struct {
	destinationStart int
	sourceStart      int
	length           int
}

type RangeMap struct {
	entries []RangeMapEntry
}

func (rm RangeMap) lookup(source int) int {
	for _, entry := range rm.entries {
		if source >= entry.sourceStart && source < entry.sourceStart+entry.length {
			return entry.destinationStart + (source - entry.sourceStart)
		}
	}
	return source
}

type Almanac struct {
	seeds                 Seeds
	seedSoilMap           RangeMap
	soilFertilizerMap     RangeMap
	fertilizerWaterMap    RangeMap
	waterLightMap         RangeMap
	lightTemperatureMap   RangeMap
	temperaturHumidityMap RangeMap
	humidityLocationMap   RangeMap
}

func readSeeds(line string) (Seeds, error) {
	seeds := []int{}
	for _, seedStr := range strings.Fields(line[len("seeds: "):]) {
		seed, err := strconv.Atoi(seedStr)
		if err != nil {
			return Seeds{}, err
		}
		seeds = append(seeds, seed)
	}
	return seeds, nil
}

func readRangeMapEntry(line string) (RangeMapEntry, error) {
	var entry RangeMapEntry
	_, err := fmt.Sscanf(line, "%d %d %d", &entry.destinationStart, &entry.sourceStart, &entry.length)
	return entry, err
}

func readRangeMap(lines []string) (RangeMap, error) {
	var rangeMap RangeMap
	for _, line := range lines {
		if line == "" {
			break
		}

		entry, err := readRangeMapEntry(line)
		if err != nil {
			return RangeMap{}, err
		}

		rangeMap.entries = append(rangeMap.entries, entry)
	}
	return rangeMap, nil
}

func readAlmanac(lines []string) (Almanac, error) {
	var almanac Almanac

	seeds, err := readSeeds(lines[0])
	if err != nil {
		return Almanac{}, err
	}
	almanac.seeds = seeds

	row := 2
	for {
		if row >= len(lines) {
			break
		}

		mapName := lines[row]
		rangeMap, err := readRangeMap(lines[row+1:])
		if err != nil {
			return Almanac{}, err
		}

		if mapName == "seed-to-soil map:" {
			almanac.seedSoilMap = rangeMap
		} else if mapName == "soil-to-fertilizer map:" {
			almanac.soilFertilizerMap = rangeMap
		} else if mapName == "fertilizer-to-water map:" {
			almanac.fertilizerWaterMap = rangeMap
		} else if mapName == "water-to-light map:" {
			almanac.waterLightMap = rangeMap
		} else if mapName == "light-to-temperature map:" {
			almanac.lightTemperatureMap = rangeMap
		} else if mapName == "temperature-to-humidity map:" {
			almanac.temperaturHumidityMap = rangeMap
		} else if mapName == "humidity-to-location map:" {
			almanac.humidityLocationMap = rangeMap
		} else {
			return Almanac{}, fmt.Errorf("unknown map name: %s", mapName)
		}

		row += 2 + len(rangeMap.entries)
	}

	return almanac, nil
}

func computeSeedResult(seed int, almanac Almanac) int {
	soil := almanac.seedSoilMap.lookup(seed)
	fertilizer := almanac.soilFertilizerMap.lookup(soil)
	water := almanac.fertilizerWaterMap.lookup(fertilizer)
	light := almanac.waterLightMap.lookup(water)
	temperature := almanac.lightTemperatureMap.lookup(light)
	humidity := almanac.temperaturHumidityMap.lookup(temperature)
	location := almanac.humidityLocationMap.lookup(humidity)
	return location
}

// Time taken: 25 minutes
func Part1() (string, error) {
	lines, err := shared.ReadFileLines("days/day5/input.txt")
	if err != nil {
		return "", err
	}

	almanac, err := readAlmanac(lines)
	if err != nil {
		return "", err
	}

	lowestSeedResult := -1
	for _, seed := range almanac.seeds {
		seedResult := computeSeedResult(seed, almanac)
		if lowestSeedResult == -1 || seedResult < lowestSeedResult {
			lowestSeedResult = seedResult
		}
	}

	return strconv.Itoa(lowestSeedResult), nil
}
