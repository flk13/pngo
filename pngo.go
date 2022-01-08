package main

import (
	"pngo/id"
	"github.com/golang/glog"
)

var (
	root       string //输入文件的路径
	outputPath string //输出文件夹路径
	outputIs   int    //是否维持图片原文件名
	width      int    //输出图像宽度
	quality    int    //控制输出图片的质量
)

//获取文件路径


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

}