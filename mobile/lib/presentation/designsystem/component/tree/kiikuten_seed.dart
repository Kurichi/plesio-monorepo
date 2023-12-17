import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/designsystem/component/dot/brown_dot.dart';

class KiikutenSeed extends StatelessWidget {
  const KiikutenSeed({
    super.key,
    required this.size,
    this.dotCount = 10,
  });

  final double size;
  final int dotCount;

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: size,
      height: size,
      child: GridView.builder(
        gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
          crossAxisCount: dotCount,
        ),
        itemCount: dotCount * dotCount,
        itemBuilder: (context, index) {
          return const BrownDot();
        },
      ),
    );
  }
}
