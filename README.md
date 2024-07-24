# Learn GO with tests

I've been wanting to learn Go for a year or two now, but I never had the time or mental energy to make any real progress. Now that I've been laid off, I have plenty of both! 

When I'm not applying to 100 jobs a day, I'm  focussing some of my time on learning Go. I discovered a great resource for learning Go using TDD ([link](https://quii.gitbook.io/learn-go-with-tests)), and this repo is meant to record my journey.


## Useful GO commands
- `go mod init <module_name>` - this will create a module file (`go.mod`) in the current directory. You need to run this command before running commands like `go test` or `go build`. This file describes the module’s properties such as:
   - The current module’s module path. This should be a location from which the module can be downloaded by Go tools, such as the module code’s repository location. This serves as a unique identifier, when combined with the module’s version number. It is also the prefix of the package path for all packages in the module. For more about how Go locates the module, see the Go Modules Reference.
    - The minimum version of Go required by the current module.
    - A list of minimum versions of other modules required by the current module.
    - Instructions, optionally, to replace a required module with another module version or a local directory, or to exclude a specific version of a required module.
