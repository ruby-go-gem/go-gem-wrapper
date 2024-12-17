# patch_for_go_gem
Patch to make a gem into a Go gem right after `bundle gem`

## Requirements
* Go 1.23+
* Ruby 3.3+

## Usage
1. Download [patch_for_go_gem.rb](patch_for_go_gem.rb)
2. Run `bundle gem <GEM_NAME> --ext=c`
    * The other options for `bundle gem` are optional
3. Run `ruby patch_for_go_gem.rb --file /path/to/GEM_NAME.gemspec --dry-run`
4. Run `ruby patch_for_go_gem.rb --file /path/to/GEM_NAME.gemspec`
5. `cd` to the same location as `ext/GEM_NAME/go.mod`
6. Run `go get -u github.com/ruby-go-gem/go-gem-wrapper@latest`
