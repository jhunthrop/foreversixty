package gg.foreversixty.app.update

import android.content.Context
import android.content.SharedPreferences
import androidx.core.content.edit
import java.util.UUID

/**
 * The `X-App-Install` id: a random UUID made on first launch. It lives in its own preferences
 * file, excluded from backup and device transfer (res/xml), so it is tied to no user and an
 * uninstall or a new phone starts a new one.
 */
object InstallId {
    const val PREFS = "forever_install"
    private const val KEY = "install_id"

    fun current(context: Context): String =
        current(context.getSharedPreferences(PREFS, Context.MODE_PRIVATE))

    fun current(prefs: SharedPreferences): String {
        prefs.getString(KEY, null)?.takeIf { isUuid(it) }?.let { return it }
        val fresh = UUID.randomUUID().toString()
        prefs.edit { putString(KEY, fresh) }
        return fresh
    }

    private fun isUuid(value: String): Boolean = runCatching { UUID.fromString(value) }.isSuccess
}
