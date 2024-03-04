/* eslint-disable */
import type { CallContext, CallOptions } from "nice-grpc-common";
import * as _m0 from "protobufjs/minimal";
import { UnaryStatus } from "./base";

export const protobufPackage = "tagioalisi";

export interface SendMessageRequest {
  jwt: string;
  channelID: string;
  content: string;
}

export interface SendMessageReply {
  status: UnaryStatus | undefined;
}

function createBaseSendMessageRequest(): SendMessageRequest {
  return { jwt: "", channelID: "", content: "" };
}

export const SendMessageRequest = {
  encode(message: SendMessageRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.jwt !== "") {
      writer.uint32(10).string(message.jwt);
    }
    if (message.channelID !== "") {
      writer.uint32(18).string(message.channelID);
    }
    if (message.content !== "") {
      writer.uint32(26).string(message.content);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SendMessageRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSendMessageRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.jwt = reader.string();
          break;
        case 2:
          message.channelID = reader.string();
          break;
        case 3:
          message.content = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  create(base?: DeepPartial<SendMessageRequest>): SendMessageRequest {
    return SendMessageRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<SendMessageRequest>): SendMessageRequest {
    const message = createBaseSendMessageRequest();
    message.jwt = object.jwt ?? "";
    message.channelID = object.channelID ?? "";
    message.content = object.content ?? "";
    return message;
  },
};

function createBaseSendMessageReply(): SendMessageReply {
  return { status: undefined };
}

export const SendMessageReply = {
  encode(message: SendMessageReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.status !== undefined) {
      UnaryStatus.encode(message.status, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SendMessageReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSendMessageReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = UnaryStatus.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  create(base?: DeepPartial<SendMessageReply>): SendMessageReply {
    return SendMessageReply.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<SendMessageReply>): SendMessageReply {
    const message = createBaseSendMessageReply();
    message.status = (object.status !== undefined && object.status !== null)
      ? UnaryStatus.fromPartial(object.status)
      : undefined;
    return message;
  },
};

export type SockpuppetDefinition = typeof SockpuppetDefinition;
export const SockpuppetDefinition = {
  name: "Sockpuppet",
  fullName: "tagioalisi.Sockpuppet",
  methods: {
    sendMessage: {
      name: "SendMessage",
      requestType: SendMessageRequest,
      requestStream: false,
      responseType: SendMessageReply,
      responseStream: false,
      options: {},
    },
  },
} as const;

export interface SockpuppetServiceImplementation<CallContextExt = {}> {
  sendMessage(
    request: SendMessageRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<SendMessageReply>>;
}

export interface SockpuppetClient<CallOptionsExt = {}> {
  sendMessage(
    request: DeepPartial<SendMessageRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<SendMessageReply>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
