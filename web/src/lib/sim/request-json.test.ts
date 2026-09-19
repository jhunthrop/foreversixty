import { describe, expect, it, vi } from 'vitest';
import envelope from '../../fixtures/sim/envelope-v2.json';
import {
  checkRequest,
  MAX_REQUEST_CHARS,
  formatRequest,
  parseRequest,
  settingsFromRequest,
} from './request-json';
import { defaultSettings } from './settings';
import type { SimRequest } from './types';

const request = (envelope as unknown as { request: SimRequest }).request;

describe('formatRequest', () => {
  it('is indented JSON a person can edit, ending in a newline', () => {
    const text = formatRequest(request);
    expect(text.startsWith('{\n  "engine_version"')).toBe(true);
    expect(text.endsWith('\n')).toBe(true);
  });

  it('round-trips through parseRequest unchanged', () => {
    const parsed = parseRequest(formatRequest(request));
    expect(parsed.ok).toBe(true);
    if (parsed.ok) expect(parsed.request).toEqual(request);
  });
});

describe('parseRequest', () => {
  it('names the syntax error rather than saying “invalid”', () => {
    const parsed = parseRequest('{ "spec": ');
    expect(parsed.ok).toBe(false);
    if (!parsed.ok) expect(parsed.message.length).toBeGreaterThan(0);
  });

  it('refuses anything that is not a JSON object', () => {
    for (const text of ['[]', '"a string"', '42', 'null']) {
      expect(parseRequest(text).ok, text).toBe(false);
    }
  });

  it('refuses a request past the size bound before parsing it', () => {
    const parsed = parseRequest('{'.padEnd(MAX_REQUEST_CHARS + 1, ' '));
    expect(parsed.ok).toBe(false);
  });
});

describe('settingsFromRequest', () => {
  it('takes the encounter verbatim', () => {
    expect(settingsFromRequest(request).encounter).toEqual(request.encounter);
  });

  it('takes the buffs, consumables and cooldowns, and calls the preset custom', () => {
    const settings = settingsFromRequest(request);
    expect(settings.buffs).toEqual(request.character.buffs);
    expect(settings.consumables).toEqual(request.character.consumes);
    expect(settings.cooldowns).toEqual(request.character.cooldowns);
    // A pasted list is nobody's preset, and calling it "raid-buffed" would let the next
    // preset change silently discard it.
    expect(settings.preset).toBe('custom');
  });

  it('fills an absent encounter field from the defaults rather than leaving it undefined', () => {
    const thin = { ...request, encounter: { ...request.encounter, target_level: undefined } };
    expect(settingsFromRequest(thin as SimRequest).encounter.target_level).toBe(
      defaultSettings().encounter.target_level,
    );
  });

  it('copies rather than sharing, so editing the page never edits the pasted request', () => {
    const settings = settingsFromRequest(request);
    expect(settings.buffs).not.toBe(request.character.buffs);
    expect(settings.encounter).not.toBe(request.encounter);
  });
});

describe('checkRequest', () => {
  it('reports a parse error without ever calling validate', async () => {
    const validate = vi.fn();
    const result = await checkRequest('{ "spec": ', validate);
    expect(result.status).toBe('parse-error');
    expect(validate).not.toHaveBeenCalled();
  });

  it("reports the engine's own field errors when it refuses the request", async () => {
    const text = formatRequest(request);
    const errors = [{ field: 'spec', message: 'spec is required' }];
    const result = await checkRequest(text, async () => ({ ok: false, errors }));
    expect(result).toEqual({ status: 'invalid', errors });
  });

  it('reports a request the engine accepts as valid, with the parsed request attached', async () => {
    const text = formatRequest(request);
    const result = await checkRequest(text, async () => ({ ok: true, errors: [] }));
    expect(result).toEqual({ status: 'valid', request });
  });

  it('treats a rejected validate call as an error to show, never as valid', async () => {
    const text = formatRequest(request);
    const result = await checkRequest(text, async () => {
      throw new Error('json: cannot unmarshal string into Go value of type int');
    });
    expect(result).toEqual({
      status: 'validate-error',
      message: 'json: cannot unmarshal string into Go value of type int',
    });
  });
});
