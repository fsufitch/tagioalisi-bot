// package: tagioalisi
// file: sockpuppet.proto

import * as jspb from "google-protobuf";
import * as base_pb from "./base_pb";

export class SendMessageRequest extends jspb.Message {
  getJwt(): string;
  setJwt(value: string): void;

  getChannelid(): string;
  setChannelid(value: string): void;

  getContent(): string;
  setContent(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SendMessageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SendMessageRequest): SendMessageRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SendMessageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SendMessageRequest;
  static deserializeBinaryFromReader(message: SendMessageRequest, reader: jspb.BinaryReader): SendMessageRequest;
}

export namespace SendMessageRequest {
  export type AsObject = {
    jwt: string,
    channelid: string,
    content: string,
  }
}

export class SendMessageReply extends jspb.Message {
  hasStatus(): boolean;
  clearStatus(): void;
  getStatus(): base_pb.UnaryStatus | undefined;
  setStatus(value?: base_pb.UnaryStatus): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SendMessageReply.AsObject;
  static toObject(includeInstance: boolean, msg: SendMessageReply): SendMessageReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SendMessageReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SendMessageReply;
  static deserializeBinaryFromReader(message: SendMessageReply, reader: jspb.BinaryReader): SendMessageReply;
}

export namespace SendMessageReply {
  export type AsObject = {
    status?: base_pb.UnaryStatus.AsObject,
  }
}

