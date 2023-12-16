abstract class GithubActionRepository {
  Future<List<String>> detectActions(String userId);
}

class GithubActionRepositoryImpl implements GithubActionRepository {
  @override
  Future<List<String>> detectActions(String userId) async {
    return ['commit', 'pull_request'];
  }
}
