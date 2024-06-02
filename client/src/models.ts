export interface Request {
	method: string;
	params: any[];
	param_type: string[];
	id: number;
}

export interface Response {
	result: any;
	result_type: string;
	id: number;
	error: string;
}
