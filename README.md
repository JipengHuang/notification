# use notification 


CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -o ./notfli main.go

CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build  -o ./notfli main.go

./notfli 

-name   指定项目名称

-url    指定webbookurl

-branch 指定branch

-user   指定发起者

-time   指定时间

-result 指定结果

-details 指定链接 

