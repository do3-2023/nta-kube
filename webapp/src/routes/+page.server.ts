import type { PageServerLoad } from './$types';
export const load = (async ({ params }) => {
    return { 
        msg: "hello world"
    };
}) satisfies PageServerLoad;