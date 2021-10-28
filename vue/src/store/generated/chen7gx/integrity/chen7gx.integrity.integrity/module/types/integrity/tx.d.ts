import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "chen7gx.integrity.integrity";
export interface MsgCreateHash {
    creator: string;
    details: string;
    hash: string;
}
export interface MsgCreateHashResponse {
}
export declare const MsgCreateHash: {
    encode(message: MsgCreateHash, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateHash;
    fromJSON(object: any): MsgCreateHash;
    toJSON(message: MsgCreateHash): unknown;
    fromPartial(object: DeepPartial<MsgCreateHash>): MsgCreateHash;
};
export declare const MsgCreateHashResponse: {
    encode(_: MsgCreateHashResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateHashResponse;
    fromJSON(_: any): MsgCreateHashResponse;
    toJSON(_: MsgCreateHashResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateHashResponse>): MsgCreateHashResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreateHash(request: MsgCreateHash): Promise<MsgCreateHashResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateHash(request: MsgCreateHash): Promise<MsgCreateHashResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
