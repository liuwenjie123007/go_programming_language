// jpeg 命令从标准输入读入 PNG 图像
// 并发它作为 JPEG 图像写到标准输出
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // 注册 PNG解码器
	"io"
	"os"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jped: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) interface{} {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
