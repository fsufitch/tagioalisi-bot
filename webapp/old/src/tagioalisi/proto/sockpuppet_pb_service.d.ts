// package: tagioalisi
// file: sockpuppet.proto

import * as sockpuppet_pb from "./sockpuppet_pb";
import {grpc} from "@improbable-eng/grpc-web";

type SockpuppetSendMessage = {
  readonly methodName: string;
  readonly service: typeof Sockpuppet;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof sockpuppet_pb.SendMessageRequest;
  readonly responseType: typeof sockpuppet_pb.SendMessageReply;
};

export class Sockpuppet {
  static readonly serviceName: string;
  static readonly SendMessage: SockpuppetSendMessage;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class SockpuppetClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  sendMessage(
    requestMessage: sockpuppet_pb.SendMessageRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: sockpuppet_pb.SendMessageReply|null) => void
  ): UnaryResponse;
  sendMessage(
    requestMessage: sockpuppet_pb.SendMessageRequest,
    callback: (error: ServiceError|null, responseMessage: sockpuppet_pb.SendMessageReply|null) => void
  ): UnaryResponse;
}

