package day5

import (
	"fmt"
	"robertbrignull/adventofcode2023/shared"
	"strconv"
	"strings"
)

type Seeds []int

type SeedRange struct {
	start  int
	length int
}

type RangeMapEntry struct {
	destinationStart int
	sourceStart      int
	length           int
}

type RangeMap struct {
	entries []RangeMapEntry
}

// Looks up a value and returns:
//   - The mapped value
//   - How many other source values after this one also map in the same way
//     i.e. if x maps to y, then x + 1 maps to y + 1
func (rm RangeMap) lookup(source int) (int, int) {
	for _, entry := range rm.entries {
		if source >= entry.sourceStart && source < entry.sourceStart+entry.length {
			return entry.destinationStart + (source - entry.sourceStart), entry.length - (source - entry.sourceStart)
		}
	}

	nextSourceStart := -1
	for _, entry := range rm.entries {
		if source < entry.sourceStart && (nextSourceStart == -1 || nextSourceStart > entry.sourceStart) {
			nextSourceStart = entry.sourceStart
		}
	}
	if nextSourceStart == -1 {
		return source, 1000000000
	}
	return source, nextSourceStart - source - 1
}

type Almanac struct {
	seeds                 Seeds
	seedRanges            []SeedRange
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

func readSeedRanges(line string) ([]SeedRange, error) {
	fields := strings.Fields(line[len("seeds: "):])
	seedRanges := []SeedRange{}
	for i := 0; i < len(fields); i += 2 {
		start, err := strconv.Atoi(fields[i])
		if err != nil {
			return []SeedRange{}, err
		}

		length, err := strconv.Atoi(fields[i+1])
		if err != nil {
			return []SeedRange{}, err
		}

		seedRanges = append(seedRanges, SeedRange{start, length})
	}
	return seedRanges, nil
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

	seedRanges, err := readSeedRanges(lines[0])
	if err != nil {
		return Almanac{}, err
	}
	almanac.seedRanges = seedRanges

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

func computeSeedResult(seed int, almanac Almanac) (int, int) {
	soil, soilFV := almanac.seedSoilMap.lookup(seed)
	fertilizer, fertilizerFV := almanac.soilFertilizerMap.lookup(soil)
	water, waterFV := almanac.fertilizerWaterMap.lookup(fertilizer)
	light, lightFV := almanac.waterLightMap.lookup(water)
	temperature, temperatureFV := almanac.lightTemperatureMap.lookup(light)
	humidity, humidityFV := almanac.temperaturHumidityMap.lookup(temperature)
	location, locationFV := almanac.humidityLocationMap.lookup(humidity)
	return location, min(soilFV, fertilizerFV, waterFV, lightFV, temperatureFV, humidityFV, locationFV)
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

	lowestResult := -1
	for _, seed := range almanac.seeds {
		result, _ := computeSeedResult(seed, almanac)
		if lowestResult == -1 || result < lowestResult {
			lowestResult = result
		}
	}

	return strconv.Itoa(lowestResult), nil
}

// Original time taken to get answer: 15 minutes
// Execution time before optimization: 352 seconds
// Execution time after optimization: 0.2 seconds
func Part2() (string, error) {
	lines, err := shared.ReadFileLines("days/day5/input.txt")
	if err != nil {
		return "", err
	}

	almanac, err := readAlmanac(lines)
	if err != nil {
		return "", err
	}

	lowestResult := -1
	for _, seedRange := range almanac.seedRanges {
		for seed := seedRange.start; seed < seedRange.start+seedRange.length; seed++ {
			result, followingValues := computeSeedResult(seed, almanac)
			if lowestResult == -1 || result < lowestResult {
				lowestResult = result
			}

			// We're only looking for the lowest result, and the following values all map linearly
			// so they must be higher than the result we just computed and we can ignore them
			seed += followingValues
		}
	}

	return strconv.Itoa(lowestResult), nil
}
