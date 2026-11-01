package gg.foreversixty.app.guides.logic

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.auth.session.StoredValue
import gg.foreversixty.app.navigation.DeepLink

/** A screen pushed on the Guides tab (spec §3.1): a class page, a guide, a changelog entry. */
sealed interface GuidesScreen {
    data class ClassPage(val classSlug: String) : GuidesScreen
    data class Guide(val spec: String) : GuidesScreen
    data class Change(val id: String) : GuidesScreen
}

/** The Guides tab's stack, and the guide links (spec §3.4) that open it. */
class GuidesNavigator {
    var stack: List<GuidesScreen> by mutableStateOf(emptyList())
        private set

    /** Set when a link asks for the Guides tab; the scaffold selects it and clears this. */
    var wantsTab: Boolean by mutableStateOf(false)

    val top: GuidesScreen? get() = stack.lastOrNull()

    fun push(screen: GuidesScreen) {
        stack = stack + screen
    }

    fun pop() {
        stack = stack.dropLast(1)
    }

    /** `/app/guides` opens the tab's root; `/app/guides/{class}/{spec}` that guide. */
    fun open(link: DeepLink): Boolean {
        stack = when (link) {
            DeepLink.Guides -> emptyList()
            is DeepLink.Guide -> listOf(GuidesScreen.Guide(Guides.specId(link.className, link.spec)))
            else -> return false
        }
        wantsTab = true
        return true
    }
}

/** The tab's segments (spec F9.1); Tier lists only behind `guide_tiers`. Persisted per device. */
enum class GuidesSegment(val key: String) {
    Guides("guides"),
    Tiers("tiers"),
    Changes("changes"),
    ;

    companion object {
        fun visible(tiersOn: Boolean): List<GuidesSegment> = if (tiersOn) entries else listOf(Guides, Changes)

        fun load(stored: StoredValue, tiersOn: Boolean = true): GuidesSegment =
            entries.firstOrNull { it.key == stored.load() }?.takeIf { it in visible(tiersOn) } ?: Guides

        fun save(stored: StoredValue, segment: GuidesSegment) = stored.save(segment.key)
    }
}
