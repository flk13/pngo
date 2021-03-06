package main

import (
	"github.com/golang/glog"
	"os"
	"path/filepath"
	"pngo/id"
	"strconv"
	"strings"
)

var (
	root       string //输入文件的路径
	outputPath string //输出文件夹路径
	outputIs   int    //是否维持图片原文件名
	width      int    //输出图像宽度
	quality    int    //控制输出图片的质量
)

//获取文件路径
func retrieveData(root string) (value chan string, err chan error) {
	err = make(chan error, 1)
	value = make(chan string)
	//开一个go程遍历根目录文件
	go func() {
		defer close(value)
		//遍历root根目录
		err <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//判断文件信息是否合法
			if !info.Mode().IsRegular() {
				return nil
			}
			//将path的制传给value管道
			value <- path
			return nil
		})
	}()
	return
}

//接收文件并将其发送到一个channel处理

//生成图片的唯一id进行识别
func onlyID1() string {
	u, err := id.NewUUID(id.Version1, nil)
	if err != nil {
		glog.Error(err)
	}
	glog.V(1).Info("use UUID")
	return u.String()
}

//默认使用雪花算法生成id
func onlyID() string {
	snow, err := id.NewSnowFlake(1) //代表本机的节点id
	if err != nil {
		glog.Error(err)
	}
	glog.V(1).Info("use snowflake")
	return strconv.FormatInt(snow.GetID(), 10)
}

//查找文件后缀名
func findName(name string) string {
	name = strings.ToLower(name)
	v := name[len(name)-4:]
	v1 := name[len(name)-3:]
	if v == "jpeg" {
		return v
	}
	if v1 == "jpg" || v1 == "png" || v1 == "gif" {
		return v1
	}
	return ""
}
