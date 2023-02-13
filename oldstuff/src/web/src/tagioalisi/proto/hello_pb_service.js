import * as hello_pb from "./hello_pb.js";
import grpcWeb from "@improbable-eng/grpc-web";
var grpc = grpcWeb.grpc;
var Greeter = (function () {
    function Greeter() { }
    Greeter.serviceName = "tagioalisi.Greeter";
    return Greeter;
}());
Greeter.SayHello = {
    methodName: "SayHello",
    service: Greeter,
    requestStream: false,
    responseStream: false,
    requestType: hello_pb.HelloRequest,
    responseType: hello_pb.HelloReply
};
function GreeterClient(serviceHost, options) {
    this.serviceHost = serviceHost;
    this.options = options || {};
}
GreeterClient.prototype.sayHello = function sayHello(requestMessage, metadata, callback) {
    if (arguments.length === 2) {
        callback = arguments[1];
    }
    var client = grpc.unary(Greeter.SayHello, {
        request: requestMessage,
        host: this.serviceHost,
        metadata: metadata,
        transport: this.options.transport,
        debug: this.options.debug,
        onEnd: function (response) {
            if (callback) {
                if (response.status !== grpc.Code.OK) {
                    var err = new Error(response.statusMessage);
                    err.code = response.status;
                    err.metadata = response.trailers;
                    callback(err, null);
                }
                else {
                    callback(null, response.message);
                }
            }
        }
    });
    return {
        cancel: function () {
            callback = null;
            client.close();
        }
    };
};
export { Greeter };
export { GreeterClient };
