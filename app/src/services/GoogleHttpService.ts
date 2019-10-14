import { HttpClient } from "aurelia-fetch-client";
import { MediaItem, GoogleResponseRootObject } from "interfaces/GoogleApiInterfaces";
import { HttpClientHelper } from "helpers/HttpClientHelper";

export class GoogleHttpService extends HttpClient {
	private readonly GoogleUrlPrefix = `https://photoslibrary.googleapis.com/v1/`;

	private apiToken: string;

	constructor() {
		super();

		this.configure(config => {
			config
				.useStandardConfiguration()
				.withBaseUrl(this.GoogleUrlPrefix)
				.withDefaults({
					headers: {
						"Authorization": () => "Bearer " + this.apiToken
					}
				})
		});
	}

	setApiToken(apiToken: string) {
		this.apiToken = apiToken;
	}

	async list(pageSize: number, pageToken: string = ""): Promise<GoogleResponseRootObject> {
		const query = {
			"pageSize": pageSize
		};

		if (pageToken)
			query["pageToken"] = pageToken;

		const queryStr = HttpClientHelper.objectToQuery(query);
		const response = await this.get("mediaItems" + queryStr, {
			headers: {
				"Content-Type": "application/json"
			},
			
		}).then(res => res.json() as Promise<GoogleResponseRootObject>);
		return response;
	}
}