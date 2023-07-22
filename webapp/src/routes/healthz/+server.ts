import type { RequestHandler } from './$types';
import { env } from '$env/dynamic/private';

export const GET: RequestHandler = async () => {
  try {
    const response = await fetch(env.API_URL ? env.API_URL : "http://localhost:3000");
    if (response.ok) {
      console.log("healthcheck: ok")
      return new Response("ok", { status: 200 });
    } else {
      console.error(Date.now(), "healthcheck: not ok")
      return new Response(null, { status: 500 });
    }
  } catch (error) {
    console.error("healthcheck: not ok")
    return new Response(null, { status: 500 });
  }
};

