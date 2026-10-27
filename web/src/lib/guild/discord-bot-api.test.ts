import { describe, expect, it } from 'vitest';
import { guildBotInstallUrl } from './discord-bot-api';

describe('guildBotInstallUrl', () => {
  it('builds the install URL on the API origin', () => {
    expect(guildBotInstallUrl(7, 'https://api.example')).toBe(
      'https://api.example/v1/guilds/7/discord/bot/install',
    );
  });
});
