# floRest
florest code repo

Go installation link
https://golang.org/doc/install

Use the download link.

##Fetch the project

```bash
cd <GOPROJECTSPATH>/
git clone https://github.com/jabong/florest-core
```

##Setup the application log

```bash
sudo mkdir /var/log/florest/
chown <user_name who is running ./florest> /var/log
```

##Make (build, install, test)

```bash
cd <GOPROJECTSPATH>/florest/
make deploy ENV=dev
cd <GOPROJECTSPATH>/bin/
./florest
```

##Example url

```
http://localhost:8080/florest/v1/hello
```

##Steps to bootstrap from floRest

```bash
make newapp NEWAPP="<your_project_name>"
```

##Steps to pull changes from floRest

* `florest` is available as library only. So just sync your `_libs/src/github.com/florest-core with github.com/jabong/florest-core`
* [Wiki](https://wiki.jira.rocket-internet.de/display/INDFAS/floRest+Framework)


##How to raise PR in floRest

Refer the how to contribute section in wiki [FAQ](https://wiki.jira.rocket-internet.de/display/INDFAS/florest-FAQ).


##Run coverage for local tests

To get test coverage for all the local tests. Execute the below command:-

```bash
make coverage
```

##Run coverage against external tests

Follow the below steps:-

1. `make coverall`
2. `cd bin`
3. `./floRest.test -test.coverprofile coverage.cov`
4. Now run your external tests against this build and it will keep on calculating coverage. Once all the
tests have run, just stop the program, `coverage.cov` will be generated in `bin/`
5. Run below command to convert coverage in html format
  
  ```bash
  go tool cover -html=coverage.cov -o coverage.html
  ```
  
