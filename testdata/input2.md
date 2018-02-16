# 1を1億回足して1億にならない場合

「1を1億回足す」ってのを [Go 言語]でも書いてみる。

```go
package main

import "fmt"

func main() {
    var d float32 = 0.0
    for i := 0; i < 100000000; i++ {
        d += 1.0
    }
    fmt.Println(d)
}
```

実行結果は予想通り

```text
$ go run loop1.go
1.6777216e+07
```

となる[^f32]。

[^f32]: `float32` は32ビットサイズの浮動小数点数型で，符号部1ビット，指数部8ビット，仮数部23ビット，という内訳になっている（仮数部は仮数の小数点以下を表す）。つまり有効桁数が24ビット（10進数で約7桁）しかない。したがって今回のような「1づつ加算する動作を繰り返す」処理では16,777,216（`=0xffffff+1`）以降は「情報落ち」が発生する。ちなみに `float64` は64ビットサイズで仮数部は52ビットあり，10進数にして約15桁の有効桁数になる。