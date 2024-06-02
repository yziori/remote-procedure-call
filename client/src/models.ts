export type ParamType = number | string;

export interface Request {
	method: string;
	params: ParamType[];
	param_types: string[];
	id: number;
}

export interface Response {
	result: ParamType | ParamType[] | boolean;
	result_type: string;
	id: number;
	error: string;
}
