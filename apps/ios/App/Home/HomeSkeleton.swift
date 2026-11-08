import DesignSystem
import DesignTokens
import Models
import SwiftUI

/// F2.2: the band and the card as skeletons at their real heights, five rows, and the line.
struct HomeSkeleton: View {
    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            VStack(alignment: .leading, spacing: 10) {
                bar(width: 220, height: 13)
                bar(width: 160, height: 11).padding(.top, 4)
                bar(width: 200, height: 30)
                bar(width: 260, height: 13)
                bar(width: 180, height: 13)
            }
            .padding(EdgeInsets(top: 14, leading: 16, bottom: 18, trailing: 16))
            .frame(maxWidth: .infinity, alignment: .leading)
            .surfaceBand()
            VStack(spacing: 0) {
                ForEach(0..<HomePresenter.visibleRows, id: \.self) { _ in
                    HStack(alignment: .top, spacing: 12) {
                        RoundedRectangle(cornerRadius: FSRadius.legacyPill).fill(FSColor.cardTop).frame(width: 44, height: 44)
                        VStack(alignment: .leading, spacing: 6) {
                            bar(width: 180, height: 14)
                            bar(width: 140, height: 11)
                            bar(width: 200, height: 11)
                        }
                        Spacer()
                    }
                    .padding(.horizontal, 16)
                    .frame(height: 72)
                }
            }
            .padding(.vertical, 8)
            .quietCard()
            .padding(.horizontal, 16)
            .padding(.top, 16)
            Text(Copy.homeAnswerLoading).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                .padding(16)
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.homeAnswerLoading)
    }

    private func bar(width: CGFloat, height: CGFloat) -> some View {
        RoundedRectangle(cornerRadius: 4).fill(FSColor.cardTop).frame(width: width, height: height)
    }
}
