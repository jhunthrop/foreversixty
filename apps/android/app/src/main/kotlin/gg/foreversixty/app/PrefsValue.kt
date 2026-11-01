package gg.foreversixty.app

import android.content.Context
import androidx.core.content.edit
import gg.foreversixty.app.auth.session.StoredValue

/** One small non-secret string in private prefs (the picked character key, spec §3.4). */
class PrefsValue(context: Context, private val key: String) : StoredValue {
    private val prefs = context.applicationContext.getSharedPreferences(FILE, Context.MODE_PRIVATE)

    override fun load(): String? = prefs.getString(key, null)

    override fun save(value: String?) = prefs.edit { if (value == null) remove(key) else putString(key, value) }

    companion object {
        private const val FILE = "forever_account"
        const val PICKED_CHARACTER = "picked_character"

        /** The Guides tab's segment (spec F9.1), kept per device. */
        const val GUIDES_SEGMENT = "guides_segment"
    }
}
