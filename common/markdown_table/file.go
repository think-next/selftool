package markdown_table

import (
	"github.com/selftool/common/file"
	"strings"
)

/*
输入文件的格式，默认使用|进行分割，可以指定分割符号
*/

const (
	DefaultTableSplitSeparator = ","
	DefaultTableJoinSeparator  = "|"
	DefaultReadFileBuffer      = 100
)

/*
	将文件进行拆分，每行数据拆分成一个数组结构
	chan 的元素是一个数组结构, 每个元素表示的是一个表中的字段
	separators 维护1个元素，表示元文件的分割符号
*/
func SplitFile(path string, pipeLine chan<- []string, separators ...string) error {
	splitSeparator := DefaultTableSplitSeparator
	if len(separators) != 0 {
		splitSeparator = separators[0]
	}

	// 判断文件是否存在
	fileUtil := file.UtilFile{}
	if isOk, err := fileUtil.IsFile(path); !isOk {
		return err
	}

	// 读取文件内容
	srcLines := make(chan string, DefaultReadFileBuffer)
	utilFile := file.UtilFile{}
	go func() {
		if err := utilFile.ReadLine(path, srcLines); err != nil {
			panic(err)
		}
	}()

	for line := range srcLines {
		pipeLine <- strings.Split(line, splitSeparator)
	}

	close(pipeLine)
	return nil
}
