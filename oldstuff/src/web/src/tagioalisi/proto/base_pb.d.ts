// package: tagioalisi
// file: base.proto

import * as jspb from "google-protobuf";

export class UnaryStatus extends jspb.Message {
  getOk(): boolean;
  setOk(value: boolean): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnaryStatus.AsObject;
  static toObject(includeInstance: boolean, msg: UnaryStatus): UnaryStatus.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UnaryStatus, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnaryStatus;
  static deserializeBinaryFromReader(message: UnaryStatus, reader: jspb.BinaryReader): UnaryStatus;
}

export namespace UnaryStatus {
  export type AsObject = {
    ok: boolean,
    message: string,
  }
}

