//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: message_contents/message.proto

package org.xmtp.proto.message.contents;

@kotlin.jvm.JvmName("-initializemessage")
public inline fun message(block: org.xmtp.proto.message.contents.MessageKt.Dsl.() -> kotlin.Unit): org.xmtp.proto.message.contents.MessageOuterClass.Message =
  org.xmtp.proto.message.contents.MessageKt.Dsl._create(org.xmtp.proto.message.contents.MessageOuterClass.Message.newBuilder()).apply { block() }._build()
public object MessageKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: org.xmtp.proto.message.contents.MessageOuterClass.Message.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: org.xmtp.proto.message.contents.MessageOuterClass.Message.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): org.xmtp.proto.message.contents.MessageOuterClass.Message = _builder.build()

    /**
     * <code>.xmtp.message_contents.MessageV1 v1 = 1;</code>
     */
    public var v1: org.xmtp.proto.message.contents.MessageOuterClass.MessageV1
      @JvmName("getV1")
      get() = _builder.getV1()
      @JvmName("setV1")
      set(value) {
        _builder.setV1(value)
      }
    /**
     * <code>.xmtp.message_contents.MessageV1 v1 = 1;</code>
     */
    public fun clearV1() {
      _builder.clearV1()
    }
    /**
     * <code>.xmtp.message_contents.MessageV1 v1 = 1;</code>
     * @return Whether the v1 field is set.
     */
    public fun hasV1(): kotlin.Boolean {
      return _builder.hasV1()
    }

    /**
     * <code>.xmtp.message_contents.MessageV2 v2 = 2;</code>
     */
    public var v2: org.xmtp.proto.message.contents.MessageOuterClass.MessageV2
      @JvmName("getV2")
      get() = _builder.getV2()
      @JvmName("setV2")
      set(value) {
        _builder.setV2(value)
      }
    /**
     * <code>.xmtp.message_contents.MessageV2 v2 = 2;</code>
     */
    public fun clearV2() {
      _builder.clearV2()
    }
    /**
     * <code>.xmtp.message_contents.MessageV2 v2 = 2;</code>
     * @return Whether the v2 field is set.
     */
    public fun hasV2(): kotlin.Boolean {
      return _builder.hasV2()
    }
    public val versionCase: org.xmtp.proto.message.contents.MessageOuterClass.Message.VersionCase
      @JvmName("getVersionCase")
      get() = _builder.getVersionCase()

    public fun clearVersion() {
      _builder.clearVersion()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun org.xmtp.proto.message.contents.MessageOuterClass.Message.copy(block: org.xmtp.proto.message.contents.MessageKt.Dsl.() -> kotlin.Unit): org.xmtp.proto.message.contents.MessageOuterClass.Message =
  org.xmtp.proto.message.contents.MessageKt.Dsl._create(this.toBuilder()).apply { block() }._build()

public val org.xmtp.proto.message.contents.MessageOuterClass.MessageOrBuilder.v1OrNull: org.xmtp.proto.message.contents.MessageOuterClass.MessageV1?
  get() = if (hasV1()) getV1() else null

public val org.xmtp.proto.message.contents.MessageOuterClass.MessageOrBuilder.v2OrNull: org.xmtp.proto.message.contents.MessageOuterClass.MessageV2?
  get() = if (hasV2()) getV2() else null

