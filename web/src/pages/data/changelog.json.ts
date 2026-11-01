// /data/changelog.json (Wave 3 contract W3-C1): the site's changelog for the apps' What changed
// list. Not per build: it is site news, not game data.
import type { APIRoute } from 'astro';
import { jsonResponse } from '../../lib/app-data/endpoint';
import { appChangelog, encodeAppFile } from '../../lib/app-data/app-files';

export const GET: APIRoute = () => jsonResponse(encodeAppFile('changelog', 'changelog.json', appChangelog()));
