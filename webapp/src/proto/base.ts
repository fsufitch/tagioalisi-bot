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

  fromJSON(object: any): UnaryStatus {
    return {
      ok: isSet(object.ok) ? Boolean(object.ok) : false,
      message: isSet(object.message) ? String(object.message) : "",
    };
  },

  toJSON(message: UnaryStatus): unknown {
    const obj: any = {};
    message.ok !== undefined && (obj.ok = message.ok);
    message.message !== undefined && (obj.message = message.message);
    return obj;
  },

  create<I extends Exact<DeepPartial<UnaryStatus>, I>>(base?: I): UnaryStatus {
    return UnaryStatus.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UnaryStatus>, I>>(object: I): UnaryStatus {
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

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
