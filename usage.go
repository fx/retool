package main

var usage = `usage: retool (add | remove | upgrade | sync | do | clean | build | help)

use retool with a subcommand:

add will add a tool
remove will remove a tool
upgrade will upgrade a tool
sync will synchronize your _tools with tools.json, downloading if necessary
build will compile all the tools in _tools
do will run stuff using your installed tools
clean will delete the repo cache stored at ~/.retool

help [command] will describe a command in more detail
`

var addUsage = `usage: retool add [repository] [commit]

eg: retool add github.com/tools/godep 3020345802e4bff23902cfc1d19e90a79fae714e

Add will mark a repository as a tool you want to use. It will rewrite
tools.json to record this fact. It will then fetch the repository,
reset it to the desired commit, and install it to _tools/bin.

You can also use a symbolic reference, like 'master' or
'origin/master' or 'origin/v1.0'. Retool will end up parsing this and
storing the underlying SHA.
`

var upgradeUsage = `usage: retool upgrade [repository] [commit]

eg: retool upgrade github.com/tools/godep 3020345802e4bff23902cfc1d19e90a79fae714e

Upgrade set the commit SHA of a tool you want to use. It will
rewrite tools.json to record this fact. It will then fetch the
repository, reset it to the desired commit, and install it to
_tools/bin.

You can also use a symbolic reference, like 'master' or
'origin/master' or 'origin/v1.0'. Retool will end up parsing this and
storing the underlying SHA.
`

var removeUsage = `usage: retool remove [repository]

eg: retool remove github.com/tools/godep

Remove will remove a tool from your tools.json. It won't delete the
underlying repo from _tools, because it might be a dependency of some
other tool. If you want to clean things up, retool sync will clear out
unused dependencies.
`

var syncUsage = `usage: retool sync

Sync will synchronize your _tools directory to match tools.json.
`

var doUsage = `usage: retool do [command and args]

retool do will make sure your _tools directory is synced, and then
execute a command with the tools installed in _tools.

This is just
  retool sync && PATH=$PWD/_tools/bin:$PATH [command and args]
That works too.
`

var cleanUsage = `usage: retool clean

retool clean will delete the repo cache stored at ~/.retool

This is just
  rm -rf ~/.retool
That works too.
`

var buildUsage = `usage: retool build

retool build will compile all the tools listed in tools.json, obeying whatever is currently
downloaded into _tools. It will not do additional network calls. This is typically useful for
compiling vendored tools so you can use them inside isolated environments.
`
