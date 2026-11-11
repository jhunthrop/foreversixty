import DesignSystem
import Models
import SwiftUI

/// An Activity row's body as the feed styles it (`segments`): names in class colour, items in
/// their quality colour, strong figures; the plain body when the row has none.
enum ActivityBody {
    enum Tone: Equatable {
        case plain
        case strong
        case className(String?)
        case quality(Int?)
    }

    struct Part: Equatable {
        let text: String
        let tone: Tone
    }

    static let nameStyle = "name"
    static let itemStyle = "item"
    static let strongStyle = "strong"

    /// The parts to draw; nil when the row carries no segments.
    static func parts(_ row: ActivityRow) -> [Part]? {
        guard let segments = row.segments, !segments.isEmpty else { return nil }
        return segments.map { segment in
            let tone: Tone = switch segment.style {
            case nameStyle: .className(segment.className)
            case itemStyle: .quality(segment.quality)
            case strongStyle: .strong
            default: .plain
            }
            return Part(text: segment.text, tone: tone)
        }
    }

    /// The row's figure (sim DPS, the guild ladder's), shown once and worded as the other figures
    /// (row 11), in the house format (apps/shared/formats.json: grouped and whole from 100): `204 DPS`.
    static func figure(_ row: ActivityRow) -> String? {
        row.figure.map { Copy.homeUpgradesGain(figure: ReportsCopy.houseFigure($0), unit: Copy.bisUnitDps) }
    }

    static func text(_ row: ActivityRow, fallback: String) -> Text {
        guard let parts = parts(row) else { return Text(fallback) }
        return parts.reduce(Text("")) { text, part in Text("\(text)\(styled(part))") }
    }

    private static func styled(_ part: Part) -> Text {
        switch part.tone {
        case .plain: Text(part.text)
        case .strong: Text(part.text).foregroundColor(FSColor.strong).fontWeight(.semibold)
        case let .className(slug): Text(part.text).foregroundColor(ClassColor.of(slug))
        case let .quality(value): Text(part.text).foregroundColor(ItemQuality(value: value).text)
        }
    }
}

/// One Activity row as Activity and Home's New for you draw it: the crest or the game icon, the
/// title, the styled body, the time, the unread dot, the figure at the right in tabular digits.
struct ActivityRowView: View {
    let row: ActivityRow
    /// The body when the row has no segments (Home words it for its line).
    let line: String
    var meta: String?
    let build: String
    let last: Bool

    var body: some View {
        PanelRow(
            // A row whose figure is its whole news (sim.finished) has no body: no empty line under the title.
            row.title, sub: ActivityBody.parts(row) == nil && line.isEmpty ? nil : ActivityBody.text(row, fallback: line), meta: meta,
            unread: row.read ? nil : Copy.activityUnread, last: last, value: ActivityBody.figure(row)
        ) {
            ActivityLead(row: row, build: build)
        }
    }
}
