# go-reveal

An small tool for those who want to impress with a shiny [reveal.js](https://github.com/hakimel/reveal.js/) slide deck by just writing a markdown file.

## The feature list

Honoring the [Readme Driven Development](http://tom.preston-werner.com/2010/08/23/readme-driven-development.html), let's describe the tool by its features:

- [x] Read markdown file and serve it as reveal.js slide deck.
- [x] Serve current working directory (`pwd`) children when the markdown file path is not specified.
- [x] Add slide and section modifiers as comments Ex.: `.element: class="fragment"` [Read more](https://github.com/hakimel/reveal.js/#element-attributes)
- [ ] Allow to override default configurations with flags, env variables or yaml/rc files.
- [ ] Offline mode - Bundle reveal.js latest version as part of the package.
- [ ] Allow to change the reveal.js theme with any of the stock ones.
- [ ] Render demo page with brief explanation.
- [ ] Improve `serve` logging format.
- [ ] Add presenter notes support
- [ ] Fix images' relative path when a markdown slide is render outside of `pwd`
- [ ] Export markdown slide deck to reveal.js html page.
- [ ] Be able to add a custom reveal.js theme (css file).
- [ ] Be able to add a custom js files.
- [ ] Override reveal.js options using tool configurations.
- [ ] (nice to have) Update reveal.js offline styles and scripts on demand.
- [ ] (nice to have) Support [multiplexer](https://github.com/hakimel/reveal.js#multiplexing) reveal.js plugin using websockets.
- [ ] (nice to have) Export slide deck in html format with inline styles and scripts.
- [ ] (nice to have) Watch for changes in markdown and livereload the presentation.
- [ ] (nice to have) Export slide deck as PDF.
- [ ] ???
- [ ] Profit

## FAQ

- Why go-reveal instead of just sticking with [deck](https://github.com/ajstarks/deck)? Short and simple, there's **nothing** wrong with deck but I'm pretty much used to [reveal-md](https://github.com/webpro/reveal-md), so I want to mimic its behavior in a self contained tool.

- Why did you write a tool instead of getting [reveal-md](https://github.com/webpro/reveal-md) inside a docker container?
  - In a scenario with a poor bandwidth, downloading a `alpine + node.js + reveal-md (+ deps)` docker image is a **bad idea** and you still need a docker-engine up 'n running in the computer that you want to use in your presentation. I just want something to carry with me in any external storage device, be able to plug-in it and start the presentation right away.
  - And for fun :octocat:
