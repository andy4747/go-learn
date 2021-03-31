read -p "  [1] Compile for Linux  [2] Compile for Windows => " platform
if [ $platform == "1" ]
then
    GOOS=linux GOARCH=amd64 go build -o executable
elif [ $platform == "2" ]
then
    GOOS=windows GOARCH=amd64 go build -o app.exe
else
    echo "Enter valid option."
fi