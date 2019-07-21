// thumbnail 包提供缩略图功能，现今暂时只支持JPEG格式
package thumbnail

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Image函数返回源图片的缩略图版本
func Image(src image.Image) image.Image {
	// 计算缩略图大小，保持纵横比
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect)
	} else {
		height = int(128 / aspect)
	}
	xScale := float64(xs) / float64(width)
	yScale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// 粗略的缩放算法
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcX := int(float64(x) * xScale)
			srcY := int(float64(y) * yScale)
			dst.Set(x, y, src.At(srcX, srcY))
		}
	}
	return dst
}

// ImageStream函数读取源图片并返回缩略图版本
func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(w, dst, nil)
}

func ImageFile2(outfile, infile string) error {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer func() {
		_ = in.Close()
	}()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	if err := ImageStream(out, in); err != nil {
		_ = out.Close()
		return fmt.Errorf("scaling %s to %s: %s", infile, outfile, err)
	}
	return out.Close()
}

// ImageFile从infile中读取一幅画像并把它的缩略图写入同一个目录中
// 返回生成的文件名，比如"foo.thumb.jpg"
func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile) // e.g., ".jpg", ".JPEG"
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}
