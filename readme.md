A command line tool to create update delete your files

Currently work for windows.

features to be implmented

1. search files using filters (name, create date, extension)
2. search folder by name and create data
3. get files and folders not used since (data, or something like last 3 months)
4. delete folder
5. delete files
6. schedule tasks?

Get File details
go run main.go -op=details -name=your-file-path
example: go run main.go -op=details -fname="C:\webworkers\webworkers\src\App.js"

Search file by name
go run main.go -op=search -name=your-file-path -dir=directory-to-be-searched
example: go run main.go -op=search -fname="App" -dir="C:\DEVEL\GITHUB"

Recursive with skip, match and multithreading opts
go run main.go -op=search -name=your-file-path -dir=directory-to-be-searched -r=true
example: go run main.go -op=search -fname="App" -dir="C:\DEVEL\GITHUB" -r=true

go run main.go -op=search -name=your-file-path -dir=directory-to-be-searched -r=true
example: go run main.go -op=search -fname="App" -dir="C:\DEVEL\GITHUB" -r=true -skipDir=node_modules

go run main.go -op=search -name=your-file-path -dir=directory-to-be-searched -r=true match=exact
example: go run main.go -op=search -fname="App" -dir="C:\DEVEL\GITHUB" -r=true -skipDir=node_modules match=exact

go run main.go -op=search -name=your-file-path -dir=directory-to-be-searched -r=true match=exact -multiThread=true
example: go run main.go -op=search -fname="App" -dir="C:\DEVEL\GITHUB" -r=true -skipDir=node_modules match=exact -multiThread=true

Compress a file
go run main.go -op=zip -fpath=your-file-path -dir=where-file-should-get-saved
go run main.go -op=zip -fpath="C:\test-file-manager\Zone_api.xlsx" -dir="C:\test-file-manager"

go run main.go -op=zip -fpath=your-file-path -dir=where-file-should-get-saved
go run main.go -op=zip -fpath="C:\test-file-manager\yy" -dir="C:\test-file-manager" -ftype="folder"
