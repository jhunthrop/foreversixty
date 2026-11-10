<!-- web/src/components/guild/setup/GuildSetup.svelte -->
<!-- /guild/<region>/<ruleset>/<name>/setup (design/specs/2026-11-09-guild-onboarding.md): the
     officer's checklist after a claim. The steps in order of value, each with its live state and,
     while it is to do, the controls that finish it in place; the first one to do carries the page's
     one primary. The Plan panel sits in the rail. Every state is the API's (GET …/setup), read
     again after each write and when the officer comes back to the window, so a bot install, a
     raid planned on the guild page or a member's sync all show without a reload. -->
<script lang="ts">
  import { guildHref, type CharacterPath } from '../../../lib/characters';
  import { onWindowFocus } from '../../../lib/data/query';
  import {
    GuildApiError,
    fetchGuildSettings,
    refetchGuildSettings,
    type GuildSettingsData,
  } from '../../../lib/guild/api';
  import { GUILD_LOADING } from '../../../lib/guild/layout';
  import { clockText, dateLabel, viewerLocale } from '../../../lib/guild/raid-time';
  import {
    clearSetupStep,
    fetchGuildSetup,
    markSetupStep,
    type GuildSetup,
    type SetupStep,
    type SetupStepKey,
  } from '../../../lib/guild/setup-api';
  import { setupCopy } from '../../../lib/guild/setup-copy';
  import { CONFIRMABLE, OPTIONAL, doneFigure, nextStep, stepMark } from '../../../lib/guild/setup-model';
  import { fetchGuild } from '../../../lib/rankings/api';
  import GuildStatus from '../../GuildStatus.svelte';
  import Figure from '../../ui/Figure.svelte';
  import Page from '../../ui/Page.svelte';
  import PageHeader from '../../ui/PageHeader.svelte';
  import Section from '../../ui/Section.svelte';
  import Steps from '../../ui/Steps.svelte';
  import SetupPlanRail from './SetupPlanRail.svelte';
  import SetupStepBody from './SetupStepBody.svelte';

  let { path }: { path: CharacterPath } = $props();

  let guildId = $state<number | null>(null);
  let guildName = $state('');
  let faction = $state<'alliance' | 'horde' | null>(null);
  let setup = $state<GuildSetup | null>(null);
  let settings = $state<GuildSettingsData | null>(null);
  let status = $state<'loading' | 'ready' | 'forbidden' | 'failed'>('loading');
  let busy = $state(false);
  let error = $state('');

  const guildPath = $derived(guildHref(path.region, path.ruleset, guildName || path.slug));
  const next = $derived(setup === null ? null : nextStep(setup.steps));
  const locale = viewerLocale();
  const viewerZone =
    typeof Intl === 'undefined' ? 'UTC' : Intl.DateTimeFormat().resolvedOptions().timeZone || 'UTC';
  const when = (iso: string): string => {
    const at = new Date(iso);
    return `${dateLabel(at, viewerZone, locale)}, ${clockText(at, viewerZone, locale)}`;
  };

  async function load(): Promise<void> {
    status = 'loading';
    try {
      const page = await fetchGuild(path);
      guildId = page.guild.id;
      guildName = page.guild.name;
      faction = page.guild.faction ?? null;
      [setup, settings] = await Promise.all([
        fetchGuildSetup(page.guild.id),
        fetchGuildSettings(page.guild.id, undefined, (fresh) => (settings = fresh)),
      ]);
      status = 'ready';
    } catch (thrown) {
      status = thrown instanceof GuildApiError && thrown.status === 403 ? 'forbidden' : 'failed';
    }
  }

  $effect(() => {
    void load();
  });

  // A bot install, a raid planned on the guild page or a member's sync happen elsewhere: the
  // checklist reads them again when the officer comes back.
  $effect(() =>
    onWindowFocus(() => {
      if (guildId === null || status !== 'ready') return;
      void refresh(guildId);
    }),
  );

  async function refresh(id: number): Promise<void> {
    try {
      setup = await fetchGuildSetup(id);
      refetchGuildSettings(id, (fresh) => (settings = fresh));
    } catch {
      // The checklist on screen stays; the next focus or write reads it again.
    }
  }

  /** Every write: busy while it runs, the checklist read again after, a failure said once. */
  function run(action: () => Promise<void>): void {
    if (guildId === null || busy) return;
    const id = guildId;
    busy = true;
    error = '';
    void action()
      .then(() => refresh(id))
      .catch((thrown: unknown) => {
        error = thrown instanceof Error && thrown.message !== '' ? thrown.message : setupCopy.actionFailed;
      })
      .finally(() => (busy = false));
  }

  const mark = (key: SetupStepKey, state: 'done' | 'skipped'): void =>
    run(async () => {
      if (guildId !== null) setup = await markSetupStep(guildId, key, state);
    });

  const undo = (key: SetupStepKey): void =>
    run(async () => {
      if (guildId !== null) setup = await clearSetupStep(guildId, key);
    });

  const items = $derived(
    (setup?.steps ?? []).map((s, i) => ({
      key: s.key,
      n: i + 1,
      mark: stepMark(s.state, s.key === next),
      title: setupCopy.titles[s.key],
    })),
  );
  const stepOf = (key: string): SetupStep | undefined => setup?.steps.find((s) => s.key === key);
</script>

{#snippet answer()}
  {#if setup !== null}
    {@const figure = doneFigure(setup)}
    <Figure value={figure.value} unit={figure.unit} align="end" testid="setup-done-figure" />
  {/if}
{/snippet}

{#snippet crumb()}<a href={guildPath} data-testid="setup-back">‹ {guildName || path.slug}</a>{/snippet}

<div class="reveal flex flex-col" data-testid="guild-setup">
  <PageHeader
    title={guildName === '' ? setupCopy.heading : guildName}
    tone={faction === null ? 'strong' : { faction }}
    {crumb}
    {answer}
    pending={status === 'loading' && guildName === ''}
    testid="guild-setup-header"
  />
  {#if status === 'loading' || status === 'failed'}
    <GuildStatus
      status={status === 'loading' ? 'loading' : 'failed'}
      error={setupCopy.failed}
      onRetry={() => void load()}
      lines={GUILD_LOADING.settings.lines}
      minHeight={GUILD_LOADING.settings.minHeight}
      testid="guild-setup"
    />
  {:else if status === 'forbidden'}
    <p class="px-(--page-inset) text-[14px]" data-testid="guild-setup-forbidden">{setupCopy.forbidden}</p>
  {:else if setup !== null && settings !== null && guildId !== null}
    {@const current = setup}
    {@const currentSettings = settings}
    {@const id = guildId}
    <Page layout="split" panels testid="guild-setup-page">
      <Section
        title={setupCopy.heading}
        meta={current.finished ? setupCopy.allDone : setupCopy.toDo(current.total - current.done)}
        panel
        testid="guild-setup-steps"
      >
        <Steps {items} testid="setup-step">
          {#snippet sub(item)}
            {@const step = stepOf(item.key)}
            {#if step !== undefined}
              {#if OPTIONAL.has(step.key)}<span class="text-muted text-[13px]">{setupCopy.optional}</span
                >{/if}
              <SetupStepBody
                {step}
                setup={current}
                settings={currentSettings}
                guildId={id}
                {guildName}
                {guildPath}
                primary={step.key === next}
                {busy}
                {when}
                {run}
                onConfirm={() => mark(step.key, 'done')}
                onSettings={(fresh) => (settings = fresh)}
              />
            {/if}
          {/snippet}
          {#snippet action(item)}
            {@const step = stepOf(item.key)}
            {#if step !== undefined}
              <span class="flex items-center gap-3">
                {#if step.state === 'skipped'}
                  <span class="text-muted text-[13px]" data-testid={`setup-step-${item.key}-state`}
                    >{setupCopy.states.skipped}</span
                  >
                {/if}
                {#if step.state === 'todo'}
                  <button
                    class="text-muted hover:text-text min-h-11 text-[13px] underline"
                    disabled={busy}
                    onclick={() => mark(step.key, 'skipped')}
                    data-testid={`setup-step-${item.key}-skip`}>{setupCopy.skip}</button
                  >
                {:else if step.marked && (step.state === 'skipped' || CONFIRMABLE.has(step.key))}
                  <button
                    class="text-muted hover:text-text min-h-11 text-[13px] underline"
                    disabled={busy}
                    onclick={() => undo(step.key)}
                    data-testid={`setup-step-${item.key}-undo`}>{setupCopy.undo}</button
                  >
                {/if}
              </span>
            {/if}
          {/snippet}
        </Steps>
      </Section>
      {#snippet rail()}
        <SetupPlanRail plan={current.plan} />
      {/snippet}
    </Page>
    <div class="min-h-[21px] px-(--page-inset) pt-3">
      {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="guild-setup-error">{error}</p>{/if}
    </div>
  {/if}
</div>
