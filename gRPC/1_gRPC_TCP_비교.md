# gRPC vs TCP
회사에서 서비스를 설계하는 도중 대용량 데이터를 마이크로서비스 끼리 주고받을때 gRPC가 좋을지 TCP가 좋을지 확신 할 수가없어서 테스트를 했습니다.

우선 전체 코드는 [gRPC Upload](https://github.com/jaden7856/go-grpcUpload)에 있으며 테스트 코드는 `AsyngRPC`폴더에 있습니다.

먼저 테스트 코드의 `.proto`파일부터 보도록 하겠습니다.

```protobuf
syntax = "proto3";

package ipcgrpc;

option go_package = "github.com/jaden7856/go-grpcUpload/AsyngRPC/AsyngRPCs/protobuf";

// 내부IPC 서비스 - 내부IPC 용으로 GRPC 사용 테스트
service Ipcgrpc {
  rpc SendData (stream IpcRequest) returns (stream IpcReply) {}
}

// 요청 메시지
message IpcRequest {
  bytes bsreq = 1;
  int64 nsize = 2;
}

// 응답 메시지
message IpcReply {
  bytes bsres = 1;
}
```

위의 코드를 보는법과 내용은 [여기](0_gRPC_Intro.md)에서 다뤘으니 참고하시길 바랍니다.

클라이언트에서 보낼 Request 와 서버에서 받을 Reply 를 구성했습니다. 그리고 클라이언트에서 임의로 사이즈를 구성하기 위해 따로 파라미터를 만들었습니다.

<br>

```go
// 초기화
pstAddress := flag.String("add", "192.168.124.131:50057", "ip:port")
pnPackSize := flag.Int("size", 512, "packet size")
pnPackCount := flag.Int("count", 1000000, "packet count")
pstLogTime := flag.String("logtime", "ztime.json", "logtime name")
pnDebugMode := flag.Int("debug", 0, "debug mode - 0,1,2")

flag.Parse()

bsBufS := make([]byte, *pnPackSize)

for ix := 0; ix < *pnPackSize; ix++ {
    bsBufS[ix] = 'a'
}

req.Bsreq = bsBufS
req.Nsize = int64(*pnPackSize)

nElapsedCnt = *pnPackCount/1000 + 30
srtTimeElapeed = make([]timeElapsed, nElapsedCnt)
fpTime, _ := os.OpenFile(*pstLogTime, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
```
명령어로 인자값을 주기위해 `flag`를 사용하였고 `pnPackSize`를 통해 pack size를 설정하고 `pnPackCount`를 가지고 몇번을 보낼지 정합니다.
예를 들어 `pnPackSize`를 `1024`로 설정하고 `pnPackCount`를 `1000000`으로 한다면 총 1GB의 용량을 보내고 `pnPackSize`를 `65536`로 한다면
64GB를 보내게 되는 방식입니다.

그런다음 보낼때마다 시간 경과를 log파일에 기록하기위해 생성합니다. 

<br>

```go
// 대기 등록
wgMain.Add(2)

// 수신
go func() {
    defer wgMain.Done()

    for {
        res, err = stream.Recv()
        if errors.Is(err, io.EOF) {
            if err != nil {
                fmt.Printf("Read ERR (%v)\n\n", err)
                return
            }
            if *pnDebugMode > 0 {
                fmt.Printf("Read EOF\n\n")
            }
            break
        }
        nReadCntTotal++
    }
}()

// 송신
go func() {
    defer wgMain.Done()

    for ix = 0; ix < *pnPackCount; ix++ {
        err = stream.Send(&req)
        if err != nil {
            fmt.Printf("[F] Send ERR (%v)\n", err)
            break
			
        } else {
            nSendCntTotal++

            // 경과 시간 저장
            if nSendCntTotal == 1 ||
                (nSendCntTotal >= 10 && nSendCntTotal < 100 && nSendCntTotal%10 == 0) ||
                (nSendCntTotal >= 100 && nSendCntTotal < 1000 && nSendCntTotal%100 == 0) ||
                (nSendCntTotal%1000 == 0) {
                if nElapsedIx < nElapsedCnt {
                    srtTimeElapeed[nElapsedIx].nSendCntTotal = nSendCntTotal
                    srtTimeElapeed[nElapsedIx].elapsedTime = time.Since(startTime)
                    nElapsedIx++
                } else {
                    fmt.Printf("[W] ElapsedIx(%d) < ELASPIX\n", nElapsedIx)
                }
            }
        }
    }
}()

// 대기
wgMain.Wait()

// 종료
conn.Close()
```

<br>

```go
// 타임로그 작성
if nElapsedIx < nElapsedCnt {
    srtTimeElapeed[nElapsedIx].nSendCntTotal = nSendCntTotal + 1
    srtTimeElapeed[nElapsedIx].elapsedTime = time.Since(startTime)
    nElapsedIx++
}
for ix = 0; ix < nElapsedCnt; ix++ {
    if srtTimeElapeed[ix].nSendCntTotal > 0 {
        fmt.Fprintf(fpTime, "{ \"index\" : { \"_index\" : \"commspeed\", \"_type\" : \"record\", \"_id\" : \"%v\" } }\n{\"sync\":\"async\", \"packsize\":%d, \"packcnt\":%d, \"escount\":%d, \"estime\":%v}\n",
            time.Now().UnixNano(), *pnPackSize, *pnPackCount, srtTimeElapeed[ix].nSendCntTotal, srtTimeElapeed[ix].elapsedTime)
    }
}

// 종료
_ = fpTime.Close()
fmt.Println("End")
```