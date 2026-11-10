package gg.foreversixty.app

import android.content.res.Resources
import gg.foreversixty.app.auth.session.RelativeAge

/** The one relative-age phrase (`2 h ago`): Home's sync age and You's device rows. */
fun RelativeAge.phrase(res: Resources): String = when (this) {
    RelativeAge.Now -> res.getString(R.string.common_age_now)
    is RelativeAge.Minutes -> res.getString(R.string.common_age_minutes, n)
    is RelativeAge.Hours -> res.getString(R.string.common_age_hours, n)
    is RelativeAge.Days -> res.getString(R.string.common_age_days, n)
}

/** Past a week a sync stamp reads stale (round 3 row 12): the warning tone, Refresh beside it. */
const val STALE_DAYS = 7

fun RelativeAge.isStale(): Boolean = this is RelativeAge.Days && n > STALE_DAYS
