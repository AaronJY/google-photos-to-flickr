export class Home {
    googleApiKey: string;

    activate(params: any) {
        if (params.googleapikey) {
            this.googleApiKey = params.googleapikey;
        }
    }
}