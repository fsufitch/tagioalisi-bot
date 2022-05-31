import * as sockpuppet_pb from "./sockpuppet_pb.js";
import grpcWeb from "@improbable-eng/grpc-web";
var grpc = grpcWeb.grpc;
var Sockpuppet = (function () {
    function Sockpuppet() { }
    Sockpuppet.serviceName = "tagioalisi.Sockpuppet";
    return Sockpuppet;
}());
Sockpuppet.SendMessage = {
    methodName: "SendMessage",
    service: Sockpuppet,
    requestStream: false,
    responseStream: false,
    requestType: sockpuppet_pb.SendMessageRequest,
    responseType: sockpuppet_pb.SendMessageReply
};
function SockpuppetClient(serviceHost, options) {
    this.serviceHost = serviceHost;
    this.options = options || {};
}
SockpuppetClient.prototype.sendMessage = function sendMessage(requestMessage, metadata, callback) {
    if (arguments.length === 2) {
        callback = arguments[1];
    }
    var client = grpc.unary(Sockpuppet.SendMessage, {
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
export { Sockpuppet };
export { SockpuppetClient };
