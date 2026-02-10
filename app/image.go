package main

import (
	"github.com/royalrick/wechatwriter/app/image"
	"github.com/spf13/cobra"
)

// imageCmd 图片处理命令组
func imageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "image",
		Short: "图片处理（上传、下载、生成）",
		Long: `图片处理命令组

支持的操作：
  upload    - 上传本地图片到微信素材库
  download  - 下载在线图片并上传到微信
  generate  - AI 生成图片并上传到微信`,
	}

	cmd.AddCommand(imageUploadCmd())
	cmd.AddCommand(imageDownloadCmd())
	cmd.AddCommand(imageGenerateCmd())

	return cmd
}

// imageUploadCmd 上传本地图片
func imageUploadCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "upload <file_path>",
		Short: "上传本地图片到微信素材库",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return initConfig()
		},
		Run: func(cmd *cobra.Command, args []string) {
			filePath := args[0]
			processor := image.NewProcessor(cfg, log)
			result, err := processor.UploadLocalImage(filePath)
			if err != nil {
				responseError(err)
				return
			}
			responseSuccess(result)
		},
	}
}

// imageDownloadCmd 下载并上传图片
func imageDownloadCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "download <url>",
		Short: "下载在线图片并上传到微信",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return initConfig()
		},
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]
			processor := image.NewProcessor(cfg, log)
			result, err := processor.DownloadAndUpload(url)
			if err != nil {
				responseError(err)
				return
			}
			responseSuccess(result)
		},
	}
}

// imageGenerateCmd AI 生成图片
func imageGenerateCmd() *cobra.Command {
	var size string

	cmd := &cobra.Command{
		Use:   "generate <prompt>",
		Short: "AI 生成图片并上传到微信",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return initConfig()
		},
		Run: func(cmd *cobra.Command, args []string) {
			prompt := args[0]
			processor := image.NewProcessor(cfg, log)

			// 如果指定了尺寸，临时覆盖配置
			if size != "" {
				result, err := processor.GenerateAndUploadWithSize(prompt, size)
				if err != nil {
					responseError(err)
					return
				}
				responseSuccess(result)
				return
			}

			result, err := processor.GenerateAndUpload(prompt)
			if err != nil {
				responseError(err)
				return
			}
			responseSuccess(result)
		},
	}

	cmd.Flags().StringVarP(&size, "size", "s", "", "Image size (e.g., 2560x1440 for 16:9)")

	return cmd
}
