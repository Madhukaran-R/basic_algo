package main

func selectionSort(a []int) []int {
    for i:=0; i<len(a)-1; i++ {
        curMin := i
        for j:=i+1;j<len(a);j++ {
            if a[curMin] > a[j] {
                curMin = j
            }
        }
        a[i],a[curMin] = a[curMin],a[i]
    }
    return a
}

func bubbleSort(a []int) []int {
    for i:=0; i < len(a); i++ {
        for j:=0; j < len(a)-i-1; j++ {
            if a[j] > a[j+1] {
                a[j],a[j+1] = a[j+1],a[j]
            }
        }
    }
    return a
}

func linearSearch(a []int,st int) int {
    for i:=0;i<len(a);i++ {
        if a[i] == st {
            return i
        }
    }
    return -1
}
