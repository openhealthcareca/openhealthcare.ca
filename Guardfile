# Watch and compile Less files
guard 'less', :all_on_start => true, :all_after_change => true do
  watch(%r{^.+\.less$})
end

guard 'livereload' do
  watch(%r{src/openhealthcare.ca/views/.+\.(tmpl)$})
  watch(%r{src/openhealthcare.ca/public/.+\.(css|js|html)})
end

# Add files and commands to this file, like the example:
#   watch(%r{file/path}) { `command(s)` }
#
# guard :shell do
#   watch(/(.*).txt/) {|m| `tail #{m[0]}` }
# end
