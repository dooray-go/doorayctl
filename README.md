# doorayctl
* dooray open api 를 사용하는 cli tool 입니다.

## 설치

### Homebrew (macOS / Linux)
```bash
brew tap dooray-go/tap
brew install doorayctl
```

업그레이드:
```bash
brew upgrade doorayctl
```

### 직접 다운로드
[GitHub Releases](https://github.com/dooray-go/doorayctl/releases)에서 플랫폼에 맞는 바이너리를 다운로드합니다.

```bash
# macOS (Apple Silicon)
sudo mv doorayctl.darwin.arm64 /usr/local/bin/doorayctl

# macOS (Intel)
sudo mv doorayctl.darwin.amd64 /usr/local/bin/doorayctl

# Linux (amd64)
sudo mv doorayctl.linux.amd64 /usr/local/bin/doorayctl
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

## 쉘 자동완성 (Completion)

doorayctl은 쉘 자동완성을 지원합니다. 사용하는 쉘에 맞게 설정하세요.

### Bash
```bash
# 현재 세션에만 적용
source <(doorayctl completion bash)

# 영구 적용 (Linux)
doorayctl completion bash > /etc/bash_completion.d/doorayctl

# 영구 적용 (macOS)
doorayctl completion bash > $(brew --prefix)/etc/bash_completion.d/doorayctl
```

### Zsh
```bash
# 현재 세션에만 적용
source <(doorayctl completion zsh)

# 영구 적용 (Linux)
doorayctl completion zsh > "${fpath[1]}/_doorayctl"

# 영구 적용 (macOS)
doorayctl completion zsh > $(brew --prefix)/share/zsh/site-functions/_doorayctl
```

> zsh에서 자동완성이 처음이라면 먼저 다음을 실행하세요:
> ```bash
> echo "autoload -U compinit; compinit" >> ~/.zshrc
> ```

### Fish
```bash
# 현재 세션에만 적용
doorayctl completion fish | source

# 영구 적용
doorayctl completion fish > ~/.config/fish/completions/doorayctl.fish
```

### PowerShell
```powershell
# 현재 세션에만 적용
doorayctl completion powershell | Out-String | Invoke-Expression

# 영구 적용 (PowerShell 프로필에 위 명령어를 추가)
```

> 설정 후 새 쉘을 열어야 적용됩니다.

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
# 특정 워크플로우만 조회 (backlog, registered, working, closed)
$ doorayctl project post list 3787724725029315943 --workflow-classes "working"

# 여러 워크플로우 조회
$ doorayctl project post list 3787724725029315943 --workflow-classes "registered,working"

# 특정 담당자의 post만 조회
$ doorayctl project post list 3787724725029315943 --to-members "1111111111111111111"

# 특정 작성자의 post만 조회
$ doorayctl project post list 3787724725029315943 --from-members "1111111111111111111"

# 작성자 이메일로 필터링
$ doorayctl project post list 3787724725029315943 --from-email "manty@manty.co.kr"

# 제목으로 검색
$ doorayctl project post list 3787724725029315943 --subjects "버그"

# 특정 태그가 붙은 post 조회
$ doorayctl project post list 3787724725029315943 --tag-ids "tag123,tag456"

# 특정 마일스톤의 post 조회
$ doorayctl project post list 3787724725029315943 --milestone-ids "milestone123"

# 하위 업무 조회
$ doorayctl project post list 3787724725029315943 --parent-post-id "9876543210"

# 특정 post 번호로 조회
$ doorayctl project post list 3787724725029315943 --post-number "42"

# 워크플로우 ID로 필터링
$ doorayctl project post list 3787724725029315943 --workflow-ids "wf111,wf222"

# 참조자로 필터링
$ doorayctl project post list 3787724725029315943 --cc-members "1111111111111111111"
```

* post 목록 조회 (날짜 필터)
```bash
# 오늘 생성된 post
$ doorayctl project post list 3787724725029315943 --created-at "today"

# 이번 주 생성된 post
$ doorayctl project post list 3787724725029315943 --created-at "thisweek"

# 최근 7일간 업데이트된 post
$ doorayctl project post list 3787724725029315943 --updated-at "prev-7d"

# 향후 3일 내 마감인 post
$ doorayctl project post list 3787724725029315943 --due-at "next-3d"

# ISO8601 범위로 조회
$ doorayctl project post list 3787724725029315943 --created-at "2024-01-01T00:00:00Z/2024-12-31T23:59:59Z"
```

* post 목록 조회 (페이징 & 정렬)
```bash
# 페이지 크기 지정
$ doorayctl project post list 3787724725029315943 --size 50

# 2번째 페이지 조회
$ doorayctl project post list 3787724725029315943 --page 1 --size 20

# 마감일 기준 정렬
$ doorayctl project post list 3787724725029315943 --order "postDueAt"

# 최신 생성순 정렬
$ doorayctl project post list 3787724725029315943 --order "-createdAt"

# 업데이트 순 정렬
$ doorayctl project post list 3787724725029315943 --order "postUpdatedAt"
```

* post 목록 조회 (필터 조합)
```bash
# 내가 담당하는 진행 중인 업무를 최신순으로
$ doorayctl project post list 3787724725029315943 \
  --to-members "1111111111111111111" \
  --workflow-classes "working" \
  --order "-createdAt"

# 이번 주 마감인 긴급 업무
$ doorayctl project post list 3787724725029315943 \
  --due-at "thisweek" \
  --workflow-classes "registered,working"

# 특정 마일스톤의 완료되지 않은 업무
$ doorayctl project post list 3787724725029315943 \
  --milestone-ids "milestone123" \
  --workflow-classes "backlog,registered,working"
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