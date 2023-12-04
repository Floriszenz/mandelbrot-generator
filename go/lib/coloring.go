package main

import "math"

func hsvToRgb(hue, saturation, value float64) (uint8, uint8, uint8) {
    v := uint8(255.0 * value)

    if saturation == 0.0 {
        return v, v, v
    }

    h := hue / 60.0
    h_int := math.Floor(h)
    f := h - h_int
    p := uint8(255.0 * value * (1.0 - saturation))
    q := uint8(255.0 * value * (1.0 - saturation * f))
    t := uint8(255.0 * value * (1.0 - saturation * (1.0 - f)))

    switch uint(h_int) {
    case 1:
        return q, v, p
    case 2:
        return p, v, t
    case 3:
        return p, q, v
    case 4:
        return t, p, v
    case 5:
        return v, p, q
    default:
        return v, t, p
    }
}

