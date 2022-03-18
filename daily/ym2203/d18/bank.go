package d18

import "sync"

type Bank struct {
	mu       sync.Mutex // 如果考虑线程安全，那么应该加锁
	accounts []int64
	n        int
}

func Constructor(balance []int64) Bank {
	return Bank{
		accounts: balance,
		n:        len(balance),
	}
}

// Transfer 从account1转账到account2
func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	account1 = account1 - 1
	account2 = account2 - 1
	if account1 >= b.n || account1 < 0 || account2 >= b.n || account2 < 0 {
		return false
	}
	if b.accounts[account1] < money {
		return false
	}

	b.accounts[account1] -= money
	b.accounts[account2] += money // 如果上一步出现了问题，那么这里应该进行回滚操作
	return true
}

// Deposit 存款
func (b *Bank) Deposit(account int, money int64) bool {
	account = account - 1

	if account >= b.n || account < 0 {
		return false
	}

	b.accounts[account] += money
	return true
}

// Withdraw 取款
func (b *Bank) Withdraw(account int, money int64) bool {
	account = account - 1
	if account >= b.n || account < 0 {
		return false
	}

	if b.accounts[account] < money {
		return false
	}

	b.accounts[account] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(accounts,money);
 * param_3 := obj.Withdraw(accounts,money);
 */
