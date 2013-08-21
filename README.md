Joke  
=====
The web cosole for [godns](https://github.com/kenshinx/godns)

### Running

#### Developemt

```
$ go get github.com/kenshinx/joke
$ go get github.com/astaxie/bee

$ cd $GOPATH/src/github.com/kenshinx/joke

$ bee run joke  

```

*http://127.0.0.1:1223*


#### Productive

Compile first

```
$ go get github.com/kenshinx/joke
$ cd $GOPATH/src/github.com/kenshinx/joke
$ go build 
```

Deploy strongly suggest supervisord.

```

[program:joke]
command=/var/joke/joke
autostart=true
autorestart=true
user=joke
directory=/var/joke
stdout_logfile_maxbytes = 50MB
stdout_logfile_backups = 20
stdout_logfile = /var/log/joke.log

```


## Configuration

`conf/app.conf` 


```

#[joke]
appname = Joke
httpaddr = "127.0.0.1"
httpport = 1223
runmode = "dev"
autorender = false
autorecover = true
viewspath = "views"


#[auth]
#username:password.
#basic_auth = "joke:hello"


#[redis]
redisaddr = "127.0.0.1:6379"
redisdb = 0
redispassword = "hello"
bindkey = "godns:hosts"



#[log]
logfile = ""
logrorate = true

```


## Auth

Support http basic auth  


## Screenshot

*http://127.0.0.1:1223*

![joke](https://raw.github.com/kenshinx/joke/master/screenshot/joke.png)


### Dependence

* [github.com/astaxie/beego](https://github.com/astaxie/beego)  
* [github.com/hoisie/redis](https://github.com/hoisie/redis)
