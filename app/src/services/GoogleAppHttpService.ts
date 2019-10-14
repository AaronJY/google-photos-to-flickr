import { HttpClient } from "aurelia-fetch-client";

export class GoogleAppHttpService extends HttpClient {
    private readonly ApiBaseUrl = "http://localhost:1337/api/google/";

    constructor() {
        super();

        this.configure(config => {
            config
                .useStandardConfiguration()
                .withBaseUrl(this.ApiBaseUrl);
        });
    }

    getAuthUrl(): string {
        return `${this.ApiBaseUrl}auth`;
    }

    download() {
        
    }
}