// web/src/lib/character-layout.ts
// The height Character.svelte's ready view renders at for the one-fight fixture at 360px
// (measured with the same route-stub character-phone.spec.ts uses; see the plan step that
// derived this number). The loading Skeleton reserves it so the reveal changes only
// opacity -- current-character-layout.ts's own reason, same pattern.
export const CHARACTER_LOADING_MIN_H = 'min-h-[799px]';
