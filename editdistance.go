package editdistance

type editDistanceDp struct {
	DPTable  [][]uint32
	First    []rune
	Second   []rune
	Distance int
}

func Compare(first string, second string) *editDistanceDp {
	var ed = &editDistanceDp{
		First:  []rune(first),
		Second: []rune(second),
	}
	var lenx = len(ed.First) + 1
	var leny = len(ed.Second) + 1
	ed.DPTable = make([][]uint32, lenx, lenx)

	var flag = 0
	for i := 0; i < lenx; i++ {
		ed.DPTable[i] = make([]uint32, leny, leny)
		if i == 0 {
			for j := 0; j < leny; j++ {
				ed.DPTable[0][j] = uint32(j)
			}
		}
		ed.DPTable[i][0] = uint32(i)
	}
	for i := 1; i < lenx; i++ {
		for j := 1; j < leny; j++ {
			if ed.First[i-1] == ed.Second[j-1] {
				flag = 0
			} else {
				flag = 1
			}
			ed.DPTable[i][j] = min(ed.DPTable[i-1][j-1]+uint32(flag), ed.DPTable[i-1][j]+1, ed.DPTable[i][j-1]+1)
		}
	}
	ed.Distance = int(ed.DPTable[lenx-1][leny-1])
	return ed
}

func (this *editDistanceDp) GetOutPut(emptyStr []rune, SubFunc func(rune, rune) ([]rune, []rune), AddFunc func(rune) []rune, DelFunc func(rune) []rune) (firstout []rune, secondout []rune) {
	var nowx = len(this.First)
	var nowy = len(this.Second)

	for nowx > 0  && nowy > 0 {
		var sub = this.DPTable[nowx-1][nowy-1]
		var add = this.DPTable[nowx][nowy-1]
		var del = this.DPTable[nowx-1][nowy]
		var subFirst, subSecond []rune
		if this.First[nowx-1] == this.Second[nowy-1] {
			subFirst = []rune{this.First[nowx-1]}
			subSecond = []rune{this.Second[nowy-1]}
			nowx--
			nowy--
		} else if sub <= add && sub <= del {
			subFirst, subSecond = SubFunc(this.First[nowx-1], this.Second[nowy-1])
			nowx--
			nowy--
		} else if add <= del {
			subFirst = emptyStr
			subSecond = AddFunc(this.Second[nowy-1])
			nowy--
		} else {
			subSecond = emptyStr
			subFirst = DelFunc(this.First[nowx-1])
			nowx--
		}
		if subFirst != nil {
			firstout = append(subFirst, firstout...)
		}
		if subSecond != nil {
			secondout = append(subSecond, secondout...)
		}
	}

	if nowx == 0 && nowy > 0{
		var tmp = make([]rune, nowy, nowy)
		for i:=0; i< nowy;i++ {
			tmp = append(tmp, AddFunc(this.Second[i])...)
			firstout = append(emptyStr, firstout...)
		}
		secondout = append(tmp, secondout...)
	}else if nowy == 0 && nowx >0 {
		var tmp = make([]rune, nowx, nowx)
		for i:=0; i< nowx;i++ {
			tmp = append(tmp, DelFunc(this.First[i])...)
			secondout = append(emptyStr, secondout...)
		}
		firstout = append(tmp, firstout...)
	}
	return firstout, secondout
}

func min(first uint32, second uint32, third uint32) uint32 {
	if first > second {
		first = second
	}
	if first > third {
		first = third
	}
	return first
}
