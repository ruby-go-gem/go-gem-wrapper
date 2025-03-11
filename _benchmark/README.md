# Benchmark
## Ruby's [Ractor](https://docs.ruby-lang.org/en/master/ractor_md.html) vs Ruby's [Fiber](https://docs.ruby-lang.org/en/master/Fiber.html) vs Go's [goroutine](https://go.dev/tour/concurrency/1)
Benchmark with [tak function](https://en.wikipedia.org/wiki/Tak_(function)) [^tak]

[^tak]: https://www.ruby-lang.org/en/news/2020/12/25/ruby-3-0-0-released/

```bash
$ ruby tarai.rb
go version go1.24.1 darwin/arm64
ruby 3.4.2 (2025-02-15 revision d2930f8e7a) +PRISM [arm64-darwin24]
Warming up --------------------------------------
          sequential     1.000 i/100ms
   parallel (Ractor)     1.000 i/100ms
    parallel (Fiber)     1.000 i/100ms
parallel (goroutine)     1.000 i/100ms
Calculating -------------------------------------
          sequential      0.018 (± 0.0%) i/s    (56.40 s/i) -      1.000 in  56.397748s
   parallel (Ractor)      0.052 (± 0.0%) i/s    (19.39 s/i) -      1.000 in  19.394641s
    parallel (Fiber)      0.018 (± 0.0%) i/s    (56.42 s/i) -      1.000 in  56.417936s
parallel (goroutine)      1.581 (± 0.0%) i/s  (632.49 ms/i) -      9.000 in   5.743819s

Comparison:
parallel (goroutine):        1.6 i/s
   parallel (Ractor):        0.1 i/s - 30.66x  slower
          sequential:        0.0 i/s - 89.17x  slower
    parallel (Fiber):        0.0 i/s - 89.20x  slower
```
