# `powerline-go-moodle`

## A Moodle plugin for [`powerline-go`](https://github.com/justjanne/powerline-go)

This plugin adds a segment which displays the [Moodle](https://moodle.org) version the current directory belongs to.

![Moodle plugin for powerline go, demonstating different versions corresponding to different git branches](powerline-go-moodle.png)

It works by finding the parent directory containing a composer.json for a project called "moodle/moodle", then parses the $release variable from the version.php file. It also works for [Totara](https://www.totara.com/).

## Installation

Clone this repository and run `go install`

```bash
git clone https://github.com/marxjohnson/powerline-go-moodle
cd powerline-go-moodle
go install
```

This will place the `powerline-go-moodle` binary in your `$GODIR/bin` directory. Make sure this is added to your $PATH.

To confirm it is working, `cd` to a directory containing a Moodle codebase and run `powerline-go-moodle`

```bash
$ cd moodle
$ powerline-go-moodle
[{"name":"moodle","content":" M4.5dev+ ","foreground":15,"background":166}]
```

## Usage

Add `moodle` to the `-modules` argument when you configure `powerline-go` for your shell. 
For example, for Bash, your `.bashrc` should contain something like:

```bash
function _update_ps1() {
    PS1="$($HOME/go/bin/powerline-go \
	    -error $? \
	    -jobs $(jobs -p | wc -l) \
	    -hostname-only-if-ssh \
	    -modules 'aws,cwd,git,root,exit,moodle')"

    # Uncomment the following line to automatically clear errors after showing
    # them once. This not only clears the error for powerline-go, but also for
    # everything else you run in that shell. Don't enable this if you're not
    # sure this is what you want.

    #set "?"
}
```
