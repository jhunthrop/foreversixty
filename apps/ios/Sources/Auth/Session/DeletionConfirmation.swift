import Foundation

/// The typed confirmation on the Delete account screen, by the server's rule for `DELETE /v1/me`:
/// the account's BattleTag or the word "delete", trimmed, in any case.
public enum DeletionConfirmation {
    /// What an account without a BattleTag types.
    public static let word = "delete"

    /// What the field asks for: the BattleTag, or the word when the account has none.
    public static func expected(battleTag: String?) -> String {
        guard let battleTag, !trimmed(battleTag).isEmpty else { return word }
        return battleTag
    }

    /// True when `typed` would be accepted, so the Delete account button can be enabled.
    public static func matches(_ typed: String, battleTag: String?) -> Bool {
        let candidate = trimmed(typed)
        guard !candidate.isEmpty else { return false }
        return [word, battleTag].compactMap { $0 }.contains { accepted in
            candidate.caseInsensitiveCompare(trimmed(accepted)) == .orderedSame
        }
    }

    private static func trimmed(_ text: String) -> String {
        text.trimmingCharacters(in: .whitespacesAndNewlines)
    }
}
