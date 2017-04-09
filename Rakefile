task :build do
	sh "go build elmpoc"
	sh "elm-make elm/examples/01-button.elm --output frontend/01-button.html"
end
task :run do
	sh "./elmpoc"
end