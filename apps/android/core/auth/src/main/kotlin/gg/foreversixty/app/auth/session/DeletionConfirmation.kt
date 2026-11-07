package gg.foreversixty.app.auth.session

/**
 * The typed confirmation on the Delete account screen, by the server's rule for `DELETE /v1/me`
 * (iOS `DeletionConfirmation`): the account's BattleTag or the word "delete", trimmed, in any case.
 */
object DeletionConfirmation {
    /** What an account without a BattleTag types. */
    const val WORD = "delete"

    /** What the field asks for: the BattleTag, or the word when the account has none. */
    fun expected(battleTag: String?): String = battleTag?.takeIf { it.isNotBlank() } ?: WORD

    /** True when [typed] would be accepted, so the Delete account button can be enabled. */
    fun matches(typed: String, battleTag: String?): Boolean {
        val candidate = typed.trim()
        if (candidate.isEmpty()) return false
        return listOfNotNull(WORD, battleTag).any { candidate.equals(it.trim(), ignoreCase = true) }
    }
}
