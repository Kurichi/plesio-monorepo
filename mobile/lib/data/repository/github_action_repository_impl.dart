import 'package:kiikuten/domain/repository/github_action_repository.dart';

class GithubActionRepositoryImpl implements GithubActionRepository {
  @override
  Future<List<String>> detectActions(String userId) async {
    return ['commit', 'pull_request'];
  }
}
