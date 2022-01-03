# 티베로 설치 - Linux

### 1. 라이센스 신청

- 라이센스 신청하는 방법은 [게시물](https://github.com/jaden7856/Tibero/blob/main/1_Tibero-License.md)을 참조합니다.





### 2. 설치파일 다운

1) [다운로드](https://technet.tmaxsoft.com/ko/front/download/viewDownload.do?cmProductCode=0301&version_seq=PVER-20150504-000001&doc_type_cd=DN#binary) 클릭후 **Linux (x86) 64-bit** 를 다운받았다

2. Linux server에 `license.xml`파일과 `tibero6....tar.gz`파일을 넣어주고 압축을 해제

   - ```
     # tar -zxvf [파일이름]
     ```

3. `tibero6`폴더안에 license폴더에 `license.xml`파일을 넣어주겠습니다.





### 3. 환경변수 셋팅

- `~/.bash_profile`을 열어 아래와 같이 입력해주겠습니다.

```shell
export TB_HOME=[티베로 설치 위치]		 # ex) /home/centos/tibero6
export TB_SID=[사용하고자하는 DB명]		# ex) tibero
export LD_LIBRARY_PATH=$TB_HOME/lib:$TB_HOME/client/lib
export PATH=$PATH:$TB_HOME/bin:$TB_HOME/client/bin
```





### 4. 티베로 환경파일 생성

```shell
$ sh tibero6/config/gen_tip.sh
```

