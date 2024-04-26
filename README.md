
https://velog.io/@vamos_eon/go-fmt

#wsl 환경 셋팅
cd /usr/local
sudo wget https://dl.google.com/go/go1.22.2.linux-amd64.tar.gz // 가장 최신버전 으로 설치
sudo tar -xvf go1.22.2.linux-amd64.tar.gz        // 압축해제
sudo rm -rf go1.22.2.linux-amd64.tar.gz          // 압축파일 삭제

vi ~/.profile    // 아래 3개 변수 추가
---------------------------------
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
---------------------------------
source ~/.profile  // 설정파일 적용

go version     // go 버전 체크 하면서 설치 잘 됐는지 확인

