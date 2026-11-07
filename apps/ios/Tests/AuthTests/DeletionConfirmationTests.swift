@testable import Auth
import Testing

/// The server's rule for the typed confirmation (Android DeletionConfirmationTest mirrors it).
@Suite("Deletion confirmation")
struct DeletionConfirmationTests {
    private let tag = "Yell#1604"

    @Test("the BattleTag matches trimmed and in any case", arguments: ["Yell#1604", "  yell#1604\n", "YELL#1604"])
    func battleTag(typed: String) {
        #expect(DeletionConfirmation.matches(typed, battleTag: tag))
    }

    @Test("the word delete matches with or without a BattleTag")
    func word() {
        #expect(DeletionConfirmation.matches("delete", battleTag: tag))
        #expect(DeletionConfirmation.matches(" DELETE ", battleTag: nil))
    }

    @Test("anything else does not match", arguments: ["", "   ", "Yell", "Yell#16045", "delete my account"])
    func mismatch(typed: String) {
        #expect(!DeletionConfirmation.matches(typed, battleTag: tag))
    }

    @Test("a BattleTag typed for an account without one does not match")
    func noTag() {
        #expect(!DeletionConfirmation.matches("Yell#1604", battleTag: nil))
    }

    @Test("the field asks for the BattleTag, or delete when there is none")
    func expected() {
        #expect(DeletionConfirmation.expected(battleTag: tag) == tag)
        #expect(DeletionConfirmation.expected(battleTag: nil) == DeletionConfirmation.word)
        #expect(DeletionConfirmation.expected(battleTag: " ") == DeletionConfirmation.word)
    }
}
