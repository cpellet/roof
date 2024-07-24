import { RoofService } from "@/lib/service";
import { Text } from "@mantine/core";
import { IconCircleFilled } from "@tabler/icons-react";

/**
 * Checks the status by sending a ping request to the RoofService.
 * @returns {Promise<Error>} A promise that resolves to an Error object if there is an error, otherwise resolves to undefined.
 */
async function checkStatus() {
    try {
        const roofService = new RoofService();
        const { res, error } = await roofService.sendPing("ping");
        return { res, error };
    } catch (error) {
        return { res: null, error };
    }
}

export default async function StatusIndicator() {
    const { res, error } = await checkStatus();
    const color = error !== null ? "red" : "green";
    return (
        <div className="flex items-center mt-12 gap-2">
            <IconCircleFilled size={15} color={color} />
            <Text ml="sm" c={color} fw={600}>{error !== null ? "Server is down" : "Ready to analyze"}</Text>
        </div>
    )
}