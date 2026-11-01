package gg.foreversixty.app.system.platform

import android.content.Intent
import android.os.Bundle
import android.widget.Toast
import androidx.activity.ComponentActivity
import androidx.core.net.toUri
import androidx.lifecycle.lifecycleScope
import gg.foreversixty.app.ForeverApplication
import gg.foreversixty.app.MainActivity
import gg.foreversixty.app.system.logic.SignupShortcutCopy
import gg.foreversixty.app.system.logic.SignupShortcutResult
import gg.foreversixty.app.widgets.WidgetRefreshWork
import kotlinx.coroutines.launch

/**
 * The `Sign up In` launcher shortcut (spec F15.1): a translucent activity with no UI, out of
 * Recents. It runs the sign-up and shows its toast, then finishes; it opens the app only for the
 * spec ask, two raids within 24 hours, or while the `system` switch is off.
 */
class SignupShortcutActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        val graph = (application as ForeverApplication).graph
        lifecycleScope.launch {
            val result = graph.system.signup.run()
            val answer = SignupShortcutCopy.answer(result, resources)
            answer.toast?.let { Toast.makeText(applicationContext, it, Toast.LENGTH_LONG).show() }
            when {
                answer.open != null -> open(Intent(Intent.ACTION_VIEW, answer.open.url().toUri(), this@SignupShortcutActivity, MainActivity::class.java))
                answer.openApp -> open(Intent(this@SignupShortcutActivity, MainActivity::class.java))
            }
            // Lane F14: the widgets, tile and countdown redraw from the server's answer.
            if (result is SignupShortcutResult.Signed && graph.widgets != null) WidgetRefreshWork.now(applicationContext)
            finish()
        }
    }

    private fun open(intent: Intent) {
        startActivity(intent.addFlags(Intent.FLAG_ACTIVITY_NEW_TASK or Intent.FLAG_ACTIVITY_CLEAR_TOP))
    }
}
