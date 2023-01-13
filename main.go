package main

import (
	"flag"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
)

//ColorGreen  Color = "\u001b[32m"
//ColorYellow Color = "\u001b[33m"
//ColorBlue   Color = "\u001b[34m"
//ColorBlack  Color = "\u001b[30m"

const (
	bSize      = 128 // размер буффера
	ColorRed   = "\u001b[31m"
	ColorReset = "\u001b[0m"
)

var (
	flagInput  string
	flagOutput string
	flagOffset int64
	flagLimit  int64
)

func init() {
	flag.StringVar(&flagInput, "input", "", "file read from")
	flag.StringVar(&flagOutput, "output", "", "file save to")
	flag.Int64Var(&flagOffset, "offset", 0, "file offset")
	flag.Int64Var(&flagLimit, "limit", 0, "file read limit")
}

// файл источник (From), файл копия (To), Отступ в источнике (Offset), по умолчанию - 0,
// Количество копируемых байт (Limit), по умолчанию - весь файл из From
func copyFile(from, to string, limit, offset int64) error {
	var bufSize, shift int64

	if limit == 0 || limit > bSize {
		bufSize = bSize
	} else {
		bufSize = limit
	}

	buf := make([]byte, bufSize)

	file, err := os.Open(from)
	if err != nil {
		return fmt.Errorf("cannot open file. Error: %v\n", err)
	}
	fsize, err := file.Stat()
	if err != nil {
		return fmt.Errorf("cannot open file. Error: %v\n", err)
	}

	if limit > fsize.Size()-offset {
		limit = fsize.Size() - offset
	}

	if offset > fsize.Size() {
		return fmt.Errorf("Offset is bigger than filesize!\n")
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	createFile, err := os.Create(to + ".txt")
	if err != nil {
		return fmt.Errorf("cannot create file. Error: %v\n", err)
	}
	defer func(createFile *os.File) {
		err := createFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(createFile)

	tmpl := ` {{string . "my_green_string" | green}} {{string . "my_blue_string" | blue}} {{percent .}} {{ bar . "[" "-" ">" "_" | green}} {{speed . | blue }}  {{etime .}}`
	bar := pb.ProgressBarTemplate(tmpl).Start64(limit / bSize)

	for shift < limit+offset {

		if shift == 0 {
			shift = offset
		}

		read, err := file.ReadAt(buf, shift)
		shift += int64(read)

		if err == io.EOF || read == 0 {
			fmt.Println("Nothing to copy. Change copy parameters")
			break
		}
		if err != nil {
			return fmt.Errorf("cannot Read file. Error: %v\n", err)
		}

		if shift > limit+offset {
			_, err = createFile.Write(buf[:bufSize-(shift-(limit+offset))])

		} else {
			bar.Increment()
			bar.Set("my_green_string", bar.Current()).Set("my_blue_string", bar.Total())
			_, err = createFile.Write(buf)
		}

		if err != nil {
			return fmt.Errorf("cannot write file. Error: %v\n", err)
		}
	}
	bar.Finish()
	return nil
}

func colorize(color string, message error) {
	fmt.Println(color, message, string(ColorReset))
}

func main() {
	flag.Parse()

	//err := copyFile("12.txt", "newfile", 1000, 100)
	err := copyFile(flagInput, flagOutput, flagLimit, flagOffset)
	if err != nil {
		colorize(ColorRed, err)
	}

}
