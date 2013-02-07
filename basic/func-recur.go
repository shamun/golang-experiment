package main

import (
    "fmt"
    "math/rand"
    "time"
)

type walkFn func(*int) walkFn

func pickRandom(fns ...walkFn) walkFn {
    return fns[rand.Intn(len(fns))]
}

func walkEqual(i *int) walkFn {
    *i += rand.Intn(7) -3
    return pickRandom(walkForward, walkBackward)
}

func walkForward(i *int) walkFn {
    *i += rand.Intn(6)
    return pickRandom(walkEqual, walkBackward)
}

func walkBackward(i *int) walkFn {
    *i += -rand.Intn(6)
    return pickRandom(walkEqual, walkForward)
}

func main() {
    rand.Seed(time.Now().Unix())

    fn, progress:= walkEqual, 0
    for i := 0; i < 20; i++ {
        fn = fn(&progress)
        fmt.Println(progress)
    }
}
