import { RouterConfiguration, Router } from "aurelia-router";
import { PLATFORM } from "aurelia-pal";
import { computedFrom } from "aurelia-binding";

export class App {
	private router: Router;

	@computedFrom("router.title")
	get title() { return this.router.currentInstruction ? this.router.currentInstruction.config.title : "(no title)" };

	configureRouter(config: RouterConfiguration, router: Router) {
		this.router = router;

		config.options.pushState = true;
		config.options.root = "/";
		config.map([
			{ route: ['', 'index'], name: 'index', moduleId: PLATFORM.moduleName("routes/home/index"), title: "GPhotos2Flickr", activationStrategy: "replace" },
		]);
	}
}
