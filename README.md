<div id="top"></div>
<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="https://i.postimg.cc/Vv01tHTQ/gopher-1.png" alt="Logo" width="120" height="120">
  </a>

  <h3 align="center">Sparkle-Todo</h3>

  <p align="center">
    An extensive application to manage your todos better and quick
    <br />
    <a href="https://github.com/hiteshhedwig/sparkle-todo"><strong>Explore the more »</strong></a>
    <br />
    <br />
    <a href="https://github.com/hiteshhedwig/sparkle-todo#Usage">View Demo</a>
    ·
    <a href="https://github.com/hiteshhedwig/sparkle-todo/issues">Report Bug</a>
    ·
    <a href="https://github.com/hiteshhedwig/sparkle-todo/issues">Request Feature</a>
  </p>
</div>

<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="https://i.postimg.cc/TwwBn1PW/carbon-9.png" alt="Logo" width="500" height="300">
  </a>
 </div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#quick-start">Quick start</a>
      <ul>
        <li><a href="#install">Install</a></li>
        <li><a href="#usage">Usage</a></li>
      </ul>
    </li>
    <li><a href="#command-support">Command Support</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>


## Quick start

### Install

#### MacOS installation :

```
brew tap hiteshhedwig/tools
brew install sparkle
```

#### Ubuntu AMD64 installation:
```
sudo sh -c "echo 'deb [trusted=yes] https://sparkle.jfrog.io/artifactory/sparkletodo-debian-local trusty main' >> /etc/apt/sources.list"
sudo apt update && sudo apt-get install sparkle-todo
```

### Windows :

Download the [windows binary](https://github.com/hiteshhedwig/sparkle/releases/tag/v0.1.1)

After downloading, extract it in your system. I would suggest create a folder in C drive Program Files named "sparkle" and paste 
the extracted .exe file there.

and follow these [steps](https://stackoverflow.com/a/41895179/8740644). They work on point


### Usage

[![carbon-7.png](https://i.postimg.cc/Bbq2Pf6t/carbon-7.png)](https://postimg.cc/ZBMWXg6h)

No grandiose setup required. You can simply get started by adding your task :
```
sparkle add "initiate project structure"
```
Add your project name :
```
sparkle projectname "sparkle-todos"
```

## Command Support 

### List 
By default below command will provide you current ongoing task.

[![carbon-10.png](https://i.postimg.cc/rpK522ZL/carbon-10.png)](https://postimg.cc/7Cr5SdyB)


📎 Note -> In order to display whole category wise list try adding "cat" 
```
> sparkle list --cat

╭────────┬───────────────────────────┬───────────────────────╮
│ INDEX  │ TASK                      │ TIME                  │
├────────┼───────────────────────────┼───────────────────────┤
│ ****** │ Current Task              │ ****************      │
├────────┼───────────────────────────┼───────────────────────┤
│ 1      │ initial project structure │ 2022-01-23 2:13:30 pm │
│ 2      │ refactor code             │ 2022-01-23 2:15:14 pm │
├────────┼───────────────────────────┼───────────────────────┤
│ ****** │ Completed Task            │ ****************      │
├────────┼───────────────────────────┼───────────────────────┤
│ ****** │ Cancelled Task            │ ****************      │
╰────────┴───────────────────────────┴───────────────────────╯
```

### Cancel 

Let's see how to properly cancel any todo

```
> sparkle cancel --help

This command helps you cancel a specific task. Just pass the index in argument.
        You can know index by : sparkle list --cat
        . For example:
                sparkle cancel 7

        NOTE : it always cancel task from current list of task. If you wanna do remove task from other categories, consider
                        sparkle purge --help

Usage:
  sparkle cancel [flags]

Flags:
  -h, --help   help for cancel
 ```
Now let's see current task we have
```
> sparkle list 
╭───────┬───────────┬──────────────────────╮
│ INDEX │ TASK      │ TIME                 │
├───────┼───────────┼──────────────────────┤
│     1 │ task 1    │ 2022-01-23 4:4:32 pm │
├───────┼───────────┼──────────────────────┤
│     2 │ task 2    │ 2022-01-23 4:4:35 pm │
├───────┼───────────┼──────────────────────┤
│     3 │ task 3    │ 2022-01-23 4:4:37 pm │
├───────┼───────────┼──────────────────────┤
│     4 │ typo task │ 2022-01-23 4:4:47 pm │
╰───────┴───────────┴──────────────────────╯  
```

If we wish to cancel the last task "typo task". We will have to provide its index 4 to the command.
```
> sparkle cancel 4
Are you sure you want to cancel this task?  typo task
Type yes/no to confirm cancellation : yes
Task cancelled
```

### Finish 

```
> sparkle --help

For example: You can use below command to change the status of a task at index 1 

sparkle finish 1

Make sure you are providing index of the task you wish to finish.

Usage:
  sparkle finish [flags]

Flags:
  -h, --help   help for finish

```
Now let's see current task we have
```
> sparkle list 
╭───────┬──────────────────┬───────────────────────╮
│ INDEX │ TASK             │ TIME                  │
├───────┼──────────────────┼───────────────────────┤
│     1 │ task 1           │ 2022-01-23 4:4:32 pm  │
├───────┼──────────────────┼───────────────────────┤
│     2 │ task 2           │ 2022-01-23 4:4:35 pm  │
├───────┼──────────────────┼───────────────────────┤
│     3 │ task 3           │ 2022-01-23 4:4:37 pm  │
├───────┼──────────────────┼───────────────────────┤
│     4 │ task to complete │ 2022-01-23 4:15:30 pm │
╰───────┴──────────────────┴───────────────────────╯
```

If we wish to finish the last task "task to complete". We will have to provide its index 4 to the command.
```
> sparkle finish 4
🥧
 You have completed the task : task to complete 

```
### Project Name 

```
> sparkle projectname "Sparkle Todos"
Setting Project name from : Sparkle Todo -> Sparkle Todos

```

### Purge 

```
> sparkle purge --help

Command helps you purging specific task from cancelled list. For example:

        Note : you will only be able to purge a task from cancelled list. If you want
                   to do so with current task list. Please first cancel it then purge.
                   
                   for example : 
                                   -> sparkle purge [index of the task]
                                                  sparkle purge 1

Usage:
  sparkle purge [flags]

Flags:
  -h, --help   help for purge
```

listing the task category wise 
```
> sparkle list --cat
╭────────┬──────────────────┬───────────────────────╮
│ INDEX  │ TASK             │ TIME                  │
├────────┼──────────────────┼───────────────────────┤
│ ****** │ Current Task     │ ****************      │
├────────┼──────────────────┼───────────────────────┤
│ 1      │ task 1           │ 2022-01-23 4:4:32 pm  │
├────────┼──────────────────┼───────────────────────┤
│ ****** │ Completed Task   │ ****************      │
├────────┼──────────────────┼───────────────────────┤
│ 1      │ task to complete │ 2022-01-23 4:15:52 pm │
│ 2      │ task 3           │ 2022-01-23 4:45:47 pm │
├────────┼──────────────────┼───────────────────────┤
│ ****** │ Cancelled Task   │ ****************      │
├────────┼──────────────────┼───────────────────────┤
│ 1      │ typo task        │ 2022-01-23 4:10:8 pm  │
│ 2      │ typo task        │ 2022-01-23 4:13:30 pm │
│ 3      │ task 2           │ 2022-01-23 4:46:15 pm │
╰────────┴──────────────────┴───────────────────────╯
```
Let's purge the task

```
> sparkle purge 1
Purging :  typo task
```
## Stash 

```
> sparkle stash --help

It is just "git add [flags]"
         For example:

                        sparkle stash "cmd/config.go cmd/root.go .sparkle"

Usage:
  sparkle stash [flags]

Flags:
  -h, --help   help for stash
```
Note : you can simply use `git add` if needed. Just make sure to stash `.sparkle` to allow git to keep track of todos.
```
> sparkle stash ".sparkle config/config.go"
╭───────┬────────┬──────────────────────╮
│ INDEX │ TASK   │ TIME                 │
├───────┼────────┼──────────────────────┤
│     1 │ task 1 │ 2022-01-23 4:4:32 pm │
╰───────┴────────┴──────────────────────╯
🗳 Files has been stashed. Please proceed with git commit -
```

## Update Task 

```
> sparkle updatetask --help
Allows you to update a particular task in current list. Just provide index to the flag For example :

                -> sparkle updatetask [flag]
                   sparkle updatetask 1

Usage:
  sparkle updatetask [flags]

Flags:
  -h, --help   help for updatetask

```
Provide index of the task you want to update 
```
> sparkle updatetask 1
Current task is : task 1 
Please type updated task  : task updated

```
## Version 

```
> sparkle version
You are using sparkle : v0.1.0 
```

## Roadmap 

📋 To be decided

## Contributing

Feel free to send any pull request !

## License

[License](https://github.com/hiteshhedwig/sparkle-todo/blob/exp1/LICENSE)





