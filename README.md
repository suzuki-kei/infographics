# infographics

テキストによるインフォグラフィックを生成するツールです.

    $ infographics 1020304050
    1020304050 => 十億 千万 千万 十万 十万 十万 千 千 千 千 十 十 十 十 十

    $ infographics -s 1020304050
    1020304050 => 10億 2000万 30万 4000 50

    $ infographics --si 1002003004005
    1002003004005 => T G G M M M k k k k 1 1 1 1 1

    $ infographics --si -s 1002003004005
    1002003004005 => 1T 2G 3M 4k 5

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

