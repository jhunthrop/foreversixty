// web/src/lib/billing/example-image.ts
// The one source for the premium example picture: /premium and the /sim landing card both read
// it, so the files, the reserved box and the alt text can never disagree.
import { serverRunCaption } from './plans-page-copy';
import { serverRunExample, type ServerRunExample } from './premium-examples';

export const EXAMPLE_IMAGE_WIDTH = 440;
export const EXAMPLE_IMAGE_HEIGHT = 330;

const EXAMPLE_IMAGE_BASE = '/premium/example-server-run';

export interface ExampleImage {
  avif: string;
  webp: string;
  width: typeof EXAMPLE_IMAGE_WIDTH;
  height: typeof EXAMPLE_IMAGE_HEIGHT;
  alt: string;
  loading: 'lazy';
  decoding: 'async';
}

export function exampleImage(ex: ServerRunExample = serverRunExample): ExampleImage {
  return {
    avif: `${EXAMPLE_IMAGE_BASE}.avif`,
    webp: `${EXAMPLE_IMAGE_BASE}.webp`,
    width: EXAMPLE_IMAGE_WIDTH,
    height: EXAMPLE_IMAGE_HEIGHT,
    alt: serverRunCaption(ex),
    loading: 'lazy',
    decoding: 'async',
  };
}
