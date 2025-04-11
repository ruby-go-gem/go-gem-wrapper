# Benchmark
## Ruby's [Ractor](https://docs.ruby-lang.org/en/master/ractor_md.html) vs Ruby's [Fiber](https://docs.ruby-lang.org/en/master/Fiber.html) vs Go's [goroutine](https://go.dev/tour/concurrency/1)
Benchmark with [tak function](https://en.wikipedia.org/wiki/Tak_(function)) [^tak]

[^tak]: https://www.ruby-lang.org/en/news/2020/12/25/ruby-3-0-0-released/

```bash
$ ruby tarai.rb
go version go1.24.1 darwin/arm64
ruby 3.4.2 (2025-02-15 revision d2930f8e7a) +PRISM [arm64-darwin24]
Warming up --------------------------------------
    Ruby: sequential     1.000 i/100ms
        Ruby: Ractor     1.000 i/100ms
         Ruby: Fiber     1.000 i/100ms
      Go: sequential     1.000 i/100ms
       Go: goroutine     1.000 i/100ms
Calculating -------------------------------------
    Ruby: sequential      0.018 (± 0.0%) i/s    (56.43 s/i) -      1.000 in  56.434220s
        Ruby: Ractor      0.052 (± 0.0%) i/s    (19.19 s/i) -      1.000 in  19.194061s
         Ruby: Fiber      0.018 (± 0.0%) i/s    (56.57 s/i) -      1.000 in  56.574037s
      Go: sequential      0.451 (± 0.0%) i/s     (2.22 s/i) -      3.000 in   6.650596s
       Go: goroutine      1.645 (± 0.0%) i/s  (607.81 ms/i) -      9.000 in   5.473278s

Comparison:
       Go: goroutine:        1.6 i/s
      Go: sequential:        0.5 i/s - 3.65x  slower
        Ruby: Ractor:        0.1 i/s - 31.58x  slower
    Ruby: sequential:        0.0 i/s - 92.85x  slower
         Ruby: Fiber:        0.0 i/s - 93.08x  slower
```
