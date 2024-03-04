/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "tagioalisi";

export interface UnaryStatus {
  ok: boolean;
  message: string;
}

function createBaseUnaryStatus(): UnaryStatus {
  return { ok: false, message: "" };
}

export const UnaryStatus = {
  encode(message: UnaryStatus, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ok === true) {
      writer.uint32(8).bool(message.ok);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnaryStatus {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnaryStatus();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ok = reader.bool();
          break;
        case 2:
          message.message = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  create(base?: DeepPartial<UnaryStatus>): UnaryStatus {
    return UnaryStatus.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<UnaryStatus>): UnaryStatus {
    const message = createBaseUnaryStatus();
    message.ok = object.ok ?? false;
    message.message = object.message ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
