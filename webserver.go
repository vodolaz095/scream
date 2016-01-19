package scream

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func notifySend(title, message string) error {
	err := os.Setenv("DISPLAY", ":0")
	if err != nil {
		return err
	}
	cmd := exec.Command("/usr/bin/notify-send", title, message)
	//other variants
	//cmd := exec.Command("/usr/bin/zenity", "--info", fmt.Sprintf("--text='%v'", message), "--display=:0")
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func handeRequest(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8") // normal header
		rw.Write([]byte(html))
		return
	}

	exampleNotification := notification{
		Type:    "Notification",
		Message: "Hello!",
		Key:     "someLongKeyToMakeSpammersSad",
	}
	jsonExampleText, _ := json.Marshal(exampleNotification)
	var currentNotification notification

	err := req.ParseForm()
	if err != nil {

		http.Error(rw, fmt.Sprintf("Unable to parse request! It have to be request invoked like this %v", fmt.Sprintf("$ curl -X POST -d type=Notification -d message=Hello! -d key=%v  http://localhost%v/", Cfg.Key, Cfg.Address)), 400)
		return
	}

	if req.Form["type"] != nil && req.Form["message"] != nil && req.Form["key"] != nil {
		currentNotification.Type = req.Form["type"][0]
		currentNotification.Message = req.Form["message"][0]
		currentNotification.Key = req.Form["key"][0]
	} else {
		decoder := json.NewDecoder(req.Body)
		err = decoder.Decode(&currentNotification)
		if err != nil {
			http.Error(rw, fmt.Sprintf("Unable to parse request! It have to be JSON like this %v", string(jsonExampleText)), 400)
			return
		}
	}

	if currentNotification.Type == "" {
		err = fmt.Errorf("Type is missing!")
	}
	if currentNotification.Message == "" {
		err = fmt.Errorf("Message is missing!")
	}
	if currentNotification.Key == "" {
		err = fmt.Errorf("Key is missing!")
	}

	if currentNotification.Key != Cfg.Key {
		http.Error(rw, fmt.Sprintf("Wrong key! Access denied!"), 403)
		return

	}

	if err != nil {
		http.Error(rw, fmt.Sprintf("Wrong request - %v", err.Error()), 400)
		return
	}

	err = notifySend(currentNotification.Type, currentNotification.Message)
	if err != nil {
		panic(err)
	}

	if err != nil {
		http.Error(rw, fmt.Sprintf("Error sending notification - %v", err.Error()), 500)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

// StartServer listens to http connections on port defined in Cfg.Address
func StartServer() error {
	http.HandleFunc("/", handeRequest)
	example1 := fmt.Sprintf("$ curl -X POST -d '{\"type\":\"lalala\",\"message\":\"that\",\"key\":\"%v\"}' http://localhost%v/", Cfg.Key, Cfg.Address)
	example2 := fmt.Sprintf("$ curl -X POST -d type=Notification -d message=Hello! -d key=%v  http://localhost%v/", Cfg.Key, Cfg.Address)
	fmt.Printf("Notification daemon is listening on 0.0.0.0:8082\nSend notification via: \n  %v\n or\n  %v\n", example1, example2)
	return http.ListenAndServe(Cfg.Address, nil)
}

var html = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Scream notification service</title>
</head>
<style>
    body {
        background-color: #ddd;
        font-family: Georgia, sans-serif;
        line-height: 140%;
        font-size: 100%;
    }

    h1 {
        font-size: 1.4em;
    }

    h2 {
        font-size: 1.2em;
    }

    div#wrap {
        margin: auto;
        width: 95%;
        background-color: #fff;
        padding: 1em;
        border: .4em solid #ccc;
        border-radius: 2em;
        max-width: 1400px;
    }

    div#header {
        background-color: #356aa0;
        color: #fff;
        padding: 1em;
        margin-bottom: 0;
        border-radius: 2em 2em 0 0;
    }

    div#header img {
        float: left;
        padding: 1.5em;
    }

    div#header h1 {
        margin: .5em;
        padding: 0;
    }

    div#header h2 {
        margin: .5em;
        padding: 0;
    }

    div#nav {
        background-color: #333;
        color: #eee;
        float: left;
        width: 100%;
        text-align: center;
        margin-bottom: .5em;
        border-radius: 0 0 2em 2em;
    }

    div#nav ul {
        margin: 0;
        padding: 0;
        float: left;
    }

    div#nav ul:first-child {
        margin-left: 2em;
        border-left: 1px solid #ccc;
    }

    div#nav li {
        display: inline;
        margin: 0;
        padding: 0;
        margin-right: .5em;
    }

    div#nav li a {
        font-size: .8em;
        padding: .25em .75em;
        text-decoration: none;
        color: #eee;
        float: left; /*width:6em;*/
        border-right: .1em solid #ccc;
    }

    div#nav li a.private {
        color: yellow;
    }

    div#nav li a:hover {
        color: #eee;
        background-color: #555;
    }

    div#content {
        padding: 2em 1em;
        float: none;
    }

    div#content h2 {
        border-bottom: .1em solid #ddd;
        margin-bottom: .5em;
    }

    div#content img.fancy {
        border: 1px solid #ddd;
        padding: .5em;
    }

    div#content img.left {
        float: left;
        margin-right: 1em;
    }

    div#content img.right {
        float: right;
        margin-left: 1em;
    }

    div#content p {
        clear: both;
    }
</style>
<body>

<div id="wrap">
    <div id="header">
        <h1>Scream notification service</h1>
        <h2>You can send notifications to user working directly on this PC.</h2>
    </div>
    <div id="nav">
        <ul>
            <li><a href="/">Main</a></li>
            <li><a target="_blank" href="https://github.com/vodolaz095/scream">Source code on Github.com</a></li>
            <li><a href="https://github.com/vodolaz095/scream/blob/master/CHANGELOG.md" target="_blank">Version: ` + VERSION + `</a></li>
            <li><a href="https://github.com/vodolaz095/scream/blob/master/README_RU.md" target="_blank">Руководство пользователя</a></li>
            <li><a href="https://github.com/vodolaz095/scream/blob/master/README.md" target="_blank">Manual</a></li>
            <li><a href="https://github.com/vodolaz095/scream/issues" target="_blank">Report a bug</a></li>
            <li><a href="https://godoc.org/github.com/vodolaz095/scream" target="_blank">Code documentation</a></li>
        </ul>
    </div>
    <div id="content">
        <h3>Test in action via form</h3>

        <form method="post" action="/" onsubmit="return sendNotification()">
            <p>
                <label for="key">Key</label>
                <input id="key" name="key" type="text" value="" placeholder="Some long secret key">
            </p>

            <p>
                <label for="type">Type</label>
                <input id="type" name="type" type="text" value="notification" placeholder="notification">
            </p>

            <p>
                <label for="message">Message</label>
                <input id="message" name="message" type="message" value="" placeholder="Hello!">
            </p>

            <p>
                <input type="submit" value="Send">
                <input type="reset" value="Cancel">
            </p>
        </form>
        <p>Response:</p>
        <pre id="response"></pre>
        <h3>Usage via CURL</h3>

        <p>You can issue this curl commands to send notification:</p>

        <p>
            <code>
                $ curl -X POST -d '{
                \"type\":\"notification\",
                \"message\":\"there is a lot of meat!\",
                \"key\":\"someVeryLongKey\"}' <span id="hostname1"></span>
            </code>
        </p>

        <p>
            <code>
                $ curl -X POST -d type=notification -d message="there is a lot of meat!" -d key="someVeryLongKey" <span id="hostname2"></span>
            </code>
        </p>
    </div>
    <div id="footer">
        <h2>The MIT License (MIT)</h2>
        <p></p>
        <p>Copyright (c) 2015 Ostroumov Anatolij <b>ostroumov095 at gmail dot com</b></p>
        <p></p>
        <p>Permission is hereby granted, free of charge, to any person obtaining a copy of
        this software and associated documentation files (the "Software"), to deal in
        the Software without restriction, including without limitation the rights to
        use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
        the Software, and to permit persons to whom the Software is furnished to do so,
        subject to the following conditions:</p>
        <p></p>
        <p>The above copyright notice and this permission notice shall be included in all
        copies or substantial portions of the Software.</p>
        <p></p>
        <p>THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
        IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
        FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
        COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
        IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
        CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.</p>
    </div>
    <script>
        var hstnm = window.location.protocol + '//' + window.location.host + '/';
        document.getElementById('hostname1').innerHTML = hstnm;
        document.getElementById('hostname2').innerHTML = hstnm;
        function sendNotification() {
            var request = new XMLHttpRequest();
            request.open('POST', '/');
            request.setRequestHeader('Content-Type', 'application/json');
            request.onreadystatechange = function () {
                if (this.readyState === 4) {
                    if (this.status === 201) {
                        document.getElementById('response').innerHTML = "Notification delivered";
                    } else {
                        document.getElementById('response').innerHTML = this.responseText;
                    }
                }
            };
            var body = {
                key: document.getElementById('key').value,
                type: document.getElementById('type').value,
                message: document.getElementById('message').value
            };
            request.send(JSON.stringify(body));
            return false;
        }
    </script>
</div>
</body>
</html>`
