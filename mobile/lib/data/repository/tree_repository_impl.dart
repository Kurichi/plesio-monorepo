import 'package:kiikuten/data/datasource/tree_datasource.dart';
import 'package:kiikuten/domain/entity/tree.dart';
import 'package:kiikuten/domain/repository/tree_repository.dart';

class TreeRepositoryImpl implements TreeRepository {
  final TreeDataSource _treeDataSource;

  TreeRepositoryImpl(this._treeDataSource);

  @override
  Future<Tree> getTree(String treeId) async {
    final treeModel = await _treeDataSource.getTree(treeId);
    return treeModel.toEntity();
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
