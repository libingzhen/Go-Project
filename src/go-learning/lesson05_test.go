package fib

import ("testing")

func TestFibList(t *testing.T){
    a := 1
    b := 1
    for i := 0; i<5; i++ {
        t.Log(" ",b)
        tmp := a
	a = b
	b = a + tmp
    }
}
