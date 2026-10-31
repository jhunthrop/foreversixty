import DesignSystem
import Models
import SwiftUI

/// F10.4, Premium only: the night's review printed as given, or the coach's own busy and error words.
struct CoachCard: View {
    let state: CoachState
    let clock: ReportsClock

    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            switch state {
            case .hidden:
                EmptyView()
            case .writing:
                line(Copy.coachWriting, color: FSColor.text)
            case .preparing:
                line(Copy.coachPreparing, color: FSColor.text)
            case .failed:
                line(Copy.coachFailed, color: FSColor.wipe)
            case let .ready(review):
                head(review)
                ForEach(Coach.sections(review.review)) { section in sectionView(section) }
                Text(ReportsCopy.coachFooter(review, clock: clock)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
            }
        }
        .padding(16)
        .frame(maxWidth: .infinity, alignment: .leading)
        .quietCard()
    }

    private func head(_ review: CoachReview) -> some View {
        ViewThatFits(in: .horizontal) {
            HStack(alignment: .firstTextBaseline) {
                title(review)
                Spacer(minLength: 8)
                TagView(text: Copy.coachPremium, color: FSColor.gold)
            }
            VStack(alignment: .leading, spacing: 6) {
                TagView(text: Copy.coachPremium, color: FSColor.gold)
                title(review)
            }
        }
    }

    private func title(_ review: CoachReview) -> some View {
        Text(Copy.coachTitle(spec: review.spec).uppercased()).font(ScaledType.sec).foregroundStyle(FSColor.strong)
            .fixedSize(horizontal: false, vertical: true)
            .accessibilityAddTraits(.isHeader)
    }

    private func sectionView(_ section: CoachSection) -> some View {
        VStack(alignment: .leading, spacing: 4) {
            MicroLabel(ReportsCopy.coachSection(section.kind))
            ForEach(Array(section.lines.enumerated()), id: \.offset) { _, line in
                VStack(alignment: .leading, spacing: 2) {
                    if let title = line.title {
                        Text(title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                            .fixedSize(horizontal: false, vertical: true)
                    }
                    Text(line.body).font(ScaledType.body).foregroundStyle(FSColor.text)
                        .fixedSize(horizontal: false, vertical: true)
                }
                .padding(.top, 2)
            }
        }
        .accessibilityElement(children: .combine)
    }

    private func line(_ text: String, color: Color) -> some View {
        Text(text).font(ScaledType.body).foregroundStyle(color).fixedSize(horizontal: false, vertical: true)
    }
}
