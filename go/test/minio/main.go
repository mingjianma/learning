package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// 创建minio对象
	endpoint := "minio-test.ctcdn.cn"
	accessKeyID := "AKIAIOSFODNN7EXAMPLE"
	secretAccessKey := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		fmt.Printf("minio.New err:%s \n", err.Error())
		return
	}
	ctx := context.Background()
	// 检测bucket是否存在
	awsBucket := "ecx"
	exists, err := minioClient.BucketExists(ctx, awsBucket)
	if err != nil {
		fmt.Printf("err:%s \n", err.Error())
		return
	}
	if !exists {
		fmt.Printf("Bucket %s not exist \n", awsBucket)
		return
	}
	fmt.Printf("Bucket %s exist \n", awsBucket)
	// 	创建文件
	objectName := "public/虚机操作日志.xlsx"
	filePath := "/Users/mamingjian/虚机操作日志_1652329637.xlsx"
	contentType := "application/octet-stream"
	info, err := minioClient.FPutObject(ctx, awsBucket, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		fmt.Printf("err:%s \n", err.Error())
		return
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info)
}
