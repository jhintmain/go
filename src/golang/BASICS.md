# Go 프로그래밍 언어 소개
- 2012 오픈 
- GO 공식 사이트: https://go.dev/
- 만든이 : google v8 javascript 엔진 개발에 참여자, 유닉스 개발 참여자, b언어 개발자가 참여하여 만들었음.

# Go 언어 특징
- 일차적으로 시스템 프로그래밍을 위해 개발되었으며, C++, Java, Python의 장점들을 뽑아 만들어졌다
- 컴파일 언어
- 정적타입 언어
- Garbage Collection 기능을 제공

# Go 모듈 초기화 하기
- go mod init [~]
- Repository에 프로젝트를 업로드하고 다른 사람과 공유하기 위해서는 보다 공식적인 모듈명을 사용한다
- ex ) go mod init jhpark1119.com/go

# Go 실행하기
1. go run hello.go : 빌드 후 실행 > 실행파일 만들어지지않음
2. go build hello.go : 빌드 실행 > 실행파일 생성
   ./hello : 실행파일 실행
   
# Go 표준 라이브러리 
https://pkg.go.dev/std