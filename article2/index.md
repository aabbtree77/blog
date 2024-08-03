# TypeScript and Golang Ecosystems

Let me provide an outsider's view to these two ecosystems.

### TypeScript

- Stuff built with it: tutorials [how to build stuff with it](https://www.youtube.com/watch?v=ucX2zXAZ1I0)... VS Code, Excalidraw, Babylon.js, [JupyterLab](https://github.com/jupyterlab/jupyterlab), [gdbgui](https://www.gdbgui.com/howitworks/). Predominantly GUIs.

- Unusually high levels of self-inflicted pain: serverless, headless, monorepo constraints (backend-frontend fusion), enshittified astroturfed 3rd party APIs and SDKs, DDOS stories, old and new modules [(.js, .cjs, .mjs)](https://codingforseo.com/blog/mjs-vs-cjs-files/), external types (.ts, .d.ts, jsdoc), not accepting JSX (.tsx, .jsx) as part of the Js/Ts standard, "mutation management" frameworks, frameworks within frameworks, metaframeworks...

- Still no compiler or cross-compiler to a binary like Go. Ts is transpiled to Js, the latter is then JIT-compiled to a binary, but these are known to be memory hogs and the compilation and bundling is not as seamless as in Go. A minimal RAM requirement for a Ts app is [100MB](https://www.youtube.com/watch?v=gNDBwxeBrF4&t=2s) which is 10x more than in Go.

- Types are always a hindrance in the presence of generics and [colored functions](https://journal.stuffwithstuff.com/2015/02/01/what-color-is-your-function/). In Ts apps, we also get an extra dimension of types vs validation, typed or validated .env parameters, routes, REST API, forms, SQL. Untyped libs, VS Code server type crashes and restarts, [type gymnastics](https://github.com/type-challenges/type-challenges), [back to Js with JSDoc stories](https://x.com/dummdidumm_/status/1639394200872529925).

- A solid helpful sharing community: [@joshtriedcoding](https://www.youtube.com/@joshtriedcoding), [@bmdavis419](https://www.youtube.com/@bmdavis419/videos), [@YourAverageTechBro](https://www.youtube.com/@YourAverageTechBro), [@marc-lou](https://www.youtube.com/@marc-lou), [syntax.fm](https://syntax.fm/), [@codewithantonio](https://www.youtube.com/@codewithantonio/videos)... They do not always emphasize that the video is actually an ad to a paid 3rd party service which the author is not using in his main projects.

- Bottom line: GUIs are harder than it seems. Ts presents a modern way to build them, with a massive user base. Notice, however, that Figma is written in C++ compiled to asm.js, and Ubuntu adds [Flutter/Dart](https://ubuntu.com/blog/flutter-and-ubuntu-so-far) on top of the GTK4.

### Go

- Stuff built with it: [PocketBase](https://pocketbase.io/), Caddy, Docker, go-libp2p, Syncthing, [Berty](https://github.com/berty/berty), [awl](https://github.com/anywherelan/awl)...

- Older engineers also use it for the SSR (MPA) with templating, ORM or raw SQL, and Docker VPS single binary deployments. YAGNI, KISS, DIY, but no standard pipeline and templating with htmx is not Go.

- Go as a language [is not perfect](https://www.reddit.com/r/typescript/comments/15y9tbm/anyone_who_tried_to_like_golang_and_tried_it_many/). Sadly, the perfect ones are either dynamic (Python, Elixir), too academic (Elm, ReasonML, PureScript, Gleam, Scala), not Linux-centric (F#, Scala), too low level (Zig), still lacking quality sum types and focus (Nim), or too overengineered (D, Rust, Scala... be my guest finding a language designer who is not into compile time gymnastics).

- Older, but also quite vocal community: [Eli Bendersky](https://eli.thegreenplace.net/), [@ThePrimeagen](https://www.youtube.com/@ThePrimeagen/videos), [@anthonygg_](https://www.youtube.com/@anthonygg_), [Adrian Hesketh](https://youtu.be/pArsmWLvZsA?t=1853), [@WebDevCody](https://www.youtube.com/@WebDevCody)...

- Bottom line: Go is probably the only practical language that does not induce a language fatigue. However, I doubt that using SSR with Go, Templ, and htmx is the best way to build web apps. We now end up with one or two obscure extra languages with no convergence around the ORM (or no ORM), unlike, say, in the case of Drizzle in TypeScript. A basic frontend with React (mostly JSX with useState and useEffect, and React Router) is a better technology than vanilla JavaScript and templates, is it not?

