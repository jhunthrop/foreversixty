<!-- web/src/components/guild/SignupSpec.svelte -->
<!-- The raid page's spec ask (guided-cards spec P.10.7). Asking (no saved spec): the label
     `Which spec tonight?` over one native select, open. Saved: a 44 px text button
     `Marksmanship · Change spec` (aria-expanded) that reveals the select and the
     `spec.lead.change` line; a pick collapses it again. The chosen spec's own icon sits beside
     it (a native select cannot draw one per option). The pick is the commit; no save button. -->
<script lang="ts">
  import { tick } from 'svelte';
  import { SEP, raidCopy, type SpecLead } from '../../lib/guild/raid-copy';
  import { specIconSrc, specOfClass, specsForClass } from '../../lib/guild/raid-spec';
  import Select from '../ui/Select.svelte';

  const ICON_PX = 20;

  let {
    id,
    classSlug,
    chosen,
    saved,
    lead,
    busy,
    disabled,
    onPick,
  }: {
    id: string;
    classSlug: string;
    /** The slug shown chosen: the saved spec, or the pick while it saves; `''` while asking. */
    chosen: string;
    /** A spec is already saved for this character: the control is collapsed. */
    saved: boolean;
    lead: SpecLead;
    /** The pick is saving: the select is disabled, and focus returns when it settles. */
    busy: boolean;
    disabled: boolean;
    onPick: (slug: string) => void;
  } = $props();

  let expanded = $state(false);
  let wasBusy = false;

  const toggleId = $derived(`${id}-toggle`);
  const panelId = $derived(`${id}-panel`);
  const leadId = $derived(`${id}-lead`);
  const options = $derived(specsForClass(classSlug).map((spec) => ({ value: spec.spec, label: spec.name })));
  const chosenSpec = $derived(specOfClass(classSlug, chosen));
  const collapsed = $derived(saved && !expanded && !busy);

  function pick(slug: string): void {
    if (slug !== '' && slug !== chosen) onPick(slug);
  }

  /** A keyboard pick disables the select and drops focus to the body: put it back when the save
   *  settles, on the toggle (the control collapses) or on the select (the save was refused). */
  $effect(() => {
    const finished = wasBusy && !busy;
    wasBusy = busy;
    if (!finished) return;
    if (saved) expanded = false;
    void tick().then(() => {
      if (document.activeElement !== document.body && document.activeElement !== null) return;
      document.getElementById(saved ? toggleId : id)?.focus();
    });
  });
</script>

{#snippet icon()}
  <span
    class="rounded-control bg-card-top inline-block shrink-0 overflow-hidden"
    style={`width:${ICON_PX}px;height:${ICON_PX}px`}
    aria-hidden="true"
  >
    {#if chosenSpec !== null}
      <img
        src={specIconSrc(chosenSpec)}
        alt=""
        width={ICON_PX}
        height={ICON_PX}
        class="h-full w-full object-cover"
        data-testid={`${id}-icon`}
      />
    {/if}
  </span>
{/snippet}

<div class="pt-2" data-testid={`${id}-field`}>
  {#if saved}
    <button
      id={toggleId}
      type="button"
      class="text-gold focus-visible:outline-gold rounded-control flex h-11 items-center gap-2 text-[13px] font-semibold focus-visible:outline-2 disabled:opacity-50"
      aria-expanded={expanded || busy}
      aria-controls={panelId}
      {disabled}
      data-testid={`${id}-toggle`}
      onclick={() => (expanded = !expanded)}
    >
      {@render icon()}
      <span>{chosenSpec?.name ?? ''}{SEP}{raidCopy.spec.change}</span>
    </button>
  {:else}
    <label for={id} class="label text-muted" data-testid={`${id}-label`}>{raidCopy.spec.title}</label>
  {/if}
  <div id={panelId} hidden={collapsed}>
    <div class="mt-1 flex items-center gap-2">
      {#if !saved}{@render icon()}{/if}
      <div class="min-w-0 flex-1">
        <Select
          {id}
          compact
          disabled={disabled || busy}
          {options}
          label={saved ? raidCopy.spec.change : undefined}
          placeholder={raidCopy.spec.placeholder}
          describedBy={leadId}
          bind:value={() => chosen, pick}
          testid={id}
        />
      </div>
    </div>
    <p id={leadId} class="text-muted mt-1 text-[12px] leading-4">
      {lead.before}<strong class="text-text font-semibold">{lead.name}</strong>{lead.after}
    </p>
  </div>
</div>
