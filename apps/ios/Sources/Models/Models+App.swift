import Foundation

// Hand-written behaviour on the generated models (Generated/APIModels.swift). Shapes live in
// api/openapi.yaml; only what the spec cannot say belongs here.

extension AccountCharacter: Identifiable {
    /// The character key, `region/ruleset/slug`.
    public var id: String { key }
}

extension AccountGuild: Identifiable {}

extension Device: Identifiable {}

public extension Me {
    /// The chosen main, else the first linked character.
    var mainCharacter: AccountCharacter? {
        characters.first { $0.key == mainCharacterKey } ?? characters.first
    }
}
