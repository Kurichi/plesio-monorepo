import 'package:kiikuten/domain/repository/item_repository.dart';
import 'package:kiikuten/domain/repository/tree_repository.dart';
import 'package:kiikuten/domain/repository/user_repository.dart';

class UseItemUseCase {
  final UserRepository userRepository;
  final ItemRepository itemRepository;
  final TreeRepository treeRepository;

  UseItemUseCase({
    required this.userRepository,
    required this.itemRepository,
    required this.treeRepository,
  });

  Future<void> execute(String userId, String itemId) async {
    final user = await userRepository.getUser(userId);
    final item = await itemRepository.getItem(itemId);
    await itemRepository.useItem(itemId);
    await treeRepository.growTree(user.tree, item.growthEffect);
  }
}
