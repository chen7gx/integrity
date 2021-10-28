/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'chen7gx.integrity.integrity'

export interface MsgCreateHash {
  creator: string
  details: string
  hash: string
}

export interface MsgCreateHashResponse {}

const baseMsgCreateHash: object = { creator: '', details: '', hash: '' }

export const MsgCreateHash = {
  encode(message: MsgCreateHash, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.details !== '') {
      writer.uint32(18).string(message.details)
    }
    if (message.hash !== '') {
      writer.uint32(26).string(message.hash)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateHash {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateHash } as MsgCreateHash
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.details = reader.string()
          break
        case 3:
          message.hash = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateHash {
    const message = { ...baseMsgCreateHash } as MsgCreateHash
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.details !== undefined && object.details !== null) {
      message.details = String(object.details)
    } else {
      message.details = ''
    }
    if (object.hash !== undefined && object.hash !== null) {
      message.hash = String(object.hash)
    } else {
      message.hash = ''
    }
    return message
  },

  toJSON(message: MsgCreateHash): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.details !== undefined && (obj.details = message.details)
    message.hash !== undefined && (obj.hash = message.hash)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateHash>): MsgCreateHash {
    const message = { ...baseMsgCreateHash } as MsgCreateHash
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.details !== undefined && object.details !== null) {
      message.details = object.details
    } else {
      message.details = ''
    }
    if (object.hash !== undefined && object.hash !== null) {
      message.hash = object.hash
    } else {
      message.hash = ''
    }
    return message
  }
}

const baseMsgCreateHashResponse: object = {}

export const MsgCreateHashResponse = {
  encode(_: MsgCreateHashResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateHashResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateHashResponse } as MsgCreateHashResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgCreateHashResponse {
    const message = { ...baseMsgCreateHashResponse } as MsgCreateHashResponse
    return message
  },

  toJSON(_: MsgCreateHashResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgCreateHashResponse>): MsgCreateHashResponse {
    const message = { ...baseMsgCreateHashResponse } as MsgCreateHashResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateHash(request: MsgCreateHash): Promise<MsgCreateHashResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CreateHash(request: MsgCreateHash): Promise<MsgCreateHashResponse> {
    const data = MsgCreateHash.encode(request).finish()
    const promise = this.rpc.request('chen7gx.integrity.integrity.Msg', 'CreateHash', data)
    return promise.then((data) => MsgCreateHashResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>
