# tass-go-sdk

[![Release](https://img.shields.io/github/v/release/woodleighschool/tass-go-sdk?display_name=tag&sort=semver)](https://github.com/woodleighschool/tass-go-sdk/releases/latest)
[![CI](https://github.com/woodleighschool/tass-go-sdk/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/woodleighschool/tass-go-sdk/actions/workflows/ci.yaml)
[![Go](https://img.shields.io/github/go-mod/go-version/woodleighschool/tass-go=sdk?logo=go)](https://github.com/woodleighschool/tass-go-sdk/blob/main/go.mod)
[![Container](https://img.shields.io/badge/container-ghcr.io-2496ED?logo=github&logoColor=white)](https://github.com/orgs/woodleighschool/packages/container/package/tass-go-sdk)
[![License](https://img.shields.io/github/license/woodleighschool/tass-go-sdk)](https://github.com/woodleighschool/tass-go-sdk/blob/main/LICENSE)

> TASS Go SDK

A SDK for Golang to interact with the RESTful TASS API
The SDK conforms mostly to TASS' OpenAPI specification with the majority of structs and functions matching their name in the specification
Some additions or modifications have been made where appropriate

It is planned to create a generator based on the OpenAPI specification to automate the code creation but currently it is hand-written

## 🔨Progress

Progress on development/testing of available operations can be viewed [here](https://github.com/woodleighschool/tass-go-sdk/blob/main/PROGRESS.md)

## 🚀 Usage

Initiate a client with `tass.NewClient(config)`

```golang
client, err := tass.NewClient(tass.Config{
	URL: "{your_tass_url_here}",
	CompanyCode: "{your_company_code}",
	ClientKey: "{client_key}",
	ClientSecret: "{client_secret}",
	Modules: tass.ModulesConfig{
		Student: true,
		Finance: false,
		Employee: true,
	},
})

if err != nil {

}

```

Then perform an operation with the sub-modules

```golang
student, err := client.Student.GetStudent(context.Background(), "001234")
fmt.Printf("Student name: %s", student.PreferredName)
```


## 🧑‍💻 Development

Mise owns the toolchain and repository checks:

```bash
mise run build
#mise run generate
mise run test
mise run lint
mise run fmt-check
mise run workflow-lint
```

## 📄 License

Licensed under the [Apache License 2.0](LICENSE).
