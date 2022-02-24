module github.gom/hashicorp/hello-vault-go/quick-start

go 1.16

replace github.com/hashicorp/vault/sdk => github.com/averche/vault/sdk v0.3.1-0.20220224225511-85221e766af8

replace github.com/hashicorp/vault/api => github.com/averche/vault/api v1.3.2-0.20220224225511-85221e766af8

require github.com/hashicorp/vault/api v1.3.1
