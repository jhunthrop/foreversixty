<!-- web/src/components/guild/SignupSpec.svelte -->
<!-- The raid page's spec ask (guided-cards spec P.10.7): one native select under the status
     segments, the class's specs by full name, the chosen spec's own icon beside it. Asking, its
     label is `Which spec tonight?`; with a saved spec the label is the `Change spec` link and the
     spec is chosen. The pick is the commit; there is no save button. -->
<script lang="ts">
  import { raidCopy, type SpecLead } from '../../lib/guild/raid-copy';
  import { specIconSrc, specOfClass, specsForClass } from '../../lib/guild/raid-spec';
  import Select from '../ui/Select.svelte';

  const ICON_PX = 20;

  let {
    id,
    classSlug,
    chosen,
    saved,
    lead,
    disabled,
    onPick,
  }: {
    id: string;
    classSlug: string;
    /** The slug shown chosen: the saved spec, or the pick while it saves; `''` while asking. */
    chosen: string;
    /** A spec is already saved for this character: the label becomes `Change spec`. */
    saved: boolean;
    lead: SpecLead;
    disabled: boolean;
    onPick: (slug: string) => void;
  } = $props();

  const options = $derived(specsForClass(classSlug).map((spec) => ({ value: spec.spec, label: spec.name })));
  const chosenSpec = $derived(specOfClass(classSlug, chosen));
  const label = $derived(saved ? raidCopy.spec.change : raidCopy.spec.title);

  function pick(slug: string): void {
    if (slug !== '' && slug !== chosen) onPick(slug);
  }
</script>

<div class="pt-2" data-testid={`${id}-field`}>
  <label for={id} class="label {saved ? 'text-gold' : 'text-muted'}" data-testid={`${id}-label`}
    >{label}</label
  >
  <div class="mt-1 flex items-center gap-2">
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
    <div class="min-w-0 flex-1">
      <Select
        {id}
        compact
        {disabled}
        {options}
        placeholder={raidCopy.spec.placeholder}
        bind:value={() => chosen, pick}
        testid={id}
      />
    </div>
  </div>
  <p class="text-muted mt-1 text-[12px] leading-4">
    {lead.before}<strong class="text-text font-semibold">{lead.name}</strong>{lead.after}
  </p>
</div>
