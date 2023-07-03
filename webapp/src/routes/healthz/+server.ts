import type { RequestHandler } from './$types';

export const GET: RequestHandler = async () => {
  try {
    const response = await fetch('http://nta-kube-api.api.svc.cluster.local:3000');
    if (response.ok) {
      return new Response("ok", { status: 200 });
    } else {
      return new Response(null, { status: 500 });
    }
  } catch (error) {
    return new Response(null, { status: 500 });
  }
};

