# project goblog (working name)
a CS50 final project, and yeah, I know, the name's really bad

it doesn't really have any functionality at the moment, it uses a mockup backend at the moment 

## How to use
### Without Docker on Linux or similar:
- Make sure you have a recent version of go toolchain installed
- Clone master / download a [release] of goblog
- Install dependencies using `go get ./..` within the directory
- Build the executable using `go build -o run .`
- Run the executable: `./run`
- Config.json with default settings will be made on first startup, you can change any of the settings and restart
### Without Docker on Windows:
- Same as on Linux, just instead of `go build -o run .` use  `go build -o run.exe` and start `run.exe` instead of doing `./run`
### With Docker:
- It's recomended to run everything as sudo/root, since only root can map ports below 1024
- Clone master / download a [release] of goblog
- Change the docker.config.json according to your needs first (except for ListenOn)
- Run this command within the directory to create docker image: `docker build -t goblog .`
- Afterward, change `<your_port>` to your own preffered port and run this command to start a container with the goblog image we created earlier: `docker run -d -p 80:<your_port> --name goblog localhost/goblog` 

[release]: https://github.com/david-sorm/goblog/releases