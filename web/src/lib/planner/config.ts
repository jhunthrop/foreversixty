// web/src/lib/planner/config.ts
// Build-time constants for the planner. PUBLIC_API_BASE_URL is inlined by Vite into both
// the Astro build and the standalone island bundle, so the island works identically when
// the API serves it from /b/:id.

/** Where POST /v1/builds and /v1/subscribe live. */
export const API_BASE_URL: string = import.meta.env.PUBLIC_API_BASE_URL ?? 'https://api.foreversixty.gg';

/** The class the planner opens on when the query string does not say otherwise. */
export const DEFAULT_CLASS_SLUG = 'warrior';
