import { expect, test, describe } from "bun:test";
import type { Request } from "../models";
import { parseArgs } from "../utils";

describe("parseArgs", () => {
	test("数値と文字列を含む有効な引数を解析する", () => {
		const args = ["floor", "10.7", "test"];
		const expectedRequest: Request = {
			method: "floor",
			params: [10.7, "test"],
			param_types: ["number", "string"],
			id: expect.any(Number),
		};

		const request: Request | null = parseArgs(args);
		expect(request).toMatchObject(expectedRequest);
	});

	test("数値のみを含む有効な引数を解析する", () => {
		const args = ["floor", "10.7", "5"];
		const expectedRequest: Request = {
			method: "floor",
			params: [10.7, 5],
			param_types: ["number", "number"],
			id: expect.any(Number),
		};

		const request: Request | null = parseArgs(args);
		expect(request).toMatchObject(expectedRequest);
	});

	test("引数が不足している場合はnullを返す", () => {
		const args = ["floor"];
		const request: Request | null = parseArgs(args);
		expect(request).toBeNull();
	});

	test("混合引数のパラメータタイプを正しく返す", () => {
		const args = ["method", "123", "string", "45.6"];
		const expectedRequest: Request = {
			method: "method",
			params: [123, "string", 45.6],
			param_types: ["number", "string", "number"],
			id: expect.any(Number),
		};

		const request: Request | null = parseArgs(args);
		expect(request).toMatchObject(expectedRequest);
	});
});
