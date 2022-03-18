package d14

func findRestaurant(list1 []string, list2 []string) (restaurants []string) {
	restaurants = make([]string, 0)

	tmpMap := make(map[string]int, len(list1))
	min := 2002

	for i, v := range list1 {
		tmpMap[v] = i
	}
	for secondIndex, v := range list2 {
		if secondIndex > min {
			break
		}
		if firstIndex, ok := tmpMap[v]; ok {
			total := firstIndex + secondIndex
			// 如果最小值更新，那么重置参数数量
			if total < min {
				min = total
				restaurants = []string{v}
			} else if total == min {
				// 如果等于最小值，那么添加到最小索引餐厅中
				restaurants = append(restaurants, v)
			} else {
				// 如果大于最小索引，则应该直接跳过
			}
		}
	}
	return
}
