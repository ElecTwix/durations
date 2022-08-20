# 1. Go Durations

## 1.1. Why
I was looking for parser easy to use and with out any order about timing on string and I couldn't find so I created one.

## 1.2. Usage

`` Duration, Err := GetDuration("1mon 1w 1h 1m 1s") ``

## 1.3. All Durations

|  Durations  |      short      |  Package |
|----------     |:-------------:|------:|
| Nanosecond |  left-aligned | [Time.Time](https://pkg.go.dev/time)  |
| Microsecond |    us & µs & μs  |   [Time.Time](https://pkg.go.dev/time) |
| Millisecond | ms |    [Time.Time](https://pkg.go.dev/time) |
| Second      | s |    [Time.Time](https://pkg.go.dev/time) |
| Minute | m |    [Time.Time](https://pkg.go.dev/time) |
| Hour | h |    [Time.Time](https://pkg.go.dev/time) |
| Day | d |    [durations](https://github.com/ElecTwix/durations) |
| Week | w |    [durations](https://github.com/ElecTwix/durations) |
| Month | mon |    [durations](https://github.com/ElecTwix/durations) |
| Year | y |    [durations](https://github.com/ElecTwix/durations) |
| Century | c |    [durations](https://github.com/ElecTwix/durations) |

## 1.4 Credit

[ElecTwix](https://github.com/ElecTwix)