package Sessions

import (
	"math/rand"
	"time"
)



const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"


var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

var m = map[string]string{}

func AddtoMap(cookie string, nick string){
	m[cookie]=nick
	for e := range m {
		println(e)
		println(m[e])
	}
}

func ExistinMap(cookie string) bool {
	_, ok := m[cookie]
	return ok

}

func GetFormMap(name string) string {
	for k, v := range m {
		if v == name {
			return k
		}
	}
	return ""
}

func DeletefromMap(cookie string) {
	delete(m, cookie)
}
