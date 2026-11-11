// web/src/lib/guides/guide-header.ts
// The guide header's "Updated 22 Oct" stamp: the day the page's numbers were last refreshed,
// which is the later of the guide's own `updated` date and the day its ranked BiS file was
// published (the nightly republishes the weights, picks and rotation every time).
/** The later of the guide's `updated` and the BiS file's `generated_at` (ISO), when it has one. */
export function pageUpdatedAt(updated: Date, bisGeneratedAt: string | undefined): Date {
  if (bisGeneratedAt === undefined) return updated;
  const generated = new Date(bisGeneratedAt);
  if (Number.isNaN(generated.getTime())) return updated;
  return generated.getTime() > updated.getTime() ? generated : updated;
}
