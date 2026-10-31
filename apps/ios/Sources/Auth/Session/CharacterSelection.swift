import Foundation
import Models

/// The current character (spec §3.4): the last one picked on this device, else the account's
/// main, else the most recently synced, else the first. Picking writes nothing to the server.
public enum CharacterSelection {
    public static func current(in me: Me, picked: String?) -> AccountCharacter? {
        let characters = me.characters
        if let picked, let match = characters.first(where: { $0.key == picked }) { return match }
        if let main = characters.first(where: { $0.key == me.mainCharacterKey }) { return main }
        let synced = characters.compactMap { character in character.build.map { (character, $0.capturedAt) } }
        return synced.max { $0.1 < $1.1 }?.0 ?? characters.first
    }
}
