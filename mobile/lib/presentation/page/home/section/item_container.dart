import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/page/home/component/item.dart';

class ItemContainer extends StatelessWidget {
  const ItemContainer({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.all(32),
      width: double.infinity,
      height: 100,
      decoration: BoxDecoration(
        color: Colors.white,
        border: Border.all(
          color: Theme.of(context).colorScheme.outline,
          width: 8,
        ),
        boxShadow: const [
          BoxShadow(
            color: Color(0x20000000),
            offset: Offset(8, 8),
          ),
        ],
      ),
      child: const SingleChildScrollView(
        padding: EdgeInsets.all(8),
        scrollDirection: Axis.horizontal,
        child: Row(
          mainAxisSize: MainAxisSize.min,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            Item(),
          ],
        ),
      ),
    );
  }
}
