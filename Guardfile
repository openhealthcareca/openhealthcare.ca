# Watch and compile Less files
guard 'less', :all_on_start => true, :all_after_change => true do
  watch(%r{^.+\.less$})
end

guard 'livereload' do
  watch(%r{src/openhealthcare.ca/views/.+\.(tmpl)$})
  watch(%r{src/openhealthcare.ca/public/.+\.(css|js|html)})
end

