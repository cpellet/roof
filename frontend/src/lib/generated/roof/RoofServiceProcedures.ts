// Original file: ../shared/roof.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { PerformAnalysisRequest as _roof_PerformAnalysisRequest, PerformAnalysisRequest__Output as _roof_PerformAnalysisRequest__Output } from '../roof/PerformAnalysisRequest';
import type { PerformAnalysisResponse as _roof_PerformAnalysisResponse, PerformAnalysisResponse__Output as _roof_PerformAnalysisResponse__Output } from '../roof/PerformAnalysisResponse';
import type { PingMessage as _roof_PingMessage, PingMessage__Output as _roof_PingMessage__Output } from '../roof/PingMessage';
import type { RetrieveAnalysisRequest as _roof_RetrieveAnalysisRequest, RetrieveAnalysisRequest__Output as _roof_RetrieveAnalysisRequest__Output } from '../roof/RetrieveAnalysisRequest';
import type { RetrieveAnalysisResponse as _roof_RetrieveAnalysisResponse, RetrieveAnalysisResponse__Output as _roof_RetrieveAnalysisResponse__Output } from '../roof/RetrieveAnalysisResponse';

export interface RoofServiceProceduresClient extends grpc.Client {
  PerformAnalysis(argument: _roof_PerformAnalysisRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_PerformAnalysisResponse__Output>): grpc.ClientUnaryCall;
  PerformAnalysis(argument: _roof_PerformAnalysisRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_roof_PerformAnalysisResponse__Output>): grpc.ClientUnaryCall;
  PerformAnalysis(argument: _roof_PerformAnalysisRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_PerformAnalysisResponse__Output>): grpc.ClientUnaryCall;
  PerformAnalysis(argument: _roof_PerformAnalysisRequest, callback: grpc.requestCallback<_roof_PerformAnalysisResponse__Output>): grpc.ClientUnaryCall;
  performAnalysis(argument: _roof_PerformAnalysisRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_PerformAnalysisResponse__Output>): grpc.ClientUnaryCall;
  performAnalysis(argument: _roof_PerformAnalysisRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_roof_PerformAnalysisResponse__Output>): grpc.ClientUnaryCall;
  performAnalysis(argument: _roof_PerformAnalysisRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_PerformAnalysisResponse__Output>): grpc.ClientUnaryCall;
  performAnalysis(argument: _roof_PerformAnalysisRequest, callback: grpc.requestCallback<_roof_PerformAnalysisResponse__Output>): grpc.ClientUnaryCall;
  
  Ping(argument: _roof_PingMessage, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  Ping(argument: _roof_PingMessage, metadata: grpc.Metadata, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  Ping(argument: _roof_PingMessage, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  Ping(argument: _roof_PingMessage, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  ping(argument: _roof_PingMessage, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  ping(argument: _roof_PingMessage, metadata: grpc.Metadata, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  ping(argument: _roof_PingMessage, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  ping(argument: _roof_PingMessage, callback: grpc.requestCallback<_roof_PingMessage__Output>): grpc.ClientUnaryCall;
  
  RetrieveAnalysis(argument: _roof_RetrieveAnalysisRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_RetrieveAnalysisResponse__Output>): grpc.ClientUnaryCall;
  RetrieveAnalysis(argument: _roof_RetrieveAnalysisRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_roof_RetrieveAnalysisResponse__Output>): grpc.ClientUnaryCall;
  RetrieveAnalysis(argument: _roof_RetrieveAnalysisRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_RetrieveAnalysisResponse__Output>): grpc.ClientUnaryCall;
  RetrieveAnalysis(argument: _roof_RetrieveAnalysisRequest, callback: grpc.requestCallback<_roof_RetrieveAnalysisResponse__Output>): grpc.ClientUnaryCall;
  retrieveAnalysis(argument: _roof_RetrieveAnalysisRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_RetrieveAnalysisResponse__Output>): grpc.ClientUnaryCall;
  retrieveAnalysis(argument: _roof_RetrieveAnalysisRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_roof_RetrieveAnalysisResponse__Output>): grpc.ClientUnaryCall;
  retrieveAnalysis(argument: _roof_RetrieveAnalysisRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_roof_RetrieveAnalysisResponse__Output>): grpc.ClientUnaryCall;
  retrieveAnalysis(argument: _roof_RetrieveAnalysisRequest, callback: grpc.requestCallback<_roof_RetrieveAnalysisResponse__Output>): grpc.ClientUnaryCall;
  
}

export interface RoofServiceProceduresHandlers extends grpc.UntypedServiceImplementation {
  PerformAnalysis: grpc.handleUnaryCall<_roof_PerformAnalysisRequest__Output, _roof_PerformAnalysisResponse>;
  
  Ping: grpc.handleUnaryCall<_roof_PingMessage__Output, _roof_PingMessage>;
  
  RetrieveAnalysis: grpc.handleUnaryCall<_roof_RetrieveAnalysisRequest__Output, _roof_RetrieveAnalysisResponse>;
  
}

export interface RoofServiceProceduresDefinition extends grpc.ServiceDefinition {
  PerformAnalysis: MethodDefinition<_roof_PerformAnalysisRequest, _roof_PerformAnalysisResponse, _roof_PerformAnalysisRequest__Output, _roof_PerformAnalysisResponse__Output>
  Ping: MethodDefinition<_roof_PingMessage, _roof_PingMessage, _roof_PingMessage__Output, _roof_PingMessage__Output>
  RetrieveAnalysis: MethodDefinition<_roof_RetrieveAnalysisRequest, _roof_RetrieveAnalysisResponse, _roof_RetrieveAnalysisRequest__Output, _roof_RetrieveAnalysisResponse__Output>
}
