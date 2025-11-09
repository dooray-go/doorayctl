# doorayctl
* dooray open api 를 사용하는 cli tool 입니다.

## 설치
* 사용할 컴퓨터의 아키텍처에 맞는 doorayctl 바이너리를 다운로드 합니다.
  * [doorayctl](https://github.com/dooray-go/doorayctl/releases)
* 실행가능한 디렉토리로 파일이 복사되면 됩니다.
  * 예) /usr/local/bin
```bash
$ sudo mv doorayctl.darwin.arm64 /usr/local/bin/doorayctl

```
## 설정
* doorayctl은 dooray의 open-api를 사용하는 cli 톨이므로, 사용하기 위해서는 dooray api token이 필요합니다.
* dooray api token은 dooray의 설정에서 발급받을 수 있습니다.
* doorayctl을 실행하기 전에 dooray api token을 다음의 파일에 저장합니다. 
    * ~/.dooray/config
```json
{
    "token":"YOUR_DOORAY_API_TOKEN"
}
```     

## 사용
* 사용자 정보 조회
```bash
$ doorayctl account 정지범

ID                  NAME EXTERNAL_EMAIL
1111111111111111111 정지범  manty@manty.co.kr
```

* 사용자에게 메시지 보내기
```bash
$ doorayctl messenger 1111111111111111111 "안녕하세요"

ID                  SUCCESS
1231321321321321321 true

```

* 캘린더 목록 조회
```bash
$ doorayctl calendar list
ID                  NAME             TYPE         CREATED_AT           OWNER
1231232132132132131 정지범1            private      2016-07-30T02:25:01Z 1111111111111111111
1231232132132132132 정지범2            private      2016-08-01T00:53:51Z 1111111111111111111
```

* 프로젝트 목록 조회
```bash
$ doorayctl project list
ID                  CODE        TYPE     SCOPE   STATE
3787724725029315943 techcenter  project  public  active
3787724725029315944 devteam     project  private active
```

* 프로젝트 목록 조회 (필터 옵션 사용)
```bash
# 타입으로 필터링
$ doorayctl project list --type project

# 범위로 필터링
$ doorayctl project list --scope public

# 상태로 필터링
$ doorayctl project list --state active

# 여러 필터 조합
$ doorayctl project list --type project --scope public --state active
```

* 프로젝트의 post 목록 조회
```bash
$ doorayctl project post list 3787724725029315943
NUMBER TASK_NUMBER  SUBJECT          WORKFLOW CLOSED CREATED_AT
1      techcenter-1 첫 번째 업무      working  false  2023-10-01T00:00:00+09:00
2      techcenter-2 두 번째 업무      closed   true   2023-10-02T00:00:00+09:00
```

* post 목록 조회 (필터 옵션 사용)
```bash
# 특정 워크플로우만 조회
$ doorayctl project post list 3787724725029315943 --workflow-classes "working"

# 특정 담당자의 post만 조회
$ doorayctl project post list 3787724725029315943 --to-members "1111111111111111111"
```

* 새 post 생성
```bash
# 기본 사용 (제목과 내용만)
$ doorayctl project post create 3787724725029315943 \
  --subject "새로운 업무" \
  --content "업무 설명입니다"

# 담당자 지정
$ doorayctl project post create 3787724725029315943 \
  -s "버그 수정" \
  -c "메인 페이지 오류 수정 필요" \
  --to "1111111111111111111,2222222222222222222"

# 전체 옵션 사용
$ doorayctl project post create 3787724725029315943 \
  --subject "긴급 업무" \
  --content "긴급 처리가 필요합니다" \
  --to "1111111111111111111" \
  --cc "2222222222222222222" \
  --priority "urgent" \
  --workflow-id "workflow123" \
  --milestone-id "milestone456"
```