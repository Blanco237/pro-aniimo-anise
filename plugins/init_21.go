package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.7" )

func D1aefkbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XIsqNWKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkNfcLFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2HZ4XYKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xL16GeHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZRLN0pekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UJBUfm5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vMtrkKnFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqnKLGYMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDlvgUXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOIzmQjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wDUELKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func klsJLkQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c95GC5iRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eSsIFCZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a7yjwWtUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U0KJwuXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nQspAcyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S1gpg487Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VDF9QJ7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VAgRR0FAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fi62eGXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PfVN23YjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KMICAbc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Abr08Y8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WlpET6Q8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pOqipudgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I7NGAYx7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJrfiQoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIHR6c9xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPEW1iSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1z2BXp2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dFkf0swJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8rNzKLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func POXhdlUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IWF0230fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4JhjatJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7RVuNfBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func znjmyibPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qqftKvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijBXbRFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BrB9UXbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OTpyH6jtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFttt0tXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w0KM0cptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pcRWJkFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SBUJZ7PfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 83xWtcQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2RSSqqdzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func At8oMh4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9tQPCrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nt8g4gwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ajt4KLGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uJevmDmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STniagLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RcegCgY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHKGB3fsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7O10B42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SK7bLJlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G97r88lKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Xnnv0VuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TTDYjioGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qbLYyC0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCBgEkJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3NyHZlrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MaKVHHkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func buchTHouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2lBOYWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ijXNs3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cs7PGyXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rgJhtcleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wgDNf2cEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xq9pcUrhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXMosneoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NGMNthOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V4dlJ4CRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h43eBrPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQXxospuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kz45xdyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8NjrzgNTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dL0N4ghLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wzZcbVIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 81K84FPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 55uDMQZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1IHozq4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q8eZC9NAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZ0IHsAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQ0wPAf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uEKIPtGRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UdTOaiUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wrw9H4jSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BiqsFYr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c3QmvrKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYTu8rtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Np22xubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gNEVVhYmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ds7iW11jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XIdlYRwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tum9kwAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLMO1dgDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Flu6B3AFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uFGHocxHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eL9JhpluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXrVh8PcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pn9RVaOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q7EtjYSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IbxPMBEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 26TghrUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J1wTlfq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xrZqwwIOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wH6wEgFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AOqDtuBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func smM2avz1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zn6E6GHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZhPclQ1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8PNHMe3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rifSlT8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UUziqpZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GH5X43UIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fI12O4B8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZWCLQULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nG5gOtXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWi6bkwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aiFeFZ1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQ59uJFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2vMKcXzWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 60Z9ZF5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUuX3NHpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rvSXwWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KldIgeehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wV1Yn60wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xMJTVB74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5csRguaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PhjN3AUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OgMxHEEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z4SqiFprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YbZBV1SGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6W3NiUFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XdYuKUSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NI22W9KpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pi3q8VRWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QEiaE438Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5Im1rQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9OD20KAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fK3fck1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zepxQ9SGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XB1vegVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGvX1fWyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func maU9YvKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aXK8r4lAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmCI6cDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4es9h3DRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TrCCHHKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YwH0RAJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J79ORAtGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 84CGBWDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yddjTj7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5qY1N5EIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZvIjEZfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NNNxHwwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nOeBKmREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iDnJO3kbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7HEkh0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8I0U723Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X10M70MsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1DWdd7hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXoN9Q5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQibeISKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 73DFEMk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0aPatthbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func suts0KquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYRlXvvJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qj8hCtJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22jb9dxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dP5xiB2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YxtzX7n2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NRs7tNIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6AywgifqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23BGhVdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IagZnF65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kiZIgQ9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u9O9JfOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VCc0eYQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func om188a7XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Klo7vbdkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59DhyK3mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08NlPKa0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vYb44V3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjiQWgMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ao57PfsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l26KvMcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pm3AoSslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n7j6n1W2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yMhd9nN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DA5r22hvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxXa2GopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LThStALaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func of1ZYcM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dXaqhUQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOjQnBKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKfpwnb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rqhpvOUrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xu0NftTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v9P9I7riWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U8sQ8BgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K4alaIhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLa9p8jBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrGHgXkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8KsuRU59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OgQi3UXkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func frkAmPaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y25MbofFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BBGfzsDGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NpJSpthJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BBG6SW3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xMg2gWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ocuuwKLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U8U7NErSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 511oEE0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U4BSEt4nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQ1yNgxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F1DklWGtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dqbjpzpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DshH9kBWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lwxLubM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dVtu5ILjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KWI8zrtNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4wYIGynPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgOgcTvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aVgpQKuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rHoFDMgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5XVsnEW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Es7Nj3O4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5yy5FFZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ng1YRQA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y5clXzL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CockahduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func liP3UPTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QGIwNFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWjdDIi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nF1mJEyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3ObGGiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ghm8a6vuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WiwRCSukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJdbEhRdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybWHdGP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9vJ0n0yYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6vqr6Q8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9kb7xKktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7IHtHahqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kjO5lEusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ck7ULprZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1j6cm43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SqydaUisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uU89wjMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S3uKjvLXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1Zk7XDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQjEFL8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H8B48WbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2HXtViNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5msPU5OKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zb5jysv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LXS4uylNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5LvynuuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zPCT9GM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j9uCWEOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a7uLKCIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6H8JJkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jiuMIh8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func spMHqDhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMmYEzeEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Km4yB2zlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lo7R2J0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VhuOQVseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLf7bVAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38nQhKxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tUS4hEYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I1kweS09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9a8vnAUxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtX6PWDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRDbD0mPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o5PxAzKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rECtscNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yCxz5jVkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MoUNG5WYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nz8b4rV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCs3HuSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5uJI9m1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IrpUElLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lzEAkeCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kOJe2m4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4eSVE2uTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MMOsV9zUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ddXFVpNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmeUms3ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xIOyIo63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t7fPWj60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ntiAbgC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ivyXF2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func edU80UprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YKJGO1R1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xV8htFxxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wYaNSwVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bhDGbwVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9tjcjKexWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9kjMtHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rp8j1nOSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DikomQK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCrXSWERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JB2xNgBlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FZwg5RIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IyapkdDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TxOGVFLtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 64NbDqe6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJldsW5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C0FcAu9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSLLd62UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LTIbjsQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrqG0oZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z0uQhIwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZerB5nBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MgJCHq2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nK2zing0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRd56xkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJU9USiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5lTajOYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dr4uNJBmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w7ai1XaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9AlF9ycEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k5lsQAqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dusjkwECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Ar01yyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZOcZrVenWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VHIwI6qeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jETOAi4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3x35JMuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LyihCNO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V2H8hFwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vdELGXGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0zYDq10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKf2SsKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9T7qSTdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qjF82HY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ml2ASpHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ehz6PfEoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oLzY5B9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I0CsR0SlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ku6COtKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hVMc3hDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouACAIvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HSuBiSAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wn2yTi3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0KWkjtvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HfjMCtL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twOX4rbFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJK1ftX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xnNt56P8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBQY7IrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCJfQXQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QBFVyqOhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W4i9YrQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eqkHB4TgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFCcNgbUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m0RRbIIvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SRZpSQ5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKaJYv5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wvwpnCbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MC8KGR0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LHNHDJWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cumoZRwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nx6r7VvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kjnf6AhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpum1eVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IuHH7PXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8VCFK6UgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sApJUNzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hEX1bsWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWpcMQO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqYCEuN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nSL1beIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wbGBfcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUtyQYcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b6RRg4vAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yTqwk808Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tWNvrc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ESickmRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQda0bqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lEXZaEJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Xx9eQY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ynqAesgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JDfZWAqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5uKBe4iKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6mahPWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 30AVy9D6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwK5RaswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Paeg2PrOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NROmW9xTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yzuZCg2KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CrgItkHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9TJOhx3ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vHjFTPbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvusEV7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t6UOG2ejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dwjJfx3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vAhiSb6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ShC9mg3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJGgGsLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HbWtW3XLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kP5bAg8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQe9gggeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nE1txF1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vtTHtqOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5esjM3pBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39ZQXuCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RlbzkeYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjxzuKhWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3TRa3XevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iTPiRViBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6CtONehRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdB3sXAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k43XBXkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xeKk171QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sATW9lENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxlZgiUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UXKGW7IfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lo2kSBXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BlH2ycAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func weReyieNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TiDJa95VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 00FC6LwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hV5mNwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1pEgFMHAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pjIxygWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N8a3yJxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1oOZOjrcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bqMul2JHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3F6ggtEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y7beLX0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xB0iUJArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZg72xeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5RGpbMGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t2qzy16wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nR2CxJM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J1Jt1ku8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LuSgooy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XqNVEDTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S1CDkO1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8PjPq2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ejENj3MiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gH1rYqBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uRH5NYOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JpclSaH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yINWorZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZVi4UtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lrRGYUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func krEL9lqXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CqVOIv3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xOG8l6A2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4cwJgl6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

