package main

import "math/rand"

const seed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const base = "http://tinyurl.com/"

type Codec struct {
	short2long map[string]string
	long2short map[string]string
}

func Constructor() Codec {
	return Codec{make(map[string]string), make(map[string]string)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	val, exist := this.long2short[longUrl]
	if exist {
		return val
	}

	for !exist {
		short := ""
		for i := 0; i < 6; i++ {
			short += string([]byte{seed[rand.Intn(len(seed))]})
		}
		_, ok := this.short2long[short]
		if !ok {
			this.short2long[base+short] = longUrl
			this.long2short[longUrl] = base + short
			exist = true
		}
	}

	return this.long2short[longUrl]
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	if val, ok := this.short2long[shortUrl]; ok {
		return val
	}

	return ""
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
