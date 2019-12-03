// go_simulation_test
package main

import (
	"testing"
)

func TestSetAliens(t *testing.T) {
	m := map[string]city{}
	m["Bar"] = city{map[string]string{"north": "Foo"}, 0}
	m["Foo"] = city{map[string]string{"north": "Bar", "south": "B", "east": "C"}, 0}
	m["C"] = city{map[string]string{"south": "Bar", "east": "Foo"}, 0}
	number_of_aliens := 2
	aliensLocations := map[int]string{}
	m, aliensLocations = SetAliens(m, number_of_aliens)
	if number_of_aliens != len(aliensLocations) {
		t.Errorf("Set aliens function does not work")
	}
}

func TestStep(t *testing.T) {
	m := map[string]city{}
	m["Bar"] = city{map[string]string{"north": "Foo", "south": "C"}, 1}
	m["Foo"] = city{map[string]string{"east": "Bar", "south": "B", "west": "C"}, 0}
	m["C"] = city{map[string]string{"west": "Bar", "south": "Foo"}, 0}
	m["B"] = city{map[string]string{"east": "Foo"}, 0}

	aliensLocations := map[int]string{1: "Bar"}
	m, aliensLocations = Step(m, aliensLocations)
	if aliensLocations[1] != "Foo" && aliensLocations[1] != "C" {
		t.Errorf("Step function does not work")
	}
}
