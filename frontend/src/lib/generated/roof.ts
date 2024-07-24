import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { RoofServiceProceduresClient as _roof_RoofServiceProceduresClient, RoofServiceProceduresDefinition as _roof_RoofServiceProceduresDefinition } from './roof/RoofServiceProcedures';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  roof: {
    PerformAnalysisRequest: MessageTypeDefinition
    PerformAnalysisResponse: MessageTypeDefinition
    PingMessage: MessageTypeDefinition
    Point2D: MessageTypeDefinition
    Polygon: MessageTypeDefinition
    RetrieveAnalysisRequest: MessageTypeDefinition
    RetrieveAnalysisResponse: MessageTypeDefinition
    RoofServiceProcedures: SubtypeConstructor<typeof grpc.Client, _roof_RoofServiceProceduresClient> & { service: _roof_RoofServiceProceduresDefinition }
  }
}

