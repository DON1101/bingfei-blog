# bingfei-blog
A blog system for Bingfei by Tang Guojing

## Install PostgreSQL
### Mac OS

**Download Mac version from [http://www.postgresql.org/download/](http://www.postgresql.org/download/)**

**Locate psql in your Mac machine**
```
locate psql | grep /bin
```
A sample output would be:
```
/Applications/Postgres.app/Contents/Versions/9.4/bin/psql
```

You may encouter below error:
```
WARNING: The locate database (/var/db/locate.database) does not exist.
To create the database, run the following command:

  sudo launchctl load -w /System/Library/LaunchDaemons/com.apple.locate.plist

Please be aware that the database can take some time to generate; once
the database has been created, this message will no longer appear.
```

Then you need run `sudo launchctl load -w /System/Library/LaunchDaemons/com.apple.locate.plist`

If the error still persists, you may need run `sudo /usr/libexec/locate.updatedb`

**Append below line to ~/.bash_profile**
```
export PATH=/Applications/Postgres.app/Contents/Versions/9.4/bin/:$PATH
```

**After having saved the file, read the file**
```
. ~/.bash_profile
```

**try psql in your terminal command**

## Install gvm for Go language
Refer to [https://github.com/moovweb/gvm](https://github.com/moovweb/gvm) and [http://www.ascent.io/blog/2014/03/11/gvm-with-golang/](http://www.ascent.io/blog/2014/03/11/gvm-with-golang/)

## Add the following to the bottom of your .bashrc or .bash_profile
```
## gvm config
[[-s "$HOME/.gvm/scripts/gvm"]] && source "$HOME/.gvm/scripts/gvm"
```

## List all Go version installed
```
gvm listall
```

## Install Go1.4 and set as default
```
gvm install go1.4
gvm use go1.4 --default
```

## Create a Project specific package set
```
gvm pkgset create bingfei
gvm pkgset use bingfei
```

## Configure your Golang workspace
`gvm pkgenv bingfei`

**Edit GOPATH and PATH environment variables**
```
export gvm_pkgset_name="bingfei"
export GOPATH; GOPATH="${HOME}/Projects/bingfei-blog:/Users/guojing/.gvm/pkgsets/go1.4/bingfei:$GOPATH"
export PATH; PATH="${HOME}/Projects/bingfei-blog/bin:/Users/guojing/.gvm/pkgsets/go1.4/bingfei/bin:$PATH"
## Package Set-Specific Overrides
export GVM_OVERLAY_PREFIX; GVM_OVERLAY_PREFIX="${GVM_ROOT}/pkgsets/go1.4/bingfei/overlay"
export PATH; PATH="/Users/guojing/.gvm/pkgsets/go1.4/bingfei/bin:${GVM_OVERLAY_PREFIX}/bin:${PATH}"
```

Every time you update your pkgenv, you need to run `gvm pkgset use bingfei` again

## Go package install from requirements.txt
```
go get github.com/beego/bee
go get github.com/lib/pq
go get github.com/astaxie/beego/orm
```

## Make copy local_settings.go
`cp local_settings.template local_settings.go`

## Run Go application
**Run directly**

1. Build Go
`./build.sh`
2. Run Go
`./main`

**Run from Beego**
```
bee run
```

## Static package management
**Initialization**
```
npm install webpack-dev-server
```

## Unit test
```
./test.sh
```
