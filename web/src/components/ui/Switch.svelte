<!-- web/src/components/ui/Switch.svelte -->
<!-- The on/off control (polish spec 4.2): `button role="switch" aria-checked`, a 44 x 24 track
     with a 44 x 44 hit area. Off is a dark track with a muted knob; on is the primary gradient
     with the gold glow and a dark knob. Space and Enter toggle it (a native button); disabled is
     not focusable; `saving` shows the new position at 60 % and ignores presses until the save
     answers. The caller names it with `labelledby` (the row's label) or `label`, and owns the
     state: this component only reports the requested value. -->
<script lang="ts">
  let {
    checked,
    onchange,
    labelledby,
    label,
    describedby,
    disabled = false,
    saving = false,
    testid,
  }: {
    checked: boolean;
    /** The value the press asks for; the caller saves it and passes the answer back as `checked`. */
    onchange: (next: boolean) => void;
    labelledby?: string;
    label?: string;
    describedby?: string;
    disabled?: boolean;
    saving?: boolean;
    testid?: string;
  } = $props();

  function press(): void {
    if (disabled || saving) return;
    onchange(!checked);
  }
</script>

<button
  type="button"
  role="switch"
  aria-checked={checked}
  aria-labelledby={labelledby}
  aria-label={label}
  aria-describedby={describedby}
  aria-busy={saving ? 'true' : undefined}
  {disabled}
  onclick={press}
  data-testid={testid}
  class="group focus-visible:outline-gold -my-2 -mr-2 inline-flex h-11 w-11 shrink-0 cursor-pointer items-center justify-center rounded-full bg-transparent p-0 focus-visible:outline-2 focus-visible:outline-offset-3 disabled:cursor-not-allowed disabled:opacity-45 {saving
    ? 'opacity-60'
    : ''}"
>
  <span
    class="group-[:not(:disabled)]:group-hover:border-gold/30 relative block h-6 w-11 rounded-full border transition-colors duration-[120ms] motion-reduce:transition-none {checked
      ? 'border-gold-deep shadow-glow-gold bg-[linear-gradient(180deg,#f0cc6c,#c99a3a)]'
      : 'border-[#303a50] bg-[#232a3a]'}"
    aria-hidden="true"
  >
    <span
      class="absolute top-1/2 block -translate-y-1/2 rounded-full transition-[left,width,background-color] duration-[120ms] group-active:w-5 motion-reduce:transition-none {checked
        ? 'bg-bg left-[22px] h-[18px] w-[18px] group-active:left-[20px]'
        : 'left-[2px] h-[18px] w-[18px] bg-[#8a8576]'}"
    ></span>
  </span>
</button>
