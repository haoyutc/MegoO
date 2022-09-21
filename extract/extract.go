package extract

import (
	"bytes"
	"compress/gzip"
	"context"
	util "github.com/dablelv/go-huge-util"
	"github.com/juju/errors"
	"github.com/klauspost/pgzip"
	"github.com/megoo/logger/zlog"
	"github.com/mholt/archiver/v3"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) CompressionWithGZip() {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	w.Write([]byte("hello world\n"))
	w.Close()
}
func (c *Client) CompressionWithPGZip() {
	var buf bytes.Buffer
	w := pgzip.NewWriter(&buf)
	p, _ := pgzip.NewReaderN(&buf, 128, 10)
	p.Multistream(true)
	w.Write([]byte("hello world\n"))
	w.Close()
}

func DownloadObject() error {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse("http://xxx-xxx-sh-xxx.cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: "XXX",
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: "XXXX",
		},
	})

	key := "setuptools-1.4.2.tar.gz"
	//file := "localfile"

	//opt := &cos.MultiDownloadOptions{
	//	ThreadPoolSize: 5,
	//}
	//response, err := client.Object.Download(
	//	context.Background(), key, file, opt,
	//)
	//if err != nil {
	//	zlog.Error("Download Object failed", zap.Error(err))
	//}
	//buf := response.Body
	//_, err = pgzip.NewReader(buf)
	//if err != nil {
	//	zlog.Error("pgzip.NewReader", zap.Error(err))
	//}

	rsp, err := client.Object.Get(
		context.Background(), key, nil,
	)
	if err != nil {
		zlog.Error("Get Object failed", zap.Error(err))
	}
	in := rsp.Body
	//in, err := pgzip.NewReader(buff)
	if err != nil {
		zlog.Error("pgzip.NewReader", zap.Error(err))
	}

	//in, err := os.Open(buff)
	//if err != nil {
	//	return err
	//}
	//defer in.Close()

	//out, err := os.Create("output/")
	//if err != nil {
	//	return err
	//}
	//defer out.Close()

	out := bytes.Buffer{}
	if err = unArchiveForGz(in, &out); err != nil {
		zlog.Error("unArchiveForGz", zap.Error(err))
	}

	// 查询元数据
	var headRsp *cos.Response
	headRsp, err = client.Object.Head(context.Background(), key, nil)
	if err != nil {
		zlog.Error("Object.Head failed", zap.Error(err))
	}
	hData, err := util.ToIndentJSON(headRsp)
	if err != nil {
		zlog.Error("ToIndentJSON failed", zap.Error(err))
	}
	zlog.Info("Object.Head rsp", zap.Any("Response", hData))

	var rsp2 *cos.Response
	rsp2, err = client.Object.Put(
		context.Background(), "test/", &out, nil,
	)
	if err != nil {
		zlog.Error("Object.Put failed", zap.Error(err))
	}
	data, err := util.ToIndentJSON(rsp2)
	if err != nil {
		zlog.Error("ToIndentJSON failed", zap.Error(err))
	}
	zlog.Info("Object.Put rsp", zap.Any("Response", data))
	return nil
}

// 解压缩.gz
func unArchiveForGz(in io.Reader, out io.Writer) error {
	gz := archiver.Gz{}
	if err := gz.Decompress(in, out); err != nil {
		return errors.Annotatef(err, "failed decompress file")
	}
	return nil
}
