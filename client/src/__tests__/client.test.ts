import { afterAll, beforeAll, describe, expect, test } from "bun:test";
import { type Server, createServer } from "node:net";
import { connectToServer } from "../client";
import { getUnixSocketPath } from "../config";
import type { Request, Response } from "../models";

describe("connectToServer", () => {
	let server: Server;

	beforeAll((done) => {
		server = createServer((socket) => {
			socket.on("data", (data) => {
				const request: Request = JSON.parse(data.toString());
				const response: Response = {
					result: Math.floor(request.params[0] as number),
					result_type: "number",
					id: request.id,
					error: "",
				};
				socket.write(JSON.stringify(response));
				socket.end();
			});
		});

		server.listen(getUnixSocketPath(), done);
	});

	afterAll((done) => {
		server.close(done);
	});

	test("should connect to server and get a response", async () => {
		const request: Request = {
			method: "floor",
			params: [10.7],
			param_types: ["number"],
			id: 1,
		};

		const response = await connectToServer(request);
		expect(response.result).toBe(10);
		expect(response.result_type).toBe("number");
		expect(response.id).toBe(1);
		expect(response.error).toBe("");
	});
});
