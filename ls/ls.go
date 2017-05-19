package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	//Step:1 -- Trying to get the count of files available in the current
	//			directory where the command is executed
	//Step:2 -- Trying to ignore the hidden files/directories and getting the
	//			count of files that are available in the folder

	var _hiddenCount int
	var _visibleCount int

	_hidden := flag.Bool("a", false, "Will return the files which are Hidden")
	flag.Parse()

	var (
		_hideNames    []string
		_visibleNames []string
		_total        []string
	)
	read, _ := ioutil.ReadDir("./")

	for _, file := range read {
		if strings.HasPrefix(file.Name(), ".") {
			_hideNames = append(_hideNames, file.Name())
			_hiddenCount++
		} else {
			_visibleNames = append(_visibleNames, file.Name())
			_visibleCount++
		}
	}

	if *_hidden == true {
		_sum := _hiddenCount + _visibleCount
		for _, _file := range _hideNames {
			_total = append(_total, _file)
		}
		for _, _file := range _visibleNames {
			_total = append(_total, _file)
		}
		fmt.Printf("total: %d\n", _sum)
		fmt.Println(_total)
	} else {
		for _, _file := range _visibleNames {
			_total = append(_total, _file)
		}
		fmt.Printf("total: %d\n", _visibleCount)
		fmt.Println(_total)
	}

}
