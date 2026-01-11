# Koffee and Kode ☕️💻

A blog about code, tech, and everything in between.

## Setup

This blog is built with [Jekyll](https://jekyllrb.com/) and automatically deployed to GitHub Pages using GitHub Actions.

### Local Development

To run the blog locally:

1. Install Ruby and Bundler (if not already installed)
2. Install dependencies:
   ```bash
   bundle install
   ```
3. Run Jekyll:
   ```bash
   bundle exec jekyll serve
   ```
4. Open [http://localhost:4000](http://localhost:4000) in your browser

### Adding New Posts

Create a new file in the `_posts` directory with the following naming convention:
```
YYYY-MM-DD-post-title.md
```

Each post should start with front matter:
```yaml
---
layout: post
title: "Your Post Title"
date: YYYY-MM-DD
categories: category-name
---
```

### Deployment

The site automatically deploys to GitHub Pages whenever you push to the `main` branch. The GitHub Actions workflow handles the build and deployment process.

## License

This blog content is open source and available under the [MIT License](LICENSE).

