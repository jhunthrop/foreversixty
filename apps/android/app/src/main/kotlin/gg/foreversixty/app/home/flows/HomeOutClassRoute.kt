package gg.foreversixty.app.home.flows

import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.bis.BisSelection
import gg.foreversixty.app.bis.SpecCatalog

/**
 * A class tapped on signed-out Home (spec R3): that class's BiS at 60 with the raid preset, the spec page
 * pushed under Gear (its guide one tap away), in the faction the list already shows.
 */
class HomeOutClassRoute(
    private val specs: SpecCatalog,
    private val faction: () -> String,
    private val request: (BisSelection) -> Unit,
    private val pushSpecPage: () -> Unit,
    private val openGear: () -> Unit,
) {
    /** [specName] is the spec the row names (`Fury`); none, or one the catalogue does not know, is the class's first. */
    fun open(classSlug: String, specName: String? = null) {
        // A class the catalogue does not know (a broken bundle) still lands on Gear's class list.
        selection(classSlug, specName)?.let {
            request(it)
            pushSpecPage()
        }
        openGear()
    }

    fun selection(classSlug: String, specName: String?): BisSelection? {
        val spec = specName?.let { specs.spec(classSlug, it) } ?: specs.specsOf(classSlug).firstOrNull() ?: return null
        return BisSelection(classSlug, spec.spec, RAID_60, faction())
    }

    private companion object {
        val RAID_60 = BisRange(BisRange.LAST_BAND, BisRange.RAID)
    }
}
