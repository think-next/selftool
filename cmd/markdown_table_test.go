package cmd

import (
	"fmt"
	"github.com/selftool/common/file"
	"github.com/selftool/common/markdown_table"
	"sort"
	"testing"
)

func TestSrc(t *testing.T) {
	srcFile := "/Users/huifu/Desktop/srctable.md"
	dstFile := "/Users/huifu/Desktop/dsttable.md"

	linePipeline := make(chan string, 100)
	go func() {
		if err := markdown_table.SplitFile(srcFile, linePipeline, "|"); err != nil {
			panic(err)
		}
	}()

	transitionResult := make([]string, 0, 3000)
	for line := range linePipeline {
		transitionResult = append(transitionResult, line)
	}
	sort.Strings(transitionResult)

	fmt.Println(len(transitionResult))

	fileUtil := file.UtilFile{}
	if err := fileUtil.WriteLineBySlice(dstFile, transitionResult); err != nil {
		panic(err)
	}
}
