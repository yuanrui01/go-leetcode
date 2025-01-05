package atm

// ATM 2241. 设计一个 ATM 机器
type ATM struct {
	deposit [5]int
	values  [5]int
}

func Constructor() ATM {
	return ATM{
		deposit: [5]int{},
		values:  [5]int{20, 50, 100, 200, 500},
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := range banknotesCount {
		this.deposit[i] += banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	result := []int{0, 0, 0, 0, 0}
	for i := 4; i >= 0; i-- {
		if amount > 0 && this.deposit[i] > 0 {
			t := amount / this.values[i]
			t = min(t, this.deposit[i])
			result[i] = t
			this.deposit[i] -= t
			amount -= t * this.values[i]
		}
	}
	if amount > 0 {
		for i := range result {
			this.deposit[i] += result[i]
		}
		return []int{-1}
	}
	return result
}
