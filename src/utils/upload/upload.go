package upload

import (
	"image"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
)

func Mkdir(filename string) (err error) {
	path := path.Dir(filename)
	//fmt.Println(path)
	_, err = os.Stat(path)
	if err != nil {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return
		}
	}
	return
}
func PutBytes(filename string, data []byte) (err error) {
	err = Mkdir(filename)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(filename, data, os.ModePerm) //写入文件(字节数组)
	return
}

func PutFileBytes(data []byte, extname string, stype string) (remoteName string, err error) {
	rand.Seed(time.Now().UnixNano())
	remoteName = stype + "/" + time.Now().Format("20060102") + "/" + strconv.FormatInt(int64(time.Now().Unix()), 10) + strconv.FormatInt(int64(rand.Intn(999999999)), 10) + extname
	objName := "uploads/" + remoteName
	err = PutBytes(objName, data)
	if err != nil {
		return "", err
	}
	return
}
func PutPictureImage(srcImage image.Image, extname string, stype string) (remoteName string, err error) {
	rand.Seed(time.Now().UnixNano())
	remoteName = stype + "/" + time.Now().Format("20060102") + "/" + strconv.FormatInt(int64(time.Now().Unix()), 10) + strconv.FormatInt(int64(rand.Intn(999999999)), 10) + extname
	objName := "uploads/" + remoteName
	err = Mkdir(objName)
	if err != nil {
		return "", err
	}
	dst := imaging.Resize(srcImage, 800, 0, imaging.Lanczos)
	err = imaging.Save(dst, objName)
	if err != nil {
		return "", err
	}
	return
}
