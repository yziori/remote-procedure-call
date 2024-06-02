import { connectToServer } from "./client.ts";
import type { Request, Response } from "./models.ts";

export async function sendRequest(request: Request): Promise<Response> {
	try {
		const response = await connectToServer(request);
		return response;
	} catch (error) {
		throw new Error(`Faild to send request: ${error}`);
	}
}
