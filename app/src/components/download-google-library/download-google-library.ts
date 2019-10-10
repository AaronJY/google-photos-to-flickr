import { computedFrom } from "aurelia-binding";
import { GoogleHttpService } from "services/GoogleHttpService";
import { inject, bindable } from "aurelia-framework";

@inject(GoogleHttpService)
export class DownloadGoogleLibrary {
    // Dependencies
    googleHttpService: GoogleHttpService;

    // Properties
    progress: number = 0;
    @bindable apiKey: string;
    
    constructor(
        googleHttpService: GoogleHttpService
    ) {
        this.googleHttpService = googleHttpService;
    }

    @computedFrom("progress")
    get progressWidthPercentage() {
        return Math.floor(this.progress);
    }

    attached() {
        this.googleHttpService.setApiToken(this.apiKey);
    }
}