import * as readline from "node:readline";
import type { Request } from "./models";
import { sendRequest } from "./requestHandler";
import { parseArgs } from "./utils";

const rl = readline.createInterface({
	input: process.stdin,
	output: process.stdout,
});

async function main(): Promise<void> {
	while (true) {
		const input = await new Promise<string>((resolve) => {
			rl.question("Enter method and parameters: ", resolve);
		});

		const args = input.split(" ");
		const request: Request | null = parseArgs(args);

		if (request) {
			try {
				const response = await sendRequest(request);
				if (!response.error) {
					console.log("Response:", response.result);
				} else {
					console.error("Error in response:", response.error);
				}
			} catch (error) {
				console.error("Error:", error);
			}
		} else {
			console.error("Invalid arguments. Please try again.");
		}
	}

	rl.close();
}

main();
