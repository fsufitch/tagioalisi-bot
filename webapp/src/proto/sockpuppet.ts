/* eslint-disable */
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

  fromJSON(object: any): SendMessageRequest {
    return {
      jwt: isSet(object.jwt) ? String(object.jwt) : "",
      channelID: isSet(object.channelID) ? String(object.channelID) : "",
      content: isSet(object.content) ? String(object.content) : "",
    };
  },

  toJSON(message: SendMessageRequest): unknown {
    const obj: any = {};
    message.jwt !== undefined && (obj.jwt = message.jwt);
    message.channelID !== undefined && (obj.channelID = message.channelID);
    message.content !== undefined && (obj.content = message.content);
    return obj;
  },

  create<I extends Exact<DeepPartial<SendMessageRequest>, I>>(base?: I): SendMessageRequest {
    return SendMessageRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SendMessageRequest>, I>>(object: I): SendMessageRequest {
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

  fromJSON(object: any): SendMessageReply {
    return { status: isSet(object.status) ? UnaryStatus.fromJSON(object.status) : undefined };
  },

  toJSON(message: SendMessageReply): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status ? UnaryStatus.toJSON(message.status) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<SendMessageReply>, I>>(base?: I): SendMessageReply {
    return SendMessageReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SendMessageReply>, I>>(object: I): SendMessageReply {
    const message = createBaseSendMessageReply();
    message.status = (object.status !== undefined && object.status !== null)
      ? UnaryStatus.fromPartial(object.status)
      : undefined;
    return message;
  },
};

export interface Sockpuppet {
  SendMessage(request: SendMessageRequest): Promise<SendMessageReply>;
}

export class SockpuppetClientImpl implements Sockpuppet {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "tagioalisi.Sockpuppet";
    this.rpc = rpc;
    this.SendMessage = this.SendMessage.bind(this);
  }
  SendMessage(request: SendMessageRequest): Promise<SendMessageReply> {
    const data = SendMessageRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "SendMessage", data);
    return promise.then((data) => SendMessageReply.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
