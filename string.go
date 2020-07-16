package main

func stringReverse(s string) string {
	bs := []byte(s)
	for i,j := 0,len(bs) - 1; i < j; i,j = i+1,j-1{
		bs[i],bs[j] = bs[j],bs[i]
	}
	return string(bs)
}