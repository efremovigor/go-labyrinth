package main

import "strconv"

//получить рандомное не четное число
func oddRandom(limit int) (n int) {
    n = r.Intn(limit)
    for n%2 == 0 {
        n = r.Intn(limit)
    }
    return
}

func getIndex(x int, y int) string {
    return strconv.Itoa(x) + "|" + strconv.Itoa(y)
}

