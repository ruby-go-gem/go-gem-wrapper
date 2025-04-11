require "bundler/inline"

gemfile do
  source "https://rubygems.org"
  gem "benchmark-ips"
end

require "benchmark/ips"
require "fiber"

# Build native extension before running benchmark
Dir.chdir(File.join(__dir__, "..", "ruby", "testdata", "example")) do
  system("bundle config set --local path 'vendor/bundle'", exception: true)
  system("bundle install", exception: true)
  system("bundle exec rake", exception: true)
end

require_relative "../ruby/testdata/example/lib/example"

# c.f. https://www.ruby-lang.org/en/news/2020/12/25/ruby-3-0-0-released/
def tarai(x, y, z) =
  x <= y ? y : tarai(tarai(x-1, y, z),
                     tarai(y-1, z, x),
                     tarai(z-1, x, y))

# Suppress Ractor warning
$VERBOSE = nil

system("go version", exception: true)

MAX_BENCH_COUNT = 4

Benchmark.ips do |x|
  # Ruby: sequential version
  x.report("Ruby: sequential"){ MAX_BENCH_COUNT.times{ tarai(14, 7, 0) } }

  # Ruby: parallel version (with Ractor)
  x.report("Ruby: Ractor"){
    MAX_BENCH_COUNT.times.map do
      Ractor.new { tarai(14, 7, 0) }
    end.each(&:take)
  }

  # Ruby: parallel version (with Fiber)
  x.report("Ruby: Fiber"){
    MAX_BENCH_COUNT.times.map do
      Fiber.new { tarai(14, 7, 0) }
    end.each(&:resume)
  }

  # Go: sequential version
  x.report("Go: sequential"){ MAX_BENCH_COUNT.times{ Example::Benchmark.tarai(14, 7, 0) } }

  # Go: parallel version (with goroutine)
  x.report("Go: goroutine"){ Example::Benchmark.tarai_goroutine(14, 7, 0, MAX_BENCH_COUNT) }

  x.compare!
end
