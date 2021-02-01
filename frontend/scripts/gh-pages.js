const ghpages = require("gh-pages");

ghpages.publish(
	"public",
	{
		branch: "gh-pages",
		repo: "https://github.com/NightlySide/TwitterTrendR.git",
		user: {
			name: "Nightlyside",
			email: "nightlyside@gmail.com"
		}
	},
	() => {
		console.log("Deploy Complete!");
	}
);
