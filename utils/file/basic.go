package file

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
)

func Copy(src, dst string) error {
	//srcStat, err := os.Stat(src)
	//if err != nil {
	//	return err
	//}

	//if srcStat.Mode().IsRegular() {
	//	return fmt.Errorf("%s is not a regular file", src)
	//}

	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	desFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer desFile.Close()

	_, err = io.Copy(desFile, sourceFile)
	return err

}

/**
可以删除文件或者空目录
*/
func Delete(path string) error {
	return os.Remove(path)
}

/**
可以删除目录
*/
func DeleteAll(path string) error {
	return os.RemoveAll(path)
}

func Move(src, dst string) error {
	return os.Rename(src, dst)
}

func Create(path string) (*os.File, error) {
	return os.Create(path)
}

func Backup() {

}

func Mkdir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

func MkdirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)

}

func ContentMd5(path string) (string, error) {
	ctx := md5.New()
	content, err := ReadFile(path)
	if err != nil {
		return "", err
	}
	ctx.Write(content)
	return hex.EncodeToString(ctx.Sum(nil)), nil
}

func StrMd5(content string) string {
	ctx := md5.New()
	ctx.Write([]byte(content))
	return hex.EncodeToString(ctx.Sum(nil))
}
