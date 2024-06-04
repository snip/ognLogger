/*
    ognLogger / OGN APRS messages logger written in Golang
    Copyright (C) 2024  Sebastien Chaumontet

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

*/

package main

import (
	"fmt"
	"os"
	"log"
	"time"
)
var lastDate string

func main() {
	fmt.Println(softwareName+" "+softwareVersion+" Copyright (C) 2024 Sebastien Chaumontet - https://github.com/snip/ognLogger")
	fmt.Println("This program comes with ABSOLUTELY NO WARRANTY.")
	fmt.Println("This is free software, and you are welcome to redistribute it under GPL v3.0 conditions.\n")
	Listen(process_message)
}

func process_message(aprs_line string) {
	currentTime := time.Now()
	currentDay := currentTime.Format("2006-01-02")
	//fmt.Print(aprs_line)
	fmt.Print(".")
	f, err := os.OpenFile(currentDay+"-aprs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(aprs_line);
	f.Close()
}

