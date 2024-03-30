
.PHONY: build
build: # ビルドする.
	bash scripts/make.sh $@ ${ARG}

.PHONY: clean
clean: # ビルドで生成したファイルを削除する.
	bash scripts/make.sh $@ ${ARG}

.PHONY: run
run: # プログラムを実行する.
	bash scripts/make.sh $@ ${ARG}

.PHONY: stats
stats: # ソースコードに関する情報 (TODO コメントやコード行数など) を表示する.
	bash scripts/make.sh $@ ${ARG}

.PHONY: test
test: # テストを実行する.
	bash scripts/make.sh $@ ${ARG}

