

// @lc code=start
func countPrimes(n int) int {
	if n == 0 || n == 1{
		return 0
	}
	isPrime:=make([]int,n)
	for i:=2;i*i<n;i++{
		if isPrime[i]==0{
			for j:=i*i;j<n;j=j+i{
				isPrime[j]=1
			}
		}
	}
	count :=0
	for _,i := range isPrime {
		if i == 0{
			count++
		}
	}
	return count-2
}
// @lc code=end

