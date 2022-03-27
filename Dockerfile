# ベースとなるDockerイメージ指定
FROM golang:latest

# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/work/sudoku-make-problem

# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/work/sudoku-make-problem

# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/work/sudoku-make-problem