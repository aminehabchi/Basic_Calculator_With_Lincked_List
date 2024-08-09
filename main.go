func calculate(s string) int {
	a, e := strconv.Atoi(s)
	if e == nil {
		return a
	}
	a = calcule(s)
	return a
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
        if s[i]==' '{
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
    
    

	b = a
	for b.Next != nil {
		if b.Next.Opr == "/" {
			b.Num = b.Num / b.Next.Next.Num
			b.Next = b.Next.Next.Next
			if b == nil {
				break
			}
		} else if b.Next.Opr == "*" {
			b.Num = b.Num * b.Next.Next.Num
			b.Next = b.Next.Next.Next
			if b == nil {
				break
			}
		} else {
			b = b.Next
		}
	}

    b = a
	for b != nil {
		fmt.Print(b.Opr, " ")
		fmt.Println(b.Num)
		b = b.Next
	}

	b = a
	for b.Next != nil {
		if b.Next.Opr == "-" {
			b.Num = b.Num - b.Next.Next.Num
			b.Next = b.Next.Next.Next
			if b == nil {
				break
			}
		} else if b.Next.Opr == "+" {
			b.Num = b.Num + b.Next.Next.Num
			b.Next = b.Next.Next.Next
			if b == nil {
				break
			}
		}else {
			b = b.Next
		}
	}
	return a.Num
}
