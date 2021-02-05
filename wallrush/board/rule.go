package board

type dir struct {
	step  int
	limit func(int) bool
}

func compose(d1, d2 dir) dir {
	return dir{step: d1.step + d2.step, limit: func(i int) bool { return d1.limit(i) || d2.limit(i) }}
}

func e(l int) dir { return dir{step: 1, limit: func(i int) bool { return i%l == l-1 }} }
func w(l int) dir { return dir{step: -1, limit: func(i int) bool { return i%l == 0 }} }
func s(l int) dir { return dir{step: l, limit: func(i int) bool { return i/l == l-1 }} }
func n(l int) dir { return dir{step: -l, limit: func(i int) bool { return i/l == 0 }} }

func ne(l int) dir { return compose(n(l), e(l)) }
func nw(l int) dir { return compose(n(l), w(l)) }
func se(l int) dir { return compose(s(l), e(l)) }
func sw(l int) dir { return compose(s(l), w(l)) }

func allDirs(l int) map[string]dir {
	return map[string]dir{"n": n(l), "e": e(l), "s": s(l), "w": w(l), "ne": ne(l), "nw": nw(l), "se": se(l), "sw": sw(l)}
}

func Flip(b *Board, pos int, owner uint8, built bool) {
	md := allDirs(Side(b))
	for _, v := range md {
		if !canFlip(NewBIT(b, pos, v.step, v.limit), owner, built) {
			continue
		}
		flip(NewBIT(b, pos, v.step, v.limit), owner)
	}
}

func canFlip(bi interface{ Next() (*Cell, error) }, owner uint8, built bool) bool {
	if !built {
		return false
	}
	var wallLen int
	curr, err := bi.Next()
	for err != ErrEndOfLine {
		if curr.Lvl() == 0 {
			return false
		}
		if curr.Own() == owner {
			return wallLen > 0
		}
		wallLen++
		curr, err = bi.Next()
	}
	return err != ErrEndOfLine
}

func flip(bi interface{ Next() (*Cell, error) }, owner uint8) {
	curr, err := bi.Next()
	for err != ErrEndOfLine {
		if curr.Own() == owner {
			return
		}
		curr.own = owner
		curr, err = bi.Next()
	}
}
