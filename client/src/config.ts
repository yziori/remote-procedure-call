import { resolve } from "node:path";

export function getUnixSocketPath(): string {
	return resolve("/tmp/unix_socket");
}
