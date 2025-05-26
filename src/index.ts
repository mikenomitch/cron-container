import { Container, getContainer } from 'cf-containers';

export class CronContainer extends Container {
  sleepAfter = '10s';
  manualStart = true;

  override onStart() {
    console.log('Starting container');
  }

  override onStop() {
    console.log('Contianer Stopped');
  }
}

export default {
  async fetch(): Promise<Response> {
    return new Response("This Worker runs a cron job to execute a container on a schedule.");
  },

  async scheduled(_controller: any, env: { CRON_CONTAINER: DurableObjectNamespace<CronContainer> }) {
    let container = getContainer(env.CRON_CONTAINER);
    await container.startContainer({
      envVars: {
        LOCATION: "The Cloudflare San Francisco Office",
        LATITUDE: "37.780259",
        LONGITUDE: "-122.390519",
      }
    })
  },
};
