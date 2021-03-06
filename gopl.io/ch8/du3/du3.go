package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// walkDir 递归地遍历以dir为根目录的整个文件树
// 并在 fileSizes 上发送每个已找到的文件的大小
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema 是一个用于显著目录并发数的计数信号量
var sema = make(chan struct{}, 20)

// dirents 返回 dir 目录中的条目
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        //获取令牌
	defer func() { <-sema }() //释放令牌
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du2: %v\n", err)
		return nil
	}
	return entries
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	// 确定初始目录
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// 不并行遍历每一个文件树
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// 定期输出结果
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	// 输出结构
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes关闭
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDisUsage(nfiles, nbytes)
		}
	}
	printDisUsage(nfiles, nbytes)
}

func printDisUsage(nfiles int64, nbytes int64) {
	fmt.Printf("%d files %1.f GB\n", nfiles, float64(nbytes)/1e9)
}
