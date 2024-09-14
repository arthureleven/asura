package database

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client

func Init() error {
	Cache = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASS"),
	})

	command := Cache.Ping(context.Background())

	if err := command.Err(); err != nil {
		return err
	}

	return nil
}

func GetCachedImage(ctx context.Context, suffix, name string) *image.Image {
	k := fmt.Sprintf("/images/%s/%s", suffix, name)
	v := Cache.Get(ctx, k)

	if b, _ := v.Bytes(); len(b) != 0 {
		reader := bytes.NewReader(b)

		if img, err := png.Decode(reader); err != nil {
			return &img
		}
	}

	return nil
}

func CacheImage(ctx context.Context, img *image.Image, suffix, name string) {
	if img != nil {
		var buf bytes.Buffer

		png.Encode(&buf, *img)

		if buf.Len() > 0 {
			k := fmt.Sprintf("/images/%s/%s", suffix, name)

			Cache.Set(ctx, k, buf.Bytes(), time.Minute*120)
		}
	}
}
