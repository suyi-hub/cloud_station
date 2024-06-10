package aliyun

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/schollz/progressbar/v3"
)

func NewDefaultProgressListener(isUpload bool) *ProgressListener {

	return &ProgressListener{
		isUpload: isUpload,
	}
}

type ProgressListener struct {
	bar      *progressbar.ProgressBar
	isUpload bool
}

// func (p *ProgressListener) ProgressChanged(event *oss.ProgressEvent) {
// 	switch event.EventType {
// 	case oss.TransferStartedEvent:
// 		p.bar = progressbar.Default(event.TotalBytes, "文件上传中")
// 	case oss.TransferDataEvent:
// 		p.bar.Add64(event.RwBytes)

// 	case oss.TransferCompletedEvent:
// 		fmt.Println("文件上传完成")
// 	case oss.TransferFailedEvent:

// 		fmt.Println("文件上传失败")
// 	default:
// 	}
// }

func (p *ProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		// 初始化进度条
		var prefix string
		if p.isUpload {
			prefix = "文件上传中: "
		} else {
			prefix = "文件下载中: "
		}
		p.bar = progressbar.Default(event.TotalBytes, prefix)
	case oss.TransferDataEvent:
		var prefix string
		if p.isUpload {
			prefix = "文件上传中: "
		} else {
			prefix = "文件下载中: "
		}
		// 更新进度条
		p.bar = progressbar.Default(event.TotalBytes, prefix)

		p.bar.Add64(event.RwBytes)

	case oss.TransferCompletedEvent:

		if p.isUpload {
			fmt.Println("文件上传完成")
		} else {
			fmt.Println("文件下载完成")
		}

	case oss.TransferFailedEvent:

		if p.isUpload {
			fmt.Println("文件上传失败")
		} else {
			fmt.Println("文件下载失败")
		}

	default:
		// 其他事件类型
	}
}
