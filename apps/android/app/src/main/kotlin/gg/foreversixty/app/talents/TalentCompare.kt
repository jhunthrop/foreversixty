package gg.foreversixty.app.talents

import androidx.annotation.PluralsRes
import androidx.annotation.StringRes
import gg.foreversixty.app.R

/**
 * What F6's compare holds your build against, and so its words (lane F15, F13.5 follow-up): a
 * band's guide build (`talents.*`), or a shared build (`plan.cmp.*`, `Compare with mine`). One
 * compare draws both; only the copy differs. Each string takes the same arguments in the same
 * order (the shared ones simply leave the band and spec out).
 */
sealed interface CompareAgainst {
    /** The segment and tag of the other side: `Guide build` or `Shared build`. */
    @get:StringRes val side: Int

    /** The segment and tag of yours. */
    @get:StringRes val yours: Int

    /** `{spec}`, `{band}`: your build is it. */
    @get:StringRes val match: Int

    @get:StringRes val matchBeyond: Int

    /** `{n}`, `{m}`, `{band}`. */
    @get:StringRes val differ: Int

    /** `{n}`, `{band}`. */
    @get:StringRes val toTake: Int

    @get:PluralsRes val missing: Int

    @get:StringRes val beyond: Int

    @get:StringRes val legend: Int

    /** The rank sheet's line: `{g}`, `{max}`. */
    @get:StringRes val rank: Int

    /** A tab's TalkBack value with differing talents: `{n}`, `{k}`. */
    @get:StringRes val a11yTab: Int

    /** A differing cell's TalkBack label: `{name}`, `{n}`, `{max}`, `{g}`. */
    @get:StringRes val a11yCell: Int

    data object Guide : CompareAgainst {
        override val side = R.string.talents_guide
        override val yours = R.string.talents_yours
        override val match = R.string.talents_match
        override val matchBeyond = R.string.talents_matchBeyond
        override val differ = R.string.talents_differ
        override val toTake = R.string.talents_toTake
        override val missing = R.plurals.talents_missing
        override val beyond = R.string.talents_beyond
        override val legend = R.string.talents_legend
        override val rank = R.string.talents_guideRank
        override val a11yTab = R.string.talents_a11yTabDiffers
        override val a11yCell = R.string.talents_a11yCellGuide
    }

    /** A shared build (F13.5's `Compare with mine`); [yourName] is the character whose export it is. */
    data class Shared(val yourName: String) : CompareAgainst {
        override val side = R.string.plan_cmp_side
        override val yours = R.string.plan_cmp_yours
        override val match = R.string.plan_cmp_match
        override val matchBeyond = R.string.plan_cmp_matchBeyond
        override val differ = R.string.plan_cmp_differ
        override val toTake = R.string.plan_cmp_toTake
        override val missing = R.plurals.plan_cmp_missing
        override val beyond = R.string.plan_cmp_beyond
        override val legend = R.string.plan_cmp_legend
        override val rank = R.string.plan_cmp_rank
        override val a11yTab = R.string.plan_cmp_a11yTab
        override val a11yCell = R.string.plan_cmp_a11yCell
    }
}
