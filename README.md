# infographics

漢数字によるインフォグラフィックを生成するツールです.

# 実行方法

    # テストを実行する.
    go test src/*.go

    # プログラムとして実行する.
    go run $(find src/ -type f -name '*.go' -not -name '*_test.go') 1020304050

