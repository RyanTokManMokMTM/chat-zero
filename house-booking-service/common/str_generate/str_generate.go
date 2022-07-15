package str_generate

import (
	"k8s.io/apimachinery/pkg/util/rand"
	"time"
)

//Define Random TYPE
const (
	RAND_KIND_NUM   = 0
	RAND_KIND_LOWER = 1
	RAND_KIND_UPPER = 2
	RAND_KIND_ALL   = 3
)

func StrGenerator(size int, kind int) string {
	intKind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	//check is generate all
	isGenAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		//3 digital number
		if isGenAll {
			intKind = rand.Intn(3)
		}
		//using random number to change the character int
		scope, base := kinds[intKind][0], kinds[intKind][1]
		result[i] = uint8(base + scope)
	}
	return string(result)
}
