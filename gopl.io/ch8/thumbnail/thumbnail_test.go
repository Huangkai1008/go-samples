package thumbnail_test

import (
	"log"
	"os"
	"sync"

	"go-starter/gopl.io/ch8/thumbnail"
)

// makeThumbnails生成指定文件的缩略图
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// NOTE: 错误并行版本!
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

// makeThumbnails3并行生成指定文件的缩略图
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			_, _ = thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}

	//	等待goroutine完成
	for range filenames {
		<-ch
	}
}

// Note: 存在goroutine泄露的风险
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
			return err
		}
	}
	return nil
}

// makeThumbnails5为指定文件并行地生成缩略图
// 它以任意顺序返回生成的文件名
// 如果任何步骤错误就返回一个错误
func makeThumbnails5(filenames []string) (thumbFiles []string, err error) {
	type item struct {
		thumbFile string
		err       error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbFile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, err
		}
		thumbFiles = append(thumbFiles, it.thumbFile)
	}
	return thumbFiles, nil
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // 工作goroutine的个数
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
			info, _ := os.Stat(thumb)
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
