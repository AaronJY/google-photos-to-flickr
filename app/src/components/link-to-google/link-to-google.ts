import { inject, bindable } from "aurelia-framework";
import { GoogleAppHttpService } from "services/GoogleAppHttpService";

@inject(GoogleAppHttpService)
export class LinkToGoogle {
  googleAppHttpService: GoogleAppHttpService;

  label: string = "Link your Google account!";
  @bindable disabled: boolean;

  constructor(googleHttpService: GoogleAppHttpService) {
    this.googleAppHttpService = googleHttpService;
  }

  onClick(event: MouseEvent) {
    const authUrl = this.googleAppHttpService.getAuthUrl();
    window.location.href = authUrl;
  }
}
