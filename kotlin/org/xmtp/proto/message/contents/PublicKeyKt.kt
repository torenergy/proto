//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: message_contents/public_key.proto

package org.xmtp.proto.message.contents;

@kotlin.jvm.JvmName("-initializepublicKey")
public inline fun publicKey(block: org.xmtp.proto.message.contents.PublicKeyKt.Dsl.() -> kotlin.Unit): org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey =
  org.xmtp.proto.message.contents.PublicKeyKt.Dsl._create(org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey.newBuilder()).apply { block() }._build()
public object PublicKeyKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey = _builder.build()

    /**
     * <code>uint64 timestamp = 1;</code>
     */
    public var timestamp: kotlin.Long
      @JvmName("getTimestamp")
      get() = _builder.getTimestamp()
      @JvmName("setTimestamp")
      set(value) {
        _builder.setTimestamp(value)
      }
    /**
     * <code>uint64 timestamp = 1;</code>
     */
    public fun clearTimestamp() {
      _builder.clearTimestamp()
    }

    /**
     * <code>optional .xmtp.message_contents.Signature signature = 2;</code>
     */
    public var signature: org.xmtp.proto.message.contents.SignatureOuterClass.Signature
      @JvmName("getSignature")
      get() = _builder.getSignature()
      @JvmName("setSignature")
      set(value) {
        _builder.setSignature(value)
      }
    /**
     * <code>optional .xmtp.message_contents.Signature signature = 2;</code>
     */
    public fun clearSignature() {
      _builder.clearSignature()
    }
    /**
     * <code>optional .xmtp.message_contents.Signature signature = 2;</code>
     * @return Whether the signature field is set.
     */
    public fun hasSignature(): kotlin.Boolean {
      return _builder.hasSignature()
    }
    public val PublicKeyKt.Dsl.signatureOrNull: org.xmtp.proto.message.contents.SignatureOuterClass.Signature?
      get() = _builder.signatureOrNull

    /**
     * <code>.xmtp.message_contents.PublicKey.Secp256k1Uncompressed secp256k1_uncompressed = 3;</code>
     */
    public var secp256K1Uncompressed: org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey.Secp256k1Uncompressed
      @JvmName("getSecp256K1Uncompressed")
      get() = _builder.getSecp256K1Uncompressed()
      @JvmName("setSecp256K1Uncompressed")
      set(value) {
        _builder.setSecp256K1Uncompressed(value)
      }
    /**
     * <code>.xmtp.message_contents.PublicKey.Secp256k1Uncompressed secp256k1_uncompressed = 3;</code>
     */
    public fun clearSecp256K1Uncompressed() {
      _builder.clearSecp256K1Uncompressed()
    }
    /**
     * <code>.xmtp.message_contents.PublicKey.Secp256k1Uncompressed secp256k1_uncompressed = 3;</code>
     * @return Whether the secp256k1Uncompressed field is set.
     */
    public fun hasSecp256K1Uncompressed(): kotlin.Boolean {
      return _builder.hasSecp256K1Uncompressed()
    }
    public val unionCase: org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey.UnionCase
      @JvmName("getUnionCase")
      get() = _builder.getUnionCase()

    public fun clearUnion() {
      _builder.clearUnion()
    }
  }
  @kotlin.jvm.JvmName("-initializesecp256k1Uncompressed")
  public inline fun secp256k1Uncompressed(block: org.xmtp.proto.message.contents.PublicKeyKt.Secp256k1UncompressedKt.Dsl.() -> kotlin.Unit): org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey.Secp256k1Uncompressed =
    org.xmtp.proto.message.contents.PublicKeyKt.Secp256k1UncompressedKt.Dsl._create(org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey.Secp256k1Uncompressed.newBuilder()).apply { block() }._build()
  public object Secp256k1UncompressedKt {
    @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
    @com.google.protobuf.kotlin.ProtoDslMarker
    public class Dsl private constructor(
      private val _builder: org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey.Secp256k1Uncompressed.Builder
    ) {
      public companion object {
        @kotlin.jvm.JvmSynthetic
        @kotlin.PublishedApi
        internal fun _create(builder: org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey.Secp256k1Uncompressed.Builder): Dsl = Dsl(builder)
      }

      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _build(): org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey.Secp256k1Uncompressed = _builder.build()

      /**
       * <pre>
       * uncompressed point with prefix (0x04) [ P || X || Y ], 65 bytes
       * </pre>
       *
       * <code>bytes bytes = 1;</code>
       */
      public var bytes: com.google.protobuf.ByteString
        @JvmName("getBytes")
        get() = _builder.getBytes()
        @JvmName("setBytes")
        set(value) {
          _builder.setBytes(value)
        }
      /**
       * <pre>
       * uncompressed point with prefix (0x04) [ P || X || Y ], 65 bytes
       * </pre>
       *
       * <code>bytes bytes = 1;</code>
       */
      public fun clearBytes() {
        _builder.clearBytes()
      }
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey.copy(block: org.xmtp.proto.message.contents.PublicKeyKt.Dsl.() -> kotlin.Unit): org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey =
  org.xmtp.proto.message.contents.PublicKeyKt.Dsl._create(this.toBuilder()).apply { block() }._build()

@kotlin.jvm.JvmSynthetic
public inline fun org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey.Secp256k1Uncompressed.copy(block: org.xmtp.proto.message.contents.PublicKeyKt.Secp256k1UncompressedKt.Dsl.() -> kotlin.Unit): org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey.Secp256k1Uncompressed =
  org.xmtp.proto.message.contents.PublicKeyKt.Secp256k1UncompressedKt.Dsl._create(this.toBuilder()).apply { block() }._build()

public val org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKeyOrBuilder.signatureOrNull: org.xmtp.proto.message.contents.SignatureOuterClass.Signature?
  get() = if (hasSignature()) getSignature() else null

public val org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKeyOrBuilder.secp256K1UncompressedOrNull: org.xmtp.proto.message.contents.PublicKeyOuterClass.PublicKey.Secp256k1Uncompressed?
  get() = if (hasSecp256K1Uncompressed()) getSecp256K1Uncompressed() else null

