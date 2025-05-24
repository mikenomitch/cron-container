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
  async scheduled(_controller: any, env: any, _ctx: any) {
    let container = getContainer(env.CRON_CONTAINER);
    container.startContainer({
      envVars: {
        LOCATION: "The Cloudflare San Francisco Office",
        LATITUDE: "37.780259",
        LONGITUDE: "-122.390519",
      }
    })
  },
};
