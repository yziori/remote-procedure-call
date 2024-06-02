import type { Request } from "./models";

/**
 * コマンドライン引数を解析し、Requestオブジェクトを生成する関数
 *
 * @param {string[]} args - コマンドライン引数の配列。最初の要素はメソッド名、以降の要素はパラメータ
 * @returns {Request | null} 解析されたRequestオブジェクト、または引数が不正な場合はnullを返す
 * @throws {Error} 引数の数が2未満の場合にエラーメッセージを出力する
 */
export function parseArgs(args: string[]): Request | null {
	if (args.length < 2) {
		console.error("Usage <method> <param1> <param2> ... <paramN>");
		return null;
	}

	const method = args[0];
	const params = args.slice(1).map((param) => {
		if (!Number.isNaN(Number(param))) {
			return Number(param);
		}
		return param;
	});

	const paramTypes = params.map((param) => typeof param);
	const id = Date.now();

	return {
		method,
		params,
		param_types: paramTypes,
		id,
	};
}
