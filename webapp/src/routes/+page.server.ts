import type { Drink } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load = (async () => {
    let drinks: Drink[] | null = await fetch(`http://localhost:3000/drinks`).then(res => res.json());
    if (drinks === null) return { drinks: [] };
    return { drinks };
}) satisfies PageServerLoad;
