package repos

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFindList(t *testing.T) {
	areaList := FindList("100000")
	areaBuf, _ := json.Marshal(areaList)
	fmt.Println(string(areaBuf))
}

func BenchmarkFindList(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindList("100000")
	}
}
