import os

class Map:
    def __init__(self, destinationStart, sourceStart, length):
        self.destinationStart = destinationStart
        self.sourceStart = sourceStart
        self.length = length

def parseInput(filename):
    # Open the file
    with open(filename, 'r') as file:
        lines = file.readlines()

    maps = {}
    seeds = []

    # Loop through each line in the file
    for line in lines:
        # Check if the line contains seeds
        if "seeds:" in line:
            seeds = [int(x) for x in line.split()[1:]]
        # Check if the line contains a map
        elif "map:" in line:
            mapName = line.split()[0]
            mapData = []
            for line in lines[lines.index(line)+1:]:
                if line == "\n":
                    break
                nums = [int(x) for x in line.split()]
                mapData.append(Map(nums[0], nums[1], nums[2]))
            maps[mapName] = mapData

    return seeds, maps

def createMap(maps, num):
    # Loop through each map
    for m in maps:
        # Check if the number is within the map's range
        if num >= m.sourceStart and num < m.sourceStart + m.length:
            return m.destinationStart + (num - m.sourceStart)
    return num

def findLowestLocation(seeds, maps):
    lowestLocation = -1
    # Loop through each seed
    for seed in seeds:
        soil = createMap(maps["seed-to-soil"], seed)
        fertilizer = createMap(maps["soil-to-fertilizer"], soil)
        water = createMap(maps["fertilizer-to-water"], fertilizer)
        light = createMap(maps["water-to-light"], water)
        temperature = createMap(maps["light-to-temperature"], light)
        humidity = createMap(maps["temperature-to-humidity"], temperature)
        location = createMap(maps["humidity-to-location"], humidity)
        # Check if the location is the lowest found so far
        if lowestLocation == -1 or location < lowestLocation:
            lowestLocation = location
    return lowestLocation

def main():
    seeds, maps = parseInput("input.txt")
    print(findLowestLocation(seeds, maps))

if __name__ == "__main__":
    main()