import type { Drink } from '$lib/types';
import type { PageLoad } from './$types';

export const load = (async () => {
    let drinks: Drink[] | null = await fetch('/drinks').then(res => res.json());
    if (drinks === null) drinks = [];
    return { drinks };
}) satisfies PageLoad;
