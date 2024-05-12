# gomock-aws-example
 Mocked AWS clients example for testing using `gomock`

## Install the `mockgen` command

```shell
go install go.uber.org/mock/mockgen@latest
```

### Generate the code that mocks the types and interfaces we need

```shell
mockgen -source=main_test.go -destination=mocks/main_test_mock.go
```