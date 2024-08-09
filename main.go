func calculate(s string) int {
	return calcule(s)
}

type Node struct {
	Num  int
	Opr  string
	Next *Node
}

func calcule(s string) int {
	a := &Node{}
	b := a
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if i == 0 {
			b.Num = int(s[i] - 48)
		} else {
			if s[i] >= '0' && s[i] <= '9' {
				b.Num = b.Num*10 + int(s[i]-48)
			} else {
				b.Next = &Node{Opr: string(s[i])}
				b = b.Next
				b.Next = &Node{}
				b = b.Next
			}
		}
	}
	a = check(a, "/", "*")
	a = check(a, "-", "+")
	return a.Num
}
func check(a *Node, opr1 string, opr2 string) *Node {
	b := a
	for b.Next != nil {
		if b.Next.Opr == opr1 {
			if opr1 == "-" {
				b.Num = b.Num - b.Next.Next.Num
			} else {
				b.Num = b.Num / b.Next.Next.Num
			}

			b.Next = b.Next.Next.Next
			if b == nil {
				break
			}
		} else if b.Next.Opr == opr2 {
			if opr2 == "*" {
				b.Num = b.Num * b.Next.Next.Num
			} else {
				b.Num = b.Num + b.Next.Next.Num
			}
			b.Next = b.Next.Next.Next
			if b == nil {
				break
			}
		} else {
			b = b.Next
		}
	}
	return a
}
