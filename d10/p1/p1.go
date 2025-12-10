//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type state bool

const (
	off state = false
	on  state = true
)

func (s *state) toggle() {
	if *s == on {
		*s = off
	} else {
		*s = on
	}
}

type lights []state

func (l lights) equal(other lights) bool {
	return slices.Equal(l, other)
}

type button []int

func (l lights) press(button button) lights {
	result := make(lights, len(l))
	copy(result, l)
	for _, b := range button {
		result[b].toggle()
	}
	return result
}

type machine struct {
	buttons      []button
	targetLights lights
}

func parseMachines(lines []string) ([]machine, error) {

	machines := make([]machine, 0, len(lines))
	for _, line := range lines {
		// lights
		start := strings.Index(line, "[")
		end := strings.Index(line, "]")
		lights := make(lights, 0, end-start-1)
		for j := start + 1; j < end; j++ {
			switch line[j] {
			case '.':
				lights = append(lights, off)
			case '#':
				lights = append(lights, on)
			}
		}

		// buttons
		parts := strings.Split(line, "(")[1:]
		buttons := make([]button, 0, len(parts))
		for _, part := range parts {
			end := strings.Index(part, ")")
			cleaned := part[:end]
			nums := strings.Split(cleaned, ",")
			button := make(button, 0, len(nums))
			for _, num := range nums {
				numInt, err := strconv.Atoi(num)
				if err != nil {
					return nil, err
				}
				button = append(button, numInt)
			}
			buttons = append(buttons, button)
		}

		machines = append(machines, machine{
			targetLights: lights,
			buttons:      buttons,
		})

	}
	return machines, nil
}

func hashLights(l lights) string {
	return fmt.Sprint(l)
}

func (m *machine) PressesToTurnOn() int {
	type State struct {
		lights  lights
		presses int
	}

	startLights := make(lights, len(m.targetLights))
	queue := []State{{lights: startLights, presses: 0}}

	checked := make(map[string]bool)
	var state State
	for {
		state, queue = queue[0], queue[1:]

		for _, button := range m.buttons {
			newLights := state.lights.press(button)
			hashed := hashLights(newLights)
			if chk, _ := checked[hashed]; chk {
				continue
			}
			if newLights.equal(m.targetLights) {
				return state.presses + 1
			}

			newState := State{lights: newLights, presses: state.presses + 1}
			queue = append(queue, newState)
			checked[hashed] = true
		}
	}
}

func P1(input []string) string {
	sum := 0
	machines, err := parseMachines(input)
	if err != nil {
		return ""
	}
	for _, machine := range machines {
		sum += machine.PressesToTurnOn()
	}
	return strconv.Itoa(sum)
}
