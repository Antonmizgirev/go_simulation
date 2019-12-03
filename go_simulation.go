// go_simulation
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type city struct {
	links map[string]string
	alien int
}

func SetAliens(m map[string]city, numberOfAliens int) (map[string]city, map[int]string) {
	//fill the array with random numbers positions for aliens
	numberOfCities := len(m)
	set := map[int]bool{}
	r := 0
	rand.Seed(time.Now().UTC().UnixNano())
	for a := 0; a < numberOfAliens; a++ {
		ok := true
		for ok == true {
			r = rand.Intn(numberOfCities)
			_, okk := set[r] // check for existence
			ok = okk
		}
		set[r] = true // add element
	}

	i := 0
	n := 1
	aliensLocations := map[int]string{}
	//put aliens to cities
	for k, _ := range m {
		_, ok := set[i]
		if ok {
			val := m[k]
			val.alien = n
			aliensLocations[n] = k
			m[k] = val
			delete(set, i) // remove element
			n++
		}
		i++
	}

	return m, aliensLocations
}

func Step(m map[string]city, aliensLocations map[int]string) (map[string]city, map[int]string) {
	for alienId, location := range aliensLocations {
		//delete link to not existing cities

		helper := m[location]
		for key, val := range helper.links {
			_, ok := m[val]
			if ok == false {
				// if not city not exist than delete it from the map of linked cites
				delete(helper.links, key)
			}
		}
		m[location] = helper
		//for each alien move to random linked city
		if len(m[location].links) == 0 {
			//no way out of this city I assume that city destroyed and alien also
			delete(m, location)
			delete(aliensLocations, alienId)
			fmt.Println("Alien destroyed city", location, " and died because there is no way out")
		} else {

			//get random
			r := rand.Intn(len(m[location].links))
			moveTo := ""
			i := 0
			for _, val := range m[location].links {
				if i == r {
					moveTo = val
					break
				}
				i += 1
			}

			//destroy the city where he was
			delete(m, location)
			//check if the city where he arrived is busy
			if m[moveTo].alien != 0 {
				//destroy everything
				fmt.Println("Alien", alienId, "destroy the city", location, "move to the city", moveTo, ",fight with alien", m[moveTo].alien, "and destroy the city", moveTo)
				delete(aliensLocations, m[moveTo].alien)
				delete(aliensLocations, alienId)
				delete(m, moveTo)
			} else {
				fmt.Println("Alien", alienId, "destroyed the city", location, "and moved to", moveTo)
				tt := m[moveTo]
				tt.alien = alienId
				m[moveTo] = tt
				aliensLocations[alienId] = moveTo
			}
		}

	}
	return m, aliensLocations
}

func main() {
	flag.Parse()
	s := flag.Arg(0)
	//parse argument
	numberOfAliens, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	//open file with map
	file, err := os.Open("input_alien_simulation.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//map from city to name and links
	m := map[string]city{}
	//scan file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//map for input
		temp := scanner.Text()
		words := strings.Fields(temp)
		road_map := map[string]string{}
		for i, w := range words {
			if i != 0 {
				road_map[w[:strings.IndexByte(w, '=')]] = w[strings.IndexByte(w, '=')+1:]
			}

		}
		m[words[0]] = city{road_map, 0}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//set_initial_random_aliens
	m, aliensLocations := SetAliens(m, numberOfAliens)

	//while aliens exist
	steps_made := 0
	for len(aliensLocations) > 0 && steps_made <= 10000 {
		m, aliensLocations = Step(m, aliensLocations)

	}

	result := ""
	for key, val := range m {
		result += key + " "
		for k, v := range val.links {
			result += k + "=" + v + " "
		}
		result += "\n"
	}
	fo, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(result))
	if err != nil {
		log.Fatal(err)
	}
}
