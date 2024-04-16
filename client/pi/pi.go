package pi

import (
    "math/rand"
)


func Calc_pi(numThrows int) float64 {

    var numInside int

    for i := 0; i < numThrows; i++ {
        x := rand.Float64()
        y := rand.Float64()

        if x*x+y*y <= 1 {
            numInside++
        }
    }

    pi := 4 * float64(numInside) / float64(numThrows)

    return pi
}