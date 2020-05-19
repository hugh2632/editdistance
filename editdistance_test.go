package editdistance

import (
	"fmt"
	"testing"
)

func TestGetDistance(t *testing.T) {
	var a = "23"
	var b = "32345"
	var res = Compare(a, b)
	var first, second = res.GetOutPut([]rune("__"), func(r rune, r2 rune) (runes []rune, runes2 []rune) {
		return []rune("[" + string(r) + "]"), []rune("{" + string(r2) + "}")
	}, func(r rune) []rune {
		return []rune("{" + string(r) + "}")
	}, func(r rune) []rune {
		return []rune("[" + string(r) + "]")
	})
	t.Log("需要操作", res.Distance, "下变成同一个字符串。")
	t.Log(string(first))
	t.Log(string(second))
}

func BenchmarkEditDistanceDp_GetOutPut(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var a = "研究:羊驼产生抗体可消灭新冠病毒 治疗效果良好"
			var b = "研究:羊驼粑粑可消灭新冠病毒 治疗脑残效果良好1"
			var res = Compare(a, b)
			var first, second = res.GetOutPut([]rune("___"), func(r rune, r2 rune) (runes []rune, runes2 []rune) {
				return []rune("[" + string(r) + "]"), []rune("{" + string(r2) + "}")
			}, func(r rune) []rune {
				return []rune("{" + string(r) + "}")
			}, func(r rune) []rune {
				return []rune("[" + string(r) + "]")
			})
			fmt.Println("需要操作", res.Distance, "下变成同一个字符串。")
			fmt.Println(string(first))
			fmt.Println(string(second))
		}
	})
}