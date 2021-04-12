package cmd

import (
	"fmt"
	"github.com/selftool/common/file"
	"github.com/selftool/common/markdown_table"
	"github.com/selftool/common/uniq"
	"sort"
	"strings"
	"sync"
	"testing"
	"unsafe"
)

func TestSrc(t *testing.T) {
	srcFile := "/Users/huifu/Desktop/srctable.md"
	dstFile := "/Users/huifu/Desktop/dsttable.md"

	linePipeline := make(chan []string, 100)
	go func() {
		if err := markdown_table.SplitFile(srcFile, linePipeline, "|"); err != nil {
			panic(err)
		}
	}()

	transitionResult := make([]string, 0, 3000)
	uniqReq := make([]string, 0, 3000)
	for line := range linePipeline {

		/*
			默认使用第一次的脚本，使用第一个字段做唯一区分
		*/

		if len(line) == 1 {
			continue
		}

		handleStr := strings.TrimSpace(line[1])
		if uniq.IsDuplicateForStrings(uniqReq, handleStr) {
			fmt.Println(line[1])
			continue
		}

		uniqReq = append(uniqReq, handleStr)

		transitionResult = append(transitionResult, strings.Join(line, markdown_table.DefaultTableJoinSeparator))
	}
	sort.Strings(transitionResult)

	fileUtil := file.UtilFile{}
	if err := fileUtil.WriteLineBySlice(dstFile, transitionResult); err != nil {
		panic(err)
	}
}

func TestParse(t *testing.T) {
	var sm sync.Map
	sm.Store(1, "first")
	if v, ok := sm.Load(1); ok {
		fmt.Println(v)
	}
}

//
func TestMap(t *testing.T) {
	type Person struct {
		Age int
	}

	var m = make(map[int]*Person)
	m[1] = &Person{Age: 0}
	m[1].Age = 10

	fmt.Println(*m[1])
}

func TestPanic(t *testing.T) {
	var result = make(map[int]int)

	// 1
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("=====1")
		}
	}()

	go func() {
		// 2
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("=====2")
			}
		}()

		for {
			_ = result[1]
		}
	}()

	for {
		result[1] = 1
	}
}

// 8 bit
func TestPadi(t *testing.T) {
	type Person struct {
		Name string // 16
		Age  byte   // 1
		Sex  byte   // 1
	}

	size := unsafe.Sizeof(Person{})
	fmt.Println(size)
}
