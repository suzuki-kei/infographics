# infographics

漢数字によるインフォグラフィックを生成するツールです.

# 各種手順

    # プログラムを実行する.
    make run ARG='1020304050'
    make run ARG='-s 1020304050'

    # テストを実行する.
    make test

    # ビルドする.
    make build

    # ビルドで生成した実行ファイルを起動する.
    ./target/infographics 1020304050
    ./target/infographics -s 1020304050

    # ビルドで生成したファイルを削除する.
    make clean

    # その他のルールを確認する.
    make help

