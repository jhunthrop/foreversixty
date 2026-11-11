package gg.foreversixty.app.raids.flows

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.RaidPingResult

/**
 * What a Discord mention did, in the words Not answered and the officer push (RAID_SHORT) both use:
 * `Mentioned 3 healers in Discord.`, then who was left out for a recent mention, or nobody to mention.
 */
class MentionWords(private val res: Resources) {
    /** [people] names who the mention was for (`3 healers`); null mentions everyone not answered. */
    fun line(result: RaidPingResult, people: String?): String {
        if (!result.posted) return res.getString(R.string.raid_na_noDiscord)
        val said = if (people != null && result.mentioned > 0) {
            res.getString(R.string.raid_na_mentioned, people)
        } else {
            res.getString(R.string.raid_na_mentionedAll, result.mentioned)
        }
        val recent = result.recentlyMentioned?.takeIf { it > 0 }?.let { res.getQuantityString(R.plurals.raid_na_recently, it, it) }
        return listOfNotNull(said, recent).joinToString(" ")
    }
}
