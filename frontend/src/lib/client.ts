import { loadPackageDefinition } from "@grpc/grpc-js";
import { loadSync } from "@grpc/proto-loader";
import { ProtoGrpcType } from "./generated/roof";

// roofPackage is the generated TypeScript type for the roof package in the gRPC service.

const PROTO_PATH = process.env.PROTO_PATH || ".";

const packageDefinition = loadSync(PROTO_PATH, {
  keepCase: true,
  defaults: true,
  oneofs: true,
});

const roofPackage = (
  loadPackageDefinition(packageDefinition) as unknown as ProtoGrpcType
).roof;

export const { RoofServiceProcedures } = roofPackage;
