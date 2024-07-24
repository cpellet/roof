import { promisify } from "util";
import { RoofServiceProcedures } from "./client";
import * as grpc from "@grpc/grpc-js";

/**
 * Represents a service for interacting with the Roof API.
 */
export class RoofService extends RoofServiceProcedures {
  constructor() {
    super(process.env.NEXT_PUBLIC_HOST, grpc.credentials.createInsecure());
  }

  /**
   * Sends a ping request to the Roof API.
   * @param payload - The payload to send with the ping request.
   * @returns A promise that resolves to the response or an error object.
   */
  public async sendPing(payload: string) {
    const pingInfo = promisify(this.Ping).bind(this);
    return await pingInfo({ message: payload })
      .then((res) => ({ res, error: null }))
      .catch((error) => ({ res: null, error }));
  }

  /**
   * Sends an analysis request to the Roof API.
   * @param cmap - The buffer containing the cmap data.
   * @param emap - The buffer containing the emap data.
   * @returns A promise that resolves to the response or an error object.
   */
  public async sendAnalysisRequest(cmap: Buffer, emap: Buffer) {
    const requestInfo = promisify(this.PerformAnalysis).bind(this);
    return await requestInfo({ cmap, emap })
      .then((res) => ({ res, error: null }))
      .catch((error) => ({ res: null, error }));
  }

  /**
   * Sends a results request to the Roof API.
   * @param id - The ID of the analysis to retrieve results for.
   * @returns A promise that resolves to the response or an error object.
   */
  public async sendResultsRequest(id: string) {
    const resultsInfo = promisify(this.RetrieveAnalysis).bind(this);
    return await resultsInfo({ id })
      .then((res) => ({ res, error: null }))
      .catch((error) => ({ res: null, error }));
  }
}
