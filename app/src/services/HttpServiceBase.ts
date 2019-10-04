import { HttpClient } from "aurelia-fetch-client";

export class HttpServiceBase {
	protected httpClient: HttpClient;

	protected readonly ApiBaseUrl = "http://localhost:1337/api/";

	constructor(httpClient: HttpClient) {
		this.httpClient = httpClient;
	}
}