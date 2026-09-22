// web/src/lib/addon/paste-layout.ts
// The height AddonPasteSave.svelte's signed-in form renders at (name field, region/ruleset
// selects, save button): a reasoned estimate from its own markup (label + input + two
// selects + button, each min-h-11 with gaps), reserved while AddonPasteBox.svelte's
// fetchMeOnce() is still resolving so the hint or the form does not push the links above it
// once it lands.
export const ADDON_PASTE_STATUS_MIN_H = 'min-h-[168px]';
