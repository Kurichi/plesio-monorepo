import 'package:kiikuten/data/datasource/tree_datasource.dart';
import 'package:kiikuten/domain/entity/tree.dart';
import 'package:kiikuten/domain/entity/tree_ranking.dart'; // Assuming this entity exists
import 'package:kiikuten/domain/repository/tree_repository.dart';

class TreeRepositoryImpl implements TreeRepository {
  final TreeDataSource _treeDataSource;

  TreeRepositoryImpl(this._treeDataSource);

  @override
  Future<Tree> getMyTree() async {
    final treeModel = await _treeDataSource.getMyTree();
    return treeModel.toEntity();
  }

  @override
  Future<Tree> getTreeByUserId(String userId) async {
    final treeModel = await _treeDataSource.getTreeByUserId(userId);
    return treeModel.toEntity();
  }

  @override
  Future<List<TreeRanking>> getTreeRanking() async {
    final treeRankingModels = await _treeDataSource.getTreeRanking();
    return treeRankingModels.map((model) => model.toEntity()).toList();
  }

  @override
  Future<void> plantTree() async {
    await _treeDataSource.plantTree();
  }

  @override
  Future<void> initTree() async {
    await _treeDataSource.initTree();
  }
}
