import 'package:kiikuten/domain/entity/tree.dart';

abstract class TreeRepository {
  Future<Tree> getTree(String treeId);
  Future<void> growTree(Tree tree, int growthAmount);
  Future<void> updateTree(Tree tree);
}
