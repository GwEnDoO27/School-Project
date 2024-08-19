package posts

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

func FormatImg(Path multipart.File, handler *multipart.FileHeader) (ImagePath string) {
	dst, err := os.Create(filepath.Join("./front/static/imgs/", filepath.Base(handler.Filename)))
	if err != nil {
		log.Println(err)
	}
	defer dst.Close()
	fmt.Println(dst)

	if _, err := io.Copy(dst, Path); err != nil {
		log.Println("Err Copy")
	}
	ImagePath = ("/static/imgs/" + handler.Filename)
	return ImagePath
}
