# gin_template
Gin structure for go devs

## Why we need it:

When building my own website: http://mroomsoft.com. It takes me hours (maybe days :) ) to design a good structure for API. I must repeat the steps I have done before, it's really bullsh*t.
That's why I need to do something. Let's check out this repository and see how long we take to make a web service.

## Requirements:

- Golang
- Dep

## How to run:

- go get -u github.com/trongdth/gin_template.git
- cd $GOPATH/src/github.com/trongdth/gin_template
- dep ensure
- mv config/conf.json.example config/conf.json (here is my sample config)

```json
    {
        "env": "localhost",
        "port": 8000,
        "db": "root:@tcp(localhost:3306)/mroom_software?charset=utf8&parseTime=True&loc=UTC",
        "token_secret_key": "mroom_software"
    }
```

Congratulation!

P/S: if there are any issues, don't hesitate to contact me at trongdinh@mroomsoft.com
