import { inject } from "aurelia-framework";
import { GoogleHttpService } from "services/GoogleHttpService";

@inject(GoogleHttpService)
export class LinkToGoogle {
  private label: string = "Link your Google account!";

  private googleHttpService: GoogleHttpService;

  constructor(googleHttpService: GoogleHttpService) {
    this.googleHttpService = googleHttpService;
  }

  private onClick(event: MouseEvent) {
    const authUrl = this.googleHttpService.getAuthUrl();
    window.location.href = authUrl;

    console.log('Assigned window location...');
  }
}
