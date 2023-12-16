import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/page/auth/loading.dart';
import 'package:kiikuten/presentation/page/auth/signin.dart';
import 'package:kiikuten/presentation/page/home/home.dart';

class KiikutenAuth extends StatelessWidget {
  const KiikutenAuth({super.key});

  @override
  Widget build(BuildContext context) {
    return StreamBuilder<User?>(
      stream: FirebaseAuth.instance.authStateChanges(),
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          return const LoadingScreen();
        } else if (snapshot.hasData) {
          return const HomeScreen();
        } else {
          return const SigninScreen();
        }
      },
    );
  }
}
