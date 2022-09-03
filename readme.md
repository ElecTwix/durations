# 1. Go Durations

[![Go Report Card](https://goreportcard.com/badge/github.com/ElecTwix/durations)](https://goreportcard.com/report/github.com/ElecTwix/durations) [![PkgGoDev](https://pkg.go.dev/badge/github.com/ElecTwix/durations)](https://pkg.go.dev/github.com/ElecTwix/durations)


## 1.1. What is this for

Go Durations easily parse string durations without any order in input order and high durations like 1 year.

## 1.2. Why
I was looking for parser with out any order on string, easy to use, support long term durations such as year, month. I couldn't find so I created one.

## 1.3. Usage

`` Duration, Err := GetDuration("1mon 1w 1h 1m 1s") ``

`` Duration, Err := GetDuration("1h 1w 1m 1y") ``

## 1.4. All Durations

|  Durations  |      short      |  Package |
|----------     |:-------------:|------:|
| Nanosecond |  ns | [Time.Time](https://pkg.go.dev/time)  |
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

## 1.5. Credit

[ElecTwix](https://github.com/ElecTwix)
