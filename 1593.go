package main

func maxUniqueSplit(s string) int {
	set := map[string]bool{}
	ans := 0
	var dfs = func(s string, index int) {}
	dfs = func(s string, index int) {
		if index == len(s) {
			if len(set) > ans {
				ans = len(set)
			}
		}
		for length := 1; index+length <= len(s); length++ {
			str := s[index : index+length]
			if set[str] == true {
				continue
			}
			set[str] = true
			dfs(s, index+length)
			delete(set, str)
		}
	}
	dfs(s, 0)
	return ans
}
