import 'package:kiikuten/domain/entity/tree.dart';
import 'package:kiikuten/domain/repository/tree_repository.dart';

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

  @override
  Future<void> updateTree(Tree tree) async {
    //
  }
}
