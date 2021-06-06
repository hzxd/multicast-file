package file

import (
	"flag"
	"fmt"
	"math"
	"os"
	"time"
)

var BufferSize = 1024 * 1024 * 10

func GetPathFromArgs() string {
	path := flag.Args()[len(flag.Args())-1]
	return path
}

func SliceFile(file *os.File, blockSize uint, index int) ([]byte, error) {
	b := make([]byte, blockSize, blockSize+2)

	offset := len(b) * index
	n, err := file.ReadAt(b, int64(offset))
	return b[:n], err
}

func ConcatFile(filename string, fileMap map[int64]string, block int, serverAddr string) error {
	file, err := os.Create("/tmp/" + filename)
	if err != nil {
		return err
	}
	for i := 0; i < block; i++ {
		f := make([]byte, BufferSize, BufferSize+2)
		if fname, ok := fileMap[int64(i)]; !ok {
			f = GetDataFromServer(serverAddr, int64(i))
		} else {
			fd, _ := os.Open(fname)
			n, _ := fd.Read(f)
			f = f[:n]
		}
		_, err := file.WriteAt(f, int64(i*BufferSize))
		if err != nil {
			return err
		}

	}
	return nil
}

func SendFile(path string, ch chan *FileMessage, stop chan bool) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	if stat, ok := file.Stat(); ok == nil {
		num := int(math.Floor(float64(stat.Size()) / float64(BufferSize)))
		fmt.Printf("num:%d", num)
		ch <- &FileMessage{
			Filename: stat.Name(),
			Block:    int64(num),
			Type:     1,
		}

		for i := 0; i < num; i++ {
			buf, err := SliceFile(file, uint(BufferSize), i)
			if err != nil {
				return err
			}
			ch <- &FileMessage{
				Index:    int64(i),
				Buf:      buf,
				Filename: stat.Name(),
				Block:    int64(num),
			}
		}
		for len(ch) > 0 {
			time.Sleep(time.Millisecond * 10)
		}
		ch <- &FileMessage{
			Filename: stat.Name(),
			Block:    int64(num),
			Type:     2,
		}
		for len(ch) > 0 {
			time.Sleep(time.Millisecond * 10)
		}
		close(stop)
	}
	return nil
}
