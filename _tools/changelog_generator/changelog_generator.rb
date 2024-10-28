# frozen_string_literal: true

require "optparse"
require "json"

# @return [String]
def latest_tag
  `git tag`.each_line.map(&:strip).max_by { |tag| Gem::Version.create(tag.delete_prefix("v")) }
end

# @param before [String]
# @param after [String]
# @return [Array<Integer>]
def search_pr_numbers(before:, after:)
  commits = `git rev-list --merges --right-only #{before}...#{after}`.each_line.map(&:strip)
  commits.map do |commit|
    commit_message = `git show -q #{commit}`
    commit_message =~ /Merge pull request #([0-9]+)/
    Regexp.last_match(1).to_i
  end
end

before = nil
after = nil

opt = OptionParser.new
opt.on("--before=BEFORE", "Before tag or sha1 (default: latest tag)") { |v| before = v }
opt.on("--after=AFTER", "After tag or sha1 (default: HEAD)") { |v| after = v }

opt.parse!(ARGV)

before ||= latest_tag
after ||= "HEAD"

pr_numbers = search_pr_numbers(before:, after:)
all_prs = pr_numbers.map do |pr_number|
  pr = JSON.parse(`gh pr view --json number,title,author,labels,url #{pr_number}`)
  pr["label_names"] = pr["labels"].map { |label| label["name"] }
  pr
end

# @param prs [Array<Hash>]
#
# @return [String]
def generate_category_changelog(prs)
  lines = []
  found_pr_numbers = []
  prs = prs.dup

  ["breaking change", "bug", "enhancement"].each do |label|
    label_prs = prs.find_all { |pr| pr["label_names"].include?(label) }

    next if label_prs.empty?

    case label
    when "breaking change"
      lines << "### :bomb: Breaking changes"
    when "bug"
      lines << "### Bugfixes"
    when "enhancement"
      lines << "### New Features"
    end

    label_prs.each do |pr|
      lines < generate_changelog_line(pr)
      found_pr_numbers << pr["number"]
    end

    found_pr_numbers.push(*label_prs.map { |pr| pr["number"] })

    prs.reject! { |pr| found_pr_numbers.include?(pr["number"]) }

    lines << ""
  end

  other_prs = prs.reject do |pr|
    found_pr_numbers.include?(pr["number"]) || pr["label_names"].include?("chore")
  end

  lines << "### Other changes"
  other_prs.each do |pr|
    lines << generate_changelog_line(pr)
  end
  found_pr_numbers.push(*other_prs.map { |pr| pr["number"] })
  lines << ""

  return "* No changes\n\n" if found_pr_numbers.empty?

  "#{lines.join("\n")}\n"
end

# @param pr [Hash]
#
# @return [String]
def generate_changelog_line(pr)
  author = pr["author"]["login"].delete_prefix("app/")

  "* #{pr["title"]} by @#{author} in #{pr["url"]}"
end

changelog_body = +""

changelog_body << "## Go\n"
go_prs = all_prs.find_all { |pr| pr["label_names"].include?("go") }
changelog_body << generate_category_changelog(go_prs)

changelog_body << "## Ruby\n"
ruby_prs = all_prs.find_all { |pr| pr["label_names"].include?("ruby") }
changelog_body << generate_category_changelog(ruby_prs)

changelog_body << "## ruby_h_to_go\n"
ruby_h_to_go_prs = all_prs.find_all { |pr| pr["label_names"].include?("ruby_h_to_go") }
changelog_body << generate_category_changelog(ruby_h_to_go_prs)

changelog_body << "## Other\n"
other_prs = all_prs.reject { |pr| pr["label_names"].any? { |label| %w[go ruby ruby_h_to_go chore].include?(label) } }
changelog_body << generate_category_changelog(other_prs)

changelog_body << "**Full Changelog**: https://github.com/ruby-go-gem/go-gem-wrapper/compare/#{before}...#{after}"

puts changelog_body
