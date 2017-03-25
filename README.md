# go-reveal

An small tool for those who want to impress with a shiny [reveal.js](https://github.com/hakimel/reveal.js/) slide deck by just writing a markdown file.

## The feature list

Honoring the [Readme Driven Development](http://tom.preston-werner.com/2010/08/23/readme-driven-development.html), let's describe the tool by its features:

- [ ] Read markdown file and serve it as reveal.js slide deck.
- [ ] Allow to override default configurations with flags, env variables or yaml/rc files.
- [ ] Render demo page with brief explanation.
- [ ] Allow to change the reveal.js theme with any of the stock ones.
- [ ] Add slide and section modifiers as metadata or comment. Ex.: `.element: class="fragment"`
- [ ] Offline mode - Bundle reveal.js latest version as part of the package.
- [ ] Serve current working directory (`pwd`) children when the markdown file path is not specified.
- [ ] Export markdown slide deck in html format as a reveal.js one.
- [ ] Provide custom reveal.js theme (css file).
- [ ] Update reveal.js offline styles and scripts on demand.
- [ ] Override reveal.js options using tool configurations.
- [ ] (nice to have) Support [multiplexer](https://github.com/hakimel/reveal.js#multiplexing) reveal.js plugin using websockets.
- [ ] (nice to have) Export slide deck in html format with inline styles and scripts.
- [ ] (nice to have) Watch for changes in markdown and livereload the presentation.
- [ ] (nice to have) Export slide deck as PDF.
- [ ] ???
- [ ] Profit


## FAQ

- Why go-reveal instead of just sticking with [deck](https://github.com/ajstarks/deck)?
Short and simple, there's **nothing** wrong with deck but I'm pretty much used to [reveal-md](https://github.com/webpro/reveal-md), so I want to mimic its behavior in a self contained tool.

- Why did you write a tool instead of getting [reveal-md](https://github.com/webpro/reveal-md) inside a docker container?
  - For fun
  - In a scenario with a poor bandwith to download a `alpine + node.js + reveal-md (+ deps)` docker image is a **bad idea** and you still need a docker-engine up 'n running in the machine that you're going to use. I just want something to carry with me in any external storage device, be able to plug-in it and start the presentation right away.
