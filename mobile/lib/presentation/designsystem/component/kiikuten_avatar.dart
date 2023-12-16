import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';

class KiikutenAvatar extends StatelessWidget {
  const KiikutenAvatar({
    super.key,
    this.size = 36,
  });

  final double? size;

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        shape: BoxShape.circle,
        color: Theme.of(context).colorScheme.primary,
      ),
      clipBehavior: Clip.hardEdge,
      child: Image.network(
        FirebaseAuth.instance.currentUser!.photoURL ?? '',
        width: size,
        height: size,
      ),
    );
  }
}
