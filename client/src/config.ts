import { resolve } from "path";

export function getUnixSocketPath():string {
    return resolve('/tmp/unix_socket')
}