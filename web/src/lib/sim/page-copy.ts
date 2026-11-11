// web/src/lib/sim/page-copy.ts
// The simulator page's frame copy (spec 2026-11-06-sim §5.1, §6): the landing, the head's answer
// in every state, the run line and More settings. Verbatim from the spec; figures arrive as
// arguments, never typed here (prose never quotes a nightly number).
import { uiCopy } from '../ui/copy';

export const simPageCopy = {
  title: 'Simulator',
  /** The landing's lead: what a sim gives the player, in the game's words. */
  lead: 'Run your gear, talents, buffs and consumables against a boss. You get your DPS, a breakdown by ability and a cast-by-cast timeline of the fight.',
  sections: {
    simYourCharacter: 'Sim your character',
    otherWaysIn: 'Other ways in',
    bisSet: 'Sim a best-in-slot set',
    yourSims: 'Your sims',
    yourSimsAll: 'All',
  },
  paste: {
    label: 'Addon export or build code',
    placeholder: 'FS1:…',
    run: 'Run',
    hint: 'In game: /fs copy, then paste here.',
    setUp: 'Get set up ›',
    setUpHref: '/setup#paste',
    invalid: 'That does not look like an addon export or a build code.',
  },
  ways: {
    characters: 'Your characters',
    signIn: 'Sign in with Battle.net',
    build: 'A planner build',
    buildPlaceholder: 'foreversixty.gg/b/…',
    fight: 'A logged fight',
    fightPlaceholder: 'foreversixty.gg/reports/…?fight=2',
    load: 'Load',
  },
  noCharacters:
    'No characters yet: they are on their way from Battle.net. Paste an export below meanwhile, or check',
  noCharactersLink: 'your account.',
  noCharactersHref: '/account#characters',
  head: {
    meta: (race: string, spec: string, level: number, split: string, itemLevel: string): string =>
      `${race} ${spec} · Level ${level} · ${split} talents · item level ${itemLevel}`,
    metaNoGear: (race: string, spec: string, level: number, split: string): string =>
      `${race} ${spec} · Level ${level} · ${split} talents`,
    source: (from: string, date: string | null): string =>
      date === null ? `Gear from ${from}` : `Gear from ${from}, ${date}`,
    sourceKind: {
      addon: 'the addon',
      armory: 'Battle.net',
      blizzard: 'Battle.net',
      manual: 'a pasted export',
      build: 'a build',
      fight: 'a logged fight',
    } as Readonly<Record<string, string>>,
    savedMeta: (
      race: string,
      level: number,
      split: string,
      fight: string,
      buffs: string,
      consumables: number,
    ): string =>
      `${race} · Level ${level} · ${split} talents · ${fight} fight, ${buffs}, ${consumables} consumables`,
    saved: (date: string): string => `Saved ${date}`,
  },
  answer: {
    notSimulated: 'Not simulated yet',
    notSimulatedCaption: (spec: string): string => `${spec} is not simulated yet; there is nothing to run.`,
    specSupport: 'Spec support ›',
    specSupportHref: '/sim/specs',
    loadingEngine: 'Loading the engine…',
    loadingEngineCaption: 'The engine is about 4 MB. It loads once and is cached after that.',
    onServers: 'On our servers',
    starting: 'Starting',
    serverRunning: (done: string, total: string): string => `Running · ${done} of ${total} runs`,
    runningBrowser: 'Running in your browser',
    runningCaption: (band: string, done: string, total: string): string =>
      `± ${band} DPS · ${done} of ${total} runs`,
    /** A finished result names no lane: the band says how sure it is. */
    resultLabel: '',
    stoppedLabel: (done: string, total: string): string => `Stopped at ${done} of ${total} runs`,
    stoppedFigure: (estimate: string): string => `about ${estimate}`,
    stoppedCaption: (done: string, estimate: string, band: string): string =>
      `Stopped at ${done} runs: about ${estimate} ± ${band}. A stopped run is not saved.`,
    serverFailed: 'On our servers · failed',
    serverFailedReason: 'Our server could not finish this run.',
    serverFailedFix:
      'Press Run again. If it fails twice, run it in your browser: choose a precision under High.',
    browserFailed: 'In your browser · failed',
    browserRefused: (reason: string): string => `The engine could not run this character: ${reason}.`,
    browserRefusedFix: 'Paste a fresh addon export, or swap that item in Top gear, then press Run again.',
    saveThis: 'Save this sim',
    signInToSave: 'Sign in to save',
    saved: 'Saved',
    copyLink: 'Copy link',
    topGear: 'Find upgrades in Top Gear ›',
    openInSim: 'Open in the sim',
  },
  bar: {
    label: 'Sim choices',
    fight: 'Fight',
    length: 'Length',
    targets: 'Targets',
    buffs: 'Buffs',
    precision: 'Precision',
    buffOptions: { 'raid-buffed': 'Raid', bare: 'None', custom: 'Custom' } as Readonly<
      Record<string, string>
    >,
    precisionOptions: {
      fast: 'Fast, 500 runs',
      normal: 'Normal, 3,000 runs',
      high: 'High, 10,000 runs',
      'target-error': 'Until ±0.5%',
    } as Readonly<Record<string, string>>,
    moreSettings: 'More settings',
    runSim: 'Run sim',
    runAgain: 'Run again',
    tryAgain: uiCopy.retry,
    stop: 'Stop',
    loadingEngine: 'Loading engine…',
    runningOnServers: 'Running on our servers…',
    runOnServers: 'Run on our servers',
    busyTitle: 'Stop the run to change it',
  },
  more: {
    title: 'More settings',
    close: 'Close',
    execute: 'Execute phase',
    executeSub: (pct: number): string => `Last ${pct}% of the fight`,
    on: 'On',
    off: 'Off',
    target: 'Target',
    targetLevel: (level: number): string => `Level ${level} boss`,
    armorSub: (armor: string): string => `${armor} armor`,
    rotation: 'Rotation',
    rotationValue: (spec: string): string => `The ${spec} rotation`,
    request: 'Request',
    requestLine: 'The whole request this run sends, as JSON.',
    showIt: 'Show it ›',
    hideIt: 'Hide it',
    notify: 'Tell me when a server run finishes',
    serverLine: (ceiling: string): string => `This request on our servers, up to ${ceiling} runs.`,
    backToPreset: 'Back to the preset',
    done: 'Done',
  },
} as const;
