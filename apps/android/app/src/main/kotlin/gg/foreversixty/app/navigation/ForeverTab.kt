package gg.foreversixty.app.navigation

import androidx.annotation.DrawableRes
import androidx.annotation.StringRes
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.R as DesignR

/**
 * The app's four jobs (spec 2026-11-08 section 3): Home, Raid, Gear, Activity. Every other screen is
 * reached from the profile menu, an Activity row or a link, and sits under the tab it was opened on.
 */
// The icons are the owner's (spec "Tab-bar icons"), generated into the design system by
// web/scripts/gen-ui-icons.mjs from design/assets/icons/nav.
enum class ForeverTab(@param:StringRes val label: Int, @param:DrawableRes val icon: Int) {
    Home(R.string.tab_home, DesignR.drawable.ic_tab_home),
    Raid(R.string.tab_raid, DesignR.drawable.ic_tab_raid),
    Gear(R.string.tab_gear, DesignR.drawable.ic_tab_gear),
    Activity(R.string.tab_activity, DesignR.drawable.ic_tab_activity),
}

/** The tabs' counts (spec 3): Raid, your unanswered next raid or an officer's open items; Activity, unread. */
data class TabBadges(val raid: Int = 0, val activity: Int = 0) {
    fun of(tab: ForeverTab): Int = when (tab) {
        ForeverTab.Raid -> raid
        ForeverTab.Activity -> activity
        ForeverTab.Home, ForeverTab.Gear -> 0
    }
}
