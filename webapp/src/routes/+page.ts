import type { Drink } from '$lib/types';
import type { PageLoad } from './$types';

export const load = (async () => {
    let drinks: Drink[] = await fetch('/drinks').then(res => res.json());
    return { drinks };
}) satisfies PageLoad;
