import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/designsystem/component/dot/brown_dot.dart';

class KiikutenSeed extends StatelessWidget {
  const KiikutenSeed({
    super.key,
    required this.size,
  });

  final double size;

  @override
  Widget build(BuildContext context) {
    int squareSize = 10;

    return SizedBox(
      width: 80,
      height: 80,
      child: GridView.builder(
        gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
          crossAxisCount: squareSize,
        ),
        itemCount: squareSize * squareSize,
        itemBuilder: (context, index) {
          return const BrownDot();
        },
      ),
    );
  }
}
