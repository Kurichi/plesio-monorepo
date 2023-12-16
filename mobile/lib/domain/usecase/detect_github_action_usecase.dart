import 'package:kiikuten/domain/repository/github_action_repository.dart';
import 'package:kiikuten/domain/repository/item_repository.dart';

class DetectGithubActionUseCase {
  final GithubActionRepository githubActionRepository;
  final ItemRepository itemRepository;

  DetectGithubActionUseCase(this.githubActionRepository, this.itemRepository);

  Future<void> execute(String userId) async {
    var actions = await githubActionRepository.detectActions(userId);
    for (var _ in actions) {
      // アクションに応じたアイテムをユーザーに付与するロジック
    }
  }
}
