package file

import (
	"bufio"
	"io"
	"os"
)

const LineSeparator = '\n'

type UtilFile struct {
}

// 判断文件是否存在
func (util *UtilFile) IsExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断是否为目录文件
func (util *UtilFile) IsDir(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return stat.IsDir(), nil
}

// 判断是否为文件
func (util *UtilFile) IsFile(path string) (bool, error) {
	isDir, err := util.IsDir(path)
	if err != nil {
		return false, err
	}

	return !isDir, err
}

// 通过channel来逐行读取文件
func (util *UtilFile) ReadLine(path string, pipeline chan<- string) error {
	fi, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		line, err := br.ReadString(LineSeparator)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		pipeline <- line
	}

	close(pipeline)
	return nil
}

// 通过channel来逐行读取文件
func (util *UtilFile) WriteLineByChannel(path string, pipeline <-chan string) error {
	fi, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer fi.Close()

	for v := range pipeline {
		_, err := fi.WriteString(v)
		if err != nil {
			return err
		}
	}

	return nil
}

// 通过channel来逐行读取文件
func (util *UtilFile) WriteLineBySlice(path string, pipeline []string) error {
	fi, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer fi.Close()

	for _, v := range pipeline {
		_, err := fi.WriteString(v)
		if err != nil {
			return err
		}
	}

	return nil
}
