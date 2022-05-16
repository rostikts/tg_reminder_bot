package utils

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var TokenCache *cache.Cache

func InitCache() {
	TokenCache = cache.New(time.Hour*1, time.Hour*1)
}
