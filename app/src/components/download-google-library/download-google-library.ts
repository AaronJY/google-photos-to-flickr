import { computedFrom } from "aurelia-binding";
import { GoogleHttpService } from "services/GoogleHttpService";
import { inject, bindable } from "aurelia-framework";
import { MediaItem, GoogleResponseRootObject } from "interfaces/GoogleApiInterfaces";
import { promises } from "fs";

@inject(GoogleHttpService)
export class DownloadGoogleLibrary {
    // Dependencies
    googleHttpService: GoogleHttpService;

    // Properties
    progress: number = 0;
    isProcessing: boolean;
    isUploading: boolean;
    mediaItems: MediaItem[] = [];
    @bindable apiKey: string;

    readonly pageSize: number = 100;

    constructor(
        googleHttpService: GoogleHttpService) {
        this.googleHttpService = googleHttpService;
    }

    @computedFrom("progress")
    get progressWidthPercentage() {
        return Math.floor(this.progress);
    }

    async attached() {
        this.googleHttpService.setApiToken(this.apiKey);

        return this.processGoogleImages();
    }

    async processGoogleImages(): Promise<void> {
        this.isProcessing = true;

        let nextPageToken = null;
        let response: GoogleResponseRootObject;

        try {
            do {
                response = await this.googleHttpService.list(this.pageSize, nextPageToken);
                nextPageToken = null;
    
                this.mediaItems.push(...response.mediaItems);
    
                console.log(response);
            } while (nextPageToken);
        } catch (err) {
            return Promise.reject(err);
        }
        
        this.isProcessing = false;

        return Promise.resolve();
    }
}