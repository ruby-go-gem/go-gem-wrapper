# frozen_string_literal: true

# D = Steep::Diagnostic

target :lib do
  signature "ruby/testdata/example/sig"
  signature "_gem/sig"

  check "ruby/testdata/example/lib"
  check "_gem/lib"

  # check "Gemfile"                   # File name
  # check "app/models/**/*.rb"        # Glob
  # ignore "lib/templates/*.rb"

  # library "pathname"              # Standard libraries
  # library "strong_json"           # Gems

  collection_config "rbs_collection.yaml"

  # configure_code_diagnostics(D::Ruby.default)      # `default` diagnostics setting (applies by default)
  # configure_code_diagnostics(D::Ruby.strict)       # `strict` diagnostics setting
  # configure_code_diagnostics(D::Ruby.lenient)      # `lenient` diagnostics setting
  # configure_code_diagnostics(D::Ruby.silent)       # `silent` diagnostics setting
  # configure_code_diagnostics do |hash|             # You can setup everything yourself
  #   hash[D::Ruby::NoMethod] = :information
  # end
end

# target :test do
#   signature "sig", "sig-private"
#
#   check "test"
#
#   # library "pathname"              # Standard libraries
# end
