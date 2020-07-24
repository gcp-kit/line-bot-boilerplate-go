# Line Bot Boilerplate

## About
line-bot-boilerplate-goはローカル開発時にはginでWebHookを受けて動作し  
クラウド開発によるデプロイ待ちなどを気にせずに開発効率を維持できるようにしたもの

## Features

boilerplateは多くあるが、ローカル開発とクラウド運用が共存しているもの作成した

## Requirement

* go1.11以上
* 外部依存モジュールはgo.modから取得  

## Installation

```shell script
git clone https://github.com/gcp-kit/line-bot-boilerplate-go.git
cd line-bot-boilerplate-go
go get -u
go mod tidy
```

## Usage

`.env.yaml.tpl`にある値を整え`.tpl`を外す  

### Local

```shell script
cd examples
go run main.go
```

### GCP(Cloud Functions)
```shell script
gcloud builds submit --config=cloudbuild.yaml .
```

# License
[MIT license](https://en.wikipedia.org/wiki/MIT_License).
