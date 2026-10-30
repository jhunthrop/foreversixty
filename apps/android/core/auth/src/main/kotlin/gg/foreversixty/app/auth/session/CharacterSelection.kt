package gg.foreversixty.app.auth.session

import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.Me

/**
 * The current character (spec §3.4): the last one picked on this device, else the account's
 * main, else the most recently synced, else the first. Picking writes nothing to the server.
 */
object CharacterSelection {
    fun current(me: Me, picked: String?): AccountCharacter? {
        val characters = me.characters
        characters.firstOrNull { it.key == picked }?.let { return it }
        characters.firstOrNull { it.key == me.mainCharacterKey }?.let { return it }
        // RFC 3339 UTC stamps sort as text.
        return characters.filter { it.build != null }.maxByOrNull { it.build?.capturedAt.orEmpty() }
            ?: characters.firstOrNull()
    }
}
