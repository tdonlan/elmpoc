task :build do
	sh "go build elmpoc"
	sh "cd elm && elm-make src/Main.elm --output ../frontend/js/main.js && cd .."
end
task :run => :build do
	sh "./elmpoc"
end