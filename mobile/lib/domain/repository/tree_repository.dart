import 'package:kiikuten/domain/entity/tree.dart';

abstract class TreeRepository {
  Future<Tree> getTree(String treeId);
  Future<void> growTree(Tree tree, int growthAmount);
}

class TreeRepositoryImpl implements TreeRepository {
  @override
  Future<Tree> getTree(String treeId) async {
    return Tree(
      id: '1',
      growthLevel: 1,
    );
  }

  @override
  Future<void> growTree(Tree tree, int growthAmount) async {
    tree.grow(growthAmount);
  }
}
