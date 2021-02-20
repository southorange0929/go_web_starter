### 项目说明

#### 项目环境
1. 项目推荐使用GO MODULE进行项目开发管理
2. 使用 `make` 进行项目构建，具体可参考Makefile文件

#### 启动环境变量说明
- ROOTDIR 程序根目录  必须    
用于读取当前执行的路径

- ENV 当前环境 非必须 默认：dev    
用于读取不同的配置文件

#### 配置文件说明
所有配置文件全部写在 `conf` 目录中    
分文件读取对应环境的配置文件内容，其配置文件规则如下：
1. 使用yaml文件格式，文件名格式为 `{ENV}_conf.yml` 。其中ENV为启动环境变量
2. 新增的配置文件需要在 `config/config.go` 中新增配置读取
3. application为必备参数，用于启动server时进行处理。
    1. port为启动端口，默认：8080。当使用HTTPS启动服务时，此配置无效，默认监听443
    2. mode为gin的启动模式，同时根据此字段会启动部分中间件。取值为：`debug` `release` `test`
    3. https是否以HTTPS启动服务。如果设置为 `true`，则使用HTTPS启动服务，对应的PEM,KEY文件放在 `ssl` 文件夹下，文件名为 `ssl.pem` `ssl.key`
    