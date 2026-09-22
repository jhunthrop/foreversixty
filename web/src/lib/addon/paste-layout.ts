// web/src/lib/addon/paste-layout.ts
// While AddonPasteBox.svelte's fetchMeOnce() is still resolving, a Skeleton reserves this
// height so neither branch it can land on -- AddonPasteSave.svelte's short signed-out hint
// or its taller signed-in form -- pushes the planner/sim links above it once it lands.
// Sized to the SHORTER of the two branches (the signed-out hint's own measured height,
// 19.5px for its one line of text-[13px], measured in a browser against
// addon-paste-signin-hint), not the taller form's:
// reserving for the taller branch (its own ~172.5px, once measured) looks safer on paper,
// but this component's real content SHRINKS to fit once fetchMeOnce resolves, so a
// too-tall reservation collapses visibly for the shorter branch -- worse than reserving
// nothing at all. Reserving the shorter branch instead means the common signed-out case
// resolves flat (no shift) and the signed-in case only ever grows downward, which reads
// as content arriving, not a skeleton collapsing.
export const ADDON_PASTE_STATUS_MIN_H = 'min-h-[19.5px]';
