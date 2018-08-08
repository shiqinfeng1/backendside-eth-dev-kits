package httpservice

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/labstack/echo"
	"github.com/nicksnyder/go-i18n/i18n"
	"golang.org/x/net/context"
	"golang.org/x/net/webdav"
)

var (
	defaultLang = "en-US"
	defaultFunc i18n.TranslateFunc
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// Init i18n initialize
func Init() {
	loadLanguage("en-US.all.yaml", defaultLang)
	loadLanguage("zh-CN.all.yaml", "zh-CN")
	defaultFunc, _ = i18n.Tfunc(defaultLang)
}

// HTTPFS implements http.FileSystem
type HTTPFS struct{}

// FileEnUSAllYaml is "en-US.all.yaml"
var FileEnUSAllYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8d\x41\x0a\xc2\x40\x0c\x45\xf7\x39\xc5\xbf\x80\x90\x2e\xed\x29\xc4\x0b\x48\x68\x3f\x1a\x98\x76\x34\xcd\x80\x78\x7a\xd1\x9d\x4e\xb7\x0f\xde\x7b\x07\xf8\x3c\x82\x11\x53\x9d\x79\x19\x54\x07\x01\x32\x6c\xdd\x8a\xa5\xd7\x75\xc4\xc9\xc2\x16\x26\x03\x7c\x4e\xbc\x7f\xa0\x00\xf2\x2b\x1e\x55\xf5\x5f\x3c\xf3\xd1\xb8\x25\xac\xe5\xad\x86\xbf\xbe\x18\xe9\x0b\x6b\x4b\xe9\x03\xdd\x79\x3f\xe0\xa5\xf0\x6a\x45\xde\x01\x00\x00\xff\xff\xc1\x81\x44\x4a\xba\x00\x00\x00")

// FileZhCNAllYaml is "zh-CN.all.yaml"
var FileZhCNAllYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x55\xc8\x4c\xb1\x52\x48\x2d\x2a\x4a\xce\x4f\x49\x8d\x37\x34\x30\x30\xe4\x52\x50\x28\x29\x4a\xcc\x2b\xce\x49\x2c\xc9\xcc\xcf\xb3\x52\x78\xb1\xa1\xf9\xf9\x94\x15\x4f\xfb\x9b\x9e\x4d\xdd\xf0\x74\x4f\xd3\xd3\x1d\x3b\xb8\x14\x14\xb8\x50\xf5\x59\x1a\x18\x18\x60\xe8\x5b\xbf\xfd\xd9\xc6\xa6\x97\x9d\x5b\x9e\xcd\x6d\x7e\xb1\xad\xf5\xd9\xf4\x6d\x5c\x98\xba\x30\x6d\x43\xd2\xf5\x72\xee\xbc\x67\x9b\xa7\x72\x01\x02\x00\x00\xff\xff\x78\x67\x5b\xfd\xa3\x00\x00\x00")

func init() {
	if CTX.Err() != nil {
		log.Fatal(CTX.Err())
	}

	var err error

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileEnUSAllYaml)
	r, err = gzip.NewReader(rb)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err = FS.OpenFile(CTX, "en-US.all.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	rb = bytes.NewReader(FileZhCNAllYaml)
	r, err = gzip.NewReader(rb)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err = FS.OpenFile(CTX, "zh-CN.all.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}

func loadLanguage(filename, lang string) {
	fileBytes, err := ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = i18n.ParseTranslationFileBytes(filename, fileBytes)
	if err != nil {
		panic(err)
	}
}

// Locate 获取对应语言的翻译方法
func Locate(lang string) i18n.TranslateFunc {
	tfunc, err := i18n.Tfunc(lang)
	if err != nil || tfunc == nil {
		return defaultFunc
	}
	return tfunc
}

// TLang 返回绑定 accept-language 的i18n方法
func TLang(c echo.Context) i18n.TranslateFunc {
	return Locate(GetAcceptLanguage(c))
}
