import { createConnection } from "node:net";
import { getUnixSocketPath } from "./config";
import type { Request, Response } from "./models";

/**
 * Unixソケットを使用してサーバーに接続し、リクエストを送信する。
 * @param {Request} request - リクエストオブジェクト
 * @returns - サーバーからのレスポンスを持つPromise
 */
export async function connectToServer(request: Request): Promise<Response> {
	return new Promise((resolve, reject) => {
		const client = createConnection(getUnixSocketPath());

		client.on("connect", () => {
			client.write(JSON.stringify(request));
		});

		let responseData = "";

		client.on("data", (data) => {
			responseData += data.toString();
		});

		client.on("end", () => {
			try {
				const response: Response = JSON.parse(responseData);
				resolve(response);
			} catch (error) {
				reject(new Error(`Faild to parse server response: ${error}`));
			}
		});

		client.on("error", (error) => {
			reject(new Error(`Faild to connect to server: ${error.message}`));
		});
	});
}
