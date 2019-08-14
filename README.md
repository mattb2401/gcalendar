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

First setup your application credentials from Google console by creating a .env file with the following details 
```shell
{"installed":
    {
        "project_id":"<project ID provided by Google>",
        "client_id":"<client ID provided by Google>",
        "client_secret":"client secret provided by Google>",
        "auth_uri":"https://accounts.google.com/o/oauth2/auth",
        "token_uri":"https://oauth2.googleapis.com/token",
        "auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs",
        "redirect_uris":["urn:ietf:wg:oauth:2.0:oob","http://localhost"]
    }
}
```

To initialize the application

```shell
gcalendar init
```

To view a list of events in your calendar

```shell
gcalendar all
```

To view a list of events in your calendar for a particular date

```shell
gcalendar all -d <date forexample "2019-06-03">
```
To view a list of events in your calendar for today or tomorrow

```shell
gcalendar all -d today
```
or 

```shell
gcalendar all -d tomorrow
```
## TODO

- Work on test scenarios :) Never worked on them 
- Work on editing events
- Work on timezone issue when creating events on the calendar

## To contribute
To contribute just do a clone and when you are done just do a merge request.