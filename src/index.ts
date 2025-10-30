import { Container, getContainer } from '@cloudflare/containers';

export class CronContainer extends Container {
  sleepAfter = '10s';

  override onStart() {
    console.log('starting container');
  }

  override onStop() {
    console.log('container stopped');
  }
}

export default {
  async fetch(): Promise<Response> {
    return new Response("This Worker runs a cron job to execute a container on a schedule.");
  },

  async scheduled(_controller: any, env: { CRON_CONTAINER: DurableObjectNamespace<CronContainer> }) {
    let container = getContainer(env.CRON_CONTAINER);
    await container.start({
      envVars: {
        LOCATION: "The Cloudflare San Francisco Office",
        LATITUDE: "37.780259",
        LONGITUDE: "-122.390519",
      }
    })
  },
};
