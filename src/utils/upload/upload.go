package upload

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

func PutBytes(filename string, data []byte) error {
	path := path.Dir(filename)
	//fmt.Println(path)
	_, err := os.Stat(path)
	if err != nil {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filename, data, os.ModePerm) //写入文件(字节数组)
}

func PutFileBytes(data []byte, oldname string, stype string) (remoteName string, err error) {
	rand.Seed(time.Now().UnixNano())
	remoteName = stype + "/" + time.Now().Format("20060102") + "/" + strconv.FormatInt(int64(time.Now().Unix()), 10) + strconv.FormatInt(int64(rand.Intn(999999999)), 10) + path.Ext(oldname)
	objName := "uploads/" + remoteName
	err = PutBytes(objName, data)
	if err != nil {
		return "", err
	}
	return
}
