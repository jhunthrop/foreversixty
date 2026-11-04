// web/src/lib/icons/placeholder.ts
// The client's own art for an id the pipeline has no real icon for (data/pipeline/icons.py).
// Every surface draws it as no icon, never a red question mark. Browser safe.
export const PLACEHOLDER_ICON = 'inv_misc_questionmark';

/** Whether `icon` is a real stem: present, not empty and not the placeholder. */
export function isRealIcon(icon: string | null | undefined): icon is string {
  return icon !== undefined && icon !== null && icon !== '' && icon !== PLACEHOLDER_ICON;
}
