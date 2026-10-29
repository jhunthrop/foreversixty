// web/src/lib/watch-media.ts
// Reads a media query now and on every change; returns the cleanup for an $effect.
export function watchMedia(query: string, onChange: (matches: boolean) => void): () => void {
  const list = window.matchMedia(query);
  const read = (): void => onChange(list.matches);
  read();
  list.addEventListener('change', read);
  return () => list.removeEventListener('change', read);
}
