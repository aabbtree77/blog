# Introduction

This is a minimal practical blog system written in vanilla Js and a bit of Go. Why? 

There is a lot of blogging software floating around. One can choose, say, Hugo or Astro. They have a lot of templates/themes, some pretty cool, like [Hugo.386](https://themes.gohugo.io/themes/hugo.386/). Astro has view transitions and React support and is virtually another full-blown Next.js. All these nice features also bring the onboarding time to waste, complexity, bugs.

Here is what I really need: 

- [x] An expandable list of articles with the first p-element displayed by default. 
- [x] An article is any folder in this root with the name starting with "article*". Each such folder must include the file "index.md".
- [x] A simple responsive layout with the chosen Google font and the background color that does not burn eye sight.
- [x] Markdown and LaTeX. This is done with "gomarkdown", see my old blog for a math example: [a raw Markdown](https://raw.githubusercontent.com/aabbtree77/aabbtree77.github.io/main/shankland/shankland.md), [the rendered result](https://github.com/aabbtree77/aabbtree77.github.io/blob/main/shankland/shankland.md). 

- [ ] Code syntax highlighter. 

What I do not care about: All the meta stuff like tags, reading time, time stamps, comments, contact forms, search bars, TOC, archive, recent posts, suggestions, internal navigation links, CMS/online rich text editor. A blog is not a corporate media system, it is about personal notes, mostly. A single github readme would suffice, but it is better to control Markdown and LaTeX rendering explicitly and locally.

# Setup

Install Go, cd into gocode, run "go build", move the resulting "md2html" to the root of this repo. Run it ("./md2html"). It will scan all the root level "article*" folders with "index.md" files, turn Markdown to HTML and then inject those HTML pieces into "index_template.html" resulting in one big "index.html". This "index.html" file, along with the "css", "js", and "imgs" folders, serves a blog. I host it using Github Pages.

# References

[How to Deploy a Static Website for Free Using Github Pages, 2022.](https://medium.com/flycode/how-to-deploy-a-static-website-for-free-using-github-pages-8eddc194853b)


