`gcalendar` is a tiny program to view, create, update and delete [Google Calendar](https://calendar.google.com) events.

## Installing

### Requirements

go 1.9.X or higher is required. See [here](https://golang.org/doc/install) for installation instructions and platform installers.

* Make sure to set your GOPATH in your env, .bashrc or .bash\_profile file. If you have not yet set it, you can do so like this:

```shell
cat << ! >> ~/.bashrc
> export GOPATH=\$HOME/gopath
> export PATH=\$GOPATH:\$GOPATH/bin:\$PATH
> !
source ~/.bashrc # To reload the settings and get the newly set ones # Or open a fresh terminal
```

### From sources

To install from the latest source, run:

```shell
go get -u github.com/mattb2401/gcalendar
```

## Usage
To initialize the application

```shell
./gcalendar init
```

To view a list of events in your calendar

```shell
./gcalendar all
```

To view a list of events in your calendar for a particular date

```shell
./gcalendar all -d <date forexample "2019-06-03">
```
To view a list of events in your calendar for today or tomorrow

```shell
./gcalendar all -d today
```
or 

```shell
./gcalendar all -d tomorrow
```

