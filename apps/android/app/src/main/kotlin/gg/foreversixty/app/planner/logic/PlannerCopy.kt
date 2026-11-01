package gg.foreversixty.app.planner.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.talents.TalentRules

/** The planner's words from the shared catalog (`plan.*`). iOS: PlannerCopy.swift. */
class PlannerCopy(private val res: Resources) {
    /** The refusal sentence, byte for byte the web planner's (`rules.ts` `messages`). */
    fun refusal(refusal: Refusal): String = when (refusal) {
        is Refusal.TierLocked -> res.getString(R.string.plan_refuse_tierLocked, refusal.tier, refusal.treeName, refusal.points)
        is Refusal.PrereqMissing ->
            if (refusal.rank == 1) res.getString(R.string.plan_refuse_prereqOne, refusal.talentName, refusal.prereqName)
            else res.getString(R.string.plan_refuse_prereq, refusal.talentName, refusal.rank, refusal.prereqName)
        is Refusal.MaxRank -> res.getString(R.string.plan_refuse_maxRank, refusal.talentName, refusal.maxRank)
        Refusal.Cap -> res.getString(R.string.plan_refuse_cap, TalentRules.MAX_POINTS)
        is Refusal.UnknownTalent -> res.getString(R.string.plan_refuse_unknownTalent, refusal.id)
        is Refusal.NoPoints -> res.getString(R.string.plan_refuse_noPoints, refusal.talentName)
        Refusal.NoneLeft -> res.getString(R.string.plan_refuse_noneLeft)
    }

    /** `30 / 3 / 18`: a build's points per tree, the way the boards and the web write them. */
    fun totals(points: List<Int>): String = points.joinToString(TOTALS_SEPARATOR)

    private companion object {
        const val TOTALS_SEPARATOR = " / "
    }
}
