
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

# Golang에는 삼항 연산자가 없다
# Golang은 기본 인코딩이 UTF-8
# 현재 패키지 외부에서도 쓰일 수 있게 하려면 대문자, 현재 패키지 내부에서만 쓰일 수 있게 하려면 그 외의 문자

# break문을 사용하지 않아도 case 한 개만 실행하고 탈출한다는 점입니다.
- break문은 컴파일 단계에서 컴파일러가 추가해 줍니다.

# Golang에서의 반복문은 오로지 for를 사용하는 방법만 존재합니다.


# array
- 배열은 초기에 지정한 크기를 변경할 수 없습니다.
- 값을 지정하지 않으면 초기값은 지정 type의 zero-value
- 권장 선언 array 선업 type 형식
 > var a [5]int = [5]int{1, 2, 3, 4}
1. 짧은 선언 & 초기화
 > b := [3]bool{true, false, true}
2. 인덱스 지정 선언 & 초기화
 > c := [5]float64{1: 1.1, 3: 3.3}
3. 요소갯수로 크기 지정
 > d := [...]string{"one", "two", "three"}
4. 다차원 배열 선언
 > var mul [2][3]int = [2][3]int{
	{1, 2, 3}, 
	{4, 5, 6},      // 콤마(,)필수
}
- 배열은 메모리 연역에 할당이 연속적
- golang은 최강타입 언어라고 했습니다. 타입이나 크기가 다르면 값의 복사를 불허합

