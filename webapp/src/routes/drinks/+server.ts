import { env } from '$env/dynamic/private';
import type { RequestHandler } from './$types';

export const POST = (async () => {
  let res = await fetch(`${env.API_URL ? env.API_URL : "http://localhost:3000"}/drinks`, {
    method: 'POST'
  });
  console.log("New drink status:", res.statusText);
  return res;
}) satisfies RequestHandler;

export const GET = (async () => {
  let res = await fetch(`${env.API_URL ? env.API_URL : "http://localhost:3000"}/drinks`);
  console.log("Fetching drinks:", res.statusText);
  return res;
}) satisfies RequestHandler;