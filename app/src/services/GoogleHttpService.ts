import { inject } from "aurelia-framework";
import { HttpClient } from "aurelia-fetch-client"
import { HttpServiceBase } from "./HttpServiceBase";

@inject(HttpClient)
export class GoogleHttpService extends HttpServiceBase {
	private readonly UrlPrefix = `${this.ApiBaseUrl}google/`;

	getAuthUrl(): string {
		return `${this.UrlPrefix}auth`;
	}
}
