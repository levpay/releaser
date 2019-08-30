## Releaser

App release manager for Levpay's applications. Allows tagging versions without all the necessary cmding and editing.

## Dependancy

It strongly depends on git and [goreleaser](https://goreleaser.com/) CLI to build your apps.

## Why

It allows us to release across multiple apps without the need of adding a `goreleaser.yml` file on each one of them.

## Install 

```
    $ go get -v github.com/levpay/releaser
```

## Usage

### New release of an app

It will ask you about your new version:

```
➜  releaser git:(master) ✗ releaser n releaser
Latest release of releaser: v1.0.1

Chose the new semantic version of the app releaser: |
```

Than it will use `goreleaser` to compile the binaries and generate everything needed (changelog, download deps etc). 

### 

## Edge cases

* It does not generate the first version of an app.
* It only works with go code.
* If it fails to release you have to fix the problems, erase the latest release from local machine and try again.

#### Fail case

In this example the git repository `tibiahouses` is in a dirty state (files are not commited or staged), so it fails. To fix it you have to commit or stash the files, erase the local tag `v1.1.0` and then retry the `releaser new <appname>` command.
```
➜  tibiahouses git:(master) ✗ releaser n tibiahouses
Latest release of tibiahouses: v1.0.0

Chose the new semantic version of the app tibiahouses: v1.1.0
v1.1.0

2019/08/30 09:56:09 
   • releasing using goreleaser 0.110.0...
   • loading config file       file=/Users/arxdsilva/go/src/github.com/levpay/releaser/goreleaser.yml
   • RUNNING BEFORE HOOKS     
      • running go mod download  
   • GETTING AND VALIDATING GIT STATE
      • releasing v1.1.0, commit ea63e2781c9df5e076265ee05322407c0a7144a4
   ⨯ release failed after 0.08s error=git is currently in a dirty state, please check in your pipeline what can be changing the following files:
?? actions/workers/houses.go

exit status 1
```