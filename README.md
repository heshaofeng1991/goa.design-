### goa 说明。

#### install
```bash
go get -u -v github.com/NextSmartShip/goa
go install github.com/NextSmartShip/goa@latest
```

#### design有更新的时候
```bash
make all
```

#### notes
1. 当design变化的时候，注意看一下makefile，大家可以自行去更新就好。
2. 如何使用使用github.com私有仓库作为项目的依赖包,可以通过设置SSH公钥来解决私有项目授权问题，这个大家自行研究一下。