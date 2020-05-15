
func addBinary(a string, b string) string {
	c := ""
	l, s := "", ""
	if len(a) > len(b) {
		l = a
		s = b
	} else {
		l = b
		s = a
	}
	cha := len(l) - len(s)
	bu := ""
	for i := 0; i < cha; i++ {
		bu = bu + "0"
	}
	s = bu + s
	var wei uint8
	for i := len(l) - 1; i >= 0; i-- {
		tmp := (l[i] - '0') + (s[i] - '0') + wei
		if tmp >= 2 {
			wei = 1
			c = string(tmp-2+'0') + c
		} else {
			c = string(tmp+'0') + c
			wei = 0
		}
	}
	if wei == byte(1) {
		c = "1"+ c
	}
	return c
}

