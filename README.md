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
go get github.com/gcp-kit/line-bot-boilerplate-go
```

## Usage
```shell script
git clone https://github.com/gcp-kit/line-bot-boilerplate-go-example.git
cd line-bot-boilerplate-go-example
go get -u
```

`.env.yaml.tpl` を `.env.yaml` にしてyaml内の値を整える  

### Local

```shell script
go run main.go
```

### GCP(Cloud Functions)
```shell script
cd functions
gcloud builds submit --config=cloudbuild.yaml .
```

# License
[MIT license](https://en.wikipedia.org/wiki/MIT_License).
