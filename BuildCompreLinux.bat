set GOOS=linux
go build -ldflags="-s -w"
del /f DiscCompressed
D:\upx\upx.exe -9 -oDiscCompressed "%~dp0DiscGo.discordgo"
PAUSE