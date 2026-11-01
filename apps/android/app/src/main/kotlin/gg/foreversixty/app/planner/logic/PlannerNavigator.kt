package gg.foreversixty.app.planner.logic

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.navigation.DeepLink
import java.util.Locale

/** A planner screen (spec §3.3): the editor (full screen), a shared build, Your builds. */
sealed interface PlannerScreen {
    data class Editor(val request: EditorRequest) : PlannerScreen
    data class Viewer(val id: String) : PlannerScreen
    data object Builds : PlannerScreen
}

/**
 * The planner's own stack over the app (spec F13): every entry point (Talents, a guide, rankings,
 * a link) pushes here, and the scaffold draws the top screen. Empty means the planner is closed.
 */
class PlannerNavigator {
    var stack: List<PlannerScreen> by mutableStateOf(emptyList())
        private set

    val top: PlannerScreen? get() = stack.lastOrNull()

    fun push(screen: PlannerScreen) {
        stack = stack + screen
    }

    fun pop() {
        stack = stack.dropLast(1)
    }

    /**
     * `/app/planner/{class}` opens the editor (a code loads its build), `/app/builds` Your builds
     * (signed in only) and `/b/{id}` the viewer. Each replaces whatever planner screens were open.
     */
    fun open(link: DeepLink, signedIn: Boolean): Boolean {
        stack = when (link) {
            is DeepLink.Planner -> listOf(PlannerScreen.Editor(request(link)))
            DeepLink.Builds -> if (signedIn) listOf(PlannerScreen.Builds) else return false
            is DeepLink.SharedBuild -> listOf(PlannerScreen.Viewer(link.id))
            else -> return false
        }
        return true
    }

    companion object {
        /** A planner link's request: with a code it loads that build, else the class's draft. */
        fun request(link: DeepLink.Planner): EditorRequest {
            val slug = link.className.lowercase(Locale.ROOT)
            val code = link.code
            return if (code.isNullOrBlank()) EditorRequest.Draft(slug, link.race) else EditorRequest.Code(slug, code, link.race)
        }
    }
}
