package ps01

type linkedList struct {
	number     string
	index      int
	prev, next *linkedList
}

func (l *linkedList) Prev() *linkedList {
	return l.prev
}

func (l *linkedList) Next() *linkedList {
	return l.next
}

func (l *linkedList) First() *linkedList {
	for {
		l = l.Prev()
		if l == nil || l.index == 0 {
			return l
		}
	}
}

func (l *linkedList) Append(num string) *linkedList {
	l.next = &linkedList{number: num, index: l.index + 1, prev: l, next: l.First()}
	return l.next
}

func (l *linkedList) Find(target string) *linkedList {
	for {
		if l.number == target {
			return l
		}
		l = l.Prev()
	}
}

var nums = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func solution(p string, s string) int {
	result := 0
	l := &linkedList{number: "0", prev: nil}
	for _, n := range nums {
		l = l.Append(n)
	}
	l.First().prev = l

Outer:
	for i := 0; i < len(p); i++ {
		P := l.Find(string(p[i]))
		if P.number == string(s[i]) {
			continue
		}
		PP := P.Prev()
		PN := P.Next()
		for j := 1; j < 10; j++ {
			if PP.number == string(s[i]) || PN.number == string(s[i]) {
				result += j
				continue Outer
			}
			PP = PP.Prev()
			PN = PN.Next()
		}
	}
	return result
}
