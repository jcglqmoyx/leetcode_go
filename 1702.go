package main

func maximumBinaryString(binary string) string {
	n := len(binary)
	k := 0
	for k < n {
		if binary[k] == '1' {
			k++
		} else {
			break
		}
	}
	if k == n {
		return binary
	}
	cnt := 0
	for i := k + 1; i < n; i++ {
		if binary[i] == '0' {
			cnt++
		}
	}
	b := []byte(binary)
	for i := 0; i < n; i++ {
		b[i] = '1'
	}
	b[k+cnt] = '0'
	return string(b)
}
