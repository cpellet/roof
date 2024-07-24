// Original file: ../shared/roof.proto


export interface RoofAnalysisRequest {
  'cmap'?: (Buffer | Uint8Array | string);
  'emap'?: (Buffer | Uint8Array | string);
}

export interface RoofAnalysisRequest__Output {
  'cmap': (Buffer);
  'emap': (Buffer);
}
