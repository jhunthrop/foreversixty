package discordbot

import (
	"context"
	"errors"
	"fmt"
)

const (
	globalScope      = "global"
	guildScopePrefix = "guild:"
)

// commandList is the commands to register: Commands() unless a test swapped it.
func (s *Service) commandList() []Command {
	if s.commandSet != nil {
		return s.commandSet()
	}
	return Commands()
}

// RegisterCommands bulk-overwrites Discord's commands: the released ones globally, the rest on
// the development server (skipped when none is configured). A scope is sent only when its hash
// differs from the stored one, and the hash is stored after Discord accepted it. A lock keeps
// two starting instances from racing. Scope errors are joined; the caller logs them and goes on.
func (s *Service) RegisterCommands(ctx context.Context) error {
	return s.d.Store.WithCommandsLock(ctx, func(ctx context.Context) error {
		cmds := s.commandList()
		errs := []error{s.registerScope(ctx, "", globalScope, commandsOf(cmds, true))}
		if s.d.CommandsGuildID != "" {
			errs = append(errs, s.registerScope(ctx, s.d.CommandsGuildID,
				guildScopePrefix+s.d.CommandsGuildID, commandsOf(cmds, false)))
		}
		return errors.Join(errs...)
	})
}

func (s *Service) registerScope(ctx context.Context, serverID, scope string, cmds []Command) error {
	hash := commandsHash(cmds)
	stored, err := s.d.Store.CommandsHash(ctx, s.d.ApplicationID, scope)
	if err != nil {
		return err
	}
	if stored == hash {
		return nil
	}
	if err := s.d.REST.PutCommands(ctx, serverID, cmds); err != nil {
		return fmt.Errorf("discordbot: register %s commands: %w", scope, err)
	}
	if err := s.d.Store.SaveCommandsHash(ctx, s.d.ApplicationID, scope, hash); err != nil {
		return err
	}
	s.log.Info("discordbot", "op", "register_commands", "scope", scope, "commands", len(cmds))
	return nil
}
