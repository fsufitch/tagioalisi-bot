/* eslint-disable */
import type { CallContext, CallOptions } from "nice-grpc-common";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "tagioalisi";

/** The request message containing the user's name. */
export interface HelloRequest {
  name: string;
}

/** The response message containing the greetings */
export interface HelloReply {
  message: string;
}

function createBaseHelloRequest(): HelloRequest {
  return { name: "" };
}

export const HelloRequest = {
  encode(message: HelloRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HelloRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHelloRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  create(base?: DeepPartial<HelloRequest>): HelloRequest {
    return HelloRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<HelloRequest>): HelloRequest {
    const message = createBaseHelloRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseHelloReply(): HelloReply {
  return { message: "" };
}

export const HelloReply = {
  encode(message: HelloReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HelloReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHelloReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.message = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  create(base?: DeepPartial<HelloReply>): HelloReply {
    return HelloReply.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<HelloReply>): HelloReply {
    const message = createBaseHelloReply();
    message.message = object.message ?? "";
    return message;
  },
};

export type GreeterDefinition = typeof GreeterDefinition;
export const GreeterDefinition = {
  name: "Greeter",
  fullName: "tagioalisi.Greeter",
  methods: {
    /** Sends a greeting */
    sayHello: {
      name: "SayHello",
      requestType: HelloRequest,
      requestStream: false,
      responseType: HelloReply,
      responseStream: false,
      options: {},
    },
  },
} as const;

export interface GreeterServiceImplementation<CallContextExt = {}> {
  /** Sends a greeting */
  sayHello(request: HelloRequest, context: CallContext & CallContextExt): Promise<DeepPartial<HelloReply>>;
}

export interface GreeterClient<CallOptionsExt = {}> {
  /** Sends a greeting */
  sayHello(request: DeepPartial<HelloRequest>, options?: CallOptions & CallOptionsExt): Promise<HelloReply>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
