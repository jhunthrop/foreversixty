// web/src/lib/ui/preload-images.ts
// Holds a reveal until its images have decoded or failed, so they do not pop in after it. The wait
// is capped: a slow image never keeps the page behind a skeleton for longer than `timeoutMs`.
export async function preloadImages(
  urls: readonly string[],
  timeoutMs: number,
  createImage: () => HTMLImageElement = () => new Image(),
): Promise<void> {
  if (urls.length === 0) return;
  const settled = Promise.allSettled(
    urls.map((url) => {
      const image = createImage();
      image.src = url;
      return image.decode();
    }),
  );
  let timer: ReturnType<typeof setTimeout> | undefined;
  const capped = new Promise<void>((resolve) => {
    timer = setTimeout(resolve, timeoutMs);
  });
  await Promise.race([settled, capped]);
  clearTimeout(timer);
}
