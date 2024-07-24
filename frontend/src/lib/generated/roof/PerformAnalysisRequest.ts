// Original file: ../shared/roof.proto


export interface PerformAnalysisRequest {
  'cmap'?: (Buffer | Uint8Array | string);
  'emap'?: (Buffer | Uint8Array | string);
}

export interface PerformAnalysisRequest__Output {
  'cmap': (Buffer);
  'emap': (Buffer);
}
