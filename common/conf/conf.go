package conf

import (
	"github.com/megoo/logger/zlog"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"go.uber.org/zap"
	"os"
)

//scanner
type scanner struct {
	paths []string
	val   interface{}
}

type Options struct {
	paths    []string
	files    []string
	scanners []scanner
}

type Option func(*Options)

const ConfigCommon = "common"

func Load(opts ...Option) {
	var options Options

	for _, o := range opts {
		o(&options)
	}

	for _, p := range options.paths {
		var sources []source.Source
		for _, f := range options.files {
			if _, err := os.Stat(p + "/" + f); err != nil && os.IsExist(err) == false {
				continue
			}
			sources = append(sources, file.NewSource(file.WithPath(p+"/"+f)))
		}
		if err := config.Load(sources...); err != nil {
			zlog.Error("Load config sources error", zap.Error(err))
		}
	}

	for _, s := range options.scanners {
		if err := config.Get(s.paths...).Scan(s.val); err != nil {
			zlog.Fatal("Load config paths error", zap.Error(err))
		}
	}
}

func Path(path ...string) Option {
	return func(o *Options) {
		o.paths = append(o.paths, path...)
	}
}

func File(file ...string) Option {
	return func(o *Options) {
		o.files = append(o.files, file...)
	}
}

func Scan(val interface{}, path ...string) Option {
	return func(o *Options) {
		o.scanners = append(o.scanners, scanner{[]string{ConfigCommon}, val}, scanner{path, val})
	}
}
