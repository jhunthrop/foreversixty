<!-- web/src/components/ui/UiSheet.svelte -->
<!-- The component sheet (polish spec 4.2): every shared control in each of its states on one
     page, served at /design-sheet. It is the living reference for design/DESIGN-SYSTEM.md and
     the page the axe, keyboard and capture specs drive (tests/e2e/ui-components.spec.ts).
     Hover, focus and active are real: use the pointer and the keyboard. -->
<script lang="ts">
  import { button } from '../../lib/ui/button';
  import ActionLink from './ActionLink.svelte';
  import CardHead from './CardHead.svelte';
  import Glyph from './Glyph.svelte';
  import SettingsRow from './SettingsRow.svelte';
  import Select from './Select.svelte';
  import Spinner from './Spinner.svelte';
  import StatTile from './StatTile.svelte';
  import Switch from './Switch.svelte';
  import Tag from './Tag.svelte';

  const TONES = ['primary', 'secondary', 'ghost', 'danger'] as const;
  const TONE_WORDS = {
    primary: 'Subscribe',
    secondary: 'Manage billing',
    ghost: 'See plans',
    danger: 'Remove',
  } as const;
  const REGIONS = [
    { value: 'us', label: 'US' },
    { value: 'eu', label: 'EU' },
  ];

  let simSwitch = $state(true);
  let privateSwitch = $state(false);
  let region = $state('us');
  let emptyRegion = $state('');
</script>

<div class="flex flex-col gap-12 pb-16" data-testid="ui-sheet">
  <section class="flex flex-col gap-4" aria-labelledby="sheet-buttons">
    <h2 id="sheet-buttons" class="heading-section">Buttons</h2>
    {#each TONES as tone (tone)}
      <div class="flex flex-wrap items-center gap-4" data-testid="sheet-buttons-{tone}">
        <button class={button(tone)}><Glyph name="card" />{TONE_WORDS[tone]}</button>
        <button class={button(tone)} disabled>{TONE_WORDS[tone]} (disabled)</button>
        <button class={button(tone)} disabled aria-busy="true"><Spinner />{TONE_WORDS[tone]} (loading)</button
        >
        <button class={button(tone, 'sm')}>{TONE_WORDS[tone]} (small)</button>
      </div>
    {/each}
  </section>

  <section class="flex flex-col gap-4" aria-labelledby="sheet-switch">
    <h2 id="sheet-switch" class="heading-section">Switch</h2>
    <div class="flex flex-wrap items-center gap-8" data-testid="sheet-switches">
      <span class="flex items-center gap-3">
        <span id="sw-off" class="text-[14px]">Off</span>
        <Switch checked={false} labelledby="sw-off" onchange={() => {}} testid="sheet-switch-off" />
      </span>
      <span class="flex items-center gap-3">
        <span id="sw-on" class="text-[14px]">On</span>
        <Switch checked={true} labelledby="sw-on" onchange={() => {}} testid="sheet-switch-on" />
      </span>
      <span class="flex items-center gap-3">
        <span id="sw-disabled" class="text-[14px]">Disabled</span>
        <Switch checked={false} disabled labelledby="sw-disabled" onchange={() => {}} />
      </span>
      <span class="flex items-center gap-3">
        <span id="sw-saving" class="text-[14px]">Saving</span>
        <Switch checked={true} saving labelledby="sw-saving" onchange={() => {}} />
      </span>
    </div>
  </section>

  <section class="flex flex-col gap-4" aria-labelledby="sheet-rows">
    <h2 id="sheet-rows" class="heading-section">Settings rows</h2>
    <div class="card-surface max-w-[520px] overflow-hidden" data-testid="sheet-rows">
      <SettingsRow label="Server sims" glyph="server" value="On · 3 runs this month" href="#sim" />
      <SettingsRow label="History charts" glyph="chart" value="Open on your main character" href="#history" />
      <SettingsRow label="Discord notifications" glyph="bell" value="Premium only" muted />
      <SettingsRow label="Loading value" glyph="star" loading />
      <SettingsRow label="Empty value" glyph="star" empty />
      <SettingsRow label="Error value" glyph="star" error="Could not load" />
      <SettingsRow label="Disabled row" glyph="star" value="Off" disabled />
      <SettingsRow label="A server sim finishes">
        {#snippet control(labelId)}
          <Switch
            checked={simSwitch}
            labelledby={labelId}
            onchange={(next) => (simSwitch = next)}
            testid="sheet-switch-live"
          />
        {/snippet}
      </SettingsRow>
      <SettingsRow label="Include private logs" sub="Off, private logs never post." nested>
        {#snippet control(labelId)}
          <Switch checked={privateSwitch} labelledby={labelId} onchange={(next) => (privateSwitch = next)} />
        {/snippet}
      </SettingsRow>
    </div>
  </section>

  <section class="flex flex-col gap-4" aria-labelledby="sheet-tags">
    <h2 id="sheet-tags" class="heading-section">Tags</h2>
    <div class="flex flex-wrap gap-3" data-testid="sheet-tags">
      <Tag>Founder</Tag>
      <Tag tone="success" dot>Connected</Tag>
      <Tag tone="muted">Sample</Tag>
      <Tag tone="danger">Needs attention</Tag>
    </div>
  </section>

  <section class="flex flex-col gap-4" aria-labelledby="sheet-links">
    <h2 id="sheet-links" class="heading-section">Links</h2>
    <ActionLink href="#specs">All 28 specs</ActionLink>
    <p class="text-[15px]">
      An inline link in prose reads like <a class="text-gold underline" href="#refunds">the refund policy</a>,
      with a soft gold underline.
    </p>
  </section>

  <section class="flex flex-col gap-4" aria-labelledby="sheet-stats">
    <h2 id="sheet-stats" class="heading-section">Stat tiles</h2>
    <div class="flex flex-wrap gap-3" data-testid="sheet-stats">
      <StatTile value="07" label="Days" />
      <StatTile value="3" label="Runs this month" />
      <StatTile value="88" label="Forever Score" />
    </div>
  </section>

  <section class="flex flex-col gap-4" aria-labelledby="sheet-select">
    <h2 id="sheet-select" class="heading-section">Select</h2>
    <div class="flex max-w-[320px] flex-col gap-4" data-testid="sheet-selects">
      <label class="label text-muted" for="sheet-region">Region</label>
      <Select id="sheet-region" bind:value={region} options={REGIONS} testid="sheet-select" />
      <label class="label text-muted" for="sheet-region-compact">Region (filter bar)</label>
      <Select id="sheet-region-compact" bind:value={region} options={REGIONS} compact />
      <label class="label text-muted" for="sheet-region-empty">Region (empty)</label>
      <Select
        id="sheet-region-empty"
        bind:value={emptyRegion}
        options={REGIONS}
        placeholder="Choose a region"
      />
      <label class="label text-muted" for="sheet-region-error">Region (error)</label>
      <Select
        id="sheet-region-error"
        value=""
        options={REGIONS}
        placeholder="Choose a region"
        error="Pick a region."
      />
      <label class="label text-muted" for="sheet-region-off">Region (disabled)</label>
      <Select id="sheet-region-off" value="us" options={REGIONS} disabled />
    </div>
  </section>

  <section class="flex flex-col gap-4" aria-labelledby="sheet-head">
    <h2 id="sheet-head" class="heading-section">Card head</h2>
    <div class="card-surface max-w-[520px]" data-testid="sheet-card-head">
      <CardHead glyph="crown" label="Plan" title="Premium">
        {#snippet titleSnippet()}<Tag>Founder</Tag>{/snippet}
        {#snippet aside()}<Tag tone="success" dot>Connected</Tag>{/snippet}
      </CardHead>
    </div>
  </section>
</div>
