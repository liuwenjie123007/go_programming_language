package thumbnail_test

import (
	"gopl.io/ch8/thumbnail"
	"log"
	"os"
	"sync"
)

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// 注意： 不正确
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // 注意： 忽略错误
	}
}

// makeThumbnails3 并行生成指定文件的缩略图
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f) // 注意： 此处忽略了可能的错误
			ch <- struct{}{}
		}(f)
	}
	// 等待goroutine完成
	for range filenames {
		<-ch
	}
}

// makeThumbnails4 为指定文件并行地生成缩略图
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			return err // 注意: 不正确，goroutine泄露
		}
	}
	return nil
}

// makeThumbnails6 为从通道接收到的每个文件生成缩略图
// 它返回其生成文件占用的字节数
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // 工作 goroutine 的个数
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // 可以忽略错误
			sizes <- info.Size()
		}(f)
	}
	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
