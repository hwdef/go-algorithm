
 func countAndSay(n int) string {
	str := "1"
	for i:=1;i<n;i++{
		res := ""
		number:=str[0]
		count:=0
		for j:=0;j<len(str);j++{
			if str[j] == number {
				count++
			 } else {
				res = res + strconv.Itoa(count) + string(number) 
				number = str[j]
				count = 1
			 }
		}
		res = res + strconv.Itoa(count) + string(number) 
		str = res
	}
	return str
}

