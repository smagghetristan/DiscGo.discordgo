set GOOS=windows
go build -ldflags="-s -w"
del /f DiscCompressed.exe
D:\upx\upx.exe -9 -oDiscCompressed.exe "%~dp0DiscGo.discordgo.exe"
PAUSE