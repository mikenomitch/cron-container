# Cron Container Starter

[![Deploy to Cloudflare](https://deploy.workers.cloudflare.com/button)](https://deploy.workers.cloudflare.com/?url=https://github.com/mikenomitch/cron-container)

<!-- dash-content-start -->

This is a template showing running [Cron Worker](https://developers.cloudflare.com/workers/configuration/cron-triggers/) using a [Container](https://developers.cloudflare.com/containers/).

<!-- dash-content-end -->

Outside of this repo, you can start a new project with this template using [C3](https://developers.cloudflare.com/pages/get-started/c3/) (the `create-cloudflare` CLI):

```bash
npm create cloudflare@latest -- --template=mikenomitch/cron-container
```

## Getting Started

First, run:

```bash
npm install
# or
yarn install
# or
pnpm install
# or
bun install
```

You can start editing your Worker by modifying `src/index.ts` and you can start
editing your Container by editing the content of `container_src`.

## Deploying To Production

| Command          | Action                                |
| :--------------- | :------------------------------------ |
| `npm run deploy` | Deploy your application to Cloudflare |

## Learn More

To learn more about Containers, take a look at the following resources:

- [Cron Triggers in Workers](mikenomitch/cron-container) - learn about cron Workers
- [Container Documentation](https://developers.cloudflare.com/containers/) - learn about Containers
- [Beta Information](https://developers.cloudflare.com/beta-info/) - learn about the Containers Beta
- [Container Class](https://github.com/cloudflare/containers) - learn about the Container helper class

Your feedback and contributions are welcome!
