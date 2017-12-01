package repository

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/disintegration/imaging"
)

// 图片服务功能，把远程图片抓取回来并显示出来

//GetMd5String 生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// SaveImg 保存图片到本地
func SaveImg(imageURL, saveName string) (n int64, err error) {
	out, err := os.Create(saveName)
	defer out.Close()
	if err != nil {
		return
	}
	resp, err := http.Get(imageURL)

	if err != nil {
		return
	}
	pix, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	if err != nil {
		return
	}
	n, err = io.Copy(out, bytes.NewReader(pix))

	if err != nil {
		return
	}
	// todo 获取图片类型
	// fmt.Println(resp.Header.Get("Content-Type"))
	return
}

// PrintErrorImageHandler 显示错误图片
func PrintErrorImageHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "images/404.png")
}

// PrintImageHandler 显示正常图片
func PrintImageHandler(u string, w http.ResponseWriter, r *http.Request) {

	imgname := GetMd5String(u)

	imgpath := fmt.Sprintf("file/%v.jpg", imgname)

	// 如果本地服务器不存在缓存，再去拿
	_, err := os.Stat(imgpath)
	if os.IsNotExist(err) {
		_, err2 := SaveImg(u, imgpath)
		if err2 != nil {
			imgpath = "images/404.png"
		} else {
			src, err := imaging.Open(imgpath)
			if err != nil {
				log.Fatalf("Open failed: %v", err)
				imgpath = "images/404.png"
			} else {
				// src = imaging.Resize(src, 256, 0, imaging.Lanczos)
				src = imaging.Resize(src, 400, 0, imaging.Lanczos)
				src = imaging.CropAnchor(src, 400, 300, imaging.Center)
				err = imaging.Save(src, imgpath)
				if err != nil {
					log.Fatalf("Save failed: %v", err)
					imgpath = "images/404.png"
				}
			}
		}
	}
	http.ServeFile(w, r, imgpath)
}
