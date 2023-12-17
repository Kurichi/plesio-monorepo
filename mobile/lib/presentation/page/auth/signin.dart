import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';

class SigninScreen extends StatelessWidget {
  const SigninScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Row(
          mainAxisSize: MainAxisSize.min,
          children: [
            const Icon(Icons.park_outlined),
            const SizedBox(width: 8),
            Text(
              '木育展',
              style: TextStyle(color: Theme.of(context).colorScheme.primary),
            ),
          ],
        ),
      ),
      body: Stack(
        alignment: Alignment.center,
        children: [
          SizedBox(
            width: MediaQuery.of(context).size.width,
            height: MediaQuery.of(context).size.height,
            child: FittedBox(
              child: Icon(
                Icons.park,
                color: Theme.of(context).colorScheme.primary.withOpacity(0.08),
              ),
            ),
          ),
          Center(
            child: SingleChildScrollView(
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  RichText(
                    text: TextSpan(
                      style:
                          Theme.of(context).textTheme.headlineSmall?.copyWith(
                                letterSpacing: 2.0,
                              ),
                      children: [
                        const TextSpan(text: '自分の'),
                        TextSpan(
                          text: '木',
                          style: TextStyle(
                            fontWeight: FontWeight.bold,
                            color: Theme.of(context).colorScheme.primary,
                            fontSize: 36,
                          ),
                        ),
                        const TextSpan(text: 'を'),
                        TextSpan(
                          text: '育',
                          style: TextStyle(
                            fontWeight: FontWeight.bold,
                            color: Theme.of(context).colorScheme.primary,
                            fontSize: 36,
                          ),
                        ),
                        const TextSpan(text: 'てる'),
                        TextSpan(
                          text: '展',
                          style: TextStyle(
                            fontWeight: FontWeight.bold,
                            color: Theme.of(context).colorScheme.primary,
                            fontSize: 36,
                          ),
                        ),
                        const TextSpan(text: '覧会'),
                      ],
                    ),
                  ),
                  const SizedBox(height: 40),
                  FilledButton(
                    onPressed: () async {
                      await signInWithGitHub();
                    },
                    child: const Row(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Icon(FontAwesomeIcons.github),
                        SizedBox(width: 16, height: 48),
                        Text('Sign in with GitHub'),
                      ],
                    ),
                  ),
                ],
              ),
            ),
          ),
        ],
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
