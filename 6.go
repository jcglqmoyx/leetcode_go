package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := ""
	rows := make([]string, numRows)
	row, t := 0, 1
	for _, c := range s {
		if row == 0 {
			t = 1
		} else if row == numRows-1 {
			t = -1
		}
		rows[row] += string(c)
		row += t
	}
	for _, s := range rows {
		res += s
	}
	return res
}
