// hello1.thrift
namespace go hello1

struct HelloReq {
    1: string Name (api.query="name");
}

struct HelloResp {
    1: string RespBody;
}

service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello1");
}(
     api.base_domain="http://127.0.0.1:8888";
)
