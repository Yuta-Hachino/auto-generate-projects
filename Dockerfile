FROM golang:latest
WORKDIR /app
RUN go mod init github.com/Yuta-Hachino/auto-generate-projects # ご自身のモジュールで
RUN go install github.com/cosmtrek/air@latest
RUN air init
# CMD [ "go", "run", "main.go" ]
