# Introduction

This is a minimal practical blog system written in vanilla Js and a bit of Go. Why? 

There is a lot of blogging software floating around. One can choose, say, Hugo or Astro. They have a lot of templates/themes, some pretty cool, like [Hugo.386](https://themes.gohugo.io/themes/hugo.386/). Astro has view transitions and React support and is virtually another full-blown Next.js. All these nice features also bring overengineering, complexity, bugs.

Here is what I really need: 

- [x] An expandable list of articles with the first p-element displayed by default. 
- [x] An article is any folder "article*" which contains "index.md". 
- [x] A simple responsive layout.
- [x] Markdown and LaTeX, via, say, [gomarkdown](https://github.com/gomarkdown/markdown).
- [x] The ability to test everything locally directly by opening index.html in a browser. No cdn, node_modules, cors, http servers.
- [ ] Code syntax highlighter. 

What I do not care about: All the meta stuff like tags, reading time, time stamps, comments, contact forms, search bars, TOC, archive, recent posts, suggestions, internal navigation links, CMS/online rich text editor. A blog is not a corporate media system, it is about personal notes, mostly. A single github readme would suffice, but it is better to control Markdown and LaTeX rendering explicitly and locally.

I do not care about LaTeX either, but see [this raw Markdown](https://raw.githubusercontent.com/aabbtree77/aabbtree77.github.io/main/shankland/shankland.md) and [the rendered result](https://github.com/aabbtree77/aabbtree77.github.io/blob/main/shankland/shankland.md) for an example if you need to add math to your html pages (not a good idea).

# Setup

Install Go, cd into gocode, run "go build", move the resulting "md2html" to the root of this repo. Run it ("./md2html"). It will scan all the root level "article*" folders with "index.md" files, turn Markdown to HTML and then inject those HTML pieces into "index_template.html" resulting in one big "index.html". This "index.html" file, along with the "css", "js", and "imgs" folders, serves as a blog. I host it using Github Pages.

# References

[How to Deploy a Static Website for Free Using Github Pages, 2022.](https://medium.com/flycode/how-to-deploy-a-static-website-for-free-using-github-pages-8eddc194853b)

