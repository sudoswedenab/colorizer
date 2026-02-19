// Copyright 2026 Sudo Sweden AB
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	RESET  = "\033[0m"
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		
		color := RESET
		switch {
		case strings.Contains(line, "ERROR"): color = RED
		case strings.Contains(line, "FAIL"):  color = RED
		case strings.Contains(line, "PASS"):  color = GREEN
		case strings.Contains(line, "WARN"):  color = YELLOW
		case strings.Contains(line, "DEBUG"): color = BLUE
		case strings.Contains(line, "RUN"):   color = BLUE
		}

		fmt.Printf("%s%s\n", color, line)
	}
}
