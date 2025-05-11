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
    * ~/.doorayctl/config
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