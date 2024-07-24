"use server"

import { Button, Card, FileInput, Input, Text } from "@mantine/core";
import { RoofService } from "@/lib/service";
import { redirect } from "next/navigation";
import FileDropzone from "./FileDropzone";
import { IMAGE_MIME_TYPE } from "@mantine/dropzone";


/**
 * Submits the form data for analysis to the gRPC service.
 * 
 * @param formData - The form data to be submitted.
 */
export async function onSubmit(formData: FormData) {
    let id = "";
    if ((formData.get("image") as File).name == 'undefined' || (formData.get("heightData") as File).name == 'undefined') {
        return;
    }
    try {
        const roofService = new RoofService();
        const imageArrayBuffer = await (
            formData.get("image") as File
        ).arrayBuffer();
        const elevationArrayBuffer = await (
            formData.get("heightData") as File
        ).arrayBuffer();
        const { res } = await roofService.sendAnalysisRequest(
            Buffer.from(imageArrayBuffer),
            Buffer.from(elevationArrayBuffer)
        );
        id = res!.id;
    } catch (error) {
        console.error(error);
        return;
    } finally {
        redirect(`/results?id=${id}`);
    }
}

export default async function NewAnalysisForm() {
    return (
        <form action={onSubmit}>
            <Card shadow="xl" padding="lg" radius="md" withBorder w={600}>
                <Text fw={500} fz={"h3"}>House details</Text>
                <FileInput required label="Upload image" name="image" id="image" className="mt-3" />
                <FileInput required label="Upload height data" name="heightData" id="heightData" className="mt-3" />
                <Button color="blue" fullWidth mt="md" radius="md" type="submit">
                    Find Southernmost Roof
                </Button>
            </Card>
        </form>
    )
}