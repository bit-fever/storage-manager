//=============================================================================
/*
Copyright © 2025 Andrea Carboni andrea.carboni71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
//=============================================================================

package backend

import (
	"github.com/bit-fever/core"
	"github.com/bit-fever/storage-manager/pkg/app"
	"os"
	"path/filepath"
	"strconv"
)

//=============================================================================

const (
	Doc    = "doc"
	Code   = "code"
	Temp   = "temp"
	Image  = "image"
	Report = "report"

	EquityChart = "equity-chart.png"
)

var Dirs = []string{ Doc, Code, Temp, Image, Report }

//=============================================================================

var folder         string
var defEquityChart []byte

//=============================================================================
//===
//=== Init functions
//===
//=============================================================================

func InitStorage(cfg *app.Config) {
	folder = cfg.Storage.Folder

	err := os.MkdirAll(folder, 0700)
	core.ExitIfError(err)

	defEquityChart, err = os.ReadFile("default/"+ EquityChart)
	core.ExitIfError(err)
}

//=============================================================================
//===
//=== Public functions
//===
//=============================================================================

func AddTradingSystem(id uint, username string) error {
	sId := strconv.Itoa(int(id))
	path:= folder +"/"+ username +"/"+ sId +"/"

	for _, dir := range Dirs {
		err := os.MkdirAll(path + dir, 0700)
		if err != nil {
			return err
		}
	}

	return nil
}

//=============================================================================

func DeleteTradingSystem(id uint, username string) error {
	sId := strconv.Itoa(int(id))
	return os.RemoveAll(folder +"/"+ username +"/"+ sId)
}

//=============================================================================

func ReadEquityChart(username string, id uint) ([]byte,error) {
	path := []string{
		folder,
		username,
		strconv.Itoa(int(id)),
		Image,
		EquityChart,
	}

	return readFile(path...)
}

//=============================================================================

func WriteEquityChart(username string, id uint, data []byte) error {
	path := []string{
		folder,
		username,
		strconv.Itoa(int(id)),
		Image,
		EquityChart,
	}

	return writeFile(data, path...)
}

//=============================================================================

func GetDefaultEquityChart() []byte {
	return defEquityChart
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

func readFile(path ...string) ([]byte, error) {
	file := filepath.Join(path...)
	return os.ReadFile(file)
}

//=============================================================================

func writeFile(data []byte, path ...string) error {
	file := filepath.Join(path...)
	err  := os.WriteFile(file +".temp", data, 0600)

	if err != nil {
		return err
	}

	_,err = os.Stat(file)

	if err == nil {
		err = os.Remove(file)
		if err != nil {
			return err
		}
	}

	return os.Rename(file +".temp", file)
}

//=============================================================================
