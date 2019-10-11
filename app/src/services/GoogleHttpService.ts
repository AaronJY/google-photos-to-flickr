import { HttpServiceBase } from "./HttpServiceBase";

export class GoogleHttpService extends HttpServiceBase {
	readonly UrlPrefix = `${this.ApiBaseUrl}google/`;
	readonly GoogleUrlPrefix = `https://photoslibrary.googleapis.com/v1/`;

	apiToken: string;

	getAuthUrl(): string {
		return `${this.UrlPrefix}auth`;
	}

	setApiToken(apiToken: string): GoogleHttpService {
		this.apiToken = apiToken;
		return this;
	}

	list(pageSize: number, pageToken: string = "") {
		this.httpClient.fetch(this.ep("mediaItems"))
			.then(resp => {
				console.log(resp);
			});
	}

	private ep(partialEndpoint: string) {
		return this.GoogleUrlPrefix + partialEndpoint;
	}
}
