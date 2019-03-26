# Play gomvpkg
This is sample to show you how `gomvpkg` tool can  helps for refactoring Go project.

# Caution!
`gomvpkg` have not been supported by Go module yet. So you need to use it in GOPATH. 

# How to play
- Folk repo to GOPATH
- Set GO111MODULE with shell `export GO111MODULE=auto`
- Get gomvpkg with command `go get golang.org/x/tools/cmd/gomvpkg`
- Checkout to branch `before-moving`
- Run command `gomvpkg -from github.com/[your git account]/play-gomvpkg/common/consumer -to github.com/[your git account]/play-gomvpkg/kafka`
# Sample
```
gomvpkg -from github.com/ballweera/play-gomvpkg/common/consumer -to github.com/ballweera/play-gomvpkg/kafka
```

After followed above steps, you should see message to confirm that `github.com/[your git account]/play-gomvpkg/common/consumer` is renamed to `github.com/[your git account]/play-gomvpkg/kafka` and related files are moved to `kafka` directory.
