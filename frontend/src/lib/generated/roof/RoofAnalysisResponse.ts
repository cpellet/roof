// Original file: ../shared/roof.proto


export interface RoofAnalysisResponse {
  'uvmap'?: (Buffer | Uint8Array | string);
  'msmap'?: (Buffer | Uint8Array | string);
  'id'?: (string);
}

export interface RoofAnalysisResponse__Output {
  'uvmap': (Buffer);
  'msmap': (Buffer);
  'id': (string);
}
