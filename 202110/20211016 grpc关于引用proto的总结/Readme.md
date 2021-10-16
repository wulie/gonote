##protobuf import 文件编译时 提示was not found or had errors问题总结

最近业务中在对接grpc 接口，其中grpc 采用了protobuf 这种结构化的数据存储格式，
可用于结构化数据的序列化，目前官方中已经支持多种语言。在一个proto 文件 import 另一个proto 文件，编译时报错 (Import "zeus/file.proto" was not found or had errors) 问题,接下来还原一下产生的这个问题;

##定义proto文件是都需要定义

    //协议版本
    syntax = "proto3";
    
    //需要引用时
    import "enums/enums.proto";

    //定义包路径，暴露第三方引用路径
    option go_package = "gitee.com/magusiiot/api/grpc_app";

    //定义包名称
    package grpc_app;


##为了不出现was not found or had errors的办法

    1.在在对文件编译的时候， 不同的协议的文件用各自的目录包含，
    2.定义的proto文件的引用也是用目录，如  import "enums/enums.proto";
    3.使用官方命令 并且添加参数  --proto_path=.  例子如下
    protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --proto_path=. helloworld/helloworld.proto


出处：https://www.jianshu.com/p/3e44464848d2
