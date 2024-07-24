// Original file: ../shared/roof.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { PingMessage as _roof_PingMessage, PingMessage__Output as _roof_PingMessage__Output } from '../roof/PingMessage';

export interface RoofServiceClient extends grpc.Client {
  Ping(argument: _roof_PingMessage, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  Ping(argument: _roof_PingMessage, metadata: grpc.Metadata, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  Ping(argument: _roof_PingMessage, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  Ping(argument: _roof_PingMessage, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  ping(argument: _roof_PingMessage, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  ping(argument: _roof_PingMessage, metadata: grpc.Metadata, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  ping(argument: _roof_PingMessage, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  ping(argument: _roof_PingMessage, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  
}

export interface RoofServiceHandlers extends grpc.UntypedServiceImplementation {
  Ping: grpc.handleUnaryCall<_roof_PingMessage__Output, _roof_PingMessage>;
  
}

export interface RoofServiceDefinition extends grpc.ServiceDefinition {
  Ping: MethodDefinition<_roof_PingMessage, _roof_PingMessage, _roof_PingMessage__Output, _roof_PingMessage__Output>
}
