# 포인터란?
- 메모리의 주소

# 인스턴스란?
- 메모리의 실체

# 포인터 사용이유
- 메모리 영역에 직접 접근하여 인스터스를 조작할수있다
 즉, 어디서든 인스턴스에 접근할 수 있으면, 메모리를 중복을 줄이고 불필요한 메모리 낭비를 막을 수 있다.

- 퍼포먼스를 신경쓸때 포인터를 고려하지 않을 수 없다.(효율적인 메모리 접근및 관리성)
ex)  인스턴를 통째로 복사할때, 인스턴스의 주소(포인터)만 넘김으로써 메모리 낭비를 줄일 수 있다.


# 포인터 변수의 기본값(zero value) : nil
- nil은 출력할 때 항상 <, >을 포함한다
- ( 포인터 변수 = nil ) === 인스턴스가 없다

# 포인터 자료형 출력 : %p

# 스택 & 힙
- 스택 메모리 : 함수 호출시 자동으로 할당되는 메모리 , 함수종료시 자동 정리
- 힙 메모리 : 프로그래머가 수동 할당하는 메모리, 수동해제 필요하지만!
            - 힙은 엑세스와 메모리 해제가 스택에 비해 느리다 (비용이 비싸다)
> go에서는 스택이든 힙이든 가비지 컬렉터가 알아서 메모리 정리

# 탈출분석
- 함수에서 함수로 인스턴스가 탈출하는지 여부를 분석하는것.
- 외부로 탈출되면 스택 메모리에 할당하는것은 적절치 않음. 힙에다 할당하는게 좋음

# 힙메모리 할당 확인
- go build -gcflags '-m -l'     // 옵티마이징 레벨을 1로 하고, 인라이닝을 없앱니다.
- go build -gcflags '-m'        // 옵티마이징 레벨을 1로 합니다.
= go build -gcflags '-m=2'      // 옵티마이징 레벨을 2로 합니다. (보다 상세하게 나옵니다.)

