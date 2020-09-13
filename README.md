cowweb-go
===
![Go build](https://github.com/hhiroshell/cowweb-go/workflows/Go%20build/badge.svg)
![Docker build and push](https://github.com/hhiroshell/cowweb-go/workflows/Docker%20build%20and%20push/badge.svg)

Cosay Web API.

```
$ curl "http://localhost:8080/say?m=Hello%20cowweb"
 ______________
< Hello cowweb >
 --------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||--WWW |
                ||     ||
```

How to build and run
--------------------

#### 1. Clone this repository

```
git clone https://github.com/hhiroshell/cowweb-go.git && cd cowweb-go
```

#### 2. Build and run

```
go build .
```
```
./cowweb serve
```

#### 3. Call the API
You can call the API via localhost:8080 .

```
curl "http://localhost:8080/say"
```

And you can specify a message using "m" query (special characters have to be URL encorded).

```
curl "http://localhost:8080/say?m=hello%20cowweb"
```
