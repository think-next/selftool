package markdown_table

import (
	"github.com/selftool/common/file"
)

/*
输入文件的格式，默认使用|进行分割，可以指定分割符号
*/

const (
	DefaultTableSeparator = "|"
)

/*
	将文件进行拆分，每行数据拆分成一个数组结构
*/
func SplitFile(path string, separators ...string) (result []string, err error) {
	separator := DefaultTableSeparator
	if len(separators) != 0 {
		separator = separators[0]
	}

	// 判断文件是否存在
	fileUtil := file.UtilFile{}
	if isOk, err := fileUtil.IsFile(path); !isOk {
		return result, err
	}

	// 读取文件内容

}
