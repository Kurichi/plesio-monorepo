import 'dart:math';

import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/designsystem/component/dot/brown_dot.dart';
import 'package:kiikuten/presentation/designsystem/component/dot/green_dot.dart';

class KiikutenTree extends StatelessWidget {
  const KiikutenTree({
    super.key,
    required this.size,
  });

  final double size;

  @override
  Widget build(BuildContext context) {
    final dotCount = size ~/ 8;
    return SizedBox(
      width: size,
      height: size * 2,
      child: GridView.builder(
        physics: const NeverScrollableScrollPhysics(),
        gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
          crossAxisCount: dotCount,
        ),
        itemCount: dotCount * dotCount * 2,
        itemBuilder: (context, index) {
          // 水平方向中心からの距離を計算
          final x = index % dotCount;
          final center = (dotCount - 1) / 2;
          final distance = (x - center).abs();

          // 葉の部分
          final y = index ~/ dotCount;
          if (index < dotCount * dotCount * 1.5) {
            final leafWidth = max(0, y / 2.8);
            if (distance < leafWidth) {
              return const GreenDot();
            } else {
              return Container();
            }
          }

          // 幹の部分
          else if (distance < center / 2.4) {
            return const BrownDot();
          } else {
            return Container();
          }
        },
      ),
    );
  }
}
