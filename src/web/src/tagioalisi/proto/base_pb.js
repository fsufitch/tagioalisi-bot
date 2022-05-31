import * as jspb from "google-protobuf";
var goog = jspb;
var global = Function('return this')();
goog.exportSymbol('proto.tagioalisi.UnaryStatus', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.tagioalisi.UnaryStatus = function (opt_data) {
    jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.tagioalisi.UnaryStatus, jspb.Message);
if (goog.DEBUG && !COMPILED) {
    /**
     * @public
     * @override
     */
    proto.tagioalisi.UnaryStatus.displayName = 'proto.tagioalisi.UnaryStatus';
}
if (jspb.Message.GENERATE_TO_OBJECT) {
    /**
     * Creates an object representation of this proto.
     * Field names that are reserved in JavaScript and will be renamed to pb_name.
     * Optional fields that are not set will be set to undefined.
     * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
     * For the list of reserved names please see:
     *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
     * @param {boolean=} opt_includeInstance Deprecated. whether to include the
     *     JSPB instance for transitional soy proto support:
     *     http://goto/soy-param-migration
     * @return {!Object}
     */
    proto.tagioalisi.UnaryStatus.prototype.toObject = function (opt_includeInstance) {
        return proto.tagioalisi.UnaryStatus.toObject(opt_includeInstance, this);
    };
    /**
     * Static version of the {@see toObject} method.
     * @param {boolean|undefined} includeInstance Deprecated. Whether to include
     *     the JSPB instance for transitional soy proto support:
     *     http://goto/soy-param-migration
     * @param {!proto.tagioalisi.UnaryStatus} msg The msg instance to transform.
     * @return {!Object}
     * @suppress {unusedLocalVariables} f is only used for nested messages
     */
    proto.tagioalisi.UnaryStatus.toObject = function (includeInstance, msg) {
        var f, obj = {
            ok: jspb.Message.getBooleanFieldWithDefault(msg, 1, false),
            message: jspb.Message.getFieldWithDefault(msg, 2, "")
        };
        if (includeInstance) {
            obj.$jspbMessageInstance = msg;
        }
        return obj;
    };
}
/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.tagioalisi.UnaryStatus}
 */
proto.tagioalisi.UnaryStatus.deserializeBinary = function (bytes) {
    var reader = new jspb.BinaryReader(bytes);
    var msg = new proto.tagioalisi.UnaryStatus;
    return proto.tagioalisi.UnaryStatus.deserializeBinaryFromReader(msg, reader);
};
/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tagioalisi.UnaryStatus} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tagioalisi.UnaryStatus}
 */
proto.tagioalisi.UnaryStatus.deserializeBinaryFromReader = function (msg, reader) {
    while (reader.nextField()) {
        if (reader.isEndGroup()) {
            break;
        }
        var field = reader.getFieldNumber();
        switch (field) {
            case 1:
                var value = /** @type {boolean} */ (reader.readBool());
                msg.setOk(value);
                break;
            case 2:
                var value = /** @type {string} */ (reader.readString());
                msg.setMessage(value);
                break;
            default:
                reader.skipField();
                break;
        }
    }
    return msg;
};
/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.tagioalisi.UnaryStatus.prototype.serializeBinary = function () {
    var writer = new jspb.BinaryWriter();
    proto.tagioalisi.UnaryStatus.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
};
/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tagioalisi.UnaryStatus} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tagioalisi.UnaryStatus.serializeBinaryToWriter = function (message, writer) {
    var f = undefined;
    f = message.getOk();
    if (f) {
        writer.writeBool(1, f);
    }
    f = message.getMessage();
    if (f.length > 0) {
        writer.writeString(2, f);
    }
};
/**
 * optional bool ok = 1;
 * @return {boolean}
 */
proto.tagioalisi.UnaryStatus.prototype.getOk = function () {
    return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 1, false));
};
/**
 * @param {boolean} value
 * @return {!proto.tagioalisi.UnaryStatus} returns this
 */
proto.tagioalisi.UnaryStatus.prototype.setOk = function (value) {
    return jspb.Message.setProto3BooleanField(this, 1, value);
};
/**
 * optional string message = 2;
 * @return {string}
 */
proto.tagioalisi.UnaryStatus.prototype.getMessage = function () {
    return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};
/**
 * @param {string} value
 * @return {!proto.tagioalisi.UnaryStatus} returns this
 */
proto.tagioalisi.UnaryStatus.prototype.setMessage = function (value) {
    return jspb.Message.setProto3StringField(this, 2, value);
};
// REPLACED BY GOOG_TO_CJS::  goog.object.extend(exports, proto.tagioalisi);
//=== CommonJS exports generated by goog_to_cjs
export const UnaryStatus = proto.tagioalisi.UnaryStatus;
