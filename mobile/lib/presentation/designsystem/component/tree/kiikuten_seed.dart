import 'dart:math';

import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/designsystem/component/dot/brown_dot.dart';

class KiikutenSeed extends StatelessWidget {
  const KiikutenSeed({
    super.key,
    required this.size,
    this.dotCount = 16,
  });

  final double size;
  final int dotCount;

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: size,
      height: size,
      child: GridView.builder(
        physics: const NeverScrollableScrollPhysics(),
        gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
          crossAxisCount: dotCount,
        ),
        itemCount: dotCount * dotCount,
        itemBuilder: (context, index) {
          // 中心からの距離を計算
          final x = index % dotCount;
          final y = index ~/ dotCount;
          final center = (dotCount - 1) / 2;
          final distance = sqrt(pow(x - center, 2) + pow(y - center, 2));

          // 中心からの距離が一定値以下の場合のみドットを表示
          if (distance < center) {
            return const BrownDot();
          } else {
            return Container();
          }
        },
      ),
    );
  }
}
