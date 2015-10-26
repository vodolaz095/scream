Scream
======================

Application for recieving notifications by POST requests via `curl` and
displaying them via `notification-daemon` service in modern `linux` distros

Requirements
======================

You need to install package `notification-daemon` to be able to start application.
It can be done via your favourite package manager, like `dnf`

```shell

	$ su -c 'dnf install notification-daemon'

```


Usage
======================

Firstly, you need to start application on you desctop PC.

```shell

	$ scream --key=someHardKeyToMakeSpammersSad --listen=0.0.0.0:8082

```

Also you need to open port for this application

```shell

	su -c 'firewall-cmd --permanent --add-port=8082/tcp'

```

On `Fedora 22` with `LXDE` environment, you can autostart your application
by using the startup scripts to start `noficiation-daemon` and `scream` as the ones
in `contrib/.config/autostart/`. You need to save them to `~/.config/autostart`.

Then, you can send notifcations to your desctop from remote\local machine using this, or analogous CURL commands

```shell

    $ curl -X POST -d '{"type":"Notification","message":"Hello!","key":"topSecret"}' http://localhost:8082/

    $ curl -X POST -d type=Notification -d message=Hello! -d key=topSecret  http://localhost:8082/

```


Parameters:
======================

-  `key` - secret key to verify that message is send from known host
-  `listen` - address to bind application to - examples - `127.0.0.1:8082`, `0.0.0.0:8082`,`192.168.1.2:8082`. Default is `0.0.0.0:8082`




The MIT License (MIT)
======================


Copyright (c) 2015 Ostroumov Anatolij <ostroumov095 at gmail dot com>

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
