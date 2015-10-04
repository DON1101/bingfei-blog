# bingfei-blog
A blog system for Bingfei by Tang Guojing

# Install gvm for Go language
Refer to [https://github.com/moovweb/gvm](https://github.com/moovweb/gvm) and [http://www.ascent.io/blog/2014/03/11/gvm-with-golang/](http://www.ascent.io/blog/2014/03/11/gvm-with-golang/)

# Add the following to the bottom of your .bashrc or .bash_profile
```
# gvm config
[[-s "$HOME/.gvm/scripts/gvm"]] && source "$HOME/.gvm/scripts/gvm"
```

## List all Go version installed
```
gvm listall
```

# Install Go1.4 and set as default
```
gvm install go1.4
gvm use go1.4 --default
```

# Create a Project specific package set
```
gvm pkgset create bingfei
gvm pkgset use bingfei
```

# Configure your Golang workspace
`gvm pkgenv bingfei`

## Edit GOPATH and PATH environment variables
```
export gvm_pkgset_name="bingfei"
export GOPATH; GOPATH="${HOME}/Projects/bingfei-blog:/Users/guojing/.gvm/pkgsets/go1.4/bingfei:$GOPATH"
export PATH; PATH="/Users/guojing/.gvm/pkgsets/go1.4/bingfei/bin:$PATH"
# Package Set-Specific Overrides
export GVM_OVERLAY_PREFIX; GVM_OVERLAY_PREFIX="${GVM_ROOT}/pkgsets/go1.4/bingfei/overlay"
export PATH; PATH="/Users/guojing/.gvm/pkgsets/go1.4/bingfei/bin:${GVM_OVERLAY_PREFIX}/bin:${PATH}"
```

# Go package install from requirements.txt
```
go get github.com/beego/bee
go get github.com/lib/pq
go get github.com/astaxie/beego/orm
```

# Make copy local_settings.go
`cp local_settings.template local_settings.go`

# Build Go
`./build.sh`

# Run Go
`./main`

# Static package management
## Initialization
```
npm install webpack-dev-server
```
