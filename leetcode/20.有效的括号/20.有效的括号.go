
 func isValid(s string) bool {
	if len(s)==0{
		return true
	}
	if string(s[0])==")" || string(s[0])=="}"||string(s[0])=="]"{
		return false
	}
	if len(s)%2!=0{
		return false
	}
	var a []string
	for _,i:=range s{
		if string(i) == "(" || string(i) == "{" || string(i) == "[" {
			a = append(a,string(i))
		} else{
			if a[len(a)-1]=="("{
				if string(i) != ")"{
					return false
				} else{
					a = a[:len(a)-1]
				}
			}else if a[len(a)-1]=="{"{
				if string(i)!="}"{
					return false
				} else{
					a = a[:len(a)-1]
				}
			}else if a[len(a)-1]=="["{
				if string(i)!="]"{
					return false
				} else{
					a = a[:len(a)-1]
				}
			}
		}
	}
	if len(a) != 0{
		return false
	} else{
		return true
	}
}

