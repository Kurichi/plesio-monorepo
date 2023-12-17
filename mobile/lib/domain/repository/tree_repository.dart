import 'package:kiikuten/domain/entity/tree.dart';
import 'package:kiikuten/domain/entity/tree_ranking.dart'; // Assuming this entity exists

abstract class TreeRepository {
  Future<Tree> getMyTree();
  Future<Tree> getTreeByUserId(String userId);
  Future<List<TreeRanking>> getTreeRanking();
  Future<void> plantTree();
  Future<void> initTree();
}
