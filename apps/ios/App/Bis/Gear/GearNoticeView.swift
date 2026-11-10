import DesignSystem
import SwiftUI

/// Gear with no list for the spec and band (round 2), inside the Gear tab: the spec, the line naming
/// the list with Change (so the player can pick another), and one row saying why, with Try again
/// when the list did not load.
struct GearNoticeView: View {
    let notice: GearNotice
    let change: () -> Void
    let retry: () -> Void

    var body: some View {
        ScrollView {
            BandPage {
                EmptyView()
            } content: {
                VStack(alignment: .leading, spacing: 0) {
                    PageHeader(title: notice.title) { EmptyView() }
                    GearListLine(facts: notice.facts, change: change)
                }
                PanelBlock {
                    PanelRow(notice.row, sub: notice.sub.map { Text($0) }, last: true) {
                        EmptyView()
                    } trail: {
                        if notice.retry {
                            Button(Copy.commonTryAgain, action: retry).buttonStyle(FSButtonStyle(fullWidth: false))
                        }
                    }
                }
                .parityRegion(ParityRegionName(notice.retry ? "section-error" : "section-empty"))
            }
        }
        .screenGround()
    }
}
