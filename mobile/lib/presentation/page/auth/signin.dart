import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';

class SigninScreen extends StatelessWidget {
  const SigninScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Signin'),
      ),
      body: Center(
        child: FilledButton(
          onPressed: () async {
            await signInWithGitHub();
          },
          child: const Text('Sign in with GitHub'),
        ),
      ),
    );
  }

  Future<UserCredential> signInWithGitHub() async {
    final GithubAuthProvider githubProvider = GithubAuthProvider();
    final user = await FirebaseAuth.instance.signInWithPopup(githubProvider);
    final accessToken = user.credential?.accessToken;
    final String idToken =
        await FirebaseAuth.instance.currentUser!.getIdToken() ?? '';
    debugPrint('accessToken: $accessToken');
    debugPrint('idToken: $idToken');
    return user;
  }
}
