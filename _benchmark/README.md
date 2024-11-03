# Benchmark
## Ruby's [Ractor](https://docs.ruby-lang.org/en/master/ractor_md.html) vs Ruby's [Fiber](https://docs.ruby-lang.org/en/master/Fiber.html) vs Go's [goroutine](https://go.dev/tour/concurrency/1)
Benchmark with [tak function](https://en.wikipedia.org/wiki/Tak_(function)) [^tak]

[^tak]: https://www.ruby-lang.org/en/news/2020/12/25/ruby-3-0-0-released/

```bash
$ ruby tarai.rb
go version go1.23.2 darwin/arm64
ruby 3.3.5 (2024-09-03 revision ef084cc8f4) [arm64-darwin23]
Warming up --------------------------------------
          sequential     1.000 i/100ms
   parallel (Ractor)     1.000 i/100ms
    parallel (Fiber)     1.000 i/100ms
parallel (goroutine)     1.000 i/100ms
Calculating -------------------------------------
          sequential      0.017 (± 0.0%) i/s    (57.84 s/i) -      1.000 in  57.844394s
   parallel (Ractor)      0.052 (± 0.0%) i/s    (19.41 s/i) -      1.000 in  19.412980s
    parallel (Fiber)      0.018 (± 0.0%) i/s    (56.78 s/i) -      1.000 in  56.780241s
parallel (goroutine)      1.478 (± 0.0%) i/s  (676.40 ms/i) -      8.000 in   5.412528s

Comparison:
parallel (goroutine):        1.5 i/s
   parallel (Ractor):        0.1 i/s - 28.70x  slower
    parallel (Fiber):        0.0 i/s - 83.95x  slower
          sequential:        0.0 i/s - 85.52x  slower
```
