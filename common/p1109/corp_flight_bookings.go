package p1109

// 这里有 n 个航班，它们分别从 1 到 n 进行编号。
//
//有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
//
//请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。

func corpFlightBookings(bookings [][]int, n int) []int {
	retSli := make([]int, n)

	for _, book := range bookings {
		first, last, total := book[0], book[1], book[2]
		// tmpSli[first-1] 为
		retSli[first-1] += total
		if last < n {
			retSli[last] -= total
		}
	}
	for i := 1; i < n; i++ {
		retSli[i] += retSli[i-1]
	}

	return retSli
}
