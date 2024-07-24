import { Group, Text, rem } from "@mantine/core";
import { Dropzone, FileRejection, FileWithPath } from "@mantine/dropzone";
import { IconProps, IconUpload, IconX } from '@tabler/icons-react'

/**
 * A component that allows users to drag and drop files or click to select files.
 *
 * @param mimeType - An array of allowed MIME types for the dropped files.
 * @param onDrop - A callback function that is called when files are dropped.
 * @param onReject - (Optional) A callback function that is called when files are rejected.
 */
export default function FileDropzone({ mimeType, onDrop, onReject, Icon, fileName }: {
    mimeType: string[],
    onDrop: (files: FileWithPath[]) => void,
    onReject?: (rejections: FileRejection[]) => void,
    Icon: React.FC<IconProps>,
    fileName?: string
}) {
    return <Dropzone
        onDrop={onDrop}
        onReject={onReject}
        maxSize={5 * 1024 ** 2}
        accept={mimeType}
        maxFiles={1}
        className="my-4"
    >
        <Group justify="center" gap="xl" mih={80} style={{ pointerEvents: 'none' }} className="px-12">
            <Dropzone.Accept>
                <IconUpload
                    style={{ width: rem(52), height: rem(52), color: 'var(--mantine-color-blue-6)' }}
                    stroke={1.5}
                />
            </Dropzone.Accept>
            <Dropzone.Reject>
                <IconX
                    style={{ width: rem(52), height: rem(52), color: 'var(--mantine-color-red-6)' }}
                    stroke={1.5}
                />
            </Dropzone.Reject>
            <Dropzone.Idle>
                <Icon
                    style={{ width: rem(52), height: rem(52), color: 'var(--mantine-color-dimmed)' }}
                    stroke={1.5}
                />
            </Dropzone.Idle>
            {fileName != undefined ? <Text>{fileName}</Text> :
                <div>
                    <Text size="md" inline>
                        Drag files here or click to select
                    </Text>
                    <Text size="sm" c="dimmed" inline mt={7}>
                        Attach one file, which should not exceed 5mb
                    </Text>
                </div>
            }
        </Group>
    </Dropzone>
}