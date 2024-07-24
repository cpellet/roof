/* eslint-disable @next/next/no-img-element */
'use server'

import { RoofService } from "@/lib/service";
import { Button } from "@mantine/core";
import Link from "next/link";

/**
 * Fetches the results of the analysis from the gRPC service.
 * 
 * @param id - The id of the analysis to fetch the results for.
 * @returns {Promise<{ cmap: string, msmap: string, error: Error }>} A promise that resolves to an object containing the cmap, msmap, and error.
 */
async function fetchResults(id: string) {
    const roofService = new RoofService();
    const { res, error } = await roofService.sendResultsRequest(id);
    const cmap = res!.cmap.toString("base64");
    const msmap = res!.msmap.toString("base64");
    return { cmap, msmap, error };
}

export default async function ResultsPage({ searchParams }: { searchParams: { [key: string]: string | string[] | undefined } }) {
    const id = searchParams["id"];
    const { cmap, msmap } = await fetchResults(id as string);
    return (
        <main className="flex min-h-screen flex-col items-center justify-center p-24">
            <Link href="/" className="absolute top-0 left-0 m-12">
                <Button color="blue" radius="md" fullWidth>New Analysis</Button>
            </Link>
            <img src={`data:image/png;base64,${cmap}`} alt="" className="absolute" />
            <img src={`data:image/png;base64,${msmap}`} alt="" className="absolute opacity-80" />
        </main>

    )
}